package main

import (
	"encoding/json"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
)

type (
	label struct {
		model    models.Label
		endpoint string
	}
)

func (label) Cmd() cli.Command {
	return cli.Command{
		Name:    "label",
		Aliases: []string{"l"},
		Usage:   "label interface",
	}
}

func (t label) Endpoint() string {
	return t.endpoint
}

func (label) Flags() apiFlags {
	return apiFlags{
		AddFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
			cli.Int64Flag{Name: "category,c"},
		},
		ModFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
			cli.Int64Flag{Name: "category,c"},
		},
	}
}

func (t *label) ParseFlags(c *cli.Context) error {
	t.model = models.Label{
		Name:        c.String("name"),
		Description: c.String("description"),
		CategoryID:  c.Int64("category"),
	}
	return nil
}

func (t *label) ParseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t.model)
}

func (t label) GetModel() interface{} {
	return t.model
}

func (label) NewStruct() interface{} {
	return new(models.Label)
}

func (label) NewSlice() interface{} {
	return new([]models.Label)
}

func (label) FinalFuncs() finalFuncs {
	return finalFuncs{}
}
