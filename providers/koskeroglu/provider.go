package koskeroglu

import (
	"baklava/providers/genericparser"

	"github.com/Rhymond/go-money"
)

const (
	fistikliBaklavaURL = "https://www.koskeroglu.com/fistikli-baklava-31"
	kuruBaklavaURL     = "https://www.koskeroglu.com/fistikli-kuru-baklava-42"
	fistikDolamaURL    = "https://www.koskeroglu.com/fistikli-sarma-32"
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
	price, err := genericparser.GenericParser{}.FromURL(`span.spanFiyat`, u)
	if err != nil {
		return nil, err
	}

	//We can't get 1kg price from website, so we multiply 500gr price with 2
	price = price.Multiply(2)
	return price, err
}
