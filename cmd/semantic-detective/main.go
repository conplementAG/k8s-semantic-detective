package main

import (
	"github.com/conplementAG/k8s-semantic-detective/pkg/common/logging"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	defer errorhandler()

	logging.Initialize("semantic-detective.log")
	Execute()
}

func errorhandler() {
	// as this is a CLI tool, and not a library with an API, panic is used for most errors that occur,
	// since they are unrecoverable and need some user intervention (or they are genuine panic programming
	// errors)
	if r := recover(); r != nil {
		color.Set(color.FgHiRed)
		log.Printf("k8s management detective -- error occured: %+v\n", r)
		color.Unset()
		os.Exit(1)
	}
}
