package main

import (
	"bc/block"
)

func main() {
	b := block.CreGenBlockchain("da")
	defer b.DB.Close()
	b.AddBlockToBlockchain("111")
	b.AddBlockToBlockchain("222")
	b.AddBlockToBlockchain("333")
	b.PrintChain()

}
