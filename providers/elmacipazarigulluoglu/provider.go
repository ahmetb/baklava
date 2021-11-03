package elmacipazarigulluoglu

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	farukGulluogluPriceList = "https://www.farukgulluoglu.com.tr/baklavalar?stock=1"
)

const (
	fistikliBaklavaURL = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=7"
	kuruBaklavaURL     = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=3"
	fistikDolamaURL    = "https://www.elmacipazarigulluoglu.com/Urun.aspx?pID=13"
)

type ElmacipazariGulluogluProvider struct{}

func (k ElmacipazariGulluogluProvider) Name() string { return "ElmacipazariGulluoglu" }

func (k ElmacipazariGulluogluProvider) FistikliBaklava() (*money.Money, error) {
	return k.parseProductPrice(fistikliBaklavaURL)
}

func (k ElmacipazariGulluogluProvider) KuruBaklava() (*money.Money, error) {
	return k.parseProductPrice(kuruBaklavaURL)
}

func (k ElmacipazariGulluogluProvider) FistikDolama() (*money.Money, error) {
	return k.parseProductPrice(fistikDolamaURL)
}

func (k ElmacipazariGulluogluProvider) parseProductPrice(u string) (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`table > tbody > tr:nth-child(4) > td:nth-child(3) > span > b`, u)
}
