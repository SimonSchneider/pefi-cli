package main

import (
	"encoding/json"
	"github.com/simonschneider/pefi/models"
	"testing"
)

func TestGetURL(t *testing.T) {
	base := "l:9090/"
	ep := "/test///"
	exp := "l:9090/test?d=2&d=24"
	query := []string{"d=2", "d=24"}
	url := getURL(base, query, ep)
	if url != exp {
		t.Error("error exp:", exp, "got", url)
		return
	}
}

func TestGetAndDecReq(t *testing.T) {
	endpoint := "/accounts/internal"
	method := "GET"
	query := []string{"one=1", "two=2"}
	exQuery := "one=1&two=2"
	var reQuery string
	exReq := []byte("")
	var req []byte
	exResp := []models.InternalAccount{
		{
			ExternalAccount: models.ExternalAccount{
				ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
			Balance: 0.1,
		},
		{
			ExternalAccount: models.ExternalAccount{
				ID: 2, Name: "test2", Description: "test2", CategoryID: 2},
			Balance: 0.2,
		},
	}
	handler := mockTest(exResp, &req, &reQuery)
	server := getTestServer(endpoint, method, handler)
	addr := server.URL
	resps := new([]models.InternalAccount)
	err := getAndDecReq(getURL(addr, query, endpoint), 1, resps)
	if err != nil {
		t.Error("error get and decoding", err)
		return
	}
	if string(req) != string(exReq) {
		t.Error("req error exp", string(exReq), "got", string(req))
		return
	}
	if reQuery != exQuery {
		t.Error("req error exp", exQuery, "got", reQuery)
		return
	}
	if len(*resps) != len(exResp) {
		t.Error("resperror exp", exResp, "got", resps)
		return
	}
	for i, a := range *resps {
		if a != exResp[i] {
			t.Error("resperror exp", exResp, "got", resps)
			return
		}
	}
}

func TestGetAllReq(t *testing.T) {
	mock := mockType{}
	method := "GET"
	exReq := []byte("")
	var req []byte
	exResp := []models.InternalAccount{
		{
			ExternalAccount: models.ExternalAccount{
				ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
			Balance: 0.1,
		},
		{
			ExternalAccount: models.ExternalAccount{
				ID: 2, Name: "test2", Description: "test2", CategoryID: 2},
			Balance: 0.2,
		},
	}
	handler := mockGetAll(exResp, &req)
	server := getTestServer(mock.Endpoint(), method, handler)
	responses, err := getAllReq(server.URL, 1, nil, &mock)
	if err != nil {
		t.Error("error get All", err)
		return
	}
	if string(req) != string(exReq) {
		t.Error("req error exp", string(exReq), "got", string(req))
		return
	}
	resps, ok := responses.(*[]models.InternalAccount)
	if !ok {
		t.Error("req error exp", exResp, "got", responses)
		return
	}
	if len(*resps) != len(exResp) {
		t.Error("resperror exp", exResp, "got", resps)
		return
	}
	for i, a := range *resps {
		if a != exResp[i] {
			t.Error("resperror exp", exResp, "got", resps)
			return
		}
	}
}

func TestGetReq(t *testing.T) {
	mock := mockType{}
	method := "GET"
	exReq := []byte("")
	var req []byte
	exResp := models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	handler := mockGet(exResp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	responses, err := getReq(server.URL, 1, nil, &mock, 1)
	if err != nil {
		t.Error("error get All", err)
		return
	}
	if string(req) != string(exReq) {
		t.Error("req error exp", string(exReq), "got", string(req))
		return
	}
	resps, ok := responses.(*models.InternalAccount)
	if !ok {
		t.Error("req error exp", exResp, "got", responses)
		return
	}
	if *resps != exResp {
		t.Error("resperror exp", exResp, "got", resps)
		return
	}
}

func TestAddReq(t *testing.T) {
	mock := mockType{}
	method := "POST"
	exReq, _ := json.Marshal(models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	})
	mock.model = models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	var req []byte
	exResp := []byte("")
	handler := mockAdd(exResp, &req)
	server := getTestServer(mock.Endpoint(), method, handler)
	err := addReq(server.URL, 1, nil, &mock)
	if err != nil {
		t.Error("error get All", err)
		return
	}
	if string(exReq) != string(req) {
		t.Error("request wrong ex", string(exReq), "got", string(req))
	}
}

func TestDelReq(t *testing.T) {
	mock := mockType{}
	method := "DEL"
	exReq := []byte("")
	var req []byte
	exResp := []byte("")
	handler := mockAdd(exResp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	err := delReq(server.URL, 1, nil, &mock, 1)
	if err != nil {
		t.Error("error get All", err)
		return
	}
	if string(exReq) != string(req) {
		t.Error("request wrong ex", string(exReq), "got", string(req))
	}
}

func TestModReq(t *testing.T) {
	mock := mockType{}
	method := "PUT"
	exReq, _ := json.Marshal(models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	})
	mock.model = models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID: 1, Name: "test1", Description: "test1", CategoryID: 1},
		Balance: 0.1,
	}
	var req []byte
	exResp := []byte("")
	handler := mockAdd(exResp, &req)
	server := getTestServer(mock.Endpoint()+"/1", method, handler)
	err := modReq(server.URL, 1, nil, &mock, 1)
	if err != nil {
		t.Error("error get All", err)
		return
	}
	if string(exReq) != string(req) {
		t.Error("request wrong ex", string(exReq), "got", string(req))
	}
}
