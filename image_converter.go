package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"image"
	
)



func Convert_image(img image.Image) {
	converter := convert.NewImageConverter()

	ascii := converter.Image2ASCIIString(img, &convert.Options{   
		FitScreen:			false,
		StretchedScreen:	false,           
		Colored:			true,
		FixedHeight:		50,
		FixedWidth:			100,
		Reversed:			false,
	
	})
	fmt.Println(ascii)
	
	
}
	



	