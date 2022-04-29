package generator

import (
	"crypto/rand"
	"io"
	"math/big"
	"strings"
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"
)

type PasswordGenerator interface {
	Generate(length, digitsCount int, withUpper, allowRepeat bool) (string, error)
}

type generator struct {
	Reader io.Reader
}

func NewGenerator() PasswordGenerator {
	return &generator{
		Reader: rand.Reader,
	}
}

func (g *generator) Generate(length, digitsCount int, withUpper, allowRepeat bool) (string, error) {
	var (
		result  string
		letters string = LowerLetters
	)

	if withUpper {
		letters += UpperLetters
	}

	for i := 0; i < length-digitsCount; i++ {
		letter, err := g.randomItem(letters)
		if err != nil {
			return "", err
		}

		if !allowRepeat && strings.Contains(result, letter) {
			i--
			continue
		}

		result, err = g.randomInsert(result, letter)
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < digitsCount; i++ {
		digit, err := g.randomItem(Digits)
		if err != nil {
			return "", err
		}

		if !allowRepeat && strings.Contains(result, digit) {
			i--
			continue
		}

		result, err = g.randomInsert(result, digit)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

func (g *generator) randomItem(source string) (string, error) {
	n, err := rand.Int(g.Reader, big.NewInt(int64(len(source))))
	if err != nil {
		return "", err
	}

	return string(source[n.Int64()]), nil
}

func (g *generator) randomInsert(result, char string) (string, error) {
	if result == "" {
		return char, nil
	}

	n, err := rand.Int(g.Reader, big.NewInt(int64(len(result))))
	if err != nil {
		return "", err
	}

	i := n.Int64()
	return result[:i] + char + result[i:], nil
}
