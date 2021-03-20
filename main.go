package main

import (
	"fmt"
	"os"
)

func main() {
	jobIndex := os.Getenv("JOB_INDEX")
	fmt.Printf("job index: %s\n", jobIndex)
	fmt.Println("done")
}
