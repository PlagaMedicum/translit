package main

import (
	"fmt"
	"os"

	translit "github.com/PlagaMedicum/translit/golang"
)

func main() {
	str := os.Args[1]
	fmt.Printf(translit.CyrToLat(str))
}
