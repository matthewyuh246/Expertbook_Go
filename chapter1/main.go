package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matthewyuh246/expertbook/chapter1/vault"
)

func main() {
	// reverce()
	// filter()
	// divid()
	// deduplication()
	// c := &syncutil.Counter{
	// 	Name: "Access",
	// }

	// fmt.Println(c.Increment())
	// fmt.Println(c.View())

	s := vault.NewSecret()

	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("failed to json encode, erro =", err)
	}
}
