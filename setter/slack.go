package setter

import (
	"fmt"
	"strings"
	"sync"

	"github.com/levigross/grequests"
	"github.com/victorgama/slack-status/parser"
)

func set(status *parser.Status, tokens map[string]string) map[string]error {
	w := sync.WaitGroup{}
	w.Add(len(tokens))
	lock := sync.Mutex{}
	errors := map[string]error{}
	for k, v := range tokens {
		go func(k, v string) {
			resp, err := grequests.Post("https://slack.com/api/users.profile.set", &grequests.RequestOptions{
				Headers: map[string]string{
					"Authorization": "Bearer " + v,
					"Content-Type":  "application/json; charset=utf-8",
				},
				JSON: map[string]interface{}{
					"profile": map[string]interface{}{
						"status_text":       status.Text,
						"status_emoji":      fmt.Sprintf(":%s:", status.Emoji),
						"status_expiration": 0,
					},
				},
			})
			if err != nil {
				lock.Lock()
				errors[k] = err
				lock.Unlock()
			}
			if resp.StatusCode != 200 {
				lock.Lock()
				errors[k] = fmt.Errorf("Unexpected response from Slack")
				lock.Unlock()
			}
			w.Done()
		}(k, v)
	}
	w.Wait()

	return errors
}

// SetStatus sets the provided Status to all teams loaded in the provided config
// object
func SetStatus(status *parser.Status, config *parser.Config) error {
	var errors map[string]error
	wait("Just a moment...", func() {
		errors = set(status, config.Tokens)
	})

	if len(errors) > 0 {
		teams := []string{}
		for k, v := range errors {
			teams = append(teams, fmt.Sprintf("  - %s: %s", k, v.Error()))
		}
		return fmt.Errorf("Failed setting status on the following teams:\n%s", strings.Join(teams, "\n"))
	}

	return nil
}
