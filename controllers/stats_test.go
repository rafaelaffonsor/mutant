package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {
	req, _ := http.NewRequest("GET", "/stats", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	GetStats(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}