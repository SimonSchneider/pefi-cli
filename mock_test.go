package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type (
	mockType struct {
		model models.InternalAccount
	}
)

func (mockType) Cmd() cli.Command { return cli.Command{} }
func (mockType) Endpoint() string { return "/dummy" }
func (m *mockType) ParseFlags(c *cli.Context) error {
	m.model = models.InternalAccount{models.ExternalAccount{1, "test1", "test1", 1}, 0.1}
	return nil
}
func (m mockType) NewAdd() (interface{}, error) {
	return m.model, nil
}
func (mockType) NewStruct() interface{} {
	return new(models.InternalAccount)
}
func (mockType) NewSlice() interface{} {
	return new([]models.InternalAccount)
}
func (mockType) FinalFuncs() finalFuncs { return finalFuncs{} }
func (mockType) Flags() apiFlags        { return apiFlags{} }

func getTestServer(endpoint string, method string, h http.Handler) *httptest.Server {
	router := mux.NewRouter()
	router.Handle(endpoint, h).Methods(method)
	return httptest.NewServer(router)
}

func mockGetAll(response interface{}, input *[]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*input, _ = ioutil.ReadAll(r.Body)
		json.NewEncoder(w).Encode(response)
	})
}

func mockGet(response interface{}, input *[]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*input, _ = ioutil.ReadAll(r.Body)
		json.NewEncoder(w).Encode(response)
	})
}

func mockAdd(response interface{}, input *[]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*input, _ = ioutil.ReadAll(r.Body)
	})
}

func mockDel(response interface{}, input *[]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*input, _ = ioutil.ReadAll(r.Body)
	})
}

func mockMod(response interface{}, input *[]byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*input, _ = ioutil.ReadAll(r.Body)
		json.NewEncoder(w).Encode(response)
	})
}
