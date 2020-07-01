package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/giantswarm/luigi/pkg"
	"github.com/giantswarm/microerror"
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

	ctx := context.Background()

	err := mainE(ctx, os.Stdin, os.Stdout, *flagGrep)
	if err != nil {
		fmt.Fprintf(os.Stderr, microerror.JSON(err))
		os.Exit(1)
	}
}

func mainE(ctx context.Context, in io.Reader, out io.Writer, grepExpr string) error {
	grep, err := pkg.NewGrep(grepExpr)
	if err != nil {
		return microerror.Mask(err)
	}

	scanner := bufio.NewScanner(in)
	scanner.Split(newSplitter().Split)
	for scanner.Scan() {
		line, err := format(scanner.Bytes(), grep)
		if IsJSONObjectParse(err) {
			line = scanner.Text()
		} else if IsSkip(err) {
			// Don't print empty line.
			line = ""
		} else if err != nil {
			return microerror.Mask(err)
		} else {
			line += "\n"
		}

		fmt.Fprintf(out, "%s", line)
	}
	if err := scanner.Err(); err != nil {
		return microerror.Mask(err)
	}

	return nil
}
