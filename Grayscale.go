/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P03
*  Title:			 Image Ascii Art           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*	
*		This Package converts an image to graycale and saves it to a new image  
*
* 
*  Usage:
*    - import to main.go file 
* 
*  Files           
*       N/A 
*****************************************************************************/
package Img_gray_scale

//necessary packages 
import (
	//fmt : allows for input and output to and from the console
	"fmt"

	//image: allows for decoding and converting of images to different formats
	"image"
	"image/color"
	"image/png"

	//os : allows to check for errors when opening a file 
	"os"
)

// Grayscale gray scales an image
func Grayscale() {
	// Open the original image
	reader, err := os.Open("downloaded_image.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image
	grayFile, err := os.Create("gray_image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer grayFile.Close()
	png.Encode(grayFile, grayImg)
	// image to indicate the new image is created 
	fmt.Println("Grayscale image saved.")

}
