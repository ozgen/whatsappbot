package main

import (
	"fmt"
	"github.com/ozgen/whatsappbot/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot, err := internal.NewBot()
	if err != nil {
		fmt.Println("Failed to initialize bot:", err)
		return
	}

	go func() {
		err := bot.Start()
		if err != nil {
			fmt.Println("Failed to start bot:", err)
		}
	}()

	go internal.StartServer(bot)

	// Handle termination signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	bot.Client.Disconnect()
	fmt.Println("WhatsApp bot stopped")
}
