package p2

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	//"../p1"
)

/**
BlockChain
Each blockchain must contain two fields described below. Don't change the name or the data type.
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
Create/update a new block chain
 */
func NewBlockChain() BlockChain{
	//create a blockchain structure, value is null
	return BlockChain{make(map[int32][]Block), 0}
}


/**
Description: This function takes a height as the argument, returns the list of blocks stored in that height or None if the height doesn't exist.
Argument: int32
Return type: []Block
 */
func (bc *BlockChain) Get(height int32) []Block {
	//find the corresponding list
	blockList := bc.Chain[height]
	//?????what type should we return
	if len(blockList) == 0 {

	}
	return blockList
}

/**
Description: This function takes a block as the argument, use its height to find the corresponding list in blockchain's Chain map.
If the list has already contained that block's hash, ignore it because we don't store duplicate blocks; if not, insert the block into the list.
Argument: block
 */
func (bc *BlockChain) Insert(block Block)  {
	height := block.Header.Height
	//blockList := bc.Chain[height]
	blockList := bc.Get(height)
	//length=0 insert
	if len(blockList) == 0 {
		blockList = append(blockList, block)
	} else {
		for _, v := range blockList {
			//same hash
			if block.Header.Hash == v.Header.Hash {
				//ignore
			} else {
				//insert into blocks
				blockList = append(blockList, block)
			}
		}
	}
	//fmt.Println("blocklist:", blockList)
	//store in map
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
Return type: string
 */
//traverse
func (bc *BlockChain) EncodeToJson() (string, error) {
	//var jsonStrings []string
	var jsonStrings []map[string]interface{}
	//k: height,
	for _,blocklist := range bc.Chain{
		for _,block := range blocklist {
			jsonStruct := block.buildJsonStruct()
			jsonStrings = append(jsonStrings, jsonStruct)
		}
	}
	lang, err := json.Marshal(jsonStrings)
	if err == nil {
		log.Println(err)
	}
	//fmt.Println("", string(lang))
	return string(lang), err
}

/**
Description:
This function is called upon a blockchain instance.
It takes a blockchain JSON string as input, decodes the JSON string back to a list of block JSON strings,
decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
Argument: self, string
 */
func DecodeJsonToBlockChain(jsonString string) (BlockChain, error) {
	//a := []int{}
	jsonArray := []interface{}{}
	//fmt.Println("json:", jsonString)
	err := json.Unmarshal([]byte(jsonString), &jsonArray)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}

	//fmt.Println("list:", list)
	bc := NewBlockChain()
	for _,v := range jsonArray {
		//
		jsonStruct := v.(map[string]interface{})
		//
		jsonString, _ := json.Marshal(jsonStruct)

		//fmt.Println("hash:",blockMap["hash"])
		block,_ := DecodeFromJson(string(jsonString))
		//insert block
		bc.Insert(block)
	}

	return bc, err
}


/**
Length equals to the highest block height.

 */

func (bc *BlockChain) getHighestLength() int {

	var keys []int

	//the biggest number of keys in map
	for key := range bc.Chain {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	len := len(bc.Chain)
	return keys[len-1]
}
