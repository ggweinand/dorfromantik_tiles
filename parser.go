package main

import (
	"regexp"
)

func QueryToRegexp(query string) *regexp.Regexp {
	regexpStr := ""
	for _, char := range query {
		switch char {
		case 'W', 'G', 'V', 'S', 'T':
			regexpStr += string(char)
		case 'A':
			regexpStr += "[WGVSTM]"
		}
	}
	return regexp.MustCompile(regexpStr)
}
