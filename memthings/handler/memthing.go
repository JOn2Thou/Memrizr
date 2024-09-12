package handler

import (
	"github.com/JOn2Thou/memrizr/memthings/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemthingHandler struct {
	memthingService *service.MemthingService
}

type MemthingContent struct {
	Content string `json:"content"`
}

// 创建Memthing
func (h *MemthingHandler) CreateMemthing(c *gin.Context) {
	var memthingContent MemthingContent

	// 从请求中获取 memthing 内容
	if err := c.ShouldBindJSON(&memthingContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用 TextCNN 分类服务
	classificationResult, err := h.memthingService.ClassifyMemthing(memthingContent.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to classify memthing"})
		return
	}

	// 返回分类结果
	c.JSON(http.StatusOK, gin.H{
		"classification": classificationResult.Classification,
		"grade":          classificationResult.Grade,
	})
}
