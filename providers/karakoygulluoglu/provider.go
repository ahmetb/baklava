package karakoygulluoglu

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = "https://www.karakoygulluoglu.com/fistikli-baklava"
	kuruBaklavaURL     = "https://www.karakoygulluoglu.com/fistikli-kuru-baklava"
	fistikDolamaURL    = "https://www.karakoygulluoglu.com/fistikli-durum"
)

type KarakoyGulluogluProvider struct{}

func (k KarakoyGulluogluProvider) Name() string { return "KarakoyGulluoglu" }

func (k KarakoyGulluogluProvider) FistikliBaklava() (*money.Money, error) {
	return k.parseProductPrice(fistikliBaklavaURL)
}

func (k KarakoyGulluogluProvider) KuruBaklava() (*money.Money, error) {
	return k.parseProductPrice(kuruBaklavaURL)
}

func (k KarakoyGulluogluProvider) FistikDolama() (*money.Money, error) {
	return k.parseProductPrice(fistikDolamaURL)
}

func (k KarakoyGulluogluProvider) parseProductPrice(u string) (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`.spanFiyat`, u)
}
