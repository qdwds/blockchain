package block

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
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

//	创建创世区块
func CreGenBlock(data string) *Block {
	return NewBlock(1, make([]byte, 32, 32), data)
}

//	序列化
func (b *Block) Serilize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()

}

//	序列化 => 将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	//	创建一个butter
	var result bytes.Buffer
	//创建一个编码器
	encoder := gob.NewEncoder(&result)
	//编码-->打包
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//	反序列化
func Deserilize(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
