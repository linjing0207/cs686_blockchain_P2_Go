package tests

import (
	"../p1"
	"../p2"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type BlockJson struct {
	Height     int32             `json:"height"`
	Timestamp  int64             `json:"timeStamp"`
	Hash       string            `json:"hash"`
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	MPT        map[string]string `json:"mpt"`
}

func TestBlockChainBasic(t *testing.T) {
	jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	bc, err := p2.DecodeJsonToBlockChain(jsonBlockChain)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	jsonNew, err := bc.EncodeToJson()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	var realValue []BlockJson
	var expectedValue []BlockJson
	err = json.Unmarshal([]byte(jsonNew), &realValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = json.Unmarshal([]byte(jsonBlockChain), &expectedValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if !reflect.DeepEqual(realValue, expectedValue) {
		fmt.Println("=========Real=========")
		fmt.Println(realValue)
		fmt.Println("=========Expcected=========")
		fmt.Println(expectedValue)
		t.Fail()
	}
}
/**
b1:

hash=3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48
timeStamp=1234567890
height=1
parentHash=genesis
size=589
JSON string='{"hash": "3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48", "timeStamp": 1234567890, "height": 1, "parentHash": "genesis", "size": 1174, "mpt": {"hello": "world", "charles": "ge"}}'

b2:
hash=24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159
timeStamp=1234567890
height=2
parentHash=e01a76a2da26aa7f64e9423937d4512785e318d63552dae467c6966a43d953a6
size=647
JSON string='{"hash": "24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159", "timeStamp": 1234567890, "height": 2, "parentHash": "3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48", "size": 1231, "mpt": {"hello": "world", "charles": "ge"}}'


blockchain:
length=2
JSON string='[{"hash": "3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48", "timeStamp": 1234567890, "height": 1, "parentHash": "genesis", "size": 1174, "mpt": {"hello": "world", "charles": "ge"}}, {"hash": "24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159", "timeStamp": 1234567890, "height": 2, "parentHash": "3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48", "size": 1231, "mpt": {"hello": "world", "charles": "ge"}}]'
 */
func Test(t *testing.T){
	fmt.Println("Example:")
	mpt := p1.MerklePatriciaTrie{}
	mpt.Initial()
	mpt.Insert("hello", "world")
	mpt.Insert("charles", "ge")
	b1 := p2.NewBlock(1, 1234567890, "genesis", mpt)
	b2 := p2.NewBlock(2, 1234567890, b1.Header.Hash, mpt)
	bc := p2.NewBlockChain()
	bc.Insert(b1)
	bc.Insert(b2)

	b1Json := `{"hash":"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48","timeStamp":1234567890,"height":1,"parentHash":"genesis","size":1174,"mpt":{"hello":"world","charles":"ge"}}`
	b2Json := `{"hash":"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159","timeStamp":1234567890,"height":2,"parentHash":"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48","size":1231,"mpt":{"hello":"world","charles":"ge"}}`
	bcJson := `[{"hash":"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48","timeStamp":1234567890,"height":1,"parentHash":"genesis","size":1174,"mpt":{"hello":"world","charles":"ge"}},{"hash":"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159","timeStamp":1234567890,"height":2,"parentHash":"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48","size":1231,"mpt":{"hello":"world","charles":"ge"}}]`

	fmt.Println("b1 json:")
	fmt.Println(b1.EncodeToJson())
	fmt.Println(b1Json)

	bb1,_ := p2.DecodeFromJson(b1Json)
	fmt.Println(bb1.EncodeToJson())

	fmt.Println("b2 json:")
	fmt.Println(b2.EncodeToJson())
	fmt.Println(b2Json)
	bb2,_ := p2.DecodeFromJson(b2Json)
	fmt.Println(bb2.EncodeToJson())

	fmt.Println("bc json:")
	fmt.Println(bc.EncodeToJson())
	fmt.Println(bcJson)
	bbc, _ := p2.DecodeJsonToBlockChain(bcJson)
	fmt.Println(bbc.EncodeToJson())
}