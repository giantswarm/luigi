package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/giantswarm/microerror"
)

func main() {
	var flagForceColor = flag.Bool("color", false, "force color output. This option overrides -no-color")
	var flagGrep = flag.String("grep", "", "grep for fields. Comma separated list of key=value pairs")
	var flagNoColor = flag.Bool("no-color", false, "disable color output")
	flag.Parse()

	if *flagNoColor {
		disable_colors(true)
	}
	if *flagForceColor {
		disable_colors(false)
	}

	grep, err := parseGrepFlag(*flagGrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line, err := format(scanner.Bytes(), grep)
		if IsGrepNotFound(err) {
			continue
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

func parseGrepFlag(s string) (map[string]string, error) {
	if len(s) == 0 {
		return nil, nil
	}

	grep := make(map[string]string)

	for _, kv := range strings.Split(s, ",") {
		pair := strings.Split(kv, "=")
		if len(pair) != 2 {
			return nil, microerror.Maskf(parseError, "invalid key=value pair: '%q'", kv)
		}

		grep[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
	}

	return grep, nil
}
