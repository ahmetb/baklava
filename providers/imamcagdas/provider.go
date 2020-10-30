package imamcagdas

import (
	"github.com/Rhymond/go-money"

	"baklava/providers/genericparser"
)

const url = `https://www.imamcagdas.com/normal-baklava-27`

type ImamCagdasFistikliBaklavaProvider struct{}

func (i ImamCagdasFistikliBaklavaProvider) UnitPrice() (*money.Money, error) {
	return genericparser.GenericParser{}.FromURL(`div.mainPrices`, url)
}
