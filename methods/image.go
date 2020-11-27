package main

import ("golang.org/x/tour/pic"
"image"
"image/color")

type Image struct{
i int
j int
}

//ColorModel() color.Model
func (im Image) ColorModel() color.Model{
return color.RGBAModel
}
    //Bounds() Rectangle
func (im Image) Bounds() image.Rectangle{
return image.Rect(0,0,im.i,im.j)
}

    //At(x, y int) color.Color
	
func (im Image) At(x,y int) color.Color{

v:= uint8(x^2*y^2) 
return color.RGBA{v, v, 255, 255}
}


func main() {
	m := Image{200,200}
	pic.ShowImage(m)
}
