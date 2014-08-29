package primaryColor

type CountPriorityQueue []*VBox
type CountVolumePriorityQueue []*VBox

func (pq CountPriorityQueue) Len() int {
  return len(pq)
}

func (pq CountPriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CountPriorityQueue) Push(x interface{}) {
  *pq = append(*pq, x.(*VBox))
}

func (pq *CountPriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  x := old[n-1]
  *pq = old[0 : n-1]
  return x
}

func (pq CountPriorityQueue) Less(i, j int) bool {
  return pq[i].Count() > pq[j].Count()
}

func (pq CountVolumePriorityQueue) Less(i, j int) bool {
  return pq[i].Count() * pq[i].Volume() > pq[j].Count() * pq[j].Volume()
}

func (pq CountVolumePriorityQueue) Len() int {
  return len(pq)
}

func (pq CountVolumePriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CountVolumePriorityQueue) Push(x interface{}) {
  *pq = append(*pq, x.(*VBox))
}

func (pq *CountVolumePriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  x := old[n-1]
  *pq = old[0 : n-1]
  return x
}
