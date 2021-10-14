package main

import (
	"log"

	"flag"

	"github.com/vladimircunichin/golang/7/goenv"
)

var (
	path    string
	program string
)

func init() {
	flag.StringVar(&path, "path", "", "env path")
	flag.StringVar(&program, "prog", "", "file to write to")
}

func main() {
	flag.Parse()
	if err := goenv.ReadDir(path); err != nil {
		log.Fatal(err)
	}
	if err := goenv.RunCmd(program); err != nil {
		log.Fatal(err)
	}
}
