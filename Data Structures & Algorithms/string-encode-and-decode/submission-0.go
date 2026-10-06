type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	var before, str string
	for len(encoded) > 0 {
		before, encoded, _ = strings.Cut(encoded, "#")
		num, _ := strconv.Atoi(before)
		str, encoded = encoded[:num], encoded[num:]
		res = append(res, str)
	}
	return res
}
