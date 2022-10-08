package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWork(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("ERROR")
	}
}
