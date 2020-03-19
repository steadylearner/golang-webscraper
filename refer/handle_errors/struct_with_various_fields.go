package main

import "fmt"

// simple HTTP error
type HttpError struct {
	status int
	method string
}

// HttpError implements Error method
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		httpError.method, httpError.status)
}

// mock function: sends HTTP request
func GetUserEmail(userId int) (string, error) {

	// request failed, return an error
	return "", &HttpError{403, "GET"}
}

func main() {

	// get user email address
	if email, err := GetUserEmail(1); err != nil {
		fmt.Println(err) // print error

		// if user is unauthorized, perform session cleaning
		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning...")
		}

	} else {
		// do something with the `email`
		fmt.Println("User email is", email)
	}
}
