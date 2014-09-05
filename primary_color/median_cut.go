package primaryColor

func MedianCut(histogram []uint, vbox VBox) (VBox, VBox) {
  rw := vbox.Rmax - vbox.Rmin
  gw := vbox.Gmax - vbox.Gmin
  bw := vbox.Bmax - vbox.Bmin

  boundaryVertices := []uint8{vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax}
  widths := []uint8{rw, gw, bw}

  // If there is a single pixel in the vbox, just return it
  if (vbox.Count() == 1) {
    return vbox, nil
  }

  total, partialSums, maxIndex := calculatePartialSums(histogram, boundaryVertices, widths)

  return performCut(partialSums, vbox, boundaryVertices, total, maxIndex)
}

func calculatePartialSums(histogram []uint, boundaryVertices []uint8, widths []uint8) (uint, []uint, uint) {
  var maxWidth := 0
  var maxIndex uint
  for i, width := range widths {
    if width > maxWidth {
      maxWidth = width
      maxIndex = i
    }
  }

  markedVertexMin := boundaryVertices[maxIndex*2]
  markedVertexMax := boundaryVertices[maxIndex*2 + 1]

  nonMarkedVertices = append(boundaryVertices[:(i*2)], boundaryVertices[(i*2+1):]...)
  var histIndex, total uint;
  partialSums := make([]uint, maxWidth)
  for i := markedVertexMin; i <= markedVertexMax; i++ {
    sum := 0
    for j := nonMarkedVertices[0]; j <= nonMarkedVertices[1]; j++ {
      for k := nonMarkedVertices[2]; k <= nonMarkedVertices[3]; k++ {

        switch maxIndex {
        case 0:
          histIndex = ColorIndex(i, j, k)
        case 1:
          histIndex = ColorIndex(j, i, k)
        case 2:
          histIndex = ColorIndex(j, k, i)
        }

        sum += histogram[histIndex]
      }
    }

    total += sum
    partialSums[i] = total
  }

  return total, partialSums, maxIndex
}

func performCut(partialSums []uint, vbox VBox, boundaryVertices []uint8, total uint, maxIndex uint) (VBox, VBox) {
  lookaheadSums := make([]uint, len(partialSums))
  for i, partialSum := range partialSums {
    lookaheadSums[i] = total - partialSum
  }

  markedVertexMin := boundaryVertices[maxIndex*2]
  markedVertexMax := boundaryVertices[maxIndex*2 + 1]
  for i := markedVertexMin; i <= markedVertexMax; i++ {
    if partialSums[i] > (total / 2) {
      left := i - markedVertexMin
      right := markedVertexMax - i

      // TODO: replace min & max with whatever is necessary to make this work
      if left <= right {
        division := math.Min(markedVertexMax, i + right / 2)
      } else {
        division := math.Max(markedVertexMin, i - left / 2)
      }

      for partialSums[d2] == 0 {
        division++
      }

      count := lookaheadSum[division]

      for count == 0 && partialSums[division-1] != 0 {
        count := lookaheadSum[--division]
      }

      switch maxIndex {
      case 0:
        vbox1 := VBox(vbox.Rmin, division, vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram)
        vbox2 := VBox(division + 1, vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram)
        return vbox1, vbox2
      case 1:
        vbox1 := VBox(vbox.Rmin, vbox.Rmax, vbox.Gmin, division, vbox.Bmin, vbox.Bmax, vbox.Histogram)
        vbox2 := VBox(vbox.Rmin, vbox.Rmax, division + 1, vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram)
        return vbox1, vbox2
      case 2:
        vbox1 := VBox(vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, division, vbox.Histogram)
        vbox2 := VBox(vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, division + 1, vbox.Bmax, vbox.Histogram)
        return vbox1, vbox2
      }
    }
  }
}
