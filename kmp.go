package main

import "fmt"

//a bc d abce

// str next 指数据前缀和后缀 相交的最长长度。
// a    0
// ab   0
// abc  0
// abcd 0
// abcda 1
// abcdab 2
// abcdabc 3
// abcdabce 0

//改进前的
func getNext(data []string) []int {
	dataLen := len(data)
	next := make([]int, dataLen)
	next[0] = 0;
	for i := 1; i < dataLen; i++ {
		preSufIndex := next[i-1]; //前一位的最长长度，那么preSufIndex就是指那个字符的后一个字符,下标从0开始
		for {
			if (data[preSufIndex] == data[i]) {
				next[i] = preSufIndex + 1
				break
			} else {
				preSufIndex = next[preSufIndex]
			}
			if (preSufIndex == 0) {
				next[i] = 0
				break
			}
		}
	}
	return next
}

/**
	* m 主串
	* s 匹配串
	主串:   abcdabcfabcesaf
	模式串: abcdabce
	如果当前字符不匹配，则回退的长度是除当前字符外，前缀后缀的最长公共子串长度 。 所以每次需要回退的长度为next[index-1]
 */
func Kmp(m []string, s []string) int {
	next := getNext(s)
	mLen := len(m)
	sLen := len(s)

	sIndex := 0
	mIndex := 0
	index := -1

	for mIndex < mLen {
		if (sIndex == 0 && (m[mIndex] != s[sIndex])) {
			mIndex++
		} else if (m[mIndex] == s[sIndex]) {
			mIndex++
			sIndex++
		} else {
			if sIndex != 0 {
				sIndex = next[sIndex-1]
			}
		}
		if (sIndex == sLen) {
			index = mIndex
			break
		}
	}

	return index
}

//改进后的next数组
func getNextNew(data []string) []int {
	dataLength := len(data)
	next := make([]int, dataLength)
	next[0] = -1
	k := -1
	index := 0

	for index < dataLength-1 {
		if k == -1 || data[index] == data[k] {
			k++
			index++
			next[index] = k
		} else {
			k = next[k]
		}
	}
	return next
}

//改进后的
func KmpNeW(m []string, s []string) int {
	next := getNextNew(s)
	mLen := len(m)
	sLen := len(s)

	sIndex := 0
	mIndex := 0
	index := -1

	for mIndex < mLen {
		if (sIndex == 0 && (m[mIndex] != s[sIndex])) {
			mIndex++
		} else if (m[mIndex] == s[sIndex]) {
			mIndex++
			sIndex++
		} else {
			if sIndex != 0 {
				sIndex = next[sIndex]
			}
		}
		if (sIndex == sLen) {
			index = mIndex
			break
		}
	}
	return index
}

func getNextNew2(data []string) []int {
	dataLength := len(data)
	next := make([]int, dataLength)
	next[0] = -1
	k := -1
	index := 0

	for index < dataLength-1 {
		if k == -1 || data[index] == data[k] {
			k++
			index++
			if data[index] == data[k] {
				next[index] = next[k]
			} else {
				next[index] = k
			}
		} else {
			k = next[k]
		}
	}
	return next
}
//改进后的
func KmpNeW2(m []string, s []string) int {
	next := getNextNew(s)
	mLen := len(m)
	sLen := len(s)

	sIndex := 0
	mIndex := 0
	index := -1

	for mIndex < mLen {
		if (sIndex == 0 && (m[mIndex] != s[sIndex])) {
			mIndex++
		} else if (m[mIndex] == s[sIndex]) {
			mIndex++
			sIndex++
		} else {
			if sIndex != 0 {
				sIndex = next[sIndex]
			}
		}
		if (sIndex == sLen) {
			index = mIndex
			break
		}
	}
	return index
}

func main()  {
	str := []string{"a","b","c","a","b","d","a","b","c","a","b","d"};
	str1 := []string{"a","b","c","a","b","d"};

	fmt.Println(getNext(str1))
	fmt.Println(getNextNew(str1))
	fmt.Println(getNextNew2(str1))

	fmt.Println(KmpNeW2(str,str1))
	fmt.Println(KmpNeW(str,str1))
	fmt.Println(Kmp(str,str1))
}
