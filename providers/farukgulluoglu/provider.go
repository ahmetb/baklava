package farukgulluoglu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"

	"baklava/util"
)

const (
	farukGulluogluPriceList = "https://www.farukgulluoglu.com.tr/baklavalar?stock=1"
)

type FarukGulluoglu struct{}

type product struct {
	Code      string  `json:"code"`             // e.g. M002204
	Currency  string  `json:"currency"`         // e.g. "TL"
	BasePrice float64 `json:"total_base_price"` // e.g. 109
	URL       string  `json:"url"`              // e.g. "cevizli-baklava"
	Name      string  `json:"name"`             // e.g. "Cevizli Baklava"
}

func (f FarukGulluoglu) FistikliBaklava() (*money.Money, error) {
	return f.findItem("Fıstıklı Baklava")
}

func (f FarukGulluoglu) KuruBaklava() (*money.Money, error) {
	return f.findItem("Fıstıklı Kuru Baklava")
}

func (f FarukGulluoglu) FistikDolama() (*money.Money, error) {
	return f.findItem("Fıstıklı Dürüm")
}

func (f FarukGulluoglu) findItem(name string) (*money.Money, error) {
	resp, err := http.Get(farukGulluogluPriceList)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %w", err)
	}

	// parses example:
	// <script>PRODUCT_DATA.push(JSON.parse('{\"id\":\"25\",\"name\":\"F\\u0131st\\u0131kl\\u0131 D\\u00fcr\\u00fcm\",\"code\":\"M002213\",\"supplier_code\":\"M002213\",\"sale_price\":\"162.962\",\"total_base_price\":192.5,\"total_sale_price\":164.590000000000003410605131648480892181396484375,\"vat\":1,\"subproduct_code\":\"\",\"subproduct_id\":0,\"image\":\"https:\\/\\/www.farukgulluoglu.com.tr\\/fistikli-durum-baklavalar-105-25-O.jpg\",\"quantity\":156,\"url\":\"fistikli-durum\",\"currency\":\"TL\",\"brand\":\"\",\"category\":\"Baklavalar\",\"category_id\":\"134\",\"category_path\":\"\"}'));</script>
	re := regexp.MustCompile(`<script>\s*PRODUCT_DATA\.push\(JSON\.parse\(\s*'*(.*)'\)\s*\)\s*;<\/script>`)
	matches, err := util.ExtractGroups(re, b)
	if err != nil {
		return nil, fmt.Errorf("regex match error: %w", err)
	}
	// unescape json-quoted string
	for i, match := range matches {
		unquoted, err := strconv.Unquote(`"` + match + `"`)
		if err != nil {
			return nil, fmt.Errorf("failed to unquote group#%d: %w; input=%s", i, err, match)
		}
		var p product
		if err := json.Unmarshal([]byte(unquoted), &p); err != nil {
			return nil, fmt.Errorf("json parse error: %w", err)
		}
		if strings.EqualFold(p.Name, name) {
			c := p.Currency
			if c == "TL" {
				c = "TRY" // correct it
			}
			return money.New(int64(p.BasePrice*100), c), nil // *100 captures the decimal point
		}
	}
	return nil, fmt.Errorf("product not found")
}
