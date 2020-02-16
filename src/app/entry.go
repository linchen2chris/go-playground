package main

import (
	"app/util"
	"consts"
	"consts/de"
	"fmt"
)

func main() {
	fmt.Println("app/entry =>> main")
	fmt.Println(morning)
	fmt.Println(consts.Greeting)
	fmt.Println(de.Greet)
	fmt.Println(util.Add(1, 2))
}
