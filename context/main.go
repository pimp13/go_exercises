package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()

	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

	fmt.Println("time is:", time.Since(now))
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond*200)
	defer cancel()

	val, err := fetchThirdPartyStuffWhichCanBeSlow()
	if err != nil {
		return 0, err
	}

	return val, nil
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 666, nil
}
