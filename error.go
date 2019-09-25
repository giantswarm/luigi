package main

import "github.com/giantswarm/microerror"

var jsonObjectParseError = &microerror.Error{
	Kind: "jsonObjectParseError",
}

// IsJSONObjectParse asserts jsonObjectParseError.
func IsJSONObjectParse(err error) bool {
	return microerror.Cause(err) == jsonObjectParseError
}

var skipError = &microerror.Error{
	Kind: "skipError",
}

// IsSkip asserts skipError.
func IsSkip(err error) bool {
	return microerror.Cause(err) == skipError
}
