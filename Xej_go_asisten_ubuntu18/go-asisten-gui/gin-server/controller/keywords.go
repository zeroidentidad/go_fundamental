package controller

import (
	"net/http"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// KeyWordsAllHTML func
func KeyWordsAllHTML(ctx *gin.Context) {
	var keywords []models.KeyWord
	Db.Find(&keywords)
	for i, j := 0, len(keywords)-1; i < j; i, j = i+1, j-1 {
		keywords[i], keywords[j] = keywords[j], keywords[i]
	}
	ctx.HTML(200, "keywords_all.html", gin.H{
		"KeyWords": keywords,
	})
}

// KeyWordsDeleteAPI func
func KeyWordsDeleteAPI(ctx *gin.Context) {
	id := ctx.Param("id")
	Db.Delete(&models.KeywordGroup{}, "key_word_id = ?", id)
	Db.Delete(&models.KeyWord{}, id)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/keywords")
}

// KeyWordsAddAPI func
func KeyWordsAddAPI(ctx *gin.Context) {
	key, _ := ctx.GetQuery("Key")
	if key == "" {
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/keywords")
		return
	}
	var keyword = new(models.KeyWord)
	keyword.Key = key
	Db.Create(keyword)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/keywords")
}
