package main

import (
	"flag"
)

func ParseCli() {

	action := ""
	flag.StringVar(&action, "action", "-h", "the action you want reverb to do")
	flag.Parse()

	switch action {
	case "-h":
		flag.PrintDefaults()
	case "map":
		target_file := "test_files/basic.java"
		MapDependencies(target_file)
	}

}
