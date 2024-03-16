package goed

import (
	"fmt"
	"sync"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
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

// editDistance calculates the Levenshtein distance between two strings
func editDistance(str1, str2 string) int {
	lenStr1 := len(str1)
	lenStr2 := len(str2)


	if lenStr1 == 0 || lenStr2 == 0 {
		return lenStr1 + lenStr2
	}

	// Initialize a 2D slice to store the edit distances
	dp := make([][]int, lenStr1+1)
	for i := range dp {
		dp[i] = make([]int, lenStr2+1)
	}

	// Fill the first row and the first column of the matrix
	for i := 0; i <= lenStr1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenStr2; j++ {
		dp[0][j] = j
	}

	// Calculate edit distances
	for i := 1; i <= lenStr1; i++ {
		for j := 1; j <= lenStr2; j++ {
			cost := 1
			if str1[i-1] == str2[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[lenStr1][lenStr2]
}

func computeTile(str1, str2 string, dp [][]int, tileStartRow, tileStartCol, tileSize, lenStr1, lenStr2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			row := tileStartRow + i
			col := tileStartCol + j
			if row <= lenStr1 && col <= lenStr2 {

				cost := 1
				if str1[row-1] == str2[col-1] {
					cost = 0
				}
				dp[row][col] = min(dp[row-1][col]+1, dp[row][col-1]+1, dp[row-1][col-1]+cost)

			}
		}
	}
}

func editDistanceParallel(str1, str2 string, tileSize int) int {
	lenStr1, lenStr2 := len(str1), len(str2)


	if lenStr1 == 0 || lenStr2 == 0 {
		return lenStr1 + lenStr2
	}


	dp := make([][]int, lenStr1+1)
	for i := range dp {
		dp[i] = make([]int, lenStr2+1)
	}

	// Fill the first row and the first column of the matrix
	for i := 0; i <= lenStr1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenStr2; j++ {
		dp[0][j] = j
	}

	var wg sync.WaitGroup
	//totalWavefronts := ((lenStr1+1)+tileSize-1)/tileSize + ((lenStr2+1)+tileSize-1)/tileSize - 1
	m := ((lenStr1) + tileSize - 1) / tileSize
	n := ((lenStr2) + tileSize - 1) / tileSize
	totalWavefronts := n + m - 1
	fmt.Println("totalWavefronts: ", totalWavefronts)

	for wave := 0; wave < totalWavefronts; wave++ {
		minStart := max2(0, wave-n+1)
		maxStart := min2(wave, m-1)

		//fmt.Printf("\t wave loop: wave = %d, minStart = %d, maxStart = %d\n", wave, minStart, maxStart)

		for start := minStart; start <= maxStart; start++ {
			//tileRow, tileCol := start*tileSize, (wave-start)*tileSize
			tileStartRow := start*tileSize + 1
			tileStartCol := (wave-start)*tileSize + 1
			//fmt.Printf("\t\t in wave loop: tileStartRow = %d, tileStartCol = %d\n", tileStartRow, tileStartCol)
			wg.Add(1)
			go computeTile(str1, str2, dp, tileStartRow, tileStartCol, tileSize, lenStr1, lenStr2, &wg)
		}
		wg.Wait()
	}

	return dp[lenStr1][lenStr2]
}





func EditDistance(str1, str2 string) int {

	return editDistance(str1, str2)

}


func EditDistanceParallel(str1 string, str2 string, tilesize ...int) int {


	var tsv int
	if len(tilesize) > 0 {
		tsv = tilesize[0]
		if tsv < 1 {
			tsv = 1
		}
	} else {
		tsv = 256
	}
	

	return editDistanceParallel(str1, str2, tsv)
	
}
