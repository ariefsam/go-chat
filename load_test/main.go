package main

import (
	"log"
	"sync"
	"time"

	"github.com/ariefsam/go-chat/configuration"
	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/implementation"
	"github.com/ariefsam/go-chat/repository"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/jinzhu/copier"
)

func main() {
	var u usecase.Usecase
	chatRepo := repository.Chat{}
	copier.Copy(&chatRepo, &configuration.Config.MySQL)
	u.ChatRepository = &chatRepo
	u.IDGenerator = &implementation.IDGenerator{}
	var wg sync.WaitGroup
	start := time.Now().Unix()
	for i := 0; i <= 10000; i++ {
		c := entity.Chat{
			ID:      u.IDGenerator.Generate(),
			Message: "Ini test message",
		}
		wg.Add(1)
		go func() {
			u.ChatRepository.Save(c)
			wg.Done()
		}()

	}
	log.Println("waiting")
	wg.Wait()
	log.Println("Done waiting")
	end := time.Now().Unix()
	log.Println("Waktu proses: ", (end - start), "detik")

}
