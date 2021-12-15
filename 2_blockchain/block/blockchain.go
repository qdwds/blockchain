package block

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbName = "blockchain.db"     //	数据库名
const blockTabName = "blockchains" //	表名
const newHash string = "newHash"   //	最新hash
type BlockChain struct {
	Tip []byte //	最新区块的hash
	DB  *bolt.DB
}

//	判断数据库是否存在
func DBExits() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//	创建区块链的 创世区块
func CreGenBlockchain(data string) *BlockChain {
	if DBExits() {
		fmt.Println("创世区块已经存在!")
		os.Exit(1)
	}

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("开始创建创世区块～")
	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		d, err := tx.CreateBucket([]byte(blockTabName))
		if err != nil {
			log.Panic(err)
		}
		//	开始创建创世区块
		if d != nil {
			cgb := CreGenBlock(data)
			//	hash 作为key值  序列化的数据作为value
			if err := d.Put(cgb.Hash, cgb.Serilize()); err != nil {
				log.Panic(err)
			}
			//	存储最新的hash
			if err = d.Put([]byte(newHash), cgb.Hash); err != nil {
				log.Panic(err)
			}
			blockHash = cgb.Hash
		}
		return nil
	})

	return &BlockChain{blockHash, db}
}

//	添加区块到区块链中
func (bc *BlockChain) AddBlockToBlockchain(data string) {
	//	TODO 缺少一个 是否存在创世区块的校样
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		d := tx.Bucket([]byte(blockTabName))
		if d != nil {
			fmt.Println(bc.Tip)
			//	获取最新区块的信息
			blockBytes := d.Get(bc.Tip)
			//	反序列化得到最新的区块信息
			block := Deserilize(blockBytes)
			//	创建新区块
			newBlock := NewBlock(block.Height+1, block.Hash, data)
			//	key当前区块hash value序列化的数据
			err := d.Put(newBlock.Hash, newBlock.Serilize())
			if err != nil {
				log.Panic(err)
			}
			//	更新最新的hash
			err = d.Put([]byte(newHash), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			//	更新最新的hash
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}
