package main

import "github.com/giantswarm/microerror"

var grepNotFoundError = &microerror.Error{
	Kind: "grepNotFoundError",
}

// IsGrepNotFound asserts grepNotFoundError.
func IsGrepNotFound(err error) bool {
	return microerror.Cause(err) == grepNotFoundError
}

var parseError = &microerror.Error{
	Kind: "parseError",
}

// IsParseError asserts parseError.
func IsParseError(err error) bool {
	return microerror.Cause(err) == parseError
}
