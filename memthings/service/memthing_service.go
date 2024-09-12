package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Memthing contains the memory item content
type Memthing struct {
	Content string `json:"content"`
}

// ClassificationResult holds the result from TextCNN
type ClassificationResult struct {
	Classification string  `json:"classification"`
	Grade          float64 `json:"grade"`
}

// MemthingService 结构体，封装对 memthing 的操作
type MemthingService struct {
	// 可以根据需要添加依赖（如数据库、外部服务客户端等）
}

// NewMemthingService 构造函数，返回一个 MemthingService 实例
func NewMemthingService() *MemthingService {
	return &MemthingService{}
}

// ClassifyMemthing sends the memthing content to the TextCNN service and returns classification result
func (s *MemthingService) ClassifyMemthing(memthingContent string) (*ClassificationResult, error) {
	memthing := Memthing{Content: memthingContent}
	jsonData, err := json.Marshal(memthing)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal memthing content: %v", err)
	}

	resp, err := http.Post("http://localhost:8000/classify", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to send classification request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("classification service returned status code %d", resp.StatusCode)
	}

	var result ClassificationResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode classification response: %v", err)
	}

	return &result, nil
}
