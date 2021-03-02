package farukgulluoglu

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	farukGulluogluPriceList = "https://www.farukgulluoglu.com.tr/baklavalar?stock=1"
)

const (
	fistikliBaklavaURL = "https://www.farukgulluoglu.com.tr/fistikli-baklava-1-kg?c=1"
	kuruBaklavaURL     = "https://www.farukgulluoglu.com.tr/fistikli-kuru-baklava-1-kg"
	fistikDolamaURL    = "https://www.farukgulluoglu.com.tr/fistikli-durum-1-kg"
)

type FarukGulluogluProvider struct{}

func (k FarukGulluogluProvider) Name() string { return "FarukGulluoglu" }

func (k FarukGulluogluProvider) FistikliBaklava() (*money.Money, error) {
	return k.parseProductPrice(fistikliBaklavaURL)
}

func (k FarukGulluogluProvider) KuruBaklava() (*money.Money, error) {
	return k.parseProductPrice(kuruBaklavaURL)
}

func (k FarukGulluogluProvider) FistikDolama() (*money.Money, error) {
	return k.parseProductPrice(fistikDolamaURL)
}

func (k FarukGulluogluProvider) parseProductPrice(u string) (*money.Money, error) {
	indirimli, err := genericparser.GenericParser{}.FromURL(`#indirimli-fiyat`, u)
	if err == nil {
		return indirimli, nil
	}
	return genericparser.GenericParser{}.FromURL(`#satis-fiyati`, u)
}
