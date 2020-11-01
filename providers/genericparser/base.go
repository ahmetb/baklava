package genericparser

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/Rhymond/go-money"
)

type GenericParser struct{}

func (_ GenericParser) FromURL(selector, url string) (*money.Money, error) {
	resp, err := http.Get(url)
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
	t := d.Find(selector).Text()
	if t == "" {
		return nil, fmt.Errorf("selector didn't match anything")
	}
	t = strings.TrimFunc(t, unicode.IsSpace)
	t = strings.ReplaceAll(t, "\n", "")

	currency := "TRY" // assume default
	if strings.HasSuffix(t, "TL") {
		currency = "TRY"
		t = strings.TrimSuffix(t, "TL")
	} else if strings.HasSuffix(t, "USD") {
		currency = "USD"
		t = strings.TrimSuffix(t, "USD")
	}

	re := regexp.MustCompile(`[.*:\s*]?(\d+)[,\.](\d+)\s*$`)
	if !re.MatchString(t) {
		return nil, fmt.Errorf("string doesn't match format for parsing: %s (%s)", t, re)
	}
	groups := re.FindStringSubmatch(t)
	dec, frac := groups[1], groups[2]
	iDec, _ := strconv.ParseInt(dec, 10, 64)
	iFrac, _ := strconv.ParseInt(frac, 10, 64)
	return money.New(iDec*100+iFrac, currency), nil
}
