package scraper

func CalculateIterations(count int) int {
	if count <= 0 {
		return 0
	}

	iterations := count / pageSize
	// We add one to the number of iterations, because if the user passes a number that is not divisible by 50 without a remainder, then we will lose the remaining part.
	if count%pageSize != 0 {
		iterations++
	}

	return iterations
}

func CalculateCarsOnPage(pageSize int, pageIndex int, iterations int, carsCount int) int {
	if pageIndex == iterations {
		// Last page, need return remaining page size
		remainder := carsCount % pageSize
		if remainder == 0 {
			return pageSize
		}
		return remainder
	}

	return pageSize
}
