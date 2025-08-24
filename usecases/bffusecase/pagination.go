package bffusecase

// totalToPages - конвертирует количество данных в количество страниц
func totalToPages(total, onPageCount int) (pageCount int) {
	pageCount = total / onPageCount

	if total%onPageCount > 0 {
		pageCount++
	}

	return
}

// generatePagination - генерирует последовательность страниц
func generatePagination(currentPage, pageCount int) []int {
	const (
		leftPageCount   = 3
		middlePageCount = 5
		rightPageCount  = 3
		maxLength       = leftPageCount + middlePageCount + rightPageCount + 2
		separator       = -1
	)

	pages := make([]int, 0, maxLength)

	// Если страниц меньше, чем максимально отображаемых, то отображаем их все
	if pageCount <= maxLength {
		for i := range pageCount {
			pages = append(pages, i+1)
		}

		return pages
	}

	// Генерация первой тройки страниц
	for i := range leftPageCount {
		pages = append(pages, i+1)
	}

	// Генерация центральных страниц
	switch {
	case currentPage < leftPageCount+middlePageCount:
		// Заполняем страницы
		for i := leftPageCount; i < leftPageCount+middlePageCount; i++ {
			pages = append(pages, i+1)
		}
		// Добавляем разделитель
		pages = append(pages, separator)
	case currentPage > pageCount+1-middlePageCount-rightPageCount:
		// Добавляем разделитель
		pages = append(pages, separator)
		// Заполняем страницы
		for i := pageCount - middlePageCount - rightPageCount; i < pageCount-rightPageCount; i++ {
			pages = append(pages, i+1)
		}
	default:
		// Добавляем разделитель
		pages = append(pages, separator)

		startIndex := currentPage - middlePageCount/2

		// Для нечетного числа страниц требуется дополнительный сдвиг
		if middlePageCount%2 > 0 {
			startIndex--
		}

		// Заполняем страницы
		for i := startIndex; i < startIndex+middlePageCount; i++ {
			pages = append(pages, i+1)
		}

		// Добавляем разделитель
		pages = append(pages, separator)
	}

	// Генерация последней тройки страниц
	for i := pageCount - rightPageCount; i < pageCount; i++ {
		pages = append(pages, i+1)
	}

	return pages
}
