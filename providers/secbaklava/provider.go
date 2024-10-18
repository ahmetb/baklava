package secbaklava

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	fistikliBaklavaURL = "https://www.secbaklava.com/urun/baklava-1-kg/"
	kuruBaklavaURL     = "https://www.secbaklava.com/urun/kuru-baklava-1-kg/"
	fistikDolamaURL    = "https://www.secbaklava.com/urun/dolama-1-kg/"
)

type SecBaklavaProvider struct{}

func (s SecBaklavaProvider) Name() string { return "Seç" }

func (s SecBaklavaProvider) FistikliBaklava() (*money.Money, error) {
	return s.parseProductPrice(fistikliBaklavaURL)
}

func (s SecBaklavaProvider) KuruBaklava() (*money.Money, error) {
	return s.parseProductPrice(kuruBaklavaURL)
}

func (s SecBaklavaProvider) FistikDolama() (*money.Money, error) {
	return s.parseProductPrice(fistikDolamaURL)
}

func (s SecBaklavaProvider) parseProductPrice(u string) (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`p.price`, u)

}
