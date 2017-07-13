package main

import (
	"bytes"
	"encoding/json"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"reflect"
	"testing"
)

func TestGetCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	mock := mockType{}
	method := "GET"
	var req []byte
	resp := models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	handler := mockGet(&resp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	cl := client{server.URL, out, 1, []reflect.Type{
		reflect.TypeOf(models.InternalAccount{}),
		reflect.TypeOf(models.ExternalAccount{}),
	}}
	get := getCmd(cl, &mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = get
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "json,j"},
	}
	app.Run([]string{"test", "-j", "1"})
	exout, _ := json.Marshal(resp)
	if string(exout)+"\n" != out.String() {
		t.Error("output differ ex, got:", len(exout), ",", len(out.Bytes()), "\n", string(exout), "\n", out.String())
	}
}
