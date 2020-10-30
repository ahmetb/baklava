package karakoygulluoglu

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Rhymond/go-money"
)

const url = "https://www.karakoygulluoglu.com/baklava-with-pistachio-en"

type KarakoyGulluogluFistikliBaklavaProvider struct{}

func (k KarakoyGulluogluFistikliBaklavaProvider) UnitPrice() (*money.Money, error) {
	return k.parseProductPrice(url)
}
	func (k KarakoyGulluogluFistikliBaklavaProvider) parseProductPrice(u string) (*money.Money, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	d, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse doc: %w", err)
	}
	t := d.Find(`div.mainPrices`).First().Text()
	if t == "" {
		return nil, fmt.Errorf("selector didn't match anything")
	}
	t = strings.TrimSpace(t)
	t = strings.ReplaceAll(t, "\n", "")
	if !strings.HasPrefix(t, "Price") {
		return nil, fmt.Errorf("unexpected prefix in string: %s", t)
	}

	t = regexp.MustCompile(`^Price\s*:\s*`).ReplaceAllString(t, "")

	re := regexp.MustCompile(`^(\d+)[\.,](\d)+\s+([A-Z]{2,3})`)
	if !re.MatchString(t) {
		return nil, fmt.Errorf("string doesn't match format for parsing: %s (%s)", t, re)
	}
	groups := re.FindStringSubmatch(t)
	dec, frac, currency := groups[1], groups[2], groups[3]
	iDec, _ := strconv.ParseInt(dec, 10, 64)
	iFrac, _ := strconv.ParseInt(frac, 10, 64)
	if currency == "TL" {
		currency = "TRY"
	}
	return money.New(iDec*100+iFrac, currency),nil
}
