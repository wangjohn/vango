package primaryColor

import (
  "fmt"
  "image"
  "image/draw"
  "container/heap"
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

  // TODO: take histogram and perform median cut
  vbox := constructVBox(pixelArray, histogram)

  return result, nil
}

func iterate(queue CountPriorityQueue, histogram []uint, iterations int) {
  numColors := 0
  for i := 0; i < iterations; i++ {
    vbox := heap.Pop(queue)
    if vbox.Count() > 0 {
      heap.Push(queue, vbox)
    } else {
      vbox1, vbox2 := applyMedianCut(histogram, vbox)

      heap.Push(queue, vbox1)
      heap.Push(queue, vbox2)
      numColors++
    }
  }
}

func constructVBox(pixelArray []Rgb, histogram []uint) VBox {
  var rval, gval, bval uint8
  rmax, gmax, bmax := uint8(0), uint8(0), uint8(0)
  rmin, gmin, bmin := ^uint8(0), ^uint8(0), ^uint8(0)

  for _, pixel := range pixelArray {
    rval = pixel.R >> rshift
    gval = pixel.G >> rshift
    bval = pixel.B >> rshift

    if rval < rmin {
      rmin = rval
    } else if rval > rmax {
      rmax = rval
    }
    if gval < gmin {
      gmin = gval
    } else if gval > gmax {
      gmax = gval
    }
    if bval < bmin {
      bmin = bval
    } else if bval > bmax {
      bmax = bval
    }
  }

  return VBox{Rmin: rmin, Rmax: rmax, Gmin: gmin, Gmax: gmax, Bmin: bmin, Bmax: bmax, Histogram: histogram}
}

func constructHistogram(pixelArray []Rgb) []uint{
  size := 1 << (3 * sigbits)
  histogram := make([]uint, size)

  for _, pixel := range pixelArray {
    i := ColorIndex(pixel.R >> rshift, pixel.G >> rshift, pixel.B >> rshift)
    histogram[i]++
  }

  return histogram
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
