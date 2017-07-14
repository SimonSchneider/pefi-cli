package main

import (
	"encoding/json"
	"github.com/buger/goterm"
	"github.com/simonschneider/dyntab"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	transaction struct {
		model    models.Transaction
		endpoint string
	}
)

func (transaction) Cmd() cli.Command {
	return cli.Command{
		Name:    "transaction",
		Aliases: []string{"t"},
		Usage:   "transaction interface",
	}
}

func (t transaction) Endpoint() string {
	return t.endpoint
}

func (transaction) Flags() apiFlags {
	return apiFlags{
		GetAllFlags: []cli.Flag{
			cli.BoolFlag{
				Name: "graph,g",
			},
		},
		AddFlags: []cli.Flag{
			cli.StringFlag{
				Name:  "time,t",
				Value: time.Now().Format(time.RFC3339),
			},
			cli.Float64Flag{Name: "amount,a"},
			cli.Int64Flag{Name: "sender,s"},
			cli.Int64Flag{Name: "receiver,r"},
			cli.Int64Flag{Name: "label,l"},
		},
		ModFlags: []cli.Flag{
			cli.StringFlag{
				Name:  "time,t",
				Value: time.Now().Format(time.RFC3339),
			},
			cli.Float64Flag{Name: "amount,a"},
			cli.Int64Flag{Name: "sender,s"},
			cli.Int64Flag{Name: "receiver,r"},
			cli.Int64Flag{Name: "label,l"},
		},
	}
}

func (t *transaction) ParseFlags(c *cli.Context) error {
	tim, err := time.Parse(time.RFC3339, c.String("time"))
	if err != nil {
		return err
	}
	t.model = models.Transaction{
		Time:       tim,
		Amount:     c.Float64("amount"),
		SenderID:   c.Int64("sender"),
		ReceiverID: c.Int64("receiver"),
		LabelID:    c.Int64("label"),
	}
	return nil
}

func (t *transaction) ParseReader(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t.model)
}

func (t transaction) GetModel() interface{} {
	return t.model
}

func (t transaction) GetSpecialize() []dyntab.ToSpecialize {
	return []dyntab.ToSpecialize{
		{
			Type: reflect.TypeOf(time.Time{}),
			ToString: func(i interface{}) (string, error) {
				t, ok := i.(time.Time)
				if !ok {
					return "", nil
				}
				return t.Format("2006-01-02"), nil
			},
		},
	}
}

func (transaction) NewStruct() interface{} {
	return new(models.Transaction)
}

func (transaction) NewSlice() interface{} {
	return new([]models.Transaction)
}

func (transaction) FinalFuncs() finalFuncs {
	return finalFuncs{
		GetAllFinal: createGraph,
	}
}

func createGraph(c *cli.Context, in interface{}) {
	trans, _ := in.([]models.Transaction)
	if !c.Bool("graph") {
		return
	}
	w := termWidth()
	chart := goterm.NewLineChart(w, 20)
	data := new(goterm.DataTable)
	data.AddColumn("past Days")
	data.AddColumn("Daily total")
	sum := map[int]float64{}
	for _, t := range trans {
		days := time.Since(t.Time)
		sum[int(days.Hours()/24)] += t.Amount
	}
	for i := 0; i <= 30; i++ {
		data.AddRow(float64(i), sum[i])
		delete(sum, i)
		if len(sum) == 0 {
			break
		}
	}
	goterm.Println(chart.Draw(data))
	goterm.Flush()
	os.Exit(0)
}

func termWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	widths := strings.Split(string(out), " ")
	width := strings.Trim(widths[1], "\n")
	w, _ := strconv.Atoi(width)
	return w
}
