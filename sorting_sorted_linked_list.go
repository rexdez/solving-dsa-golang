// You can edit this code!
// Click here and start typing.
package main


// Merge two sorted linked list

type LinkedList struct {
	num  int
	next *LinkedList
}

func SortLinkedList(list1, list2 []int) *LinkedList{
	ll1 := initLL(list1)
	ll2 := initLL(list2)

	return MergeLL(ll1, ll2)

}

func MergeLL(ll1, ll2 *LinkedList) *LinkedList {
	start := &LinkedList{}
	point := start
	for ll1 != nil && ll2 != nil {
		if ll1.num < ll2.num {
			point.next= ll1
			ll1 = ll1.next
		} else {
			point.next= ll2
			ll2 = ll2.next
		}
		point = point.next
	}
	if ll1 != nil {
		point.next = ll1
	}
	if ll2 != nil {
		point.next = ll2
	}
	return start.next
}

func initLL(list []int) *LinkedList {
	buffer := &LinkedList{num: list[0]}
	ll1 := buffer
	for _, val := range list[1:] {
		ll1.next = &LinkedList{num: val}
		ll1 = ll1.next 
	}
	return buffer
}
