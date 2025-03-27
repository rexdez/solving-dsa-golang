package main

type KthLargest struct {
    k int
	heapScores []int
}

func Constructor(k int, nums []int) KthLargest {
	data := KthLargest{k : k, heapScores: make([]int, 0)}
	for _, val := range nums {
		data.Add(val)
	}
	return data
}


func (kl *KthLargest) Add(val int) int {
	kl.heapScores = append(kl.heapScores, val)
	n := len(kl.heapScores)
	for i := n-2; i >-1 ; i-- {
		if kl.heapScores[i] > kl.heapScores[i+1] {
			kl.heapScores[i], kl.heapScores[i+1] = kl.heapScores[i+1], kl.heapScores[i]
		}
	}
	if len(kl.heapScores) > kl.k {
		kl.heapScores = kl.heapScores[(len(kl.heapScores)-kl.k):]
	}
	return kl.heapScores[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
