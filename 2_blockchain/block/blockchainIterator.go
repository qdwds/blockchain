package block

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	CurHash []byte
	DB      *bolt.DB
}

//	迭代器
func (bc *BlockChain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Tip, bc.DB}
}

//	获取下一个 区块内容
func (bci *BlockchainIterator) Next() *Block {
	var block *Block
	err := bci.DB.View(func(tx *bolt.Tx) error {
		d := tx.Bucket([]byte(blockTabName))
		if d != nil {
			//	存储的时候是根据hash作为key存储 取得时候根据hash取
			curBlockBytes := d.Get(bci.CurHash)
			block = Deserilize(curBlockBytes)
			//	更新hash
			bci.CurHash = block.PreHash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}

//	循环输出
func (bc *BlockChain) PrintChain() {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PreHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println("--------------------------------------------------------------")

		var hashInt big.Int
		hashInt.SetBytes(block.PreHash)
		//	找到最后一个结束
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}
