package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("reset", func(ctx *gin.Context) {
		if err := GenerateJetCA(); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err})
			return
		}
		if err := GeneratePowerResult(); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": "gogogo"})
	})
	r.POST("/generate", func(ctx *gin.Context) {
		body, _ := ctx.GetRawData()
		ctx.JSON(http.StatusOK, gin.H{"license": GenerateLicense(string(body))})
	})
	r.Run("0.0.0.0:2333")
}
