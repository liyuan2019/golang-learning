package sort

//冒泡排序
// func main() {
//     arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
//     fmt.Println(GetMax(arr))
//     fmt.Println(BubbleSort(arr))
// }

//GetMax 冒泡排序获取最大值
func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr[len(arr)-1]
}

//BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
