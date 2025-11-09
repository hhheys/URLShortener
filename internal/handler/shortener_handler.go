package handler

import (
	"URLShortener/internal/dto"
	"URLShortener/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	s *service.ShortenerService
}

func NewShortenerHandler(s *service.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{s: s}
}

func (s *ShortenerHandler) MakeShortURL(c *gin.Context) {
	var linkDto dto.LinkDto
	err := c.ShouldBindJSON(&linkDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing validation"})
		return
	}

	if !strings.HasPrefix(linkDto.Link, "http://") || !strings.HasPrefix(linkDto.Link, "https://") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect url. use absolute path"})
		return
	}

	linkEntity, err := s.s.CreateShortLink(linkDto.Link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldnt create short link"})
		return
	}
	c.JSON(http.StatusOK, linkEntity)
}

func (h *ShortenerHandler) UseShortURL(c *gin.Context) {
	shortUrl := c.Param("url")
	link, err := h.s.GetLinkFromShort(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.Redirect(http.StatusMovedPermanently, link.Link)
	c.Abort()
	return
}
