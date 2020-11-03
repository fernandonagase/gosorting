package sorting

func gap(size int) int {
	gap := 1
	for 3 * gap + 1 <= size {
		gap = 3 * gap + 1
	}
	return gap
}