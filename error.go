package main

import "github.com/giantswarm/microerror"

var jsonParseError = &microerror.Error{
	Kind: "jsonParseError",
}

// IsJsonParse asserts jsonParseError.
func IsJsonParse(err error) bool {
	return microerror.Cause(err) == jsonParseError
}
