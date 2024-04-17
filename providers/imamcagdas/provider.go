package imamcagdas

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = `https://www.imamcagdas.com/normal-baklava-27`
	kuruBaklavaURL     = `https://www.imamcagdas.com/fistikli-kuru-baklava`
	fistikDolamaURL    = "https://www.imamcagdas.com/fistik-dolama-1"
)

type ImamCagdasProvider struct{}

func (i ImamCagdasProvider) Name() string {
	return "ImamCagdas"
}

func (i ImamCagdasProvider) FistikliBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`.product-price`, fistikliBaklavaURL)
}

func (i ImamCagdasProvider) FistikDolama() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`.product-price`, fistikDolamaURL)
}

func (i ImamCagdasProvider) KuruBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`.product-price`, kuruBaklavaURL)
}
