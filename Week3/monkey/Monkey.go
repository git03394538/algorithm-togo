package monkey

func Do(days int, divider int, extra int, left int) int {
	if days == 1 {
		return left
	}
	return Do(days - 1 , divider, extra, divider * (left + extra))
}
