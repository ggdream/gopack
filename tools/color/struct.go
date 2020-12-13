package color

import "fmt"

func Color(way, front, back int) string {
	//return fmt.Sprintf("\033[%d;%d;%dm", way, front, back)
	return fmt.Sprintf("\033[%d;%dm", way, front)
}
