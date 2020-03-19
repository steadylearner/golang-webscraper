package main

import "fmt"

// Make a struct
type MyError struct{}

// Impl "Error" method to be an error?
func (myErr *MyError) Error() string {
    return "Error happend."
}

func main() {
   // Make an error
   myErr := &MyError{}

   // Print error message
   fmt.Println(myErr)
}
