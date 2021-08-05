package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getBooks)

	handler.ServeHTTP(rr, req)
	fmt.Println(rr)
	fmt.Println(req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"_id":"610912e479b1daa4bdb03a15","isbn":"1111","title":"Hello","author":{"firstname":"amal","lastname":"nath"}},{"_id":"6109130279b1daa4bdb03a16","isbn":"1112","title":"Wings of fire","author":{"firstname":"abdul","lastname":"kalam"}},{"_id":"61093a6a84ee5f063bb32ac5","isbn":"1113","title":"aim","author":{"firstname":"abdul","lastname":"raj"}}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
