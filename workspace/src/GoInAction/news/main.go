package main

import (
	"log"
	_ "news/matchers"
	"news/search"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("movie")
}
