package utils

import "time"

func Help() string {
	channel := make(chan string)
	go func() {
		channel <- "hello"
		channel <- "world"
	}()
	<-time.After(10 * time.Second)
	return <-channel
}
