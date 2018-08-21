package pkg

import (
	"reflect"
	"testing"
)

func Test_Grep(t *testing.T) {
	testCases := []struct {
		name        string
		expr        string
		in          map[string]string
		expectedOut bool
		isError     bool
	}{
		{
			name:        "case 0",
			expr:        "a=b,a=c",
			in:          map[string]string{"a": "b"},
			expectedOut: true,
		},
		{
			name:        "case 1",
			expr:        "a=b,a=c",
			in:          map[string]string{"a": "c", "x": "y"},
			expectedOut: true,
		},
		{
			name:        "case 2",
			expr:        "a=b,a=c",
			in:          map[string]string{"a": "d", "x": "y"},
			expectedOut: false,
		},
		{
			name:        "case 3",
			expr:        "a=b,a=d,x=y,x=z",
			in:          map[string]string{"a": "d", "x": "y"},
			expectedOut: true,
		},
		{
			name:        "case 4",
			expr:        "a!=b,x=y",
			in:          map[string]string{"a": "d", "x": "y"},
			expectedOut: true,
		},
		{
			name:        "case 5",
			expr:        "a!=b,x=y",
			in:          map[string]string{"a": "b", "x": "y"},
			expectedOut: false,
		},
		{
			name:        "case 6",
			expr:        "x!=y,x!=z",
			in:          map[string]string{"a": "d", "x": "y"},
			expectedOut: false,
		},
		{
			name:        "case 7",
			expr:        "x!=y,x!=z",
			in:          map[string]string{"a": "d", "x": "z"},
			expectedOut: false,
		},
		{
			name:        "case 8",
			expr:        "abc=xyz",
			in:          map[string]string{"a": "d", "abc": "xyz", "x": "z"},
			expectedOut: true,
		},
		{
			name:        "case 9",
			expr:        "abc!=xyz",
			in:          map[string]string{"a": "d", "abc": "xyz", "x": "z"},
			expectedOut: false,
		},
		{
			name:        "case 10",
			expr:        "abc=xyz,abc!=xyz",
			in:          map[string]string{"a": "d", "abc": "xyz", "x": "z"},
			expectedOut: false,
		},
		{
			name:        "case 11",
			expr:        "abc=xyz,abc!=xyz",
			in:          map[string]string{"a": "d", "x": "z"},
			expectedOut: false,
		},
		{
			name:        "case 12",
			expr:        "x=y,z!=t",
			in:          map[string]string{"a": "b", "c": "d"},
			expectedOut: false,
		},
		{
			name:        "case 13",
			expr:        "",
			in:          map[string]string{"a": "b", "c": "d"},
			expectedOut: true,
		},
		{
			name:        "case 14",
			expr:        "a=",
			in:          map[string]string{"a": ""},
			expectedOut: true,
		},
		{
			name:        "case 15",
			expr:        "a=",
			in:          map[string]string{"a": "b"},
			expectedOut: false,
		},
		{
			name:        "case 16",
			expr:        "a=,b=c",
			in:          map[string]string{"a": ""},
			expectedOut: false,
		},
		{
			name:        "case 17",
			expr:        "a=,b=c",
			in:          map[string]string{"a": "", "b": "c"},
			expectedOut: true,
		},
		{
			name:        "case 18: spaces between commas",
			expr:        " a   = b, x   != y ",
			in:          map[string]string{"a": "b", "x": "z"},
			expectedOut: true,
		},
		{
			name:    "case 19",
			expr:    "_",
			in:      map[string]string{},
			isError: true,
		},
		{
			name:    "case 20",
			expr:    ",",
			in:      map[string]string{},
			isError: true,
		},
		{
			name:    "case 21",
			expr:    "a",
			in:      map[string]string{},
			isError: true,
		},
		{
			name:    "case 22",
			expr:    "a,b",
			in:      map[string]string{},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := NewGrep(tc.expr)

			switch {
			case err == nil && !tc.isError:
				// correct; carry on
			case err != nil && !tc.isError:
				t.Fatalf("error == %s, want nil", err)
			case err == nil && tc.isError:
				t.Fatalf("error == nil, want non-nil")
			}

			if tc.isError {
				return
			}

			out := g.Filter(tc.in)
			if !reflect.DeepEqual(out, tc.expectedOut) {
				t.Fatalf("out == %t, want %t", out, tc.expectedOut)
			}
		})
	}
}
