// run.go is the entry point for cli application
package cli

import (
	argsLib "github.com/taubyte/tau-cli/cli/args"
	"github.com/taubyte/tau-cli/i18n"
)

func Run(args ...string) error { // func Run allows to pass string arguments 
	app, err := New() //new application instance is being created
	if err != nil {  //if error in creating the application ; 
		return i18n.AppCreateFailed(err) //returns error message and exits. i18n is for handling error messages.
	}

	if len(args) == 1 {   //if len of argument is 1;calls Run method
		return app.Run(args)
	}
	//passes command line arguments based on flags and commands defined in CLI application
	args = argsLib.ParseArguments(app.Flags, app.Commands, args...)

	return app.Run(args)
}
