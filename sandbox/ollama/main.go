package main

/* 流式处理 ollma 生成的文本
api refer: https://github.com/ollama/ollama
*/

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ModelQwen25  = "qwen2.5"
	ModelLlama32 = "llama3.2"
)

func main() {
	// 请求参数
	requestBody := map[string]interface{}{
		"model":  "qwen2.5",
		"prompt": "请问Qwen模型可以做什么？",
		"format": "json",
		"options": map[string]interface{}{
			"temperature": 0.7,
			"seed":        42,
		},
		"stream":     true,
		"keep_alive": "5m",
	}

	// 将请求数据编码为 JSON 格式
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// 创建 HTTP 请求
	url := "http://localhost:11434/api/generate"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %v", resp.Status)
	}

	// 读取响应流
	scanner := bufio.NewScanner(resp.Body)
	fmt.Println("Streaming response:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading stream: %v", err)
	}
}
