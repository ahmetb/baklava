package main

import (
	"fmt"

	"github.com/Rhymond/go-money"

	"baklava/providers/farukgulluoglu"
	"baklava/providers/imamcagdas"
	"baklava/providers/karakoygulluoglu"
	"baklava/providers/kocakbaklava"
)

type BaklavaProvider interface {
	FistikliBaklava() (*money.Money, error)
	KuruBaklava() (*money.Money, error)
	FistikDolama() (*money.Money, error)
}

func main() {
	for _, v := range []BaklavaProvider{
		karakoygulluoglu.KarakoyGulluogluProvider{},
		farukgulluoglu.FarukGulluoglu{},
		kocakbaklava.KocakProvider{},
		imamcagdas.ImamCagdasProvider{},
	} {
		cost, err := v.FistikliBaklava()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T fistikli baklava: %s\n", v, cost.Display())
		cost, err = v.KuruBaklava()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T kuru baklava: %s\n", v, cost.Display())
		cost, err = v.FistikDolama()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T fistik dolama: %s\n", v, cost.Display())
	}
}
