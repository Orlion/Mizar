package utils

type SortByLength []string

func (s SortByLength) Len() int {
	return len(s)
}
func (s SortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
