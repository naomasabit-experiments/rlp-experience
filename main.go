package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

var case1Inputs = []uint32{1, 2, 127}
var case2Inputs = []string{"dog", "doge", "cat", "0123456789012345678901234567890123456789012345678901234"}
var case3Inputs = []string{"01234567890123456789012345678901234567890123456789012345"}
var case4Inputs = []uint{1,2,3}
var case5Inputs = []uint{1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6,7,8,9,10,1,2,3,4,5,6}

func rlpEnc(val interface{}) {
	encoded, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	fmt.Println("input:", val, "=> encoded:", hex.EncodeToString(encoded))

}

func main() {
	fmt.Println("case1")
	for _, input := range case1Inputs {
		rlpEnc(input)
	}

	fmt.Println("case2")
	for _, input := range case2Inputs {
		rlpEnc(input)
	}

	fmt.Println("case3")
	for _, input := range case3Inputs {
		rlpEnc(input)
	}

	fmt.Println("case4")
	rlpEnc(case4Inputs)

	fmt.Println("case5")
	rlpEnc(case5Inputs)
}
