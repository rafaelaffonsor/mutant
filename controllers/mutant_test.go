package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMutant(t *testing.T) {

	var jsonStr = []byte(`{"dna":["CCGCGA","CCGTGC","TTAGTT","AGACGG","ACATCA","TCAAGG"]}`)

	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	IsMutant(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "true\n", rec.Body.String())
}

func TestIsHuman(t *testing.T) {

	var jsonStr = []byte(`{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`)

	req, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	IsMutant(rec, req)

	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "false\n", rec.Body.String())
}