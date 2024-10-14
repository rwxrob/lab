package main

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/chart/loader"
)

func main() {
	chart, err := loader.Load(os.Args[1])
	if err != nil {
		log.Fatalf("Error loading chart: %v", err)
	}
	fmt.Println(chart)
}
