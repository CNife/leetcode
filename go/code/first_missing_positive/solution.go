package first_missing_positive

func FirstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}

// bitmap 解法，内存爆了
//import "math"
//
//func FirstMissingPositive(nums []int) int {
//	b := newBitMap()
//	for _, num := range nums {
//		if num > 0 {
//			b.set(num)
//		}
//	}
//	return b.firstZero()
//}
//
//type bitMap struct {
//	inner []uint64
//}
//
//func newBitMap() bitMap {
//	return bitMap{inner: []uint64{1}}
//}
//
//func (b *bitMap) set(index int) {
//	block, offset := index/64, index%64
//	if oldBlockSize := len(b.inner); block >= oldBlockSize {
//		for i := oldBlockSize; i <= block; i++ {
//			b.inner = append(b.inner, 0)
//		}
//	}
//
//	b.inner[block] |= uint64(1) << offset
//}
//
//func (b *bitMap) firstZero() int {
//	block, blockIndex := uint64(0), 0
//	for i, blk := range b.inner {
//		if blk != math.MaxUint64 {
//			block, blockIndex = blk, i
//		}
//	}
//	if blockIndex == 0 && block == 0 {
//		return len(b.inner) * 64
//	}
//	for i := 0; i < 64; i++ {
//		if (block>>i)&1 == 0 {
//			return i + 64*blockIndex
//		}
//	}
//	panic("unreachable")
//}
