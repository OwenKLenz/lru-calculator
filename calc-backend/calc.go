package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type responseBody struct {
	Num1     int
	Operator string
	Num2     int
}

func performCalc(mathStuff responseBody) string {
	var result string

	num1 := mathStuff.Num1
	num2 := mathStuff.Num2
	operator := mathStuff.Operator

	switch operator {
	case "+":
		result = strconv.Itoa(num1 + num2)
	case "-":
		result = strconv.Itoa(num1 - num2)
	case "/":
		result = fmt.Sprintf("%f", float64(num1)/float64(num2))
	case "*":
		result = strconv.Itoa(num1 * num2)
	}

	return result
}

func parseJSON(body io.Reader) responseBody {
	decoder := json.NewDecoder(body)

	var b responseBody

	decoder.Decode(&b)

	return b
}

func calc(response http.ResponseWriter, req *http.Request) {
	mathStuff := parseJSON(req.Body)

	fmt.Println("Calculating:", mathStuff.Num1, mathStuff.Operator, mathStuff.Num2)
	result := performCalc(mathStuff)
	fmt.Println("Calculated:", result)
	fmt.Fprintf(response, result)
}

func main() {
	http.HandleFunc("/calc", calc)
	fmt.Println("Running Calculator Backend on Docker")
	log.Panic(http.ListenAndServe(":8080", nil))
}
