package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age int
	Sex string
	IsMan bool
}

func main()  {
	// struct转json
	hxh := person{"hxh", 21, "男", true}
	bytes, err := json.Marshal(hxh)
	if err == nil{
		fmt.Println("转换成功")
		fmt.Println(string(bytes))
	}else {
		fmt.Println("转换为json失败，原因：", err)
	}


	// map转json
	personMap := make(map[string]interface{})
	personMap["name"]  = "wxz"
	personMap["age"] = 21
	personMap["sex"] = "女"
	personMap["is_girl"] = true
	bytes2, err2 := json.Marshal(personMap)
	if err2 == nil{
		fmt.Println("转换成功")
		fmt.Println(string(bytes2))
	}else {
		fmt.Println("转换为json失败，原因：", err2)
	}

	// map切片转json
	dataSlice := make([]map[string]interface{}, 0)
	wxz1 := make(map[string]interface{})
	wxz1["name"]  = "wxz"
	wxz1["age"] = 21
	wxz1["sex"] = "女"
	wxz1["is_girl"] = true
	hxh1 := make(map[string]interface{})
	hxh1["name"]  = "hxh"
	hxh1["age"] = 21
	hxh1["sex"] = "男"
	hxh1["is_girl"] = false
	dataSlice = append(dataSlice, hxh1, wxz1)
	bytes3, err3 := json.Marshal(dataSlice)
	if err3 == nil{
		fmt.Println("转换成功")
		fmt.Println(string(bytes3))
	}else {
		fmt.Println("转换为json失败，原因：", err3)
	}

	// 反了反了，json转map了
	jsonStr := `{"age":21,"is_girl":true,"name":"wxz","sex":"女"}`
	jsonBytes := []byte(jsonStr)	// 反序列化需要bytes
	wxzMap := make(map[string]interface{})	// 准备好map来装
	err4 := json.Unmarshal(jsonBytes, &wxzMap)
	if err4 == nil{
		fmt.Println("成功转为map")
		fmt.Println(wxzMap)
	}

	// jsom转struct
	jsonBytes2 := []byte(`{"Name":"hxh","Age":21,"Sex":"男","IsMan":true}`)
	hxhStruct := new(person)
	err5 := json.Unmarshal(jsonBytes2, &hxhStruct)
	if err5 == nil{
		fmt.Println("成功转为struct")
		fmt.Println(*hxhStruct)
	}

	// json转map切片
	jsonBytes3 := []byte(`[{"age":21,"is_girl":false,"name":"hxh","sex":"男"},{"age":21,"is_girl":true,"name":"wxz","sex":"女"}]`)
	personSlice := make([]map[string]interface{}, 0)
	err6 := json.Unmarshal(jsonBytes3, &personSlice)
	if err6 == nil{
		fmt.Println("成功转为map slice")
		fmt.Println(personSlice)
	}else {
		fmt.Println("失败，原因：", err6)
	}

	// json转struct切片
	jsonBytes4 := []byte(`[{"age":21,"is_girl":false,"name":"hxh","sex":"男"},{"age":21,"is_girl":true,"name":"wxz","sex":"女"}]`)
	structSlice := make([]person, 0)
	err7 := json.Unmarshal(jsonBytes4, &structSlice)
	if err7 == nil{
		fmt.Println("成功转为struct slice")
		fmt.Println(structSlice)
	}else {
		fmt.Println("失败，原因：", err7)
	}

	// 开始写入
	dstFile,_ := os.OpenFile("/usr/local/var/go/src/go-learn/jsonOperate/jsonOperate.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	encode := json.NewEncoder(dstFile)
	defer func() {
		_ = dstFile.Close()
	}()
	err8 := encode.Encode(dataSlice) 	// encode.Encode(v interface{}) 里面可以传map，切片或结构体
	if err8 == nil{
		fmt.Println("编码并写入成功")
	}else {
		fmt.Println("编码或写入失败，原因：", err8)
	}


	// 解码器，开始读json解码为map
	srcFile,_ := os.OpenFile("/usr/local/var/go/src/go-learn/jsonOperate/jsonOperate.json", os.O_CREATE|os.O_RDONLY, 0666)
	decoder := json.NewDecoder(srcFile)
	dataSlice2 := make([]map[string]interface{}, 0)
	defer func() {
		_ = srcFile.Close()
	}()
	_ = decoder.Decode(&dataSlice2) 	// encode.Encode(v interface{}) 里面可以传map，切片或结构体
	fmt.Println("解码成功，map为")
	fmt.Println(dataSlice2)
}
