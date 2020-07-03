# Go-Chat
Golang Chat is modular chat based on Uncle Bob's Clean Architecture principle. 

# Clean Architecture
Focus on usecase. All dependency interface located in folder usercase/dependency. We can change dependency implementation (database, sms sender, token generator, etc) via ioc.

# Using Go-Chat with default Implementation
Copy config.go.example to config.go in folder configuration.
```bash
go build && ./go-chat
```

# Using Go-Chat with custom Implementation
```go
package main

import (
    "log"
    "yourpackage"
    "github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
)

func main() {
    var u usecase.Usecase
    // Your Implementation for Chat Repository
    // Must implement usecase/dependency/ChatRepository interface
    chat := yourpackage.Chat{}
    u.ChatRepository = &chat

    inputChat := entity.Chat{
            SenderID:   "xxx",
            Timestamp:  34400,
            ChatType:   "group",
            ReceiverID: "groupid",
            Message:    "Testing Message",
        }
        
    savedChat, err := u.AddChat(inputChat)
    log.Println(savedChat,err)
}
```

# Telegram Group
Please join https://t.me/joinchat/I3ig_RhopGul_UKGjDhl8g