package main

import (
	"fmt"

	"github.com/Rhymond/go-money"

	"baklava/providers/farukgulluoglu"
	"baklava/providers/imamcagdas"
	"baklava/providers/karakoygulluoglu"
	"baklava/providers/kocakbaklava"
)

type FistikliBaklavaProvider interface {
	UnitPrice() (*money.Money, error)
}

func main() {

	for _, v := range []FistikliBaklavaProvider{
		kocakbaklava.KocakFistikliBaklavaProvider{},
		imamcagdas.ImamCagdasFistikliBaklavaProvider{},
		karakoygulluoglu.KarakoyGulluogluFistikliBaklavaProvider{},
		farukgulluoglu.FarukGulluoglu{},
	}{
		cost, err := v.UnitPrice()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T: %s\n", v, cost.Display())
	}
}
