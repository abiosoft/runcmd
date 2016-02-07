package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/abiosoft/goutils/env"
	"github.com/abiosoft/ishell"
)

type runCmd struct {
	cmds  []string
	vars  env.EnvVar
	shell *ishell.Shell
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command missing.")
		os.Exit(1)
	}
	app := &runCmd{
		cmds: os.Args[1:],
	}

	shell := ishell.New()

	// remove defaults
	shell.Unregister("exit")
	shell.Unregister("help")
	shell.Unregister("clear")

	// handle all inputs
	shell.RegisterGeneric(func(args ...string) (string, error) {
		vars := env.EnvVar(append(os.Environ(), app.vars...))
		for i, a := range args {
			if strings.HasPrefix(a, "$") {
				if val := vars.Get(a[1:]); val != "" {
					args[i] = val
				}
			}
		}
		return "", app.run(args)
	})
	// environment variables
	shell.Register(".env", func(args ...string) (string, error) {
		switch len(args) {
		case 0:
			return app.vars.String(), nil
		case 1:
			app.vars.SetStr(args[0])
			return "", nil
		case 2:
			app.vars.Set(args[0], args[1])
			return "", nil
		}
		app.vars.Set(args[0], args[1])
		return "", fmt.Errorf(".env accepts at most 2 arguments")
	})
	// exit
	shell.Register(".exit", func(args ...string) (string, error) {
		shell.Stop()
		return "", nil
	})
	// clear screen
	shell.Register(".clear", func(args ...string) (string, error) {
		return "", shell.ClearScreen()
	})
	// switch command
	shell.Register(".switch", func(args ...string) (string, error) {
		app.cmds = args
		app.resetPrompt()
		return "", shell.ClearScreen()
	})

	app.shell = shell
	app.resetPrompt()

	shell.Start()
}

func (r *runCmd) run(args []string) error {
	args = append(r.cmds[1:], args...)
	cmd := exec.Command(r.cmds[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), r.vars...)
	return cmd.Run()
}

func (r *runCmd) resetPrompt() {
	r.shell.SetPrompt(strings.Join(r.cmds, " ") + "> ")
}
