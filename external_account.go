package main

import (
	"encoding/json"
	"github.com/simonschneider/dyntab"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
)

type (
	externalAccount struct {
		model    models.ExternalAccount
		endpoint string
	}
)

func (externalAccount) Cmd() cli.Command {
	return cli.Command{
		Name:    "account-external",
		Aliases: []string{"ae"},
		Usage:   "external account interface",
	}
}

func (t externalAccount) Endpoint() string {
	return t.endpoint
}

func (externalAccount) Flags() apiFlags {
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

func (t *externalAccount) ParseFlags(c *cli.Context) error {
	t.model = models.ExternalAccount{
		Name:        c.String("name"),
		Description: c.String("description"),
		CategoryID:  c.Int64("category"),
	}
	return nil
}

func (t *externalAccount) ParseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t.model)
}

func (t externalAccount) GetModel() interface{} {
	return t.model
}

func (externalAccount) GetSpecialize() []dyntab.ToSpecialize {
	return nil
}

func (externalAccount) NewStruct() interface{} {
	return new(models.ExternalAccount)
}

func (externalAccount) NewSlice() interface{} {
	return new([]models.ExternalAccount)
}

func (externalAccount) FinalFuncs() finalFuncs {
	return finalFuncs{}
}
