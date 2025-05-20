package main

import (
	"fmt"

	"github.com/mindfarm/codename"
)

func main() {
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 20; i++ {
		name := codename.Generate(rng, 0)
		fmt.Println(name)
	}
}
