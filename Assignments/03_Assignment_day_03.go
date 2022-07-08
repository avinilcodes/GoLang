package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	Id                      uint         `json:"id"`
	Uid                     string       `json:"uid"`
	Password                string       `json:"password"`
	First_name              string       `json:"first_name"`
	Last_name               string       `json:"last_name"`
	Username                string       `json:"username"`
	Email                   string       `json:"email"`
	Avatar                  string       `json:"avatar"`
	Gender                  string       `json:"gender"`
	Phone_number            string       `json:"phone_number"`
	Social_insurance_number string       `json:"social_insurance_number"`
	Date_of_birth           string       `json:"date_of_birth"`
	Employment              Employeeinfo `json:"employment"`
	Address                 Address      `json:"address"`
	Credit_card             Credit_card  `json:"credit_card"`
	Subscription            Subscription `json:"subscription"`
}
type Employeeinfo struct {
	title     string
	key_skill string
}
type Credit_card struct {
	cc_number string
}
type Coordinates struct {
	lat float64
	lng float64
}
type Address struct {
	city           string
	street_name    string
	street_address string
	zip_code       string
	state          string
	country        string
	coordinates    Coordinates
}

type Subscription struct {
	plan           string
	status         string
	payment_method string
	term           string
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	file1, err := os.Create("output.csv")
	if err != nil {
		return
	}
	defer file1.Close()
	for i := 0; i < n; i++ {
		resp, err := http.Get("https://random-data-api.com/api/users/random_user")
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		var person Person
		json.Unmarshal(body, &person)
		if person.Gender == "Male" || person.Gender == "Female" {
			// enc := struct2csv.New()
			// rows, err := enc.Marshal(person)
			// if err != nil {
			// 	return
		}
	}
}
