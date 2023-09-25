package color

import "fmt"

func Blue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}
