package s13_04_parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicate an error converting a word into an integer
type ParseError struct {
	index int    // The index into the space-separated list of words
	Word  string // The word that generated the parse error.
	Error error  // The raw error that precipitated this error, if any.
}

// String returns a human-readable error messag3
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in input as integers
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields) // here panic can occur
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no word to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
