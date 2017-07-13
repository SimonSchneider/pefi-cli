package main

import (
	"github.com/urfave/cli"
)

type (
	command interface {
		Cmd() cli.Command
		Endpoint() string
		ParseFlags(*cli.Context) error
		NewAdd() (interface{}, error)
		NewStruct() interface{}
		NewSlice() interface{}
		FinalFuncs() finalFuncs
		Flags() apiFlags
	}

	apiFlags struct {
		GetAllFlags []cli.Flag
		AddFlags    []cli.Flag
	}

	finalFuncs struct {
		GetAllFinal func(in interface{})
	}
)
