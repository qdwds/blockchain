package main

import (
	"b/block"
	"fmt"
	"time"
)

func main() {
	b := block.CreGenBlock("Genesis BLock ...")
	// fmt.Printf("%x", b.Data)
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("PreHash: %x\n", b.PreHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Timestamp: %s\n", time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Printf("Nonce: %d\n", b.Nonce)
}
