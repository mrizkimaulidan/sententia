package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

var (
	// path to original image file
	path string

	// location converted image to a new location
	location string
)

func main() {
	// parsing flags
	flag.StringVar(&path, "path", "", "-path=path/to/image path to original image")
	flag.StringVar(&location, "location", "", "-location=path/to/new-image path to a new grayscale image")
	flag.Parse()

	f, err := open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	grayscaleImage, result, err := createGrayscaleImage(f, location)
	if err != nil {
		log.Fatalln(err)
	}
	defer grayscaleImage.Close()

	err = saveGrayscaleImage(grayscaleImage, result)
	if err != nil {
		log.Fatalln(err)
	}

	c := setImagePath(location)
	defer c.Close()

	// Extract text on images
	text, err := c.Text()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Result :", text)
}

// Open image file based on path in flag
func open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Converting original image file to grayscale image
// will be created into location path on flag
func createGrayscaleImage(f *os.File, location string) (*os.File, *image.Gray, error) {
	err := createLocationPath(location)
	if err != nil {
		return nil, nil, err
	}

	grayscaleImage, err := os.Create(location)
	if err != nil {
		return nil, nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, nil, err
	}

	result := image.NewGray(img.Bounds())
	draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)

	return grayscaleImage, result, nil
}

// Save the image that has been converted to grayscale
func saveGrayscaleImage(createdFile io.Writer, m image.Image) error {
	err := jpeg.Encode(createdFile, m, nil)
	if err != nil {
		return err
	}

	return nil
}

// Create location path based on flag provided
func createLocationPath(location string) error {
	split := strings.Split(location, "/")
	folderPath := split[0 : len(split)-1]

	path := filepath.Join(folderPath...)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Set the image path to be processed on OCR
func setImagePath(imagePath string) *gosseract.Client {
	c := gosseract.NewClient()

	c.SetImage(imagePath)

	return c
}
