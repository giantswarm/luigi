package main

import (
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update .golden reference log files")

// Test_format uses golden files.
//
//  go test . -run Test_format -update
//
func Test_format(t *testing.T) {
	testCases := []struct {
		name               string
		inputFile          string
		inputGrep          string
		expectedGoldenFile string
		errorMatcher       func(error) bool
	}{
		{
			name:               "case 0: error",
			inputFile:          "error.in",
			expectedGoldenFile: "error.golden",
			inputGrep:          "",
			errorMatcher:       nil,
		},
		{
			name:               "case 1: legacy error",
			inputFile:          "legacy_error.in",
			expectedGoldenFile: "legacy_error.golden",
			inputGrep:          "",
			errorMatcher:       nil,
		},
		{
			name:               "case 2: non JSON",
			inputFile:          "non_json.in",
			expectedGoldenFile: "non_json.golden",
			inputGrep:          "",
			errorMatcher:       nil,
		},
		{
			name:               "case 3: grep hit",
			inputFile:          "grep_hit.in",
			expectedGoldenFile: "grep_hit.golden",
			inputGrep:          "level=warning,function=GetCurrentState",
			errorMatcher:       nil,
		},
		{
			name:               "case 4: grep miss",
			inputFile:          "grep_miss.in",
			expectedGoldenFile: "grep_miss.golden",
			inputGrep:          "level=debug",
			errorMatcher:       nil,
		},
		{
			name:               "case 5: deal with numbers values",
			inputFile:          "numbers.in",
			expectedGoldenFile: "numbers.golden",
			inputGrep:          "",
			errorMatcher:       nil,
		},
		{
			name:               "case 6: deal with null values",
			inputFile:          "null.in",
			expectedGoldenFile: "null.golden",
			inputGrep:          "",
			errorMatcher:       nil,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			disableColors(true)

			ctx := context.Background()

			inFile, err := os.Open(filepath.Join("testdata", tc.inputFile))
			if err != nil {
				t.Fatalf("err = %#q, want %#v", microerror.JSON(err), nil)
			}

			buf := new(bytes.Buffer)

			err = mainE(ctx, inFile, buf, tc.inputGrep)
			if err != nil {
				t.Fatalf("err = %#q, want %#v", microerror.JSON(err), nil)
			}

			out := buf.Bytes()

			p := filepath.Join("testdata", tc.expectedGoldenFile)

			if *update {
				err := ioutil.WriteFile(p, out, 0644) // nolint:gosec
				if err != nil {
					t.Fatalf("err = %#q, want %#v", microerror.JSON(err), nil)
				}
			}

			goldenFile, err := ioutil.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(out, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(out)))
			}
		})
	}
}
