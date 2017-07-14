package main

import (
	"bytes"
	"encoding/json"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"testing"
)

func TestGetAllCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	mock := &mockType{}
	method := "GET"
	var req []byte
	resp := []models.InternalAccount{
		{
			ExternalAccount: models.ExternalAccount{
				ID: 2, Name: "test1", Description: "test1", CategoryID: 1},
			Balance: 0.1,
		},
		{
			ExternalAccount: models.ExternalAccount{
				ID: 3, Name: "test2", Description: "test2", CategoryID: 2},
			Balance: 0.2,
		},
	}
	handler := mockGetAll(resp, &req)
	server := getTestServer(mock.Endpoint(), method, handler)
	cl := &client{server.URL, out, 1, nil}
	fun := getAllCmd(cl, mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = fun
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "json,j"},
	}
	app.Run([]string{"test", "-j"})
	exout, _ := json.Marshal(resp)
	if string(exout)+"\n" != out.String() {
		t.Error("output differ ex, got:", len(exout), ",", len(out.Bytes()), "\n", string(exout), "\n", out.String())
	}
}

func TestGetCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	mock := mockType{}
	method := "GET"
	var req []byte
	resp := models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 2, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	handler := mockGet(resp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	cl := &client{server.URL, out, 1, nil}
	fun := getCmd(cl, &mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = fun
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "json,j"},
	}
	app.Run([]string{"test", "-j", "1"})
	exout, _ := json.Marshal(resp)
	if string(exout)+"\n" != out.String() {
		t.Error("output differ ex, got:", len(exout), ",", len(out.Bytes()), "\n", string(exout), "\n", out.String())
	}
}

func TestAddCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	exout := ""
	mock := mockType{}
	method := "POST"
	var req []byte
	resp := models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	handler := mockAdd(resp, &req)
	server := getTestServer(mock.Endpoint(), method, handler)
	cl := &client{server.URL, out, 1, nil}
	fun := addCmd(cl, &mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = fun
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "json,j"},
		cli.Int64Flag{Name: "id,i"},
		cli.StringFlag{Name: "name,n"},
		cli.StringFlag{Name: "description,d"},
		cli.Int64Flag{Name: "categoryID,c"},
		cli.Float64Flag{Name: "balance,b"},
	}
	app.Run([]string{"test",
		"-i", "1",
		"-n", "test1",
		"-d", "test1",
		"-c", "1",
		"-b", "0.1",
	})
	exreq, _ := json.Marshal(resp)
	if exout != out.String() {
		t.Error("output differ ex, got:", len(exout), ",", len(out.Bytes()), "\n", string(exout), "\n", out.String())
	}
	if string(exreq) != string(req) {
		t.Error("differ ex", string(exreq), "got", string(req))
	}
}

func TestDelCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	mock := mockType{}
	method := "DEL"
	var req []byte
	resp := ""
	handler := mockDel(resp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	cl := &client{server.URL, out, 1, nil}
	fun := delCmd(cl, &mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = fun
	app.Run([]string{"test", "1"})
	if resp != out.String() {
		t.Error("output differ ex, got:", len(resp), ",", len(out.Bytes()), "\n", resp, "\n", out.String())
	}
}

func TestModCmd(t *testing.T) {
	var b []byte
	out := bytes.NewBuffer(b)
	exout := ""
	mock := mockType{}
	method := "PUT"
	var req []byte
	resp := models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	handler := mockMod(resp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	cl := &client{server.URL, out, 1, nil}
	fun := modCmd(cl, &mock)
	app := cli.NewApp()
	app.Name = "test"
	app.Action = fun
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "json,j"},
		cli.Int64Flag{Name: "id,i"},
		cli.StringFlag{Name: "name,n"},
		cli.StringFlag{Name: "description,d"},
		cli.Int64Flag{Name: "categoryID,c"},
		cli.Float64Flag{Name: "balance,b"},
	}
	app.Run([]string{"test",
		"-i", "1",
		"-n", "test1",
		"-d", "test1",
		"-c", "1",
		"-b", "0.1",
		"1",
	})
	exreq, _ := json.Marshal(resp)
	if exout != out.String() {
		t.Error("output differ ex, got:", len(exout), ",", len(out.Bytes()), "\n", string(exout), "\n", out.String())
	}
	if string(exreq) != string(req) {
		t.Error("differ ex", string(exreq), "got", string(req))
	}
}
