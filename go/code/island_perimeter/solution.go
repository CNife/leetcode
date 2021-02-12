//package island_perimeter
//
//import "sync"
//
//func IslandPerimeter(grid [][]int) int {
//	m, n := len(grid), len(grid[0])
//	rowIterator := func(i int) func() (int, bool) {
//		row, j := grid[i], 0
//		return func() (int, bool) {
//			if j < n {
//				j++
//				return row[j-1], true
//			} else {
//				return 0, false
//			}
//		}
//	}
//	columnIterator := func(i int) func() (int, bool) {
//		j := 0
//		return func() (int, bool) {
//			if j < m {
//				j++
//				return grid[j-1][i], true
//			} else {
//				return 0, false
//			}
//		}
//	}
//	countPerimeter := func(itr func() (int, bool)) int {
//		flag, count := 0, 0
//		for {
//			if curr, valid := itr(); valid {
//				if flag != curr {
//					flag = curr
//					count++
//				}
//			} else {
//				break
//			}
//		}
//		if flag == 1 {
//			return count + 1
//		} else {
//			return count
//		}
//	}
//
//	result := 0
//	mutex := &sync.Mutex{}
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < m; i++ {
//			p := countPerimeter(rowIterator(i))
//			mutex.Lock()
//			result += p
//			mutex.Unlock()
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 0; i < n; i++ {
//			p := countPerimeter(columnIterator(i))
//			mutex.Lock()
//			result += p
//			mutex.Unlock()
//		}
//	}()
//	wg.Wait()
//	return result
//}

package island_perimeter

func IslandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowCount, columnCount := make(chan int), make(chan int)

	go func() {
		rowResult := 0
		for i := 0; i < m; i++ {
			flag := 0
			for j := 0; j < n; j++ {
				if v := grid[i][j]; v != flag {
					flag = v
					rowResult++
				}
			}
			if flag == 1 {
				rowResult++
			}
		}
		rowCount <- rowResult
	}()

	go func() {
		columnResult := 0
		for j := 0; j < n; j++ {
			flag := 0
			for i := 0; i < m; i++ {
				if v := grid[i][j]; v != flag {
					flag = v
					columnResult++
				}
			}
			if flag == 1 {
				columnResult++
			}
		}
		columnCount <- columnResult
	}()

	return <-rowCount + <-columnCount
}
