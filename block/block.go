package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号 快速定位
	Timestamp     int64  //时间戳 区块创建的时间
	PrevBlockHash string //上一个区块的哈希值
	Hash          string //当前区块哈希值
	Data          string // 区块数据

}

//计算区块哈希值 根据上一个区块的hash和当前区块的数据生成当前区块的hash
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

//生成新的区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock

}

//生成创世区块
func GenerateGenesisBlock() Block {

	//虚拟一个区块 index为-1 hash为 空
	//当生成新的区块的时候 新区块 index 为 0 prevBlockHash 为 空
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""

	return GenerateNewBlock(preBlock, "Genesis Block")

}
