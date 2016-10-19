// Slightly more complex usage of flags command.
// Thank you @ralch
// http://blog.ralch.com/tutorial/golang-subcommands/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	askCommand := flag.NewFlagSet("ask", flag.ExitOnError)
	questionFlag := askCommand.String("question", "", "Question that you are asking for")

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
	messageFlag := sendCommand.String("message", "", "Text message")

	if len(os.Args) == 1 {
		fmt.Println("usage: siri <command> [<args>]")
		fmt.Println("<commands> are:")
		fmt.Println("\task\t\tAsk questions")
		fmt.Println("\tsend\t\tSend messages to your contacts")
		return
	}

	switch os.Args[1] {
	case "ask":
		askCommand.Parse(os.Args[2:])
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	// Now we run anything that was parsed.
	if askCommand.Parsed() {
		if *questionFlag == "" {
			fmt.Println("Please supply the question using -question option.")
			return
		}
		fmt.Printf("You asked: %q\n", *questionFlag)
	}

	if sendCommand.Parsed() {
		if *recipientFlag == "" {
			fmt.Println("Please supply the recipient using -recipient option.")
			return
		}

		if *messageFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}

		fmt.Printf("Your message is send to %q.\n", *recipientFlag)
		fmt.Printf("Message: %q.\n", *messageFlag)
	}
}
