package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"fmt"
	"log"
	"math/big"
)

//	挖矿难度
var targetBig = 6

type ProofOfWork struct {
	Block  *Block //	当前需要验证的区块
	Target *big.Int
}

//	创建工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-targetBig))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	hashInt := new(big.Int)
	var hash [32]byte
	//	TODO 这里需要debugger看一下
	for {
		dataBytas := pow.prepareData(nonce)
		hash = sha256.Sum256(dataBytas)
		fmt.Printf("生成的哈希：\r %x ", hash)
		hashInt.SetBytes(hash[:])
		if pow.Target.Cmp(hashInt) == 1 {
			break
		}
		nonce += 1
	}
	return hash[:], int64(nonce)
}

//	拼接数据 => 字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreHash,
			pow.Block.Data,
			intTuHex(pow.Block.Timestamp),
			intTuHex(int64(pow.Block.Height)),
			intTuHex(int64(targetBig)),
			intTuHex(int64(nonce)),
		}, []byte{},
	)
	return data
}

//	int => [int]byte
func intTuHex(num int64) []byte {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, num); err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
