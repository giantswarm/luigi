package main

import (
	"encoding/json"
	"sort"
	"strings"
)

func format(text []byte) string {
	var m map[string]string
	err := json.Unmarshal(text, &m)
	if err != nil {
		return string(text)
	}

	var line string

	var msg string
	if len(m["debug"]) > 0 {
		line += "D "
		msg = m["debug"]
		delete(m, "debug")
	} else if len(m["info"]) > 0 {
		line += "I "
		msg = m["info"]
		delete(m, "info")
	} else if len(m["warning"]) > 0 {
		line += "W "
		msg = m["warning"]
		delete(m, "warning")
	} else if len(m["error"]) > 0 {
		line += "E "
		msg = m["error"]
		delete(m, "error")
	} else {
		line += "U "
	}

	line += m["time"][2 : len(m["time"])-4] // Skip '20' in year and millis.
	delete(m, "time")

	obj := m["object"] // Set by operatorkit framework.
	delete(m, "object")
	if len(obj) > 0 {
		line += " " + obj
	}

	if len(msg) > 0 {
		line += " " + msg
	}

	caller := m["caller"]
	delete(m, "caller")
	{
		s := "/vendor/"
		i := strings.LastIndex(caller, s)
		if i > 0 {
			caller = caller[i+len(s):]
		}
		caller = strings.TrimPrefix(caller, "github.com/giantswarm/")
	}
	line += " | " + caller

	keys := make([]string, len(m))
	{
		var i int
		for k, _ := range m {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
	}

	for _, k := range keys {
		line += " | " + k + "=" + m[k]
	}

	// TODO object

	return line
}
