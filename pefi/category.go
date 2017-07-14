package main

import (
	"encoding/json"
	"github.com/simonschneider/dyntab"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
)

type (
	category struct {
		model    models.Category
		endpoint string
	}
)

func (category) Cmd() cli.Command {
	return cli.Command{
		Name:    "category",
		Aliases: []string{"c"},
		Usage:   "category interface",
	}
}

func (t category) Endpoint() string {
	return t.endpoint
}

func (category) Flags() apiFlags {
	return apiFlags{
		AddFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
		},
		ModFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
		},
	}
}

func (t *category) ParseFlags(c *cli.Context) error {
	t.model = models.Category{
		Name:        c.String("name"),
		Description: c.String("description"),
	}
	return nil
}

func (t *category) ParseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t.model)
}

func (t category) GetModel() interface{} {
	return t.model
}

func (t category) GetSpecialize() []dyntab.ToSpecialize {
	return nil
}

func (category) NewStruct() interface{} {
	return new(models.Category)
}

func (category) NewSlice() interface{} {
	return new([]models.Category)
}

func (category) FinalFuncs() finalFuncs {
	return finalFuncs{}
}
