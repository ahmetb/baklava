package gaziantepgulluoglu

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = "https://www.elmacipazarigulluoglu.com/fistikli-yas-baklava"
	kuruBaklavaURL     = "https://www.elmacipazarigulluoglu.com/fistikli-kuru-baklava"
	fistikDolamaURL    = "https://www.elmacipazarigulluoglu.com/fistikli-dolama-sarma"
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
	return genericparser.GenericParser{}.FromURL(`#lblUrunFiyatiKDVDahil`, u)
}
