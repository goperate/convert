package main

import (
	"encoding/json"
	"fmt"
	"github.com/goperate/convert/core/types"
	"log"
)

type ArrayTest struct {
	ArrayNumber1  types.ArrayInt     `json:"array_number_1"`
	ArrayNumber2  types.ArrayInt32   `json:"array_number_2"`
	ArrayNumber3  types.ArrayInt64   `json:"array_number_3"`
	ArrayNumber4  types.ArrayUint    `json:"array_number_4"`
	ArrayNumber5  types.ArrayUint32  `json:"array_number_5"`
	ArrayNumber6  types.ArrayUint64  `json:"array_number_6"`
	ArrayNumber7  types.ArrayFloat32 `json:"array_number_7"`
	ArrayNumber8  types.ArrayFloat64 `json:"array_number_8"`
	ArrayNumber9  types.ArrayInt     `json:"array_number_9"`
	ArrayNumber10 types.ArrayInt     `json:"array_number_10"`

	ArrayString1  types.ArrayKeyword `json:"array_string_1"`
	ArrayString2  types.ArrayKeyword `json:"array_string_2"`
	ArrayString3  types.ArrayKeyword `json:"array_string_3"`
	ArrayString4  types.ArrayKeyword `json:"array_string_4"`
	ArrayString5  types.ArrayKeyword `json:"array_string_5"`
	ArrayString6  types.ArrayKeyword `json:"array_string_6"`
	ArrayString7  types.ArrayKeyword `json:"array_string_7"`
	ArrayString8  types.ArrayKeyword `json:"array_string_8"`
	ArrayString9  types.ArrayKeyword `json:"array_string_9"`
	ArrayString10 types.ArrayKeyword `json:"array_string_10"`

	ArrayString11 types.ArrayString `json:"array_string_11"`
	ArrayString12 types.ArrayString `json:"array_string_12"`
}

const testJsonString = `{
	"array_number_1": 1,
	"array_number_2": [2, 3],
	"array_number_3": "4",
	"array_number_4": ["5", "6"],
	"array_number_5": ["7", 8],
	"array_number_6": "[9,10]",
	"array_number_7": "11,12 ,13",
	"array_number_8": ["14.23, 15.1, 16, 17", "18, 19, 20"],
	"array_number_9": [[21, 22, 23], [24, 25, 26]],
	"array_number_10": {"xx": 27, "cc": {"x": [28, 29]}},

	"array_string_1": "1zcxzv",
	"array_string_2": [2, 3],
	"array_string_3": "4",
	"array_string_4": ["5", "xxx6"],
	"array_string_5": ["7zxx", 8],
	"array_string_6": "[9,10]",
	"array_string_7": "11,12 ,13",
	"array_string_8": ["14.23, 15.1, 16, 17", "18, 19, 20"],
	"array_string_9": [[21, 22, 23], [24, 25, 26]],
	"array_string_10": {"xx": 27, "cc": {"x": [28, 29]}},

	"array_string_11": "最艰难的[asdk\"cee",
	"array_string_12": ["xsssss", "埃松加大\"家乘客都"]
}`

func main() {
	var data ArrayTest
	err := json.Unmarshal([]byte(testJsonString), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", data)
}
