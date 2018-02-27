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

	line, msg := levelMessage(m)

	line += " " + m["time"][2:len(m["time"])-4] // Skip '20' in year and millis.
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

func levelMessage(m map[string]string) (level string, message string) {
	switch m["level"] {
	case "debug":
		level = "D"
	case "info":
		level = "I"
	case "warning":
		level = "W"
	case "error":
		level = "E"
	case "":
		// Catch empty string to run fallback.
	default:
		level = "U"
	}

	message = m["message"]

	if len(level) > 0 && len(message) > 0 {
		delete(m, "level")
		delete(m, "message")
		return
	}

	// Fallback to old handling.

	switch {
	case len(m["debug"]) > 0:
		level = "D"
		message = m["debug"]
		delete(m, "debug")
	case len(m["info"]) > 0:
		level = "I"
		message = m["info"]
		delete(m, "info")
	case len(m["warning"]) > 0:
		level = "W"
		message = m["warning"]
		delete(m, "warning")
	case len(m["error"]) > 0:
		level = "E"
		message = m["error"]
		delete(m, "error")
	default:
		level = "U"
	}

	return
}
