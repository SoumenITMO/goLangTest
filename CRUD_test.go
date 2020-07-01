package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type frmFields struct{
	FieldName string `json:"field"`
	Msg string `json:"msg"`
}

func TestGetCustomers(t *testing.T) {
	req, err := http.NewRequest("GET", "/home", nil)
	fmt.Print(req.Method)

	if err != nil {
		t.Fatal(err)
		fmt.Print(err.Error())
	}
	httptest := httptest.NewRecorder()
	handler := http.HandlerFunc(searchCustomers)
	handler.ServeHTTP(httptest, req)
	if(httptest.Code == 200) {
		t.Logf("Test Pass Cause it support only GET method")
	}
}

func TestGetCustomerMethod(t *testing.T) {

	req, err := http.NewRequest("POST", "/home", nil)

	if err != nil {
		t.Fatal(err)
		fmt.Print(err.Error())
	}
	httptest := httptest.NewRecorder()
	handler := http.HandlerFunc(searchCustomers)
	handler.ServeHTTP(httptest, req)
	if(httptest.Code == 405) {
		t.Logf("API Does not support POST method")
	}
}

func TestCustomerSavingDataValidation(t *testing.T) {

	var transactions []frmFields

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader("firstname=189&lastname=189&address=189&dob=2020-02-01&gender=X&email=189"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	handler := http.HandlerFunc(createCustomerDetails)
	handler(w, req)
	_ = json.Unmarshal([]byte(w.Body.String()), &transactions)
	for _ , errorField := range transactions {
		if(errorField.FieldName == "firstname") {
			fmt.Print("\n")
			fmt.Print("First Name is invalid")
		} else if(errorField.FieldName == "lastname") {
			fmt.Print("\n")
			fmt.Print("Last Name is invalid")
		} else if(errorField.FieldName == "gender") {
			fmt.Print("\n")
			fmt.Print("Gender is invalid")
		} else if(errorField.FieldName == "email") {
			fmt.Print("\n")
			fmt.Print("Email is invalid")
		}
	}
}