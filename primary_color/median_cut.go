package primaryColor

func MedianCut(histogram Histogram, vbox VBox) (VBox, VBox) {
  rw := vbox.Rmax - vbox.Rmin
  gw := vbox.Gmax - vbox.Gmin
  bw := vbox.Bmax - vbox.Bmin
  maxw := multipleMaximum(rw, gw, bw)

  // If there is a single pixel in the vbox, just return it
  if (vbox.Count() == 1) {
    return vbox, nil
  }

  return vbox1, vbox2
}

func calculatePartialSums(histogram Histogram, vbox VBox) ([]uint, uint) {

}

func multipleMaximum(a, b, c uint8) uint8 {
  var i1, i2 uint8
  if a > b {
    i1 = a
  } else {
    i1 = b
  }

  if b > c {
    i2 = b
  } else {
    i2 = c
  }

  if i1 > i2 {
    return i1
  } else {
    return i2
  }
}
