package main

import (
	"chain/block"
)

func main() {
	//新增区块链
	//blockchain := block.NewBlockchain()
	//blockchain.SendData("Send 1 BTC To Wanzi")
	//blockchain.SendData("Send 1 EOS To Wanzi")
	//blockchain.Print()

	//b := block.GetBlockChainByKey("aaa")
	//
	//fmt.Println(b)

	b := block.NewBlockchain()

	b = block.SetBlockChainByKey("wanzi", b)
	b.Print()
}
