package news

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CreateNewsHandler(c *gin.Context) {
	var news models.News

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if news.FileInfo.FileData == "" || news.FileInfo.FileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File data or name missing"})
		return
	}

	fileName := strings.TrimSuffix(news.FileInfo.FileName, filepath.Ext(news.FileInfo.FileName))
	transliteratedName := utils.Transliterate(fileName) + filepath.Ext(news.FileInfo.FileName)

	fileBytes, err := base64.StdEncoding.DecodeString(news.FileInfo.FileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode file"})
		return
	}

	filePath := filepath.Join("storage", transliteratedName)
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post := models.News{
		Body:      news.Body,
		CreatedBy: news.CreatedBy,
		IsActual:  news.IsActual,
		Link:      "/storage/" + transliteratedName,
	}

	id, err := queries.CreatePost(post)

	c.JSON(http.StatusOK, gin.H{
		"message": "News created successfully",
		"postID":  id,
	})

}

func GetNewsListHandler(c *gin.Context) {

	posts, err := queries.GetNews()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, posts)
	}

}

func UpdatePost(c *gin.Context) {

	var post models.News

	err := c.BindJSON(&post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}

	err = queries.UpdatePost(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"IsUpdated": "Successfully",
		})
	}
}

func DeleteNewsHandler(c *gin.Context) {

	var post models.News

	err := c.BindJSON(&post)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error binding JSON. Error: %v", err.Error()))
	}

	err = queries.UpdatePost(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"IsUpdated": "Successfully",
		})
	}
}

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to receive file"})
		return
	}

	dst := filepath.Join("./storage", file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "path": dst})
}
