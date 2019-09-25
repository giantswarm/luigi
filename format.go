package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/giantswarm/luigi/pkg"
	"github.com/giantswarm/microerror"
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

func disableColors(v bool) {
	color.NoColor = v

	black = color.New(color.FgBlack).SprintFunc()
	red = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	blue = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan = color.New(color.FgCyan).SprintFunc()
	white = color.New(color.FgWhite).SprintFunc()

	separator = red(" | ")
}

func format(text []byte, grep *pkg.Grep) (string, error) {
	if len(text) == 0 {
		return "", microerror.Maskf(jsonObjectParseError, "empty string")
	}
	if text[0] != '{' {
		return "", microerror.Maskf(jsonObjectParseError, "text must start with %#q", "{")
	}

	var m map[string]interface{}
	err := json.Unmarshal(text, &m)
	if err != nil {
		return "", microerror.Maskf(jsonObjectParseError, err.Error())
	}

	if !grep.Filter(m) {
		return "", microerror.Maskf(skipError, "line does not match grep criteria")
	}

	line, msg := getLevelMessage(m)

	timeString := getString(m, "time")
	delete(m, "time")
	if timeString != "" {
		t, err := time.Parse("2006-01-02T15:04:05.999999-07:00", timeString)
		if err == nil {
			timeString = t.Format("01/02 15:04:05")
		}
	}
	line += " " + cyan(timeString)

	obj := getString(m, "object") // Set by operatorkit framework.
	delete(m, "object")
	if len(obj) > 0 {
		line += " " + yellow(obj)
	}

	resource := getString(m, "resource") // Set by operatorkit framework.
	delete(m, "resource")
	if len(resource) > 0 {
		line += " " + resource
	}

	function := getString(m, "function") // Set by operatorkit framework.
	delete(m, "function")
	if len(function) > 0 {
		if len(resource) > 0 {
			line += "." + function
		} else {
			line += " " + function
		}
	}

	if len(msg) > 0 {
		line += " " + white(msg)
	}

	caller := getString(m, "caller")
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
		line += separator + green(k+"="+getString(m, k))
	}

	line += stack

	// TODO object

	return line, nil
}

func getLevelMessage(m map[string]interface{}) (level string, message string) {
	switch getString(m, "level") {
	case "debug":
		level = white("D")
	case "info":
		level = cyan("I")
	case "warning":
		level = yellow("W")
	case "error":
		level = red("E")
	default:
		level = white("U")
	}
	delete(m, "level")

	message = getString(m, "message")

	if len(level) > 0 && len(message) > 0 {
		delete(m, "message")
		return
	}

	// Fallback to old handling.

	switch {
	case len(getString(m, "debug")) > 0:
		level = white("D")
		message = getString(m, "debug")
		delete(m, "debug")
	case len(getString(m, "info")) > 0:
		level = cyan("I")
		message = getString(m, "info")
		delete(m, "info")
	case len(getString(m, "warning")) > 0:
		level = yellow("W")
		message = getString(m, "warning")
		delete(m, "warning")
	case len(getString(m, "error")) > 0:
		level = red("E")
		message = getString(m, "error")
		delete(m, "error")
	}

	if len(level) == 0 {
		level = white("U")
	}

	return
}

func getStack(m map[string]interface{}) string {
	stack := getString(m, "stack")
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
	stack = strings.Replace(stack, " } {", "\n\t", -1)
	stack = strings.Replace(stack, "} {", "\n\t", -1)

	return stack
}

func getString(m map[string]interface{}, key string) string {
	v := m[key]
	if v == nil {
		return ""
	}

	s, ok := v.(string)
	if ok {
		return s
	}

	return fmt.Sprintf("%v", v)
}
