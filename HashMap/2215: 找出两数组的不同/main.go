package main

func main() {

}

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := map[int]int{}
	nums2Map := map[int]int{}

	for _, num := range nums1 {
		nums1Map[num]++
	}
	for _, num := range nums2 {
		nums2Map[num]++
	}

	var ans1, ans2 []int

	for key, _ := range nums1Map {
		if nums2Map[key] == 0 && nums1Map[key] > 0 {
			ans1 = append(ans1, key)
		}
	}

	for key, _ := range nums2Map {
		if nums1Map[key] == 0 && nums2Map[key] > 0 {
			ans2 = append(ans2, key)
		}
	}

	return [][]int{ans1, ans2}
}
