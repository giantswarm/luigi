package main

import "github.com/giantswarm/microerror"

var jsonObjectParseError = &microerror.Error{
	Kind: "jsonObjectParseError",
}

// IsJSONObjectParse asserts jsonObjectParseError.
func IsJSONObjectParse(err error) bool {
	return microerror.Cause(err) == jsonObjectParseError
}
