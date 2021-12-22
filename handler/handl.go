package handler

import (
	"errors"
	//"text/template"
)

type Data struct {
	NumFib  []int64
	FibPost []int
}

func NewData() Data {
	return Data{
		NumFib:  nil,
		FibPost: nil,
	}
}

func (d *Data) Fib(x int, y int) (Data, error) {
	var table []int64

	if err := d.validFib(x, y); err != true {
		return Data{}, errors.New("invalid param")
	}
	table2 := make([]int, 0, x+1)
	table = make([]int64, 0, x+1)

	table2 = append(table2, 0, 1)
	table = append(table, 0, 1)
	for i := 2; i <= x; i += 1 {
		table2 = append(table2, i)
		table = append(table, table[i-1]+table[i-2])
	}
	table2 = table2[y : x+1]
	table = table[y : x+1]
	res := Data{NumFib: table, FibPost: table2}
	return res, nil
}

func (d *Data) validFib(lst int, frst int) bool {
	if lst < 0 || frst < 0 || lst > 92 || frst > lst {
		return false
	}
	return true
}
