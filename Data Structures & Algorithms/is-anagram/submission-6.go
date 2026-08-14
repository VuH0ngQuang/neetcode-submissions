func isAnagram(s string, t string) bool {
    var maps = make(map[rune]int)
    var mapt = make(map[rune]int)

    for _,ch := range s {
        maps[ch] = maps[ch] + 1
    }

    for _,ch := range t {
        mapt[ch] = mapt[ch] + 1
    }

    if len(maps) != len(mapt) {
        return false
    }

    for k,v := range maps {
        if v != mapt[k] {
            return false;
        }
    }

    return true;
}
