package main

import (
	"fmt"

	"github.com/g4j4/sensed/version"
)

func main() {
	fmt.Println("Version:", version.Version)
	fmt.Println("Commit hash:", version.CommitHash)
	fmt.Println("Build timestamp:", version.BuildTime)
}
