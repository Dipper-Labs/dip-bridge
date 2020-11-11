package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/Dipper-Labs/dip-bridge/config"
	sdkconfig "github.com/Dipper-Labs/go-sdk/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 2 {
		log.Fatal("wrong args. should use by: dip-bridge [config file path]")
	}

	log.Printf("Enter Keystore Password: ")
	keystorePassword, err := bufio.NewReader(os.Stdout).ReadString('\n')
	if err != nil {
		log.Fatalf("get keystore password failed:[%v]", err)
	}
	keystorePassword = keystorePassword[:len(keystorePassword)-1]

	config.Init(os.Args[1])
	sdkconfig.SetKeystorePassword(keystorePassword)
	bridge := NewBridge()

	ctx := context.Background()
	go bridge.RunBridge(ctx)

	headerChan, sub := bridge.SubscribeNewHead(ctx)
	for {
		select {
		case newHeader := <-headerChan:
			bridge.UpdateEthHeaderBlock(newHeader.Number.Int64())

		case err := <-sub.Err():
			log.Fatalf("Ethereum new header subscription broken:[%v]\n", err)
		}
	}
}
