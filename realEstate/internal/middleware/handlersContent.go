package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"realEstate/internal/db"
	"realEstate/internal/db/redis"
	"realEstate/internal/models"
	"strconv"
)

func GetContent(c *gin.Context) {
	token := c.Query("Token")
	IdContent := c.Query("IdContent")
	IdContent2, err := strconv.Atoi(IdContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not convert string id content to int",
		})
		return
	}
	id := redis.InitRedis().Get(ctx, token)
	id2 := id.Val()
	id3, err2 := strconv.Atoi(id2)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The token does not exist",
		})
		return
	}
	//if id3 == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "Token not found",
	//	})
	//}
	var user models.User
	row := db.InitDB().QueryRow(`SELECT "Id_user","Name", "Role" FROM public."Users" where 
    "Id_user"=$1`, id3)
	err3 := row.Scan(&user.Id_user, &user.Name, &user.Role)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Scan not complited",
		})
		return
	}
	if user.Role != "ContentMaker" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "No authority",
		})
		return
	}

	var content models.Content
	row2 := db.InitDB().QueryRow(`SELECT "IdContent","Article","DateCreation","AuthorID", 
	"Text","MiniContent" FROM public."Content" where 
    "IdContent"=$1`, IdContent2)
	err4 := row2.Scan(&content.IdContent, &content.Article, &content.DateCreation,
		&content.AuthorID, &content.Text, &content.MiniContent)
	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Scan not complited",
		})
		return
	}
	AuthorID := content.AuthorID
	switch {
	case AuthorID != id3:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not owner of the content",
		})
	case AuthorID == id3:
		c.JSON(http.StatusOK, gin.H{
			"Content": content,
		})
	}
}
