package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type result struct {
	code int64
	msg  string
	data map[string]interface{}
	err  error
}

// 替换为您的API密钥
var apiKey = ""

// ChatGPT API的URL
var apiURL = "https://api.openai.com/v1/chat/completions"

func main() {

	if apiKey == "" {
		fmt.Println("请替换您的API秘钥后再执行")
		os.Exit(0)
	}

	closeHandler()

	// 消息数组
	messages := []interface{}{}

	for true {
		// 输入问题
		var input string

		fmt.Print("[Me]: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		// 消息列表
		messages = append(messages, map[string]string{"role": "user", "content": input})

		// 请求结果
		result := requestHandler(messages)

		if result.code == 200 {
			// 成功
			messages = append(messages, map[string]string{"role": "assistant", "content": result.msg})
			fmt.Printf("[ChatGPT]: %v\n", result.msg)
		} else if result.code == 500 {
			// API错误
			fmt.Printf("[ChatGPT Error]: %v\n[Response Data]: %v\n", result.msg, result.data)
			break
		} else if result.err != nil {
			// 请求过程中错误
			fmt.Printf("[System]: %v\n[Response Data]: %v\n", result.err, result.data)
			break
		}

	}
}

// 构建请求
func requestHandler(messages []interface{}) result {

	// 构造请求的数据
	requestData := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    messages,
		"temperature": 0.7,
	}

	// 将请求的数据编码为JSON格式
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return result{err: err}
	}

	// 创建POST请求
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return result{err: err}
	}

	// 添加请求头信息
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{Timeout: 5000 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return result{err: err}
	}

	// 读取响应的数据
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result{err: err}
	}

	// 解码响应的数据
	var responseData map[string]interface{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return result{err: err}
	}

	// 处理错误
	if errorData, ok := responseData["error"]; ok {
		if errorMessgae, ok := errorData.(map[string]interface{})["message"]; ok {
			return result{code: 500, msg: errorMessgae.(string), data: requestData}
		}

	}
	// 解析响应
	if choices, ok := responseData["choices"]; ok {
		if message, ok := choices.([]interface{}); ok {
			if content, ok := message[0].(map[string]interface{}); ok {
				return result{code: 200, msg: content["message"].(map[string]interface{})["content"].(string)}
			}
		}
	}
	return result{code: 500, msg: "出现未知错误", data: requestData}
}

// 处理关闭提示
func closeHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\rGood bye!~")
		os.Exit(0)
	}()
}
