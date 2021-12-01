package main

import (
	"GeekBrains-Level-1/lesson-8/config"
	"fmt"
	"os"
)

func main() {

	conf, err := config.LoadConfig("configuration.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(conf)

}
