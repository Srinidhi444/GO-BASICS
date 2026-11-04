package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()

	id := getUserByName("john")
	fmt.Println(id)

	ch := make(chan *Message, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go getUserChats(id, ch, &wg)
	go getUserFriends(id, ch, &wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println("Total time:", time.Since(now))
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)

	ch <- &Message{
		chats: []string{"chat-1", "chat-2", "chat-3"},
	}
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	ch <- &Message{
		friends: []string{"friend-1", "friend-2", "friend-3"},
	}
}

func getUserByName(name string) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("%s-2", name)
}
