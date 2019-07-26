package orderMap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOrderMap(t *testing.T) {
	oMap := NewOrderMap()
	oMap.Set("111", 111)
	oMap.Set("a", "aahjkh")

	oMap.SetKeys([]string{"time", "data", "a"})
	fmt.Println(oMap)

	oMap.UpdateKeys([]string{"time", "a", "data"})
	fmt.Println(oMap)

	data := `{"uid":"1111","time":"2019-07-26T16:32:30+08:00","level":"debug","type":10001,"method":"LoginRequest","msg":"OnMessage","data":{"name":"1102","uid":"1102"}}`
	oMap = NewOrderMap()
	json.Unmarshal([]byte(data), &oMap)
	fmt.Println(oMap)

	oMap.SetKeys([]string{"time", "level", "uid", "type", "method", "data"})

	fmt.Println(oMap)
}
