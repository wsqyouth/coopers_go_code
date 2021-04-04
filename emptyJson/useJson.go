package main

import (
	"encoding/json"
	"fmt"
)

type CrtInfo struct {
	TagType *uint32 `json:"use_button,omitempty"`
}

// note: 测试marshal对0默认值的线上问题
func main() {
	var crtInfo CrtInfo
	defaultUint32 := new(uint32)
	crtInfo.TagType = defaultUint32
	tagType := uint32(1)
	crtInfo.TagType = &tagType
	b, e := json.Marshal(crtInfo)
	if e != nil {
		fmt.Println("marshal error")
	}
	fmt.Println(string(b))
}
