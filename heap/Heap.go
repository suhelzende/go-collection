package heap

import "fmt"

type Heap[T any] struct {
	heap []T

	// Comparator defins if it is min heap or max heap
	//
	// for min Heap ele1 <= ele2
	//
	// for max heap ele1 >= ele2
	comparetor func(ele1, ele2 T) bool
}

func CreateHeap[T any](comparator func(T, T) bool) HeapInterface[T] {
	return &Heap[T]{
		heap:       make([]T, 0),
		comparetor: comparator,
	}
}

func (mh *Heap[T]) Insert(element T) {
	mh.heap = append(mh.heap, element)
	mh.upHeapify()
}

func (mh *Heap[T]) upHeapify() {
	idx := mh.Len() - 1
	parent := (idx - 1) / 2
	for idx > 0 && parent >= 0 {
		if mh.Less(mh.heap[idx], mh.heap[parent]) {
			mh.swap(idx, parent)
			idx = parent
			parent = (idx - 1) / 2
		} else {
			break
		}
	}
}

func CreateHeapFromList[T any](list []T, comparator func(t1, t2 T) bool) HeapInterface[T] {
	heap := &Heap[T]{
		heap:       make([]T, 0),
		comparetor: comparator,
	}

	for _, ele := range list {
		heap.Insert(ele)
	}
	return heap
}

func (mh *Heap[T]) downHeapify() {
	idx := 0
	lc := idx*2 + 1
	rc := lc + 1
	for lc < mh.Len() && rc < mh.Len() {
		if mh.Less(mh.heap[idx], mh.heap[lc]) && mh.Less(mh.heap[idx], mh.heap[rc]) {
			return
		}

		if mh.Less(mh.heap[lc], mh.heap[rc]) {
			mh.swap(lc, idx)
			idx = lc
		} else {
			mh.swap(rc, idx)
			idx = rc
		}
		lc = idx*2 + 1
		rc = lc + 1
	}

	if lc < mh.Len() && mh.Less(mh.heap[lc], mh.heap[idx]) {
		mh.swap(lc, idx)
	}
}

func (mh *Heap[T]) Get() (T, bool) {
	if len(mh.heap) == 0 {
		var None T
		return None, false
	}
	ele := mh.heap[0]
	mh.swap(0, mh.Len()-1)
	if mh.Len() > 1 {
		mh.heap = mh.heap[0 : mh.Len()-1]
	} else {
		mh.heap = []T{}
	}
	mh.downHeapify()
	return ele, true
}

func (mh *Heap[T]) Root() (T, bool) {
	if len(mh.heap) == 0 {
		var None T
		return None, false
	}
	return mh.heap[0], true
}

func (mh *Heap[T]) Less(ele1, ele2 T) bool {
	return mh.comparetor(ele1, ele2)
}

func (mh *Heap[T]) Len() int {
	return len(mh.heap)
}

// Swap : swaps element at i and j indexes in heap
func (mh *Heap[T]) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func (mh *Heap[T]) Print() {
	for _, ele := range mh.heap {
		fmt.Printf("%v ", ele)
	}
}
