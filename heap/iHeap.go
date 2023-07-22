package heap

type HeapInterface[T any] interface {
	swap(i, j int)
	Insert(element T)
	Get() (T, bool)
	Len() int
	Root() (T, bool)
	Less(ele1, ele2 T) bool
	Print()
}
