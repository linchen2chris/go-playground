package main

import (
	"fmt"
	"os"
)

func main() {
	for n, arg := range os.Args {
		fmt.Println(n, " ", arg)
	}
	// fmt.Println(strings.Join(os.Args, "-"))

}
