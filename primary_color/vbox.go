package primaryColor

type VBox struct {
  Rmin uint8
  Rmax uint8
  Gmin uint8
  Gmax uint8
  Bmin uint8
  Bmax uint8
  Histogram []uint
}

func (v *VBox) Count() uint {
  var points uint = 0
  var index uint
  for i := v.Rmin; i <= v.Rmax; i++ {
    for j := v.Gmin; j <= v.Gmax; j++ {
      for k := v.Bmin; k <= v.Bmax; k++ {
        index = ColorIndex(i, j, k)
        points += v.Histogram[index]
      }
    }
  }

  return points
}

func (v *VBox) Volume() uint {
  return uint(v.Rmax - v.Rmin) * uint(v.Gmax - v.Gmin) * uint(v.Bmax - v.Bmin)
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

