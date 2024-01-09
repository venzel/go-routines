package main

import (
	"routines/channels"
	"routines/mutex"
)

func main() {
	mutex.Execute()

	channels.Execute()
}
