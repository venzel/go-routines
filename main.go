package main

import "routines/channels"

func main() {
	japa := channels.NewAgent("Japa")
	erickson := channels.NewAgent("Erickson")
	joab := channels.NewAgent("Joab")
	marcos := channels.NewAgent("Marcos")

	tasks := 10
	agents := []*channels.Agent{japa, erickson, joab, marcos}

	channels.StartRound(tasks, agents)
}
