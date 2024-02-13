package genericparser_test

import (
	"baklava/providers/genericparser"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericParser(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test/mockTestPage.html")
	})

	// Create a new mock server
	server := httptest.NewServer(handler)
	defer server.Close()
	priceOne, err := genericparser.GenericParser{}.FromURL(`#price_one`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, priceOne.Display(), "₺1,045.00")

	priceTwo, err := genericparser.GenericParser{}.FromURL(`#price_two`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, priceTwo.Display(), "₺650.00")

	priceThree, err := genericparser.GenericParser{}.FromURL(`#price_three`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, priceThree.Display(), "₺880.00")

	priceFour, err := genericparser.GenericParser{}.FromURL(`#price_four`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, priceFour.Display(), "₺77.00")

	priceFive, err := genericparser.GenericParser{}.FromURL(`#price_five`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, priceFive.Display(), "₺480.00")

	price_six, err := genericparser.GenericParser{}.FromURL(`#price_six`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_six.Display(), "₺1,070.00", "The two price should be the same.")

	price_seven, err := genericparser.GenericParser{}.FromURL(`#price_seven`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_seven.Display(), "₺907.00", "The two price should be the same.")
}
