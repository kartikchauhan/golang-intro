package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Okay                = 200
	Created             = 201
	BadRequest          = 400
	NotFound            = 404
	InternalServerError = 500
)

type customError struct {
	statusCode int
	msg        string
}

func (ce customError) Error() string {
	return fmt.Sprintf("{message: %v, status: %v}", ce.msg, ce.statusCode)
}

func sqr(f float64) (float64, error) {
	if f < 0 {
		return 0, customError{
			statusCode: 400,
			msg:        "Can't pass a negative number",
		}
		// return 0, fmt.Errorf("Can't pass a negative number, status: %v", BadRequest)
	}

	return f * f, nil
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sq, err := sqr(float64(num))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Square of", num, "=>", sq)
}
