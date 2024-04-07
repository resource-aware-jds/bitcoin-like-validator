package handler

import (
	"bitcoin-like-validator/config"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	cfg config.Config
}

func ProvideHandler(cfg config.Config) Handler {
	return Handler{
		cfg: cfg,
	}
}

type SubmitAnswerReq struct {
	Answer string `json:"answer"`
}

func (h *Handler) SubmitSuccessTask(c *gin.Context) {
	var data SubmitAnswerReq
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"name": "Bind Err",
			"err":  err,
		})
	}

	hasher := sha256.New()
	hasher.Write([]byte(data.Answer))

	result := hasher.Sum(nil)

	resultBase64 := base64.StdEncoding.EncodeToString(result)

	if h.cfg.ExpectedHash != resultBase64 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "incorrect answer",
		})
		return
	}

	logrus.Info("Someone found the correct answer! at ", time.Now().String())

	c.JSON(http.StatusOK, gin.H{
		"message": "Correct!",
	})

}
