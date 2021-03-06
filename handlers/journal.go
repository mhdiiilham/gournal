package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/models"
)

// PostJournal ...
func PostJournal(c *gin.Context) {
	var req models.JournalInput
	var image models.Image

	c.ShouldBindJSON(&req)
	image.Find(req.ImgurID)
	journal := models.Journal{
		Title:       req.Title,
		Description: req.Description,
		Image:       image,
	}
	journal.Save()

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "journal posted!",
		"data":    journal,
	})
}

// GetJournals ...
// return list of journal
func GetJournals(c *gin.Context) {
	var journals models.ListJournal
	cp := c.Query("pages")
	field := c.Query("field")
	value := c.Query("value")
	if field == "" && value == "" {
		if cp == "" {
			journals.Find("1")
		} else {
			journals.Find(cp)
		}
	} else {
		if cp == "" {
			journals.Find("1", field, value)
		} else {
			journals.Find(cp, field, value)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success fetching list of journal",
		"data":    journals,
	})
}

// GetOneJournal ...
func GetOneJournal(c *gin.Context) {
	var journal models.Journal
	journal.First(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf(`Success fecthing journal "%v"`, journal.Title),
		"data":    journal,
	})	
}

// UpdateJournal ...
func UpdateJournal(c *gin.Context) {
	var journal models.Journal
	journal.First(c.Param("id"))

	if journal.ID < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"message": "Journal Not Found!",
			"data": journal,
		})
		return
	}

	
}
