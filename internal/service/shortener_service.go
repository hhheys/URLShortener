package service

import (
	"URLShortener/internal/model"
	"URLShortener/internal/repository"
	"URLShortener/internal/util"
)

type ShortenerService struct {
	r *repository.ShortenerRepository
	g *util.Generator
}

func NewShortenerService(r *repository.ShortenerRepository) *ShortenerService {
	g := util.NewGenerator()
	return &ShortenerService{r: r, g: g}
}

func (s *ShortenerService) CreateShortLink(link string) (*model.Link, error) {
	shortLink := s.g.GenerateShortUrl(7)
	linkModel, err := s.r.CreateShortLink(link, shortLink)
	if err != nil {
		return nil, err
	}
	return linkModel, nil
}

func (s *ShortenerService) GetLinkFromShort(shortValue string) (*model.Link, error) {
	link, err := s.r.GetLinkFromShort(shortValue)
	if err != nil {
		return &model.Link{}, err
	}
	return link, err
}
