[top]
 
#golang学习

## 算法相关
### kmp算法，next数组解析
#### 根据定义写出获取next数组的方法
定义网上很多，就不在描述了。

```go
func getNext(data []string) []int {
	dataLen := len(data)
	next := make([]int, dataLen)
	next[0] = 0;
	for i := 1; i < dataLen; i++ {
		preSufIndex := next[i-1]; //前缀后缀的最长公共子串长度，那么preSufIndex就是指那个字符的后一个字符,下标从0开始
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
```

根据以上代码，可以得出next 求的是前缀后缀的最长公共子串长度 。
举个例子， 
例如 
主串:      abcdabcfabcesaf
模式串:  abcdabce
当主串f 和模式串e 不匹配时，则回退的长度是除当前字符外，前缀后缀的最长公共子串长度 。 所以每次需要回退的长度为next[index-1]。
出于这种情况，next数组可以优化为：保存除当前字符外，前缀后缀的最长公共子串长度。
同时，根据前面的代码，会发现next[0]是特殊情况，可以使用一个特别的数字来代替，优化代码

```
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
```
这个代码，也是网上很多博客晒出获取next数组的代码。

在举个例子

主串:      abacedas
模式串:  abab
当主串c和模式串b不匹配时, b会回退到下标1，这时候1上的字符也是b。和主串c肯定是不匹配的，针对这种情况，模式串回退到下标0，是最优的。
优化后的代码
```go
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
```
