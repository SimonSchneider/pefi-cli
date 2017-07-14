package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/simonschneider/dyntab"
	"github.com/simonschneider/pefi/models"
	"github.com/urfave/cli"
	"io"
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
	m.model = models.InternalAccount{
		ExternalAccount: models.ExternalAccount{
			ID:          c.Int64("id"),
			Name:        c.String("name"),
			Description: c.String("description"),
			CategoryID:  c.Int64("categoryID")},
		Balance: c.Float64("balance")}
	return nil
}
func (m *mockType) ParseReader(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&m.model)
	return err
}
func (m mockType) GetModel() interface{} {
	return m.model
}
func (mockType) GetSpecialize() []dyntab.ToSpecialize { return nil }
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

func mockTest(response interface{}, input *[]byte, query *string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		*query = r.URL.RawQuery
		*input, _ = ioutil.ReadAll(r.Body)
		json.NewEncoder(w).Encode(response)
	})
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
