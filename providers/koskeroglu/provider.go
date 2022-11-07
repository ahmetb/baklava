package koskeroglu

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	fistikliBaklavaURL = "https://www.koskeroglu.com/urun/fistikli-baklava/"
	kuruBaklavaURL     = "https://www.koskeroglu.com/urun/fistikli-kuru-baklava/"
	fistikDolamaURL    = "https://www.koskeroglu.com/urun/fistikli-sarma/"
)

type KoskerogluProvider struct{}

func (k KoskerogluProvider) Name() string { return "Koskeroglu" }

func (k KoskerogluProvider) FistikliBaklava() (*money.Money, error) {
	return k.parseProductPrice(fistikliBaklavaURL)
}

func (k KoskerogluProvider) KuruBaklava() (*money.Money, error) {
	return k.parseProductPrice(kuruBaklavaURL)
}

func (k KoskerogluProvider) FistikDolama() (*money.Money, error) {
	return k.parseProductPrice(fistikDolamaURL)
}

func (k KoskerogluProvider) parseProductPrice(u string) (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`p.price`, u)

}
