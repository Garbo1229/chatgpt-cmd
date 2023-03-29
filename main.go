package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 提出的问题
	var content string

	//注册参数
	flag.StringVar(&content, "m", "", "提出的问题,默认为空")

	flag.Parse()

	if content == "" {
		fmt.Printf("请加入 -m 参数提问")
		return
	}
	// 替换为您的API密钥
	apiKey := ""

	// ChatGPT API的URL
	apiURL := "https://api.openai.com/v1/chat/completions"

	// 构造请求的数据
	requestData := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    [...]interface{}{map[string]string{"role": "user", "content": content}},
		"temperature": 0.7,
	}

	// 将请求的数据编码为JSON格式
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 创建POST请求
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 添加请求头信息
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// 读取响应的数据
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 解码响应的数据
	var responseData map[string]interface{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// 处理错误
	if errorData, ok := responseData["error"]; ok {
		if errorMessgae, ok := errorData.(map[string]interface{})["message"]; ok {
			fmt.Printf("[Error Message]: %s \n", errorMessgae)
			return
		}

	}
	// 解析响应
	if choices, ok := responseData["choices"]; ok {
		if message, ok := choices.([]interface{}); ok {
			if content, ok := message[0].(map[string]interface{}); ok {
				fmt.Printf("[ChatGPT]: %v \n", content["message"].(map[string]interface{})["content"])
				return
			}
		}
	}
	fmt.Printf("[System Error]: ResponData  --- %v --- \n", requestData)
}
