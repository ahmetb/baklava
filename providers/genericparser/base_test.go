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
	price_one, err := genericparser.GenericParser{}.FromURL(`#price_one`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_one.Display(), "₺1,045.00", "The two price should be the same.")

	price_two, err := genericparser.GenericParser{}.FromURL(`#price_two`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_two.Display(), "₺650.00", "The two price should be the same.")

	price_three, err := genericparser.GenericParser{}.FromURL(`#price_three`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_three.Display(), "₺880.00", "The two price should be the same.")

	price_four, err := genericparser.GenericParser{}.FromURL(`#price_four`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_four.Display(), "₺77.00", "The two price should be the same.")

	price_five, err := genericparser.GenericParser{}.FromURL(`#price_five`, server.URL)
	assert.Nil(t, err)
	assert.Equal(t, price_five.Display(), "₺480.00", "The two price should be the same.")

}
