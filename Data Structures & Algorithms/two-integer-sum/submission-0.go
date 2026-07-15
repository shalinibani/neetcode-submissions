func twoSum(nums []int, target int) []int {
    set := make(map[int]int)

    for i, n := range nums{
        rem := target - n
        if remIdx, ok := set[rem]; ok{
            return []int{remIdx, i}
        }else{
            set[n] = i
        }
    }

    return []int{}
}
