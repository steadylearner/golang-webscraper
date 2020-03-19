// You can also use fmt.Errorf instead

package main

import (
    "fmt"
    "errors"
)

func main() {
    myErr := errors.New("Error happend!")

    fmt.Printf("Type of myErr is %T \n" , myErr)
    fmt.Printf("Value of myErr is %#v \n", myErr)
    fmt.Println(myErr)
}
