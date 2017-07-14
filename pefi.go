package main

import (
	"github.com/simonschneider/dyntab"
	"github.com/urfave/cli"
	"os"
	"reflect"
)

func main() {
	app := cli.NewApp()
	app.Name = "pefi"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "addr,a",
			Value:       "http://127.0.0.1:22400",
			EnvVar:      "PEFI_ADDR",
			Destination: &cl.addr,
		},
		cli.Int64Flag{
			Name:        "user,u",
			Value:       1,
			EnvVar:      "PEFI_USER",
			Destination: &cl.user,
		},
	}
	cl.w = os.Stdout

	coms := []command{
		&externalAccount{endpoint: "/accounts/external"},
		&internalAccount{endpoint: "/accounts/internal"},
		&label{endpoint: "/labels"},
		&category{endpoint: "/categories"},
		&transaction{endpoint: "/transactions"},
	}

	recurse := []reflect.Type{}
	specialize := []dyntab.ToSpecialize{}
	for _, com := range coms {
		subcmd := com.Cmd()
		subcmd.Subcommands = getAPICmd(com)
		app.Commands = append(app.Commands, subcmd)
		recurse = append(recurse, reflect.TypeOf(com.GetModel()))
		specialize = append(specialize, com.GetSpecialize()...)
	}
	cl.table = dyntab.NewTable().
		Specialize(specialize).
		Recurse(recurse)
	app.Run(os.Args)
}

func getAPICmd(com command) []cli.Command {
	return []cli.Command{
		{
			Name:   "ls",
			Usage:  "list all",
			Flags:  append(flags.GetAllFlags, com.Flags().GetAllFlags...),
			Action: getAllCmd(cl, com),
		},
		{
			Name:   "get",
			Usage:  "get `id`",
			Flags:  append(flags.GetFlags, com.Flags().GetFlags...),
			Action: getCmd(cl, com),
		},
		{
			Name:   "del",
			Usage:  "del `id`",
			Flags:  append(flags.DelFlags, com.Flags().DelFlags...),
			Action: delCmd(cl, com),
		},
		{
			Name:   "add",
			Usage:  "add `id`",
			Flags:  append(flags.AddFlags, com.Flags().AddFlags...),
			Action: addCmd(cl, com),
		},
		{
			Name:   "mod",
			Usage:  "mod `id`",
			Flags:  append(flags.ModFlags, com.Flags().ModFlags...),
			Action: modCmd(cl, com),
		},
	}
}

var (
	cl    = new(client)
	flags = apiFlags{
		GetAllFlags: []cli.Flag{
			cli.StringSliceFlag{Name: "query,q"},
			cli.BoolFlag{Name: "json,j"},
		},
		GetFlags: []cli.Flag{
			cli.StringSliceFlag{Name: "query,q"},
			cli.BoolFlag{Name: "json,j"},
		},
	}
)
