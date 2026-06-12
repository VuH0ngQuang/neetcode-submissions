func hasDuplicate(nums []int) bool {
        var seen = make(map[int]struct{})
        for _,i := range nums {
                if _,exist := seen[i]; exist {
                        return true
                }
                seen[i]=struct{}{}
        }
        return false
}
