package p2

import (
	"encoding/json"
	"fmt"
	"log"
	//"../p1"
)

/**
BlockChain
Each blockchain must contain two fields described below.Don't change the name or the data type.
(1) Chain: map[int32][]Block
This is a map which maps a block height to a list of blocks. The value is a list so that it can handle the forks.
(2) Length: int32
Length equals to the highest block height.
 */
type BlockChain struct {
	Chain map[int32][]Block
	Length int32
}

/**
Create a new block chain
Return type: BlockChain
 */
func NewBlockChain() BlockChain{
	//create a blockchain structure
	return BlockChain{make(map[int32][]Block), 0}
}


/**
Description: This function takes a height as the argument, returns the list of blocks stored in that height or None if the height doesn't exist.
Argument: int32
Return type: []Block
Return nil if the key doesn't exist.
 */
func (bc *BlockChain) Get(height int32) []Block {
	//find the corresponding list
	blockList := bc.Chain[height]
	if len(blockList) == 0 {
		return nil
	}
	return blockList
}

/**
Description: This function takes a block as the argument, use its height to find the corresponding list in blockchain's Chain map.
If the list has already contained that block's hash, ignore it because we don't store duplicate blocks;
if not, insert the block into the list.
Argument: block
 */
func (bc *BlockChain) Insert(block Block) {
	height := block.Header.Height
	//blockList := bc.Chain[height]
	blockList := bc.Get(height)
	//length=0 insert
	if blockList == nil {
		blockList = append(blockList, block)
	} else {
		for _, v := range blockList {
			//same hash
			if block.Header.Hash != v.Header.Hash {
				//insert into blockList
				blockList = append(blockList, block)
			}
		}
	}
	//update in map
	bc.Chain[height] = blockList
	//compare current block height with previous block's length
	if block.Header.Height > bc.Length {
		bc.Length = block.Header.Height
	}
}

/**
Description: This function iterates over all the blocks,
generate blocks' JsonString by the function you implemented previously,
and return the list of those JsonStrings.
Return type: string，error
 */
func (bc *BlockChain) EncodeToJson() (string, error) {

	var jsonArray []BlockJson
	//k: height,
	for _,blocklist := range bc.Chain{
		for _,block := range blocklist {
			//fmt.Println("root:", block.Value.GetRoot())
			jsonStruct := block.blockToBlockJson()
			jsonArray = append(jsonArray, jsonStruct)
		}
	}
	lang, err := json.Marshal(jsonArray)
	if err == nil {
		log.Println(err)
	}
	return string(lang), err
}

/**
Description:
This function is called upon a blockchain instance.
It takes a blockchain JSON string as input, decodes the JSON string back to a list of block JSON strings,
decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
Argument: string
Return type: BlockChain, error
 */
//This function is same as blockchain (4) DecodeFromJSON(self, jsonString)
func DecodeJsonToBlockChain(jsonString string) (BlockChain, error) {
	jsonArray := []BlockJson{}
	//fmt.Println("json:", jsonString)
	err := json.Unmarshal([]byte(jsonString), &jsonArray)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}

	bc := NewBlockChain()
	for _,blockJson := range jsonArray {
		block := blockJsonToBlock(blockJson)
		//insert block
		//fmt.Println("block:", block)
		bc.Insert(block)
	}

	return bc, err
}

