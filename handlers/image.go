package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mhdiiilham/gournal/models"
	log "github.com/sirupsen/logrus"
)

type imgur struct {
	Data struct {
		ID         string `json:"id"`
		Link       string `json:"link"`
		DeleteHash string `json:"deletehash"`
	}
}

type responImg struct {
	ImgID uint64 `json:"image_id"`
	Link  string `json:"image_link"`
}

// UploadImage to imgur
// Response with it url
// and delete hash
func UploadImage(c *gin.Context) {
	var res imgur
	var buf = new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	formFile, err := c.FormFile("image")
	if err != nil {
		log.Warn("Error when choosing image!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error!",
			"data":    "-",
		})
		return
	}

	imageToUpload, err := writer.CreateFormFile("image", formFile.Filename)
	if err != nil {
		log.Warn("Error CreateFormFile!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error!",
			"data":    "-",
		})
		return
	}

	imageFile, _ := formFile.Open()
	io.Copy(imageToUpload, imageFile)
	writer.Close()

	req, err := http.NewRequest("POST", os.Getenv("IMGUR_URI"), buf)
	if err != nil {
		log.Warn("Error Requesting to imgur!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error!",
			"data":    "-",
		})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+os.Getenv("IMGUR_CLIENT_ID"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn("Error on send request!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error!",
			"data":    "-",
		})
		return
	}
	imgRes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(imgRes), &res)

	image := models.Image{
		ImgurID:    res.Data.ID,
		Link:       res.Data.Link,
		DeleteHash: res.Data.DeleteHash,
	}
	image.Save()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "image uploaded!",
		"data": responImg{
			ImgID: image.ID,
			Link:  image.Link,
		},
	})

}
