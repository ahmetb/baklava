package koskeroglu

import (
	"baklava/providers/genericparser"
	"github.com/Rhymond/go-money"
	"net/url"
)

const (
	fistikliBaklavaURL = "https://koskeroglu.com/urun/fistikli-baklava/"
	kuruBaklavaURL     = "https://koskeroglu.com/urun/kuru-baklava/"
	fistikDolamaURL    = "https://koskeroglu.com/urun/sarma/"
	rendererProxyURL   = "https://renderer-proxy-sc2owh7ynq-uc.a.run.app/?selector=span.amount&url="
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
	return genericparser.GenericParser{}.FromURL(`bdi`, rendererProxyURL + url.QueryEscape(u))
}
