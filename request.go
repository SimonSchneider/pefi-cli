package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getAllReq(base string, user int64, c command) (out interface{}, err error) {
	model := c.NewSlice()
	addr := getURL(base, nil, c.Endpoint())
	err = getAndDecReq(addr, user, model)
	return model, err
}

func getReq(base string, user int64, c command, id int64) (out interface{}, err error) {
	model := c.NewStruct()
	addr := getURL(base, nil, c.Endpoint(), strconv.FormatInt(id, 10))
	err = getAndDecReq(addr, user, model)
	return model, err
}

func addReq(base string, user int64, c command) (err error) {
	mod, err := c.NewAdd()
	if err != nil {
		return err
	}
	buf, err := json.Marshal(mod)
	addr := getURL(base, nil, c.Endpoint())
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	_, err = handleRequest(req, user)
	return nil
}

func delReq(base string, user int64, c command, id int64) (err error) {
	addr := getURL(base, nil, c.Endpoint(), strconv.FormatInt(id, 10))
	req, err := http.NewRequest("DEL", addr, nil)
	if err != nil {
		return err
	}
	_, err = handleRequest(req, user)
	return err
}

func modReq(base string, user int64, c command, id int64) (err error) {
	mod, err := c.NewAdd()
	if err != nil {
		return err
	}
	buf, err := json.Marshal(mod)
	addr := getURL(base, nil, c.Endpoint(), strconv.FormatInt(id, 10))
	req, err := http.NewRequest("PUT", addr, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	_, err = handleRequest(req, user)
	req.Header.Set("user", strconv.FormatInt(user, 10))
	return nil
}

func getAndDecReq(addr string, user int64, model interface{}) (err error) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return err
	}
	body, err := handleRequest(req, user)
	if err != nil {
		return err
	}
	if err = json.NewDecoder(body).Decode(model); err != nil {
		return err
	}
	return nil
}

func handleRequest(req *http.Request, user int64) (r io.Reader, err error) {
	req.Header.Set("user", strconv.FormatInt(user, 10))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			b = []byte("")
		}
		s := fmt.Sprintf("\nStatus: %s\nBody: %s\n",
			resp.Status, string(b))
		return nil, errors.New(s)
	}
	return resp.Body, err
}

func getURL(base string, query []string, endpoint ...string) string {
	url := strings.TrimRight(base, "/")
	for _, e := range endpoint {
		url += "/" + strings.Trim(e, "/")
	}
	if query == nil {
		return url
	}
	url += "?"
	for _, q := range query[:len(query)-1] {
		url += q + "&"
	}
	url += query[len(query)-1]
	return url
}
