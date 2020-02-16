package main

import (
	"app/util"
	_ "consts"
	. "consts/de"
	"fmt"
)

func main() {
	fmt.Println("app/entry =>> main")
	fmt.Println(morning)
	fmt.Println(Greet)
	fmt.Println(util.Add(1, 2))
}
