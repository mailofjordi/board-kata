package parser

import "testing"

func TestParse(t *testing.T) {
	tests := map[string]struct{ inputMessage, expected string }{
		"plain text":               {"vamos ultimos", "vamos ultimos"},
		"text has one URL":         {"aaaa https://friendsofgo.tech", "aaaa <a href='https://friendsofgo.tech'>https://friendsofgo.tech</a>"},
		"text has two URL":         {"aaaa https://friendsofgo.tech and http://pepito.com", "aaaa <a href='https://friendsofgo.tech'>https://friendsofgo.tech</a> and <a href='http://pepito.com'>http://pepito.com</a>"},
		"text has two author tags": {"aaaa @pepe and @popo", "aaaa <a href='https://fogo-parser.dev/pepe'>@pepe</a> and <a href='https://fogo-parser.dev/popo'>@popo</a>"},
		"text has two hashtags":    {"aaaa #pepe and #popo", "aaaa <a href='https://fogo-parser.dev/hash/pepe'>#pepe</a> and <a href='https://fogo-parser.dev/hash/popo'>#popo</a>"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Parse(tc.inputMessage)
			if got != tc.expected {
				t.Fatalf(
					"Testcase %s fail. Expected: %s, got: %s",
					name,
					tc.expected,
					got,
				)
			}
		})
	}
}
