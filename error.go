package main

import "github.com/giantswarm/microerror"

var grepNotFoundError = microerror.New("grep not found")

// IsGrepNotFound asserts grepNotFoundError.
func IsGrepNotFound(err error) bool {
	return microerror.Cause(err) == grepNotFoundError
}

var parseError = microerror.New("parsing")

// IsParseError asserts parseError.
func IsParseError(err error) bool {
	return microerror.Cause(err) == parseError
}
