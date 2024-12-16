// go routines managed in go runtime not OS thread level
// hence faster and more memory saved

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message = map[string]string

func main() {
	fmt.Println(time.Now().Format("13:02:12"))
	// create a wait group
	wg := &sync.WaitGroup{}
	// we are waiting for 2
	wg.Add(2)
	//ch := make(chan string)
	chM := make(chan Message, 2)
	id := getUserByName("John")
	//println(id)
	go getUserChats(id, chM, wg)
	go getUserFriends(id, chM, wg)
	wg.Wait()
	// now close after receiving done from wg
	close(chM)
	// deadlock, we don't know when the goroutine will finish
	for msg := range chM {
		log.Println(msg)
	}

	//fmt.Printf("ID: %v\n, Chats: %v\n, Friends: %v\n", id, chats, friends)
	log.Println(id)
	// log.Println(chats)
	// log.Println(friends)

}

func getUserChats(id string, ch chan<- Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	ch <- Message{
		"John": "Hey",
	}
	// don't close chan here
	// but signal donw
	wg.Done()
}

func getUserFriends(id string, ch chan<- Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	ch <- Message{
		"John": "Kamaa",
	}
	wg.Done()
}

func getUserByName(name string) string {
	// time.Sleep(time.Second * 3)
	return fmt.Sprintf("%s-2", name)
}
