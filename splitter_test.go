package main

import (
	"bufio"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"
)

func Test_split(t *testing.T) {

	testCases := []struct {
		name            string
		inputReaderFunc func() io.Reader
		expectedScans   []string
	}{
		{
			name: "case 0: read everything at once",
			inputReaderFunc: func() io.Reader {
				lines := []string{
					"a" + "\n",
					"b" + "\n",
					`{"a": "b"}` + "\n",
					`text {"a": "b"}` + "\n",
					`{"a": "b"} text` + "\n",
				}

				var r io.Reader
				r = strings.NewReader(strings.Join(lines, ""))

				return r
			},
			expectedScans: []string{
				"a" + "\n",
				"b" + "\n",
				`{"a": "b"}` + "\n",
				`text {"a": "b"}` + "\n",
				`{"a": "b"} text` + "\n",
			},
		},
		{
			name: "case 1: read byte by byte",
			inputReaderFunc: func() io.Reader {
				lines := []string{
					"a" + "\n",
					"b" + "\n",
					`{"a": "b"}` + "\n",
					`text {"a": "b"}` + "\n",
					`{"a": "b"} text` + "\n",
				}

				var r io.Reader
				r = strings.NewReader(strings.Join(lines, ""))

				// Instrument the reader to return byte by byte.
				r = iotest.OneByteReader(r)

				return r
			},
			expectedScans: []string{
				"a",
				"\n",
				"b",
				"\n",
				`{"a": "b"}` + "\n",
				"t",
				"e",
				"x",
				"t",
				" ",
				`{`,
				`"`,
				`a`,
				`"`,
				`:`,
				` `,
				`"`,
				`b`,
				`"`,
				`}`,
				"\n",
				`{"a": "b"} text` + "\n",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var scanner *bufio.Scanner
			{
				scanner = bufio.NewScanner(tc.inputReaderFunc())

				// `split` is the function to be tested.
				scanner.Split(newSplitter().Split)
			}

			var scans []string
			{
				for scanner.Scan() {
					scans = append(scans, scanner.Text())
				}
			}

			if !reflect.DeepEqual(scans, tc.expectedScans) {
				t.Fatalf("\nscans = %#v\n   want %#v", scans, tc.expectedScans)
			}
		})
	}
}
