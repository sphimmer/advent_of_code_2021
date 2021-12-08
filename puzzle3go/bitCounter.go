package main

type bitCounter struct {
	zero int
	one  int
}

func (b bitCounter) highValue() string {
	if b.one > b.zero {
		return "1"
	} else {
		return "0"
	}
}

func (b bitCounter) lowValue() string {
	if b.one < b.zero {
		return "1"
	} else {
		return "0"
	}
}
