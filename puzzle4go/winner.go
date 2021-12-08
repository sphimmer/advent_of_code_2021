package main

type winner struct {
	id            int
	Board         board
	lastCalledNum string
}

func winnersContain(winners []winner, id int) bool {
	result := false
	for _, w := range winners {
		if w.id == id {
			result = true
			break
		}
	}
	return result
}
