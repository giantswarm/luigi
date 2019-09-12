package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/giantswarm/luigi/pkg"
)

func main() {
	var flagForceColor = flag.Bool("color", false, "force color output. This option overrides -no-color")
	var flagGrep = flag.String("grep", "", "grep for fields. Comma separated list of key=value pairs")
	var flagNoColor = flag.Bool("no-color", false, "disable color output")
	flag.Parse()

	if *flagNoColor {
		disableColors(true)
	}
	if *flagForceColor {
		disableColors(false)
	}

	grep, err := pkg.NewGrep(*flagGrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(newSplitter().Split)
	for scanner.Scan() {
		line, err := format(scanner.Bytes(), grep)
		if IsJsonParse(err) {
			line = scanner.Text()
		} else if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		} else {
			line += "\n"
		}

		fmt.Printf("%s", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
