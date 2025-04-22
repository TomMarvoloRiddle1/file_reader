package main

import (
	"filer/env_reader"
	"fmt"
)

func main() {
	fmt.Println(env_reader.Getval_five("a"))
}
