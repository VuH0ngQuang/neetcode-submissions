func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var setS = make(map[rune]int, len(s))
    var setT = make(map[rune]int, len(t))

    for _, r := range s {
        setS[r] = setS[r]  + 1;
    }
    for _, r := range t {
        setT[r] = setT[r]  + 1;
    }

    for _, r := range s {
        if _, exist := setT[r]; !exist {
            return false;
        }
        setS[r] = setS[r] - 1
        setT[r] = setT[r] - 1
        if setT[r] == 0 {
            delete(setT, r)
        }
        if setT[r] == 0 {
            delete(setT, r)
        }
    }
    return true
}
