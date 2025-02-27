package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type OTPRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func StartServer(bot *Bot) {
	http.HandleFunc("/send-otp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req OTPRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = bot.SendMessage(req.Phone, req.Message)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OTP sent successfully"))
	})

	fmt.Println("Starting HTTP server on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
