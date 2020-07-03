package repository

func filterLimit(inputLimit *int) (limit int) {
	if inputLimit == nil {
		limit = 10
	} else if *inputLimit > 10000 {
		limit = 10000
	} else {
		limit = *inputLimit
	}

	return
}
