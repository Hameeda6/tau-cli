package main

import (
	"log"
	"os"

	"github.com/taubyte/tau-cli/cli"
	"github.com/taubyte/tau-cli/i18n"
)
//first func to be executed when the program is run
func main() {
	err := cli.Run(os.Args...) //calls function Run from cli package
	if err != nil {  //if error occurs during the execution of cli.Run ; 
		log.Fatal(i18n.AppCrashed(err)) // logs the error and passes through i18n.AppCrashedc func ; and exits the program.
	}
}
