package block

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

//生成一条新的区块链
func NewBlockchain() *Blockchain {
	//生成创世区块
	genesisBlock := GenerateGenesisBlock()
	//空的链
	blockchain := Blockchain{}
	//加入创世区块
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}

//数据进去区块链
func (bc *Blockchain) SendData(data string) {
	//获取当前最后一条区块
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	//生成新的区块
	newBlock := GenerateNewBlock(*preBlock, data)
	//插入区块链当中
	bc.ApendBlock(&newBlock)
}

//插入新的区块进链
func (bc *Blockchain) ApendBlock(newBlock *Block) {

	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	//判断新区块是否正确
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		//失败输出错误
		log.Fatal("invalid block")
	}

}

//校验新区块是否正确
func isValid(newBlock Block, oldBlcok Block) bool {
	//新区块index 未旧区块index++
	if newBlock.Index-1 != oldBlcok.Index {
		return false
	}
	//新区块 prehash 为旧区块hash
	if newBlock.PrevBlockHash != oldBlcok.Hash {
		return false
	}
	//校验新区块hash的有效性
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true

}

//打印数据
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Pre Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)

	}
}
