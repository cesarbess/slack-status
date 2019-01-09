package setter

import (
	"fmt"
	"strings"
	"time"

	spin "github.com/tj/go-spin"
)

var waiting = false

func wait(message string, runner func()) {
	waiting = true
	s := spin.New()
	s.Set(spin.Box2)
	go func() {
		runner()
		waiting = false
	}()
	for waiting {
		fmt.Printf("\r %s  %s", s.Next(), message)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r%s\r", strings.Repeat(" ", len(message)+5))
}
