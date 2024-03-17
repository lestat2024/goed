package goed

import (
//	"fmt"
	"sync"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}


func editDistance(first, second string) int {
	lenFirst := len(first)
	lenSecond := len(second)


	if lenFirst == 0 || lenSecond == 0 {
		return lenFirst + lenSecond
	}


	dp := make([][]int, lenFirst+1)
	for i := range dp {
		dp[i] = make([]int, lenSecond+1)
	}


	for i := 0; i <= lenFirst; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenSecond; j++ {
		dp[0][j] = j
	}


	for i := 1; i <= lenFirst; i++ {
		for j := 1; j <= lenSecond; j++ {
			cost := 1
			if first[i-1] == second[j-1] {
				cost = 0
			}
			dp[i][j] = min3(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[lenFirst][lenSecond]
}

func computeTile(first, second string, dp [][]int, tileStartRow, tileStartCol, tileSize, lenFirst, lenSecond int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			row := tileStartRow + i
			col := tileStartCol + j
			if row <= lenFirst && col <= lenSecond {

				cost := 1
				if first[row-1] == second[col-1] {
					cost = 0
				}
				dp[row][col] = min3(dp[row-1][col]+1, dp[row][col-1]+1, dp[row-1][col-1]+cost)

			}
		}
	}
}

func editDistanceParallel(first, second string, tileSize int) int {
	lenFirst, lenSecond := len(first), len(second)


	if lenFirst == 0 || lenSecond == 0 {
		return lenFirst + lenSecond
	}


	dp := make([][]int, lenFirst+1)
	for i := range dp {
		dp[i] = make([]int, lenSecond+1)
	}


	for i := 0; i <= lenFirst; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenSecond; j++ {
		dp[0][j] = j
	}

	var wg sync.WaitGroup

	m := ((lenFirst) + tileSize - 1) / tileSize
	n := ((lenSecond) + tileSize - 1) / tileSize
	totalWavefronts := n + m - 1

	//fmt.Println("totalWavefronts: ", totalWavefronts)

	for wave := 0; wave < totalWavefronts; wave++ {
		minStart := max(0, wave-n+1)
		maxStart := min(wave, m-1)

		//fmt.Printf("\t wave loop: wave = %d, minStart = %d, maxStart = %d\n", wave, minStart, maxStart)

		for start := minStart; start <= maxStart; start++ {

			tileStartRow := start*tileSize + 1
			tileStartCol := (wave-start)*tileSize + 1
			//fmt.Printf("\t\t in wave loop: tileStartRow = %d, tileStartCol = %d\n", tileStartRow, tileStartCol)

			wg.Add(1)
			go computeTile(first, second, dp, tileStartRow, tileStartCol, tileSize, lenFirst, lenSecond, &wg)
		}
		wg.Wait()
	}

	return dp[lenFirst][lenSecond]
}





func EditDistance(first, second string) int {

	return editDistance(first, second)

}


func EditDistanceParallel(first string, second string, tilesize ...int) int {


	var tsv int
	if len(tilesize) > 0 {
		tsv = tilesize[0]
		if tsv < 1 {
			tsv = 1
		}
	} else {
		tsv = 256
	}
	

	return editDistanceParallel(first, second, tsv)
	
}
