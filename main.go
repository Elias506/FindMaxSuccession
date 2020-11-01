package main

//Алгоритм работает за время O[n*log(n)]

//FindMaxSuccession ...
func FindMaxSuccession(inputData []int) (int, []int) {
	
	l := len(inputData)
	array := make([]int, l)
	var k int = 0
	for i, element := range inputData {
		element = inputData[i]
		index := searchIndex(array[0:k], element)
		if index != -1 {
			array[index] = element
			if k - 1 < index {
				k++
			}
		}
	}
	return k, array[0:k]
}

func searchIndex(array []int, q int) int{
	if len(array) == 0 {
		return 0
	}
	var left, right = 0, len(array) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if q == array[mid] || left == right {
			break
		} 
		if q > array[mid] {
			left = mid + 1
		}
		if q < array[mid] {
			right = mid - 1
		}
	}
	if q > array[mid] || array[mid] == 0{
		return mid + 1
	}
	return mid 
}	
