package ziraffe

import "fmt"

func createArray(s1, s2 []rune) [][]int {
  array := make([][]int, len(s1) + 1)
  for i := 0; i <= len(s1); i++ {
    array[i] = make([]int, len(s2) + 1)
    array[i][0] = i
  }
  for j:= 0; j <= len(s2); j++ {
    array[0][j] = j
  }
  return array
}

func min(v1, v2 int) int {
  if v1 < v2 {
    return v1
  }
  return v2
}

func debugPrint(array [][]int, len1, len2 int) {
  for i := 0; i <= len1; i++ {
    for j := 0; j <= len2; j++ {
      fmt.Printf(" %-3d", array[i][j])
    }
    fmt.Println()
  }
}

func levenshteinImpl(s1, s2 []rune, insertFunc, removeFunc, updateFunc func(c1, c2 rune) int) int {
  array := createArray(s1, s2)
  for i := 1; i <= len(s1); i++ {
    for j := 1; j <= len(s2); j++ {
      d1 := array[i - 1][j] + insertFunc(s1[i - 1], s2[j - 1])
      d2 := array[i][j - 1] + removeFunc(s1[i - 1], s2[j - 1])
      d3 := array[i - 1][j - 1] + updateFunc(s1[i - 1], s2[j - 1])
      array[i][j] = min(d1, min(d2, d3))
    }
  }
  // debugPrint(array, len(s1), len(s2))
  return array[len(s1)][len(s2)]
}

/*
LevenshteinS function calculates the edit distance (levenshtein distance) of two given strings.
This function calls Levenshtein(s1, s2, nil, nil, nil)
*/
func LevenshteinS(s1, s2 string) int {
  return Levenshtein(s1, s2, nil, nil, nil)
}

func findDefaultFuncIfOriginIsNil(origin, defaultFunc func (c1, c2 rune) int) (func (c1, c2 rune) int) {
  if origin == nil {
    return defaultFunc
  }
  return origin
}

/*
Levenshtein function calculates the edit distance (levenshtein distance) of two given strings.
Insert, remove, and update costs are computed by given functions.
If the given functions are nil, default compute algorithm is applied.
Default cost compute algorithms are:
insert, and remove cost is 1, and update cost is 0 if two runes are same, unless 1.
*/
func Levenshtein(s1, s2 string, insertFunc, removeFunc, updateFunc func(c1, c2 rune) int) int {
  insertFunc = findDefaultFuncIfOriginIsNil(insertFunc, func (c1, c2 rune) int { return 1 })
  removeFunc = findDefaultFuncIfOriginIsNil(removeFunc, insertFunc)
  updateFunc = findDefaultFuncIfOriginIsNil(updateFunc,
    func (c1, c2 rune) int {
      if c1 == c2 {
        return 0
      }
      return 1
    })
  return levenshteinImpl([]rune(s1), []rune(s2), insertFunc, removeFunc, updateFunc)
}
