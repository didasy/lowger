package lowger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type ConsoleLogger struct{}

func (c *ConsoleLogger) Error(messages ...interface{}) {
	fmt.Println(c.createMessage("[ERROR]", time.Now(), messages...))
}

func (c *ConsoleLogger) Warning(messages ...interface{}) {
	fmt.Println(c.createMessage("[WARNING]", time.Now(), messages...))
}

func (c *ConsoleLogger) Info(messages ...interface{}) {
	fmt.Println(c.createMessage("[INFO]", time.Now(), messages...))
}

func (c *ConsoleLogger) Fatal(messages ...interface{}) {
	fmt.Println(c.createMessage("[FATAL]", time.Now(), messages...))
	os.Exit(1)
}

func (c *ConsoleLogger) createMessage(prefix string, now time.Time, messages ...interface{}) string {
	data := []string{
		prefix,
		now.Format(time.RFC822Z),
		"-",
	}
	for _, msg := range messages {
		data = append(data, fmt.Sprint(msg))
	}

	return strings.Join(data, " ")
}
