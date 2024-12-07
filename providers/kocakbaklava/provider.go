package kocakbaklava

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const (
	fistikliBaklavaURL = "https://www.kocakbaklava.com.tr/tr/baklava-1-kg-paket"
	kuruBaklavaURL     = "https://www.kocakbaklava.com.tr/tr/kuru-baklava-1-kg-paket"
	fistikDolamaURL    = "https://www.kocakbaklava.com.tr/tr/dolama-1-kg-paket"
)

type KocakProvider struct{}

func (s KocakProvider) Name() string { return "Kocak" }

func (s KocakProvider) FistikliBaklava() (*money.Money, error) {
	return s.parseProductPrice(fistikliBaklavaURL)
}

func (s KocakProvider) KuruBaklava() (*money.Money, error) {
	return s.parseProductPrice(kuruBaklavaURL)
}

func (s KocakProvider) FistikDolama() (*money.Money, error) {
	return s.parseProductPrice(fistikDolamaURL)
}

func (s KocakProvider) parseProductPrice(u string) (*money.Money, error) {
	indirimli, err := genericparser.GenericParser{}.FromURL(`#indirimli-fiyat`, u)
	if err == nil {
		return indirimli, nil
	}

	return genericparser.GenericParser{}.FromURL(`div#salePrice`, u)
}
