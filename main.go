package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var flagNoColor = flag.Bool("no-color", false, "disable color output")
	var flagForceColor = flag.Bool("color", false, "force color output. This option overrides -no-color")
	flag.Parse()

	if *flagNoColor {
		disable_colors(true)
	}
	if *flagForceColor {
		disable_colors(false)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := format(scanner.Bytes())
		fmt.Printf("%s\n", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
