package block

var ChainMap map[string]*Blockchain

func init() {
	ChainMap = make(map[string]*Blockchain)
}

func GetBlockChainByKey(key string) *Blockchain {

	return ChainMap[key]
}

func SetBlockChainByKey(key string, b *Blockchain) *Blockchain {

	ChainMap[key] = b

	return ChainMap[key]
}
