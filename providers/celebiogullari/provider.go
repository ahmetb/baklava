package celebiogullari

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = "https://www.celebiogullari.com.tr/celebiogullari-baklava-1-kg-paket"
	kuruBaklavaURL     = `https://www.celebiogullari.com.tr/1-kg-paket`
	fistikDolamaURL    = `https://www.celebiogullari.com.tr/celebiogullari-fistik-sarma-dolama-1-kg-paket`
)

type CelebiogullariProvider struct{}

func (k CelebiogullariProvider) Name() string {
	return "Celebiogullari"
}

func (k CelebiogullariProvider) FistikliBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`#satis`, fistikliBaklavaURL)
}

func (k CelebiogullariProvider) KuruBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`#satis`, kuruBaklavaURL)
}

func (k CelebiogullariProvider) FistikDolama() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`#satis`, fistikDolamaURL)
}
