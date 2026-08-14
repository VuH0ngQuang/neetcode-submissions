func hasDuplicate(nums []int) bool {
    tmp := make(map[int]struct{})

    for _, n := range nums {
        if _, ok := tmp[n]; ok {
            return true
        }
        tmp[n] = struct{}{}
    }
    return false
}