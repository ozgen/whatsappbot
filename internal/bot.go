package internal

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	_ "go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Bot struct {
	Client *whatsmeow.Client
}

func NewBot() (*Bot, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:whatsup.db?_foreign_keys=on", dbLog)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device store: %v", err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	return &Bot{Client: client}, nil
}

func (b *Bot) Start() error {
	if b.Client.Store.ID == nil {
		qrChan, _ := b.Client.GetQRChannel(context.Background())
		err := b.Client.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect client: %v", err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan the QR Code:", evt.Code)
				err := SaveQRCode(evt.Code)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := b.Client.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect client: %v", err)
		}
	}
	return nil
}

func (b *Bot) SendMessage(number, message string) error {
	recipientJID := types.NewJID(number, "s.whatsapp.net")

	_, err := b.Client.SendMessage(context.Background(), recipientJID, &waE2E.Message{
		Conversation: proto.String(message),
	})
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}
