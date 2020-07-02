package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/models"
	"net/http"
)

// PostJournal ...
func PostJournal(c *gin.Context) {
	var req models.JournalInput
	var image models.Image

	c.ShouldBindJSON(&req)
	image.Find(req.ImgurID)
	journal := models.Journal{
		Title: req.Title,
		Description: req.Description,
		Image: image,
	}
	journal.Save()

	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"message": "journal posted!",
		"data": journal,
	})
}

// GetJournals ...
// return list of journal
func GetJournals(c *gin.Context) {
	var journals models.ListJournal
	journals.Find()

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Success fetching list of journal",
		"data": journals,
	})
}
