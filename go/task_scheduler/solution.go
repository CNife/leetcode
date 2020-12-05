package task_scheduler

import (
	"sort"
)

func LeastInterval(tasks []byte, n int) int {
	st := makeSortedTasks(tasks)
	time, lastLen := 0, len(st)
	for {
		time += n + 1
		st.consumeTasks(n + 1)
		if len(st) > 0 {
			lastLen = len(st)
		} else {
			break
		}
	}
	return time - (n + 1 - lastLen)
}

type sortedTasks []int

func makeSortedTasks(tasks []byte) sortedTasks {
	var counter [26]int
	for _, task := range tasks {
		index := task - 'A'
		counter[index]++
	}

	var st sortedTasks
	for _, count := range counter {
		if count > 0 {
			st = append(st, count)
		}
	}
	//goland:noinspection GoNilness
	st.sort()
	return st
}

func (s sortedTasks) sort() {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}

func (s *sortedTasks) consumeTasks(n int) {
	for i := 0; i < n && i < len(*s); i++ {
		(*s)[i]--
	}
	s.sort()
	i := len(*s) - 1
	for ; i >= 0; i-- {
		if (*s)[i] > 0 {
			break
		}
	}
	*s = (*s)[:i+1]
}
