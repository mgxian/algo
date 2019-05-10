package sort

type sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

// MyInts is a wrap of []int
type MyInts []int

func (p MyInts) Len() int           { return len(p) }
func (p MyInts) Less(i, j int) bool { return p[i] < p[j] }
func (p MyInts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
