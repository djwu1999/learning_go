package main

import (
	"container/heap"
	"fmt"
	rand2 "math/rand"
)

type Phone struct {
	number string
	id int
}

func (phone *Phone) call(number string) {
	fmt.Printf("I'm %s, please call the number %s, %d", phone.number, number, phone.id)
}

func findTargetSumWays(nums []int, target int) (count int) {
	var helper func(index int, sum int)
	helper = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		helper(index + 1, sum + nums[index])
		helper(index + 1, sum - nums[index])
	}
	helper(0, 0)
	return
}

//func helper(nums []int, target int, index int) {
//	if index == len(nums) {
//		if sum == target {
//			res++
//		}
//		return
//	}
//	sum += nums[index]
//	helper(nums, target, index + 1)
//	sum -= nums[index]
//	sum -= nums[index]
//	helper(nums, target, index + 1)
//	sum += nums[index]
//}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	//heap := new(heap.Interface)=
	h := &IntHeap{12, 123, 312}
	for i := 0; i < 12; i++ {
		num := rand2.Intn(1000)
		heap.Push(h, num)
	}
	for h.Len() > 0{
		num := heap.Pop(h)
		fmt.Println(num)
	}
	fmt.Println(h.Len())
}