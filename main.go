package main

import (
	"context"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("wrong args")
	}

	ctx := context.Background()

	bridge := NewBridge(os.Args[1])
	go bridge.RunBridge(ctx)

	headerChan, sub := bridge.SubscribeNewHead(ctx)
	for {
		select {
		case newHeader := <-headerChan:
			bridge.UpdateEthHeaderBlock(newHeader.Number.Int64())

		case err := <-sub.Err():
			log.Fatal(err)
		}
	}
}
