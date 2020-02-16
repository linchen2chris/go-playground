package consts

import "fmt"

var Greeting = "hello"

// init ...
func init() {
	fmt.Println("greet/parent => init()")
}
