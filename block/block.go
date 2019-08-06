package block

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index 			int64 		//区块编号 快速定位
	Timestamp 		int64  		//时间戳 区块创建的时间
	PrevBlockHash 	string		//上一个区块的哈希值
	Hash 			string 		//当前区块哈希值
	Data 			string 		// 区块数据

}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data

	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}