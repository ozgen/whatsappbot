
# WhatsApp Bot 📩

This is a **Go-based WhatsApp bot** built using [`whatsmeow`](https://github.com/tulir/whatsmeow). It allows sending messages via WhatsApp Web and supports **SQLite-based session storage**.

---

## 🚀 Features
- ✅ **QR Code Authentication** for WhatsApp Web.
- ✅ **Send Messages** to WhatsApp users.
- ✅ **SQLite Storage** for session persistence.
- ✅ **Docker Support** for containerized deployment.

---

## ⚙️ Installation & Setup

### **1️⃣ Prerequisites**
- **Go** (≥ 1.23.5) installed: [Download](https://go.dev/dl/)
- **Docker** installed: [Install](https://docs.docker.com/get-docker/)
- **WhatsApp account** for authentication

### **2️⃣ Clone the Repository**
```sh
git clone https://github.com/ozgen/whatsappbot.git
cd whatsappbot
```

### **3️⃣ Install Dependencies**
```sh
go mod tidy
```

### **4️⃣ Run the Bot**
```sh
go run cmd/main.go
```
- This will generate a **QR Code** (`whatsapp_qr.png`).
- Open WhatsApp **Web > Linked Devices**, scan the QR code.
- Copy code in the log and generate qr code image with this website https://www.qr-code-generator.com
---

## 🐳 Running with Docker

### **1️⃣ Build & Run the Container**
```sh
docker build -t whatsappbot:v1 .
docker run -d --name whatsappbot -p 9090:9090 -v "$(pwd)/data:/app/data" whatsappbot
```

### **2️⃣ Using Docker Compose**
```sh
docker-compose up -d
```

**Ensure your `docker-compose.yml` includes the volume:**
```yaml
version: '3.1'
services:
  whatsappbot:
    image: whatsappbot:v1
    platform: linux/amd64
    ports:
      - "9090:9090"
    volumes:
      - ./data:/app/data
```

---

## 🔧 API Usage (Send Message)

Once the bot is running, **you can send messages via API**:

### **Request**
```sh
curl -X POST "http://localhost:9090/send-message" \
     -H "Content-Type: application/json" \
     -d '{"phone": "491234567890", "message": "123456"}'
```

### **Response**
```json
{
  "message": "Message sent successfully"
}
```

## 🛠️ Troubleshooting

### **1️⃣ QR Code Not Scanning**
- Ensure WhatsApp **Web** is open and working.
- Delete `whatsapp_qr.png` and restart the bot.

### **2️⃣ Docker Error: `GLIBC_2.34 not found`**
If you get `glibc` issues, change **Dockerfile base image**:
```dockerfile
FROM golang:1.23.5 AS builder
```
Rebuild:
```sh
docker build --no-cache -t whatsappbot:v1 .
```

### **3️⃣ `failed to initialize database: CGO_ENABLED=0`**
- The bot **requires CGO** for SQLite.
- Rebuild **without** `CGO_ENABLED=0`:
```sh
CGO_ENABLED=1 go build -o whatsappbot ./cmd/main.go
```

---

## 🔗 References
- [WhatsMeow Library](https://github.com/tulir/whatsmeow)
- [WhatsApp Web API](https://web.whatsapp.com/)

---
