package termui

import (
	"encoding/json"
	"errors"
	"image"
	"io"
	"log"
	"os"
)

type Sprite struct {
	Points []image.Point `json:"points"`
}

func NewSprite() *Sprite {
	return &Sprite{
		Points: make([]image.Point, 0),
	}
}

func SaveSpriteToFile(fileName string, sprite Sprite) {
	file, err := os.OpenFile(fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Could not open file '%s' for reading\n", fileName)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(sprite)
	if err != nil {
		log.Println(err)
		return
	}
}

func LoadSpriteFromFile(fileName string) Sprite {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	sprite := Sprite{}
	err = decoder.Decode(&sprite)

	if errors.Is(err, io.EOF) {
		return sprite
	}

	if err != nil {
		log.Fatalf("Failed to load json from file. Error is '%v'", err)
	}

	return sprite
}
