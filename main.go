package main

import (
	"fmt"
	"os"
)

func main() {
	cron := os.Args[1]
	fmt.Printf("Cron entered: %s\n\n", cron)
	cronParser := NewCronParser()
	err := cronParser.Parse(cron)
	if err != nil {
		fmt.Println("Error parsing cron string: ", err)
		return
	}
	fmt.Println("Output: ")
	cronParser.Print()
}
