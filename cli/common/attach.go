package common

import (
	"github.com/urfave/cli/v2"
)

type basicFunction func() Basic //returns a Basic interface (define func used to generate instances of Basic)

// Used to attach subCommands to their relative base commands
// Ex: `tau new project`  project is attached to the new command
func Attach(app *cli.App, commands ...basicFunction) {
	for _, cmdFunc := range commands {
		attachCommand(cmdFunc())
	}
	//checks these specific commands to see if they have subcommands and appends to app's commands.
	for _, cmd := range []*cli.Command{
		_new,
		_edit,
		_delete,
		_query,
		_list,
		_select,
		_clone,
		_push,
		_pull,
		_checkout,
		_import,
	} {
		if len(cmd.Subcommands) > 0 {
			app.Commands = append(app.Commands, cmd)
		}
	}
}

func attachCommand(cmd Basic) { //attach subcommands to base command. 
	baseCmd, baseOps := cmd.Base() //gets the base command and base options.

	for _cmd, method := range map[*cli.Command]func() Command{
		_new:      cmd.New,
		_edit:     cmd.Edit,
		_delete:   cmd.Delete,
		_query:    cmd.Query,
		_list:     cmd.List,
		_select:   cmd.Select,
		_import:   cmd.Import,
		_clone:    cmd.Clone,
		_push:     cmd.Push,
		_pull:     cmd.Pull,
		_checkout: cmd.Checkout,
	} {
		_method := method()
		
		if _method != NotImplemented { 
			cliCmd := _method.Initialize(_cmd, baseCmd, baseOps) //if method is NotImplemented ; it initializes using Initialize command of Basic instance. 
			if _cmd == _list { //if subcommand is _list ; 
				pluralAlias(cliCmd) //adds pluralAlias ie provide alternatice command names
			}
			//appends to base command's subcommands
			_cmd.Subcommands = append(_cmd.Subcommands, cliCmd)
		}
	}
}

func pluralAlias(command *cli.Command) {
	if command.Aliases == nil {
		command.Aliases = make([]string, 0)
	}

	switch command.Name {
	case "messaging", "smartops":
		return
	case "library":
		command.Aliases = append(command.Aliases, "libraries")
	case "application":
		command.Aliases = append(command.Aliases, "apps", "applications")
	default:
		command.Aliases = append(command.Aliases, command.Name+"s")
	}
}

//something to be improved
//1.pluralAlias function could be more generic to handle different command names.
//2.unitTesting for the codes
