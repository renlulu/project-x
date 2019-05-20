package project_x

func distributeCandies(candies []int) int {
	most := len(candies) / 2
	m := make(map[int]interface{})
	for i := 0; i < len(candies); i++ {
		_, ok := m[candies[i]]
		if !ok {
			m[candies[i]] = nil
		}
	}

	if len(m) > most {
		return most
	} else {
		return len(m)
	}
}
