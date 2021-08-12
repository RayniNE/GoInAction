package main

import (
	"log"
	"os"

	_ "github.com/RayniNE/Chapter2/matchers"
	"github.com/RayniNE/Chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run()
}
