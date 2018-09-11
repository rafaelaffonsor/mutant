package main_test

import (
	"testing"
	"log"
	"meli/db"
	"net/http"
	"encoding/json"
	"net/http/httptest"
	"github.com/gorilla/mux"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS mutants
(
id integer unique NOT NULL AUTO_INCREMENT,
dna varchar(200) not null,
is_mutant BOOLEAN default 0
)`

func TestMain(m *testing.M) {
	ensureTableExists()
}

func ensureTableExists() {
	if _, err := db.DbConn().Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestPostWithoutDna(t *testing.T) {
	req, _ := http.NewRequest("POST", "/mutant", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	}
}