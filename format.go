package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/fatih/color"
)

var (
	black   = color.New(color.FgBlack).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	white   = color.New(color.FgWhite).SprintFunc()

	separator = red(" | ")
)

func format(text []byte) string {
	var m map[string]string
	err := json.Unmarshal(text, &m)
	if err != nil {
		return string(text)
	}

	line, msg := getLevelMessage(m)

	line += " " + cyan(m["time"][2:len(m["time"])-4]) // Skip '20' in year and millis.
	delete(m, "time")

	obj := m["object"] // Set by operatorkit framework.
	delete(m, "object")
	if len(obj) > 0 {
		line += " " + yellow(obj)
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
	line += separator + yellow(caller)

	stack := getStack(m)

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
		line += separator + green(k+"="+m[k])
	}

	line += stack

	// TODO object

	return line
}

func getLevelMessage(m map[string]string) (level string, message string) {
	switch m["level"] {
	case "debug":
		level = white("D")
	case "info":
		level = cyan("I")
	case "warning":
		level = yellow("W")
	case "error":
		level = red("E")
	case "":
		// Catch empty string to run fallback.
	default:
		level = white("U")
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
		level = white("D")
		message = m["debug"]
		delete(m, "debug")
	case len(m["info"]) > 0:
		level = cyan("I")
		message = m["info"]
		delete(m, "info")
	case len(m["warning"]) > 0:
		level = yellow("W")
		message = m["warning"]
		delete(m, "warning")
	case len(m["error"]) > 0:
		level = red("E")
		message = m["error"]
		delete(m, "error")
	default:
		level = white("U")
	}

	return
}

func getStack(m map[string]string) string {
	stack := m["stack"]
	if len(stack) == 0 {
		return ""
	}

	if !strings.HasPrefix(stack, "[{") {
		return ""
	}
	if !strings.HasSuffix(stack, "}]") {
		return ""
	}

	delete(m, "stack")

	stack = stack[2 : len(stack)-2]
	stack = "\n\t" + stack
	stack = strings.Replace(stack, "} {", "\n\t", -1)

	return stack
}
