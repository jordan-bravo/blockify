package main

import (
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
)

func binarySearch(client *rpcclient.Client, maxBlockHeight int64, targetTime int64) int64 {
	var leftBlockHeight, rightBlockHeight int64 = 0, maxBlockHeight

	for leftBlockHeight <= rightBlockHeight {
		midBlockHeight := (leftBlockHeight + rightBlockHeight) / 2

		midBlockHash, err := client.GetBlockHash(midBlockHeight)
		if err != nil {
			log.Fatal(err)
		}
		midBlock, err := client.GetBlock(midBlockHash)
		if err != nil {
			log.Fatal(err)
		}
		leftBlockHash, err := client.GetBlockHash(leftBlockHeight)
		if err != nil {
			log.Fatal(err)
		}
		leftBlock, err := client.GetBlock(leftBlockHash)
		if err != nil {
			log.Fatal(err)
		}
		rightBlockHash, err := client.GetBlockHash(rightBlockHeight)
		if err != nil {
			log.Fatal(err)
		}
		rightBlock, err := client.GetBlock(rightBlockHash)
		if err != nil {
			log.Fatal(err)
		}

		midBlockTime := midBlock.Time
		leftBlockTime := leftBlock.Time
		rightBlockTime := rightBlock.Time

		if midBlockTime == targetTime {
			return midBlockHeight
		} else if midBlockTime < targetTime {
			leftBlockHeight = midBlockHeight + 1
		} else {
			rightBlockHeight = midBlockHeight - 1
		}
	}
	return leftBlockHeight
}

func main() {
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:8332",
		User:         "yourrpcuser",
		Pass:         "yourrpcpass",
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	bestBlockHash, err := client.GetBestBlockHash()
	if err != nil {
		log.Fatal(err)
	}
	bestBlock, err := client.GetBlock(bestBlockHash)
	if err != nil {
		log.Fatal(err)
	}
	maxBlockHeight := bestBlock.Height

	var year, month, day, hour, minute, sec int
	fmt.Println("Enter year:")
	fmt.Scan(&year)
	fmt.Println("Enter month:")
	fmt.Scan(&month)
	fmt.Println("Enter day:")
	fmt.Scan(&day)
	fmt.Println("Enter hour:")
	fmt.Scan(&hour)
	fmt.Println("Enter minute:")
	fmt.Scan(&minute)
	fmt.Println("Enter second:")
	fmt.Scan(&sec)

	givenDateTime := time.Date(year, time.Month(month), day, hour, minute, sec, 0, time.UTC)
	targetTime := givenDateTime.Unix()

	fmt.Println("Finding block height...")
	result := binarySearch(client, maxBlockHeight, targetTime)
	fmt.Printf("The block height at this date and time was: %d\n", result)
}

