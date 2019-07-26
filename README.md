# orderMap

Golang有序的map, 类似python的 collections.OrderedDict

### 使用

```go
package main

import (
    "encoding/json"
    "github.com/luxingwen/orderMap"
)

func main() {
	oMap := orderMap.NewOrderMap()
	oMap.Set("111", 111)
	oMap.Set("a", "aahjkh")
	fmt.Println(oMap)

	// use Get instead of i, ok := oMap["a"]
	val, ok := oMap.Get("a")

	fmt.Println(oMap)

	// use o.Delete instead of delete(oMap, key)
	err := oMap.Delete("a")

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

```