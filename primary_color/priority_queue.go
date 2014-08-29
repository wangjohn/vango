package primaryColor

type PriorityQueue struct {
  Array []*VBox
}

func (pq PriorityQueue) Len() int {
  return len(pq.Array)
}

func (pq PriorityQueue) Swap(i, j int) {
  pq.Array[i], pq.Array[j] = pq.Array[j], pq.Array[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
  *pq.Array = append(*pq.Array, x.(*VBox))
}

func (pq *PriorityQueue) Pop() interface{} {
  old := *pq.Array
  n := len(old)
  x := old[n-1]
  *pq.Array = old[0 : n-1]
  return x
}

type CountPriorityQueue struct {
  PriorityQueue
}

func (pq CountPriorityQueue) Less(i, j int) bool {
  return pq.Array[i].Count() > pq.Array[j].Count()
}

type CountVolumePriorityQueue struct {
  PriorityQueue
}

func (pq CountVolumePriorityQueue) Less(i, j int) bool {
  return pq.Array[i].Count() * pq.Array[i].Volume() > pq.Array[j].Count() * pq.Array[j].Volume()
}
