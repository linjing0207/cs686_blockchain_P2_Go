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
	//jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	jsonBlockChain := "[{\"height\":1,\"timeStamp\":1551025401,\"hash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"parentHash\":\"genesis\",\"size\":2089,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"f367b7f59c651e69be7e756298aad62fb82fddbfeda26cb06bfd8adf9c8aa094\",\"parentHash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"05ac44dd82b6cc398a5e9664add21856ae19d107d9035af5fc54c9b0ffdef336\",\"parentHash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}}]"
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

	//fmt.Println("jsonNew:", jsonNew)
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

	bb1, err := p2.DecodeFromJson(b1Json)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(bb1.EncodeToJson())

	fmt.Println("b2 json:")
	fmt.Println(b2.EncodeToJson())
	fmt.Println(b2Json)
	bb2, err := p2.DecodeFromJson(b2Json)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(bb2.EncodeToJson())

	fmt.Println("bc json:")
	fmt.Println(bc.EncodeToJson())
	fmt.Println(bcJson)
	bbc, err := p2.DecodeJsonToBlockChain(bcJson)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(bbc.EncodeToJson())
}