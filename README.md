
A simple Go wrapper for the Telegram Bot API

This just implements a few of the Telegram Bot API methods.

I am writing this first as a learning task for Go, but also leveraging
it to write a simple program that sends updates about a trading platform
error to an Instagram Channel.

**Still at a young stage. Will be improving it with time**

A simple, contrived example.

This just sends `msg` to the channel name specified.
You will have to add your bot as an administrator to the channel.

Once added, it can send messages(text, files, etc) to the channel.


```go
package main

import (
	"fmt"
	"os"
	"github.com/mekicha/telebot"
	"time"
	"log"
)


func main() {

	bot, err := telebot.NewBot(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Millisecond * 500)


	go func() {
		for t := range ticker.C {
			msg := fmt.Sprintf("Ticking at %d", t)
			err := bot.SendToChannel("@channelName", msg)

			if err != nil {
				break
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()

	fmt.Println("tired of sending")


}
```



