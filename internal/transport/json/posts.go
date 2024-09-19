package json

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func CreatePost(c *gin.Context) {
	var req config.PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.FileData == "" || req.FileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File data or name missing"})
		return
	}

	fileName := strings.TrimSuffix(req.FileName, filepath.Ext(req.FileName))
	transliteratedName := utils.Transliterate(fileName) + filepath.Ext(req.FileName)

	fileBytes, err := base64.StdEncoding.DecodeString(req.FileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode file"})
		return
	}

	filePath := filepath.Join("files", transliteratedName)
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post := config.Post{
		Body:       req.Body,
		CreatedBy:  req.CreatedBy,
		IsActual:   req.IsActual,
		CreateDate: req.CreateDate,
		Link:       "/files/" + transliteratedName,
	}

	id, err := postgresql.CreatePost(post)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"postID":  id,
	})

}

func GetPosts(c *gin.Context) {

	posts, err := postgresql.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, posts)
	}

}

func UpdatePost(c *gin.Context) {

	var post config.Post

	err := c.BindJSON(&post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}

	err = postgresql.UpdatePost(post)

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
