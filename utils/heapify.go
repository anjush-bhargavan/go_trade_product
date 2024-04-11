package utils

import "github.com/anjush-bhargavan/go_trade_product/pkg/model"

func Push(Bids *[]model.Bids,val *model.Bids) {
	*Bids = append(*Bids, *val)
	ShiftUp(Bids,len(*Bids)-1)
}

func Pop(Bids *[]model.Bids) *model.Bids {
	max := (*Bids)[0]
	Swap(Bids,0,len(*Bids)-1)
	*Bids = (*Bids)[:len(*Bids)-1]
	ShiftDown(Bids,0)
	return &max
}

func ShiftUp(Bids *[]model.Bids,currIdx int) {
	parentIdx := ParentIdx(currIdx)
	for currIdx >= 0 && (*Bids)[currIdx].Amount > (*Bids)[parentIdx].Amount {
		Swap(Bids,currIdx,parentIdx)
		currIdx = parentIdx
		parentIdx = ParentIdx(currIdx)
	}
}


func ShiftDown(Bids *[]model.Bids,currIdx int) {
	endIdx := len(*Bids) - 1
	leftIdx := LeftChild(currIdx)

	for leftIdx <= endIdx {
		shiftIdx := leftIdx
		if (*Bids)[RightChild(currIdx)].Amount > (*Bids)[shiftIdx].Amount {
			shiftIdx = RightChild(currIdx)
		}

		if (*Bids)[shiftIdx].Amount > (*Bids)[currIdx].Amount {
			Swap(Bids,currIdx,shiftIdx)
			currIdx = shiftIdx
			leftIdx = LeftChild(currIdx)
		}else{
			break
		}
	}

}



func ParentIdx(i int) int {
	return (i-1)/2
}

func LeftChild(i int) int {
	return (2*i) + 1
}


func RightChild(i int) int {
	return (2*i) + 2
}

func Swap(Bids *[]model.Bids,i,j int) {
	(*Bids)[i],(*Bids)[j] = (*Bids)[j], (*Bids)[i]
}
