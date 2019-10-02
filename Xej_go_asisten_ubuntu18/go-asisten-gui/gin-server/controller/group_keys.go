package controller

import (
	"log"
	"net/http"
	"strconv"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// GroupKeysAllHTML func
func GroupKeysAllHTML(ctx *gin.Context) {
	// Find GroupKeys
	var groupKeys []models.GroupKey
	Db.Find(&groupKeys, "action_id = ?", ctx.Query("action_id"))
	for i, j := 0, len(groupKeys)-1; i < j; i, j = i+1, j-1 {
		groupKeys[i], groupKeys[j] = groupKeys[j], groupKeys[i]
	}
	// Add More values groupKeys
	for ind := range groupKeys {
		var kgroup []models.KeywordGroup
		Db.Preload("KeyWord").Find(&kgroup, "group_key_id = ?", groupKeys[ind].ID)
		groupKeys[ind].KeywordGroups = kgroup
	}
	// Find KeyWords
	var keywords = new([]models.KeyWord)
	Db.Find(keywords)
	ctx.HTML(200, "group_keys_all.html", gin.H{
		"Elements": groupKeys,
		"ActionID": ctx.Query("action_id"),
		"KeyWords": keywords,
	})
}

// GroupKeysAddAPI func
func GroupKeysAddAPI(ctx *gin.Context) {
	actionID := ctx.Query("ActionID")
	keyWordsIDSString := ctx.QueryArray("KeyWordsIDS")
	if len(keyWordsIDSString) == 0 {
		log.Println("ERROR: Not KeyWords")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/group_keys?action_id="+actionID)
		return
	}
	var gkey = new(models.GroupKey)
	el, err := strconv.ParseUint(actionID, 0, 64)
	if err != nil {
		log.Println("ERROR: ParseUINT ActionID")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/group_keys?action_id="+actionID)
		return
	}
	gkey.ActionID = uint(el)
	Db.Create(gkey)
	for _, val := range keyWordsIDSString {
		keyWordID, _ := strconv.ParseUint(val, 0, 64)
		Db.Create(&models.KeywordGroup{
			GroupKeyID: gkey.ID,
			KeyWordID:  uint(keyWordID),
		})
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/group_keys?action_id="+actionID)
}

// GroupKeysDeleteAPI func
func GroupKeysDeleteAPI(ctx *gin.Context) {
	actionID := ctx.Query("action_id")
	groupKeyID := ctx.Param("id")
	Db.Delete(&models.KeywordGroup{}, "group_key_id = ?", groupKeyID)
	Db.Delete(&models.GroupKey{}, groupKeyID)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/group_keys?action_id="+actionID)
}
