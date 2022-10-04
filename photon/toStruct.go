package photon

import (
	"encoding/json"
	"fmt"
)

func ToStruct[C any](data interface{}, container *C) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		// do error check
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(jsonbody, container); err != nil {
		// do error check
		fmt.Println(err)
		return
	}
}
