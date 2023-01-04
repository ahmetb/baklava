package koskeroglu

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	fistikliBaklavaURL = "https://koskeroglu.com/baklavalar/fistikli-baklava.html"
	kuruBaklavaURL     = "https://koskeroglu.com/baklavalar/fistikli-kuru-baklava.html"
	fistikDolamaURL    = "https://koskeroglu.com/baklavalar/fistikli-sarma-baklava.html"
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
	return genericparser.GenericParser{}.FromURL(`div.product-price`, u)

}
