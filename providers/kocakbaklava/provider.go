package kocakbaklava

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const url = "https://www.kocakbaklava.com.tr/baklava-1-kg-paket"

type KocakFistikliBaklavaProvider struct{}

func (k KocakFistikliBaklavaProvider) UnitPrice() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`div#satis-fiyati`, url)
}
