package main

import (
	"bytes"
)

// splitter is a struct exposing Split method compatible with bufio.Scanner.
type splitter struct {
	isTextLine bool
}

func newSplitter() *splitter {
	return &splitter{}
}

// Split is designed for bufio.Scanner to optimize printing interactive input
// piped trough luigi. E.g. if the input has a JSON line luigi needs the full
// line to be able to parse it but if the input are dots "." indicating that
// operation is in progress they should be printed immediately for good UX.
//
// Split returns a full line if data contains "\n". It returns what is
// available in the data if line in the input does not start with "{". If the
// line in the input starts with "{" it waits for the line termination and
// scans the full line.
func (s *splitter) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// We have full line. Return it.
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		s.isTextLine = false
		return i + 1, data[0 : i+1], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// This line is potential non-terminated JSON. Request more data.
	if data[0] == '{' && !s.isTextLine {
		return 0, nil, nil
	}

	// This is not a JSON line so return the data ASAP to make things more
	// interactive.
	s.isTextLine = true
	return len(data), data, nil
}
