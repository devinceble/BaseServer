package Helpers

import (
	"errors"
	"log"
	"os"
	"time"
)

var (
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white  = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

//BaseLog Logger
func BaseLog(Title string, Intro string, Path string, SubTitle string, Code int64, color int64, Error error) {
	var colors string
	if color == 0 {
		colors = green
	} else if color == 1 {
		colors = white
	} else if color == 2 {
		colors = yellow
	} else if color == 3 {
		colors = red
	}

	if Error == nil {
		Error = errors.New("")
	}

	stdlogger := log.New(os.Stdout, "", 0)
	stdlogger.Printf("[%s] %s\t%s %v | %s |%s %d %s| %v",
		Title,
		Intro,
		Path,
		time.Now().Format("2006/01/02 - 15:04:05"),
		SubTitle,
		colors, Code, reset,
		Error)
}

//BaseLog Logger
func BaseInfo(Title string, Content string) {

	stdlogger := log.New(os.Stdout, "", 0)
	stdlogger.Printf("[%s %s %s] %s", green, Title, reset, Content)
}
