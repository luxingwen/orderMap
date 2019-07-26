package orderMap

import (
	"encoding/json"
	"strings"
)

type OrderMap struct {
	keys   []string
	values map[string]interface{}
}

func NewOrderMap() *OrderMap {
	return &OrderMap{keys: make([]string, 0), values: make(map[string]interface{}, 0)}
}

func (o *OrderMap) Set(key string, value interface{}) {
	_, ok := o.values[key]
	if !ok {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *OrderMap) Get(key string) (r interface{}, ok bool) {
	r, ok = o.values[key]
	return
}

func (o *OrderMap) Delete(key string) {
	delete(o.values, key)
	keys := make([]string, 0)
	for _, k := range o.keys {
		if k == key {
			continue
		}
		keys = append(keys, k)
	}
	o.keys = keys
}

func (o *OrderMap) Keys() []string {
	return o.keys
}

func (o *OrderMap) ValueKeys() (r []string) {
	for k, _ := range o.values {
		r = append(r, k)
	}
	return
}

func (o *OrderMap) SetKeys(keys []string) {
	o.keys = keys
}

func (o *OrderMap) UpdateKeys(keys []string) {
	mkeys := make(map[string]string, 0)
	for _, v := range keys {
		mkeys[v] = v
	}
	for _, key := range o.keys {
		if _, ok := mkeys[key]; !ok {
			delete(o.values, key)
		}
	}
	o.SetKeys(keys)
}

func (o *OrderMap) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		o.Set(k, v)
	}
	return nil
}

func (o *OrderMap) MarshalJSON() ([]byte, error) {
	s := "{"
	for _, k := range o.keys {
		s = s + `"` + strings.Replace(k, `"`, `\"`, -1) + `":`
		v, ok := o.values[k]
		if !ok {
			v = ""
		}
		vBytes, err := json.Marshal(v)
		if err != nil {
			return []byte{}, err
		}
		s = s + string(vBytes) + ","
	}
	if len(o.keys) > 0 {
		s = s[0 : len(s)-1]
	}
	s = s + "}"
	return []byte(s), nil
}

func (o *OrderMap) String() string {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(b)
}
