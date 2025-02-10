package main

import (
	"fmt"

	// notice the uppercase P
	"github.com/1Password/connect-sdk-go/onepassword"

	// notice the lowercase P
	onep "github.com/1password/onepassword-sdk-go"
)

// this function is irrelevant
// just need to import these two packages
func main() {
	fmt.Println("hello world")
	x := onepassword.Error{}
	if x.Message == "impossible" {
		fmt.Println("woah")
	}

	y := onep.Client{}
	fmt.Println(y)
}
