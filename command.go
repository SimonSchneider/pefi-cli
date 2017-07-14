package main

import (
	"encoding/json"
	"github.com/simonschneider/dyntab"
	"github.com/urfave/cli"
	"io"
	"os"
	"reflect"
	"strconv"
)

const (
	errNumArg = "incorrect number of args "
)

type (
	client struct {
		addr   string
		w      io.Writer
		user   int64
		models []reflect.Type
	}

	command interface {
		Cmd() cli.Command
		Endpoint() string
		ParseFlags(*cli.Context) error
		ParseReader(io.Reader) error
		GetModel() interface{}
		NewStruct() interface{}
		NewSlice() interface{}
		FinalFuncs() finalFuncs
		Flags() apiFlags
	}

	apiFlags struct {
		GetAllFlags []cli.Flag
		GetFlags    []cli.Flag
		AddFlags    []cli.Flag
		DelFlags    []cli.Flag
		ModFlags    []cli.Flag
	}

	finalFuncs struct {
		GetAllFinal func(in interface{})
	}
)

func getAllCmd(cl *client, com command) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != 0 {
			i := strconv.Itoa(len(c.Args()))
			return cli.NewExitError(errNumArg+i, 1)
		}
		//err := com.ParseFlags(c)
		//if err != nil {
		//return cli.NewExitError(err, 1)
		//}
		query := c.StringSlice("query")
		ans, err := getAllReq(cl.addr, cl.user, query, com)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		if c.Bool("json") {
			if err = json.NewEncoder(cl.w).Encode(ans); err != nil {
				return cli.NewExitError(err, 1)
			}
			return nil
		}
		if com.FinalFuncs().GetAllFinal != nil {
			com.FinalFuncs().GetAllFinal(ans)
		}
		dyntab.PrintTable(cl.w, ans, cl.models, nil)
		return nil
	}
}

func getCmd(cl *client, com command) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != 1 {
			i := strconv.Itoa(len(c.Args()))
			return cli.NewExitError(errNumArg+i, 1)
		}
		//err := com.ParseFlags(c)
		//if err != nil {
		//return cli.NewExitError(err, 1)
		//}
		id, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		query := c.StringSlice("query")
		ans, err := getReq(cl.addr, cl.user, query, com, int64(id))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		if c.Bool("json") {
			if err = json.NewEncoder(cl.w).Encode(ans); err != nil {
				return cli.NewExitError(err, 1)
			}
			return nil
		}
		dyntab.PrintTable(cl.w, ans, cl.models, nil)
		return nil
	}
}

func addCmd(cl *client, com command) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != 0 {
			i := strconv.Itoa(len(c.Args()))
			return cli.NewExitError(errNumArg+i, 1)
		}
		err := com.ParseFlags(c)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		if path := c.String("file"); path != "" {
			file, err := os.Open(path)
			if err != nil {
				return cli.NewExitError(err, 1)
			}
			com.ParseReader(file)
		}
		query := c.StringSlice("query")
		err = addReq(cl.addr, cl.user, query, com)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	}
}

func delCmd(cl *client, com command) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != 1 {
			i := strconv.Itoa(len(c.Args()))
			return cli.NewExitError(errNumArg+i, 1)
		}
		//err := com.ParseFlags(c)
		//if err != nil {
		//return cli.NewExitError(err, 1)
		//}
		id, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		query := c.StringSlice("query")
		err = delReq(cl.addr, cl.user, query, com, int64(id))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	}
}

func modCmd(cl *client, com command) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(c.Args()) != 1 {
			i := strconv.Itoa(len(c.Args()))
			return cli.NewExitError(errNumArg+i, 1)
		}
		id, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		err = com.ParseFlags(c)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		if path := c.String("file"); path != "" {
			file, err := os.Open(path)
			if err != nil {
				return cli.NewExitError(err, 1)
			}
			com.ParseReader(file)
		}
		query := c.StringSlice("query")
		err = modReq(cl.addr, cl.user, query, com, int64(id))
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	}
}
