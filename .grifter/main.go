package main

import _ "github.com/postcert/entitlements/grifts"
import "os"
import "log"
import "github.com/markbates/grift/grift"
import "path/filepath"

func main() {
	grift.CommandName = "buffalo task"
	if err := os.Chdir(filepath.Dir("/home/postcert/go/src/github.com/postcert/entitlements/grifts")); err != nil {
		log.Fatal(err)
	}
	err := grift.Exec(os.Args[1:], false)
	if err != nil {
		log.Fatal(err)
	}
}
