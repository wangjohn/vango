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
