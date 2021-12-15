package block

import (
	"fmt"
	"time"
)

type Block struct {
	Height    int64
	Hash      []byte
	PreHash   []byte
	Data      []byte
	Timestamp int64
	Nonce     int64 //	工作量证明
}

//	创建新区块
func NewBlock(height int64, preHash []byte, data string) *Block {
	block := &Block{
		Height:    height,
		Hash:      nil,
		PreHash:   preHash,
		Data:      []byte(data),
		Timestamp: time.Now().Unix(),
		Nonce:     0,
	}
	//	当前区块 创建一个工作量证明
	pow := NewProofOfWork(block)

	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	fmt.Println()
	return block
}

//	创建传世区块
func CreGenBlock(data string) *Block {
	return NewBlock(1, make([]byte, 32, 32), data)
}
