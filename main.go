package main

import (
	"context"
	"fmt"
	irc "irc"
)

func main() {

	sampleCallback := func(ircc *irc.IrcClient, msg string) error {
		parsed_message, err := irc.ParseIrcMessage(msg)
		if err != nil {
			fmt.Println("Error while parsing message: ", err)
		}

		fmt.Printf("%+v\n", parsed_message)
		return nil
	}

	ircc := irc.NewTwitchIrcClient("justinfan123", "bruh")

	ircc.RegisterMessageCallback(sampleCallback)

	ctx := context.Background()
	client_status := make(chan error)
	go func() {
		client_status <- ircc.ClientLoop(ctx)
	}()

	// Send test.

	ircc.SendCapabilityRequest(irc.CapabilityTags)
	ircc.SendRawMessage("JOIN #alveussanctuary")
	ircc.SendRawMessage("PING tmi.twitch.tv")

	<-client_status // Wait for client to exit.
	fmt.Println("Client exited")

}
