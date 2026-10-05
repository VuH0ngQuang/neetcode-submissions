import "slices"

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, s:= range strs {
		sortedS := sortString(s)
		res[sortedS] = append(res[sortedS], s)
	}

	result := make([][]string, 0, len(res))

	for _, group := range res {
		result = append (result, group)
	}

	return result
}

func sortString(str string) string {
	r := []rune(str)
	slices.Sort(r)
	return string(r)
}