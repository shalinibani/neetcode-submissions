func hasDuplicate(nums []int) bool {
    setNum := make(map[int]struct{})

    for _, n := range nums{
        if _, ok := setNum[n];ok{
            return true
        }else{
            setNum[n] = struct{}{}
        }
    }

    return false
}
