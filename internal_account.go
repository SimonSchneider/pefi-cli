package main

import (
	"encoding/json"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
)

type (
	internalAccount struct {
		model    models.InternalAccount
		endpoint string
	}
)

func (internalAccount) Cmd() cli.Command {
	return cli.Command{
		Name:    "account-internal",
		Aliases: []string{"ai"},
		Usage:   "internal account interface",
	}
}

func (t internalAccount) Endpoint() string {
	return t.endpoint
}

func (internalAccount) Flags() apiFlags {
	return apiFlags{
		AddFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
			cli.Int64Flag{Name: "category,c"},
			cli.StringFlag{Name: "balance,b"},
		},
		ModFlags: []cli.Flag{
			cli.StringFlag{Name: "name,n"},
			cli.StringFlag{Name: "description,d"},
			cli.Int64Flag{Name: "category,c"},
			cli.Float64Flag{Name: "balance,b"},
		},
	}
}

func (t *internalAccount) ParseFlags(c *cli.Context) error {
	t.model = models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			Name:        c.String("name"),
			Description: c.String("description"),
			CategoryID:  c.Int64("category"),
		},
		Balance: c.Float64("balance"),
	}
	return nil
}

func (t *internalAccount) ParseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t.model)
}

func (t internalAccount) GetModel() interface{} {
	return t.model
}

func (internalAccount) NewStruct() interface{} {
	return new(models.InternalAccount)
}

func (internalAccount) NewSlice() interface{} {
	return new([]models.InternalAccount)
}

func (internalAccount) FinalFuncs() finalFuncs {
	return finalFuncs{}
}
