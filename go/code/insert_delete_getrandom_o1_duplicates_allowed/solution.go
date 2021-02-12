package insert_delete_getrandom_o1_duplicates_allowed

import "math/rand"

type RandomizedCollection struct {
	indexMap map[int]map[int]struct{}
	data     []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		indexMap: make(map[int]map[int]struct{}),
	}
}

func (rc *RandomizedCollection) Insert(value int) bool {
	index := len(rc.data)
	rc.data = append(rc.data, value)
	if indexes, exists := rc.indexMap[value]; exists {
		indexes[index] = struct{}{}
		return false
	} else {
		indexes := map[int]struct{}{index: {}}
		rc.indexMap[value] = indexes
		return true
	}
}

func (rc *RandomizedCollection) Remove(value int) bool {
	if indexes, exists := rc.indexMap[value]; exists {
		var index int
		for i := range indexes {
			index = i
			break
		}
		delete(indexes, index)
		if len(indexes) < 1 {
			delete(rc.indexMap, value)
		}

		newLen := len(rc.data) - 1
		if index != newLen {
			last := rc.data[newLen]
			rc.data[index] = last
			rc.data = rc.data[:newLen]
			lastIndexes := rc.indexMap[last]
			delete(lastIndexes, newLen)
			lastIndexes[index] = struct{}{}
		} else {
			rc.data = rc.data[:newLen]
		}
		return true
	} else {
		return false
	}
}

func (rc *RandomizedCollection) GetRandom() int {
	randomIndex := rand.Intn(len(rc.data))
	return rc.data[randomIndex]
}
