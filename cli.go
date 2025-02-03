package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("TODO: help text")
}

func lenCheck(i int, nArgs int) {
	if i+1 >= nArgs {
		fmt.Println("No input file specified")
		os.Exit(0)
	}
}

func readTimeframe(tf string) [6]int {
	var ret = [6]int{0, 0, 0, 0, 0, 0}
	var ss []string = strings.Split(tf, ":")
	var startString []string = strings.Split(ss[0], ".")
	var endString []string = strings.Split(ss[1], ".")
	var err error
	ret[dayStart], err = strconv.Atoi(startString[0])
	ret[monthStart], err = strconv.Atoi(startString[1])
	ret[yearStart], err = strconv.Atoi(startString[2])
	ret[dayEnd], err = strconv.Atoi(endString[0])
	ret[monthEnd], err = strconv.Atoi(endString[1])
	ret[yearEnd], err = strconv.Atoi(endString[2])
	if err != nil {
		log.Fatal("Error parsing timeframe: ", err)
	}
	return ret
}

func parseArgs(args []string) (string, [6]int) {
	args = args[1:]
	nArgs := len(args)
	var file string
	var tf = [6]int{0, 0, 0, 0, 0, 0}
	for i := range args {
		switch args[i] {
		case "-h", "--help":
			printHelp()
		case "-f", "--file":
			lenCheck(i, nArgs)
			file = args[i+1]
		case "-t", "--timeframe":
			lenCheck(i, nArgs)
			tf = readTimeframe(args[i+1])
		}
	}
	return file, tf
}
