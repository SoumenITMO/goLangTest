package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")
var database *sql.DB

const (
	host = "localhost"
    port = 4012
	user = "postgres"
	password = "admin"
	dbname = "customers"
)

type errorField struct{
	FieldName string `json:"field"`
	Msg string `json:"msg"`
}

type RspData struct {
	Data []customerDtoSearch `json:"data"`
	RecordsTotal    int  `json:"recordsTotal"`
	RecordsFiltered int  `json:"recordsFiltered"`
}

type customerDto struct {
	CustomerId int `json:"id"`
	Firstname string `json:"fname"`
	Lastname string `json:"lname"`
	Dob string `json:"dob"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Address string `json:"address"`
}
type customerDtoSearch struct {
	CustomerId int `json:"id"`
	Firstname string `json:"fname"`
	Lastname string `json:"lname"`
	Dob string `json:"dob"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Address string `json:"address"`
	Buttons string `json:"actions"`
}

func dbConnection() (*sql.DB) {
	postgreConnStrint := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := sql.Open("postgres", postgreConnStrint)
	err = database.Ping()
	if err != nil {
		fmt.Println("Failed to connect Database ...")
	} else {
		fmt.Println("Successfully connected!")
	}
	return database
}

func init() {
	database = dbConnection()
}

func seedDB() {

	var firstname string
	var lastname string
	var date time.Time
	var gender string
	var email string
	var address string
	var SQL string

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Soumen"
	lastname = "Banerjee"
	gender = "M"
	email = "soumen@gmail.com"
	address = "India"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Matt"
	lastname = "Damon"
	gender = "M"
	email = "matt@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Veronika"
	lastname = "Escacena"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Veronika1"
	lastname = "Escacena"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Veronika2"
	lastname = "Escacena"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Denise"
	lastname = "Richards"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Denise"
	lastname = "Jonas"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)

	date = time.Date(2020, 8, 25, 12, 20, 25,54, time.UTC)
	firstname = "Denise"
	lastname = "Crosby"
	gender = "F"
	email = "veronia@damon.com"
	address = "US"
	SQL = "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5,$6)"
	_,_ = database.Exec(SQL, firstname, lastname, date, gender, email, address)
}

func fetchCustomerDetailsById(response http.ResponseWriter, request *http.Request) {

	var dob time.Time
	customer := customerDto{}
	customerId := request.URL.Query().Get("customerId")
	getCustomerDetailQuery := "SELECT * FROM customer where id = $1"
	customerData := database.QueryRow(getCustomerDetailQuery, customerId)

	err := customerData.Scan(&customer.CustomerId, &customer.Firstname, &customer.Lastname, &dob, &customer.Gender,
		&customer.Email, &customer.Address)

	customer.Dob = dob.Format("2006-01-02")
	singleCustomerData,_ := json.Marshal(customer)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	if err != nil {
		_, _ = fmt.Fprint(response, string(singleCustomerData))
	} else {
		_, _ = fmt.Fprint(response, string(singleCustomerData))
	}
}

func updateCustomerDetails(response http.ResponseWriter, request *http.Request) {

	if(request.Method != "POST") {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.WriteHeader(http.StatusMethodNotAllowed)
	}

	var customerId = request.PostFormValue("customerId")
	var firstName = request.PostFormValue("firstname")
	var lastName = request.PostFormValue("lastname")
	var email = request.PostFormValue("email")
	var address = request.PostFormValue("address")
	var dob = request.PostFormValue("dob")
	var gender = request.PostFormValue("gender")

	var valid = true

	match := rxEmail.Match([]byte(email))

	errorData := errorField{}
	errorList := make([]errorField, 0)

	firstNameValid, _ := strconv.ParseInt(firstName, 0, 2)
	lastNameValid, _ := strconv.ParseInt(lastName, 0, 2)

	if match == false {
		errorData.Msg = "Invalid Email"
		errorData.FieldName = "email"
		errorList = append(errorList, errorData)
		valid = false

	}
	if(firstNameValid != 0) {
		errorData.Msg = "First Name Invalid"
		errorData.FieldName = "firstname"
		errorList = append(errorList, errorData)
		valid = false

	}
	if(lastNameValid != 0) {
		errorData.Msg = "Last Name Invalid"
		errorData.FieldName = "lastname"
		errorList = append(errorList, errorData)
		valid = false
	}
	if(address == "") {
		errorData.Msg = "Address is empty"
		errorData.FieldName = "address"
		errorList = append(errorList, errorData)
		valid = false
	}

	genderMaleCheck := gender != "M"
	genderFemaleCheck := gender != "F"

	if(genderMaleCheck && genderFemaleCheck) {
		errorData.Msg = "Invalid Gender"
		errorData.FieldName = "gender"
		errorList = append(errorList, errorData)
		valid = false
	}
	val,_ := time.Parse("2006-01-02", dob)

	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")

	if(!valid) {
		responseData,_ := json.Marshal(errorList)
		fmt.Fprint(response, string(responseData))
	} else {
		updateSQL := "UPDATE customer set firstname = $1 ,lastname = $2, email = $3, address = $4, dob = $5, gender = $6 where id = $7"
		_,_ = database.Exec(updateSQL, firstName, lastName, email, address, val, gender, customerId)
		fmt.Fprint(response, 1)
	}
}

