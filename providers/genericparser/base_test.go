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
	assert.Equal(t, "₺1,045.00", priceOne.Display(), "The two price should be the same.")

	priceTwo, err := genericparser.GenericParser{}.FromURL(`#price_two`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺650.00", priceTwo.Display(), "The two price should be the same.")

	priceThree, err := genericparser.GenericParser{}.FromURL(`#price_three`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺880.00", priceThree.Display(), "The two price should be the same.")

	priceFour, err := genericparser.GenericParser{}.FromURL(`#price_four`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺77.00", priceFour.Display(), "The two price should be the same.")

	priceFive, err := genericparser.GenericParser{}.FromURL(`#price_five`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺480.00", priceFive.Display(), "The two price should be the same.")

	price_six, err := genericparser.GenericParser{}.FromURL(`#price_six`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺1,070.00", price_six.Display(), "The two price should be the same.")

	price_seven, err := genericparser.GenericParser{}.FromURL(`#price_seven`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺907.00", price_seven.Display(), "The two price should be the same.")

	price_eight, err := genericparser.GenericParser{}.FromURL(`#price_eight`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, "₺1,310.00", price_eight.Display(), "The two price should be the same.")
}
