package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"image"
)



func Convert_image(img image.Image) {
	converter := convert.NewImageConverter()

	ascii := converter.Image2ASCIIString(img, &convert.Options{
		Ratio: 1.0,          
		FixedWidth:  50,   
		FixedHeight: 25,   
		FitScreen:    true,  
	}) 
	fmt.Println(ascii)
}
