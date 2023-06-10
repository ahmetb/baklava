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

func (k KocakProvider) Name() string {
	return "Kocak"
}

func (k KocakProvider) FistikliBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`div#satis-fiyati`, fistikliBaklavaURL)
}

func (k KocakProvider) KuruBaklava() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`div#satis-fiyati`, kuruBaklavaURL)
}

func (k KocakProvider) FistikDolama() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`div#satis-fiyati`, fistikDolamaURL)
}
