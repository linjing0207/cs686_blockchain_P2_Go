package p2

import (
	"../p1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"log"
)

/**
Each block must contain a header, and in the header there are the following fields:
(1) Height: int32
(2) Timestamp: int64
The value must be in the UNIX timestamp format such as 1550013938
(3) Hash: string.
(4) ParentHash: string
(5) Size: int32
Value: mpt MerklePatriciaTrie
Here's the summary of block structure:
The size is the length of the byte array of the block value
Block: Block{Header{Height, Timestamp, Hash, ParentHash, Size}, value}
 */
type Block struct {
	Header Header `json:"header"`
	Value p1.MerklePatriciaTrie `json:"mpt"`
}

/**
Header
 */
type Header struct {
	Height int32 `json:"height"`
	Timestamp int64 `json:"timestamp"`
	Hash string `json:"hash"`
	ParentHash string `json:"parenthash"`
	Size int32 `json:"size"`
}


/**
Description: This function takes arguments(such as height, parentHash, and value of MPT type) and forms a block.
This is a method of the block struct.
Block: Block{Header{Height, Timestamp, Hash, ParentHash, Size}, value}
 */
func (b *Block) Initial(height int32, timestamp int64, parentHash string, value p1.MerklePatriciaTrie)  {

	size := len(value.MptToByteArray())
	hash := b.hash_block()
	b.Header = Header{height, timestamp, hash, parentHash, int32(size)}
	b.Value = value
}

/**
Create a new block
 */
func NewBlock(height int32, timestamp int64, parentHash string, value p1.MerklePatriciaTrie) Block {
	//create a block structure, value is null
	block := Block{}
	block.Initial(height, timestamp, parentHash, value)
	return block
}


/**
Description: This function takes a string that represents the JSON value of a block as an input, and decodes the input string back to a block instance.
Argument: a string of JSON format
Return value: a block instance
 */
func DecodeFromJson(jsonString string) (Block, error) {
	//fmt.Println(jsonString)
	var blockMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &blockMap)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}
	mpt := p1.MerklePatriciaTrie{}
	mpt.Initial()

	mptMap :=  blockMap["mpt"].(map[string]interface{})
	for k,v := range mptMap{
		//fmt.Println(k,v.(string))
		mpt.Insert(k, v.(string))
	}
	height := int32(blockMap["height"].(float64))
	timeStamp := int64(blockMap["timeStamp"].(float64))
	hash := blockMap["hash"].(string)
	parentHash := blockMap["parentHash"].(string)
	size := int32(blockMap["size"].(float64))

	header := Header{height, timeStamp, hash, parentHash, size}
	block := Block{header, mpt}

	return block, err
}

/**
Description: This function encodes a block instance into a JSON format string.
Argument: a block or you may define this as a method of the block struct
Return value: a string of JSON format
Example
{
    "hash":"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48",
    "timeStamp":1234567890,
    "height":1,
    "parentHash":"genesis",
    "size":1174,
    "mpt":{
        "charles":"ge",
        "hello":"world"
    }
}
*/
func (b *Block) EncodeToJson() (string ,error) {
	var jsonBlock string
	jsonStruct := b.buildJsonStruct()

	buffer, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(string(buffer[:]))
	jsonBlock = string(buffer[:])

	return jsonBlock, err
}

func (b *Block) buildJsonStruct() map[string]interface{} {
	var jsonStruct = make(map[string]interface{})

	//traverse
	//map{"charles":"ge","hello":"world"}
	mptMap := b.Value.GetMpt(b.Value.GetRoot(), []uint8{})

	jsonStruct["hash"] = b.Header.Hash
	jsonStruct["timeStamp"] = b.Header.Timestamp
	jsonStruct["height"] = b.Header.Height
	jsonStruct["parentHash"] = b.Header.ParentHash
	jsonStruct["size"] = b.Header.Size
	jsonStruct["mpt"] = mptMap

	return jsonStruct
}

/**
Block’s hash is the SHA3-256 encoded value of this string(note that you have to follow this specific order):
hash_str := string(b.Header.Height) + string(b.Header.Timestamp) + b.Header.ParentHash + b.Value.Root + string(b.Header.Size)
 */
func (b *Block) hash_block() string {
	hash_str := string(b.Header.Height) + string(b.Header.Timestamp) + b.Header.ParentHash + b.Value.GetRoot() + string(b.Header.Size)
	sum := sha3.Sum256([]byte(hash_str))

	return hex.EncodeToString(sum[:])
}


