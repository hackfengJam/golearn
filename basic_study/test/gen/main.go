package main

import (
	"fmt"
	"math/rand"
	"time"
)

// genTicket
func genTicket(ticketType int, uid int64, rewardNo int32) string {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := 10000 + randGen.Intn(10000)
	return fmt.Sprintf("%s-%d-%d-%d-%d-%d", "111", ticketType, uid, rewardNo, time.Now().UnixNano()/1000, seed)
}

// Gen
func main() {
	fmt.Println(genTicket(1, 100000000, 10000))
	fmt.Println(len(genTicket(2, 100000000, 10000)))
}
