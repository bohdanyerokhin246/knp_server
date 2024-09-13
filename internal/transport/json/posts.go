package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post config.Post

	e := c.BindJSON(&post)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": e.Error(),
		})
		return
	}

	postID, err := postgresql.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"PostID": postID,
		})
	}
}

func GetPosts(c *gin.Context) {

	posts, err := postgresql.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Posts": posts,
		})
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
