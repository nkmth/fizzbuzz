package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/nkmth/fizzbuzz"
)

func main() {
	flag.Parse()

	arg, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Argument is only number.")
		os.Exit(1)
	}

	ret, _ := fizzbuzz.Convert(arg)
	fmt.Println(ret)
}
