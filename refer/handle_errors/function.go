// Just make a wrapper function to handle errors to wrap the complexity.

package main

import "fmt"
import "errors"

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not devide by Zero!")
	} else {
		return (a / b), nil
	}
}

func main() {

	// divide 4 by 0
	if result, err := divide(4, 0); err != nil {
		fmt.Println("Error occured: ", err)
	} else {
		fmt.Println("4/0 is", result)
        }
}
