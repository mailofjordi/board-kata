package parser

import (
	"regexp"
)

const (
	baseURL    = "https://fogo-parser.dev/"
	hashTagURL = "https://fogo-parser.dev/hash/"
)

func Parse(msg string) (output string) {
	pattern := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	msg = pattern.ReplaceAllString(msg, "<a href='$1://$2'>$1://$2</a>")

	pattern = regexp.MustCompile(`(@)(\w*)`)
	msg = pattern.ReplaceAllString(msg, "<a href='"+baseURL+"$2'>$1$2</a>")

	pattern = regexp.MustCompile(`(#)(\w*)`)
	msg = pattern.ReplaceAllString(msg, "<a href='"+hashTagURL+"$2'>$1$2</a>")
	return msg
}
