package levenshtein

import (
    "strings"
)

func Compare(from, to string) int {
    from = strings.ToLower(from)
    to = strings.ToLower(to)

    // initialization
    matrix := [][]int{}
    for i:=0; i<len(from)+1; i++ {
        m := make([]int, len(to)+1)
        m[0] = i
        matrix = append(matrix, m)
    }

    for i, _ := range matrix[0] {
        matrix[0][i] = i
    }

    // levenstein
    for i:=1; i<=len(from); i++ {
        for j:=1; j<=len(to); j++ {
            val1 := matrix[i-1][j]+1
            val2 := matrix[i][j-1]+1
            val3 := matrix[i-1][j-1]

            if from[i-1] != to[j-1] {
                val3 += 1
            }

            val := val1
            if val2 < val {
                val = val2
            }
            if val3 < val {
                val = val3
            }

            matrix[i][j] = val
        }
    }

	return matrix[len(from)][len(to)]
}
