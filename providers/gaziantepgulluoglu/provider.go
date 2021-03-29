package gaziantepgulluoglu

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=7"
	kuruBaklavaURL     = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=3"
	fistikDolamaURL    = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=13"
)

type GaziantepGulluogluProvider struct{}

func (k GaziantepGulluogluProvider) Name() string { return "GaziantepGulluoglu" }

func (k GaziantepGulluogluProvider) FistikliBaklava() (*money.Money, error) {
	return k.parseProductPrice(fistikliBaklavaURL)
}

func (k GaziantepGulluogluProvider) KuruBaklava() (*money.Money, error) {
	return k.parseProductPrice(kuruBaklavaURL)
}

func (k GaziantepGulluogluProvider) FistikDolama() (*money.Money, error) {
	return k.parseProductPrice(fistikDolamaURL)
}

func (k GaziantepGulluogluProvider) parseProductPrice(u string) (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`table > tbody > tr:nth-child(4) > td:nth-child(3) > span > b`, u)
}
