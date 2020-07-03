# Go-Chat
Golang Chat

# Clean Architecture
Focus to usecase. We can change dependency (database, sms sender, etc) via ioc.

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