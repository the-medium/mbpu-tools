package main

import (
  "os"
  
  "github.com/the-medium/mbpu-tools/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}