func createCustomerDetails(response http.ResponseWriter, request *http.Request) {
	if(request.Method != "POST") {

		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
	} else {

		var firstName = request.PostFormValue("firstname")
		var lastName = request.PostFormValue("lastname")
		var address = request.PostFormValue("address")
		var dob = request.PostFormValue("dob")
		var gender = request.PostFormValue("gender")
		var email = request.PostFormValue("email")

		var valid = true

		match := rxEmail.Match([]byte(email))

		errorData := errorField{}
		errorList := make([]errorField, 0)

		firstNameValid, _ := strconv.ParseInt(firstName, 0, 2)
		lastNameValid, _ := strconv.ParseInt(lastName, 0, 2)

		if match == false {
			errorData.Msg = "Invalid Email"
			errorData.FieldName = "email"
			errorList = append(errorList, errorData)
			valid = false

		}
		if (firstNameValid != 0) {
			errorData.Msg = "First Name Invalid"
			errorData.FieldName = "firstname"
			errorList = append(errorList, errorData)
			valid = false

		}
		if (lastNameValid != 0) {
			errorData.Msg = "Last Name Invalid"
			errorData.FieldName = "lastname"
			errorList = append(errorList, errorData)
			valid = false
		}
		if (address == "") {
			errorData.Msg = "Address is empty"
			errorData.FieldName = "address"
			errorList = append(errorList, errorData)
			valid = false
		}

		genderMaleCheck := gender != "M"
		genderFemaleCheck := gender != "F"

		if (genderMaleCheck && genderFemaleCheck) {
			errorData.Msg = "Invalid Gender"
			errorData.FieldName = "gender"
			errorList = append(errorList, errorData)
			valid = false
		}
		if (valid == true) {
			val, _ := time.Parse("2006-01-02", dob)
			insertSQL := "INSERT into customer(firstname, lastname, dob, gender, email, address) values($1 ,$2, $3, $4, $5, $6)"
			_, _ = database.Exec(insertSQL, firstName, lastName, val, gender, email, address)
		}

		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Content-Type", "application/json")
		responseData, _ := json.Marshal(errorList)
		_, _ = fmt.Fprint(response, string(responseData))
	}
}

func deleteCustomerData(response http.ResponseWriter, request *http.Request) {

	if(request.Method != "POST") {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.WriteHeader(http.StatusMethodNotAllowed)
	}

	customerId := request.URL.Query().Get("customerId")
	deleteCustomerQuery := "DELETE FROM customer where id = $1"
	_,_ = database.Exec(deleteCustomerQuery, customerId)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(response, 1)
}

func searchCustomers(response http.ResponseWriter, request *http.Request) {

	var count int

	if(request.Method != "GET") {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		countrows, err := database.Query("SELECT count(*) FROM customer")
		if err != nil {
			fmt.Println("Query failed.....")
		}

		for countrows.Next() {
			countrows.Scan(&count)
		}
		if (count == 0) {
			seedDB()
		}

		rows, err := database.Query("SELECT * FROM customer")
		if err != nil {
			fmt.Println("Query failed.....")
		}
		customer := customerDtoSearch{}
		customerList := make([]customerDtoSearch, 0)
		var dob time.Time

		for rows.Next() {
			rows.Scan(&customer.CustomerId, &customer.Firstname, &customer.Lastname, &dob, &customer.Gender,
				&customer.Email, &customer.Address)
			customer.Buttons = "<a href=edit.html?ID=" + strconv.Itoa(customer.CustomerId) + "> " +
				"EDIT </a> | <a href='#' onclick='deleteCustomer(" + strconv.Itoa(customer.CustomerId) + ")'> DELETE"
			customer.Dob = dob.Format("02/01/2006")
			if (customer.Gender == "F") {
				customer.Gender = "Female"
			} else {
				customer.Gender = "Male"
			}
			customerList = append(customerList, customer)
		}
		result := RspData{
			Data:            customerList,
			RecordsTotal:    len(customerList),
			RecordsFiltered: len(customerList),
		}

		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Content-Type", "application/json")
		responseData, _ := json.Marshal(result)
		_, _ = fmt.Fprint(response, string(responseData))
	}
}

func main() {

	http.HandleFunc("/register", createCustomerDetails)
	http.HandleFunc("/home", searchCustomers)
	http.HandleFunc("/getcustomer", fetchCustomerDetailsById)
	http.HandleFunc("/updatecustomer", updateCustomerDetails)
	http.HandleFunc("/deletecustomer", deleteCustomerData)

	http.ListenAndServe(":8084", nil)
}
