package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/dnlo/struct2csv"
)

type Person struct {
	Id                      uint   `json:"id"`
	Uid                     string `json:"uid"`
	Password                string `json:"password"`
	First_name              string `json:"first_name"`
	Last_name               string `json:"last_name"`
	Username                string `json:"username"`
	Email                   string `json:"email"`
	Avatar                  string `json:"avatar"`
	Gender                  string `json:"gender"`
	Phone_number            string `json:"phone_number"`
	Social_insurance_number string `json:"social_insurance_number"`
	Date_of_birth           string `json:"date_of_birth"`
	Employment              string `json:"employment"`
	Address                 string `json:"address"`
	Credit_card             string `json:"credit_card"`
	Subscription            string `json:"subscription"`
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
	csvWriter := csv.NewWriter(file1)
	var person Person
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

		json.Unmarshal(body, &person)

		fmt.Println(person.Gender == "Female" || person.Gender == "Male")
		if person.Gender == "Male" || person.Gender == "Female" {
			enc := struct2csv.New()
			var rows [][]string
			headers, err := enc.GetColNames(person)
			row, err := enc.GetRow(person)

			if err != nil {
				return
			}
			rows = append(rows, headers)
			rows = append(rows, row)

			for i := 0; i < len(rows); i++ {
				fmt.Println(rows[i])
				_ = csvWriter.Write(rows[i])
			}
		}
	}
	csvWriter.Flush()
}
