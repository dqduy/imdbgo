package imdbgo

import "fmt"

type IMDb struct {
	Name string
}

func NewIMDb() *IMDb {
	return &IMDb{""}
}

func (o *IMDb) Hello() {
	fmt.Println("Hello")
}
