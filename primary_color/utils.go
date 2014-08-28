package primaryColor

func ColorIndex(r, g, b uint8) uint {
  return (uint(r) << (2 * sigbits)) + (uint(g) << sigbits) + uint(b)
}

