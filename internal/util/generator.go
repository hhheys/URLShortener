package util

import (
	"math/rand"
	"strings"
)

type Generator struct {
	characters []rune
}

func NewGenerator() *Generator {
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrtuvw" +
		"1234567890")
	return &Generator{characters: characters}
}

func (g *Generator) GenerateShortUrl(size int) string {
	builder := strings.Builder{}
	for i := 0; i < size; i++ {
		builder.WriteRune(g.characters[rand.Intn(len(g.characters))])
	}
	return builder.String()
}
