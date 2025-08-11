package main

import (
	"github.com/captainmio/go-crud-post/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.Routes()
}
