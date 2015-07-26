package primaryColor

import "math"

func MedianCut(histogram []uint, vbox VBox) (VBox, VBox) {
	rw := vbox.Rmax - vbox.Rmin
	gw := vbox.Gmax - vbox.Gmin
	bw := vbox.Bmax - vbox.Bmin

	boundaryVertices := []uint8{vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax}
	widths := []uint8{rw, gw, bw}

	// If there is a single pixel in the vbox, just return it
	if vbox.Count() == 1 {
		return vbox, vbox
	}

	total, partialSums, maxIndex := calculatePartialSums(histogram, boundaryVertices, widths)

	return performCut(partialSums, vbox, boundaryVertices, total, maxIndex)
}

func calculatePartialSums(histogram []uint, boundaryVertices []uint8, widths []uint8) (uint, []uint, uint) {
	var maxWidth, width uint8
	maxWidth = 0
	var maxIndex uint
	var i int
	for i, width = range widths {
		if width > maxWidth {
			maxWidth = width
			maxIndex = uint(i)
		}
	}

	markedVertexMin := boundaryVertices[maxIndex*2]
	markedVertexMax := boundaryVertices[maxIndex*2+1]

	nonMarkedVertices := append(boundaryVertices[:(i*2)], boundaryVertices[(i*2+1):]...)
	var histIndex, total uint
	partialSums := make([]uint, maxWidth)
	for i := markedVertexMin; i <= markedVertexMax; i++ {
		var sum uint
		sum = 0
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
	markedVertexMax := boundaryVertices[maxIndex*2+1]
	for i := markedVertexMin; i <= markedVertexMax; i++ {
		if partialSums[i] > (total / 2) {
			left := i - markedVertexMin
			right := markedVertexMax - i

			// TODO: replace min & max with whatever is necessary to make this work
			var divisionFloat float64
			if left <= right {
				divisionFloat = math.Min(float64(markedVertexMax), float64(i+right/2))
			} else {
				divisionFloat = math.Max(float64(markedVertexMin), float64(i-left/2))
			}
			division := int(divisionFloat)
			for partialSums[division-2] == 0 {
				division++
			}

			count := lookaheadSums[division]

			for count == 0 && partialSums[division-1] != 0 {
				division -= 1
				//count := lookaheadSums[division]
			}

			switch maxIndex {
			case 0:
				vbox1 := VBox{vbox.Rmin, uint8(division), vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram}
				vbox2 := VBox{uint8(division + 1), vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram}
				return vbox1, vbox2
			case 1:
				vbox1 := VBox{vbox.Rmin, vbox.Rmax, vbox.Gmin, uint8(division), vbox.Bmin, vbox.Bmax, vbox.Histogram}
				vbox2 := VBox{vbox.Rmin, vbox.Rmax, uint8(division + 1), vbox.Gmax, vbox.Bmin, vbox.Bmax, vbox.Histogram}
				return vbox1, vbox2
			case 2:
				vbox1 := VBox{vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, vbox.Bmin, uint8(division), vbox.Histogram}
				vbox2 := VBox{vbox.Rmin, vbox.Rmax, vbox.Gmin, vbox.Gmax, uint8(division + 1), vbox.Bmax, vbox.Histogram}
				return vbox1, vbox2
			}
		}
	}
	return VBox{}, VBox{}
}
