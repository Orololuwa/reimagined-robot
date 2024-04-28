package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestCreateAUser(t *testing.T){
	type UserBody struct {
		FirstName string `json:"firstName" faker:"first_name"`
		LastName string `json:"lastName" faker:"last_name"`
		Email string `json:"email" faker:"email"`
		Password string `json:"password" faker:"password"`
	}

	body := UserBody{}
	err := faker.FakeData(&body)
	if err != nil {
		t.Log(err)
	}

	jsonBody, err := json.Marshal(body)
    if err != nil {
        t.Log("Error:", err)
        return
    }

	// Test for success
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.CreateAUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("CreateAUser handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusCreated)
	}

	// Test for missing request body
	req, _ = http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(``)))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.CreateAUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("CreateAUser handler returned wrong response code for missing body: got %d, wanted %d", rr.Code, http.StatusInternalServerError)
	}

	// Test for failed DB insert
	body.Password = "invalid"
	jsonBody, err = json.Marshal(body)
    if err != nil {
        t.Log("Error:", err)
        return
    }

	req, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.CreateAUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("CreateAUser handler returned wrong response code for failed UserRepo db function: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}
}

func TestGetAUser(t *testing.T){
	// test for success
	req, _ := http.NewRequest("GET", "/user/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.RequestURI = "/user/1"	
	res := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.GetAUser)	
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("GetAUser handler returned wrong response code: got %d, wanted %d", res.Code, http.StatusOK)
	}

	// test valid id in the path variable
	req, _ = http.NewRequest("GET", "/room", nil)
	req.Header.Set("Content-Type", "application/json")
	req.RequestURI = "/room/one"

	res = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.GetAUser)

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("GetAUser handler returned wrong response code for invalid query param 'id': got %d, wanted %d", res.Code, http.StatusInternalServerError)
	}

	// test for failed db operation
	req, _ = http.NewRequest("GET", "/room", nil)
	req.Header.Set("Content-Type", "application/json")	
	req.RequestURI = "/room/0"


	res = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.GetAUser)

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusNotFound {
		t.Errorf("GetAUser handler returned wrong response code for failed UserRepo function: got %d, wanted %d", res.Code, http.StatusNotFound)
	}
}