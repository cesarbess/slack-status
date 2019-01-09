package parser

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"time"

	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v2"
)

var (
	ErrNotFound           = fmt.Errorf("Cannot find ~/.slack-status. Please refer to the README")
	ErrNoToken            = fmt.Errorf("Cannot find 'token' section in your ~/.slack-status. Please refer to the README")
	ErrInvalidTokenFormat = fmt.Errorf("Invalid 'token' section. Please refer to the README")
	ErrEmptyToken         = fmt.Errorf("Your 'token' section is empty. Please refer to the README")
	ErrNoDefault          = fmt.Errorf("To use the parameterless form of slack-status, set a 'default' status in your ~/.slack-status file. Please refer to the README for more information")
)

var statusRegexp = regexp.MustCompile(`^:([^:]+):\s?(.*)$`)

type Status struct {
	Emoji string
	Text  string
}

func (s *Status) String() string {
	return fmt.Sprintf("(%s) %s", s.Emoji, s.Text)
}

type StatusArray []Status

func (a StatusArray) PickRandom() *Status {
	len := len(a)
	if len < 1 {
		return nil
	}
	rand.Seed(time.Now().Unix())
	idx := rand.Int() % len
	return &a[idx]
}

type Config struct {
	Tokens   map[string]string
	Statuses map[string]StatusArray
}

func (c Config) TeamNames() []string {
	names := make([]string, len(c.Tokens))
	i := 0
	for k := range c.Tokens {
		names[i] = k
		i++
	}
	return names
}

func (c Config) PrettyTeamNames() []string {
	names := c.TeamNames()
	pretty := make([]string, len(names))
	colorize := color.New(color.FgCyan)

	for i := 0; i < len(names); i++ {
		pretty[i] = colorize.Sprintf(names[i])
	}

	return pretty
}

func parseStatus(str string) *Status {
	result := statusRegexp.FindStringSubmatch(str)
	if len(result) == 0 {
		fmt.Printf("Warning: Skipping status with invalid format: %s\n", str)
		return nil
	}
	return &Status{
		Emoji: result[1],
		Text:  result[2],
	}
}

func LoadConfig() (*Config, error) {
	usr, _ := user.Current()
	statusFile := filepath.Join(usr.HomeDir, ".slack-status")
	data, err := ioutil.ReadFile(statusFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	var parsed map[string]interface{}
	err = yaml.Unmarshal(data, &parsed)
	if err != nil {
		return nil, fmt.Errorf("Invalid ~/.slack-status format. Please refer to the README.\n\n%s", err)
	}

	config := &Config{
		Tokens:   map[string]string{},
		Statuses: map[string]StatusArray{},
	}

	// Parsing token...
	rawToken, ok := parsed["token"]
	if !ok {
		return nil, ErrNoToken
	}

	token, ok := rawToken.(map[interface{}]interface{})
	if !ok {
		return nil, ErrInvalidTokenFormat
	}

	for rk, rv := range token {
		if k, ok := rk.(string); ok {
			if v, ok := rv.(string); ok {
				config.Tokens[k] = v
			} else {
				return nil, ErrInvalidTokenFormat
			}
		} else {
			return nil, ErrInvalidTokenFormat
		}
	}

	if len(config.Tokens) == 0 {
		return nil, ErrEmptyToken
	}

	// Parse statuses
	for k, rv := range parsed {
		if k == "token" {
			continue
		}

		statuses := []Status{}
		parseValue(rv, &statuses)
		if len(statuses) > 0 {
			config.Statuses[k] = statuses
		}
	}

	return config, nil
}

func parseValue(value interface{}, statuses *[]Status) {
	switch v := value.(type) {
	case string:
		status := parseStatus(v)
		if status != nil {
			*statuses = append(*statuses, *status)
		}
	case []interface{}:
		for _, rs := range v {
			if s, ok := rs.(string); ok {
				parseValue(s, statuses)
			}
		}
	}
}
