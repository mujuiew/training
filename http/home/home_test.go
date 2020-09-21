package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	HomePageHandle(res, req)
	if res.Code != 200 {
		t.Fatalf("Expected status == 200, but got %d", res.Code)
	} else {
		t.Fatalf("****************Get %d", res.Code)
	}

}
func Test_Post(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	HomePageHandle(res, req)
	if res.Code == 200 {
		t.Fatalf("Expected status != 200, but got %d", res.Code)
	}
}
