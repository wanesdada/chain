package basis

import (
	"chain/block"
	"chain/merror"
	"chain/proto"
	"encoding/json"
	"fmt"
)

type BlockContrillers struct {
	BlockBaseControllers
}

// @router /getblock [post]
func (c *BlockContrillers) GetBlockchain() {

	defer func() {
		c.ServeJSON()
		c.StopRun()
	}()

	var blockchain = block.GetBlockChainByKey(c.m.CoinName)

	if blockchain == nil {
		c.Data["json"] = &proto.BaseResponseBody{Ret: 403, Msg: "未查询到CoinName", Err: "blockchain == nil"}
		return
	}

	blockchain.Print()

	c.Data["json"] = &proto.BaseResponseBody{Ret: 200, Data: blockchain}

	return
}

// @router /writeblock [post]
func (c *BlockContrillers) WriteBlockchain() {
	defer func() {
		c.ServeJSON()
		c.StopRun()
	}()

	var values proto.BusinessData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &values)

	fmt.Println("打印一下", values.Data, "打印一下::", string(c.Ctx.Input.RequestBody))

	if err != nil {
		merror.Log("URI:", c.Ctx.Input.URI(), "参数有误")
		c.Data["json"] = &proto.BaseResponseBody{Ret: 403, Msg: "参数有误,请稍后再试", Err: err.Error()}
		return
	}

	var blockchain = block.GetBlockChainByKey(c.m.CoinName)

	if blockchain == nil {
		c.Data["json"] = &proto.BaseResponseBody{Ret: 403, Msg: "未查询到CoinName", Err: "blockchain == nil"}
		return
	}

	blockchain.SendData(values.Data)
	blockchain.Print()

	c.Data["json"] = &proto.BaseResponseBody{Ret: 200, Data: block.GetBlockChainByKey(c.m.CoinName)}

	return
}

// @router /newblock [post]
func (c *BlockContrillers) NewBlockchain() {
	defer func() {
		c.ServeJSON()
		c.StopRun()
	}()

	var blockchain = block.GetBlockChainByKey(c.m.CoinName)

	if blockchain != nil {
		c.Data["json"] = &proto.BaseResponseBody{Ret: 403, Msg: "CoinName已经存在", Err: "blockchain != nil"}
		return
	}

	b := block.NewBlockchain()

	b = block.SetBlockChainByKey(c.m.CoinName, b)

	b.Print()

	c.Data["json"] = &proto.BaseResponseBody{Ret: 200, Data: b}

	return
}
