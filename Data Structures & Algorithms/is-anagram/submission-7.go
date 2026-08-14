func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make(map[rune]int)

    for _, ch := range s {
        count[ch]++
    }
    for _, ch := range t {
        count[ch]--
    }

    for _, v := range count {
        if v != 0 {
            return false
        }
    }
    return true
}