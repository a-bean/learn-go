package search

func Linear(array []int, query int) (int, error) {
	for i, item := range array {
		if item == query {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
