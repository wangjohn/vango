package primaryColor

import (
  "fmt"
  "image"
  "image/draw"
)

type Rgb struct {
  R uint8
  G uint8
  B uint8
}

const (
  minTransparency uint8 = 125
  maxOpaqueVal uint8 = 250
  paletteSize uint8 = 10

  sigbits uint = 5
  rshift uint = 8 - sigbits
)

func PrimaryColor(img image.Image) {
  rgbaImg := image.NewRGBA(img.Bounds())
  draw.Draw(rgbaImg, rgbaImg.Bounds(), img, image.ZP, draw.Src)

  pixelArray := constructPixelArray(rgbaImg)
  quantize(pixelArray, paletteSize)
}

func quantize(pixelArray []Rgb, size uint8) ([]Rgb, error) {
  result := make([]Rgb, size)

  if size < 2 || size > 255 {
    return result, fmt.Errorf("Invalid size: cannot quantize image into %d colors", size)
  }

  histogram := constructHistogram(pixelArray)
  fmt.Println(histogram)

  # TODO: take histogram and perform median cut

  return result, nil
}

func constructHistogram(pixelArray []Rgb) []int{
  size := 1 << (3 * sigbits)
  histogram := make([]int, size)

  for _, pixel := range pixelArray {
    i := colorIndex(pixel.R >> rshift, pixel.G >> rshift, pixel.B >> rshift)
    histogram[i]++
  }

  return histogram
}

func colorIndex(r, g, b uint8) uint16 {
  return (uint16(r) << (2 * sigbits)) + (uint16(g) << sigbits) + uint16(b)
}

func constructPixelArray(img *image.RGBA) []Rgb {
  rgbaImg := image.NewRGBA(img.Bounds())
  draw.Draw(rgbaImg, rgbaImg.Bounds(), img, image.ZP, draw.Src)

  pixelArray := make([]Rgb, 0, 50)
  var rgbVal = Rgb{}

  for i, pix := range rgbaImg.Pix {
    switch i % 4 {
    case 0:
      rgbVal.R = pix
    case 1:
      rgbVal.G = pix
    case 2:
      rgbVal.B = pix
    case 3:
      if pix >= minTransparency && isOpaque(rgbVal) {
        pixelArray = append(pixelArray, rgbVal)
        rgbVal = Rgb{}
      }
    }
  }

  return pixelArray
}

func isOpaque(r Rgb) bool {
  return r.R <= maxOpaqueVal && r.G <= maxOpaqueVal && r.B <= maxOpaqueVal
}
