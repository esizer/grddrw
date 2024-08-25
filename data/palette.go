package data

import (
	"errors"
	"slices"
)

type Palette []string

var (
	HexRequired         = errors.New("hex code is required")
	HexInvalidLength    = errors.New("hex must be of length 3 or 6")
	HexInvalidCharacter = errors.New("hex contains an invalid char")
	HexExists           = errors.New("hex already exists")
)

func NewPalette() *Palette {
	return &Palette{}
}

var validChars = []string{
	"a", "b", "c", "d", "e", "f",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func ValidateHex(hex string) error {
	if len(hex) <= 0 {
		return HexRequired
	}

	h := hex
	if string(hex[0]) == "#" {
		h = hex[1:]
	}

	if !(len(h) == 3 || len(h) == 6) {
		return HexInvalidLength
	}

	for _, b := range h {
		if !slices.Contains(validChars, string(b)) {
			return HexInvalidCharacter
		}
	}

	return nil
}

func (p *Palette) Add(hex string) (string, error) {
	err := ValidateHex(hex)
	if err != nil {
		return "", err
	}

	if slices.Contains(*p, hex) {
		return "", HexExists
	}

	*p = append(*p, hex)

	return hex, nil
}
