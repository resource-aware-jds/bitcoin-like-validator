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

	round map[string]bool
}

func ProvideHandler(cfg config.Config) Handler {
	return Handler{
		cfg:   cfg,
		round: make(map[string]bool),
	}
}

func (h *Handler) SubmitSuccessTask(c *gin.Context) {
	data := c.Param("answer")

	resultBase64 := h.generateBase64Hash(data)

	if h.cfg.ExpectedHash != resultBase64 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "incorrect answer",
		})
		return
	}

	result := c.GetHeader("X-NODE-ID")
	logrus.Infof("%s found the correct answer! at %s", result, time.Now().String())

	roundID := c.GetHeader("X-ROUND-ID")
	h.round[roundID] = true

	c.JSON(http.StatusOK, gin.H{
		"message": "Correct!",
	})

}

func (h *Handler) generateBase64Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	result := hasher.Sum(nil)

	resultBase64 := base64.StdEncoding.EncodeToString(result)
	return resultBase64
}

func (h *Handler) GetTheHashBase64(c *gin.Context) {
	data := c.Param("data")

	c.JSON(http.StatusOK, gin.H{
		"message": h.generateBase64Hash(data),
	})
}

func (h *Handler) CheckRoundWinner(c *gin.Context) {
	roundID := c.GetHeader("X-ROUND-ID")

	ok := h.round[roundID]
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Already got the winner"})
		return
	}

	c.Status(http.StatusOK)
}
