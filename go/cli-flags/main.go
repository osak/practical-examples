package main

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	verbose  bool
	numParam int
	strParam string
	positionals []string
	showHelp bool
}

func parseArgs() Args {
	var args Args
	flag.BoolVar(&args.verbose, "v", false, "verbose")
	flag.IntVar(&args.numParam, "num-param", 0, "number param")
	flag.StringVar(&args.strParam, "str-param", "", "string param")
	flag.BoolVar(&args.showHelp, "help", false, "show help")
	flag.Parse()
	args.positionals = flag.Args()
	return args
}

// `mycmd -v --num-param=1 --str-param=foo /tmp/a.txt`
func main() {
	args := parseArgs()
	if args.showHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.Printf("%v\n", args)
}
