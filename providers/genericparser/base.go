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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "curl/7.64.1") // with Go UA, or a real ChromeUA, karakoygulluoglu returns price in USD (wtf)
	req.Header.Set("Accept-Language", "tr-TR, tr, en-US, en")
	resp, err := http.DefaultClient.Do(req)
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

	re := regexp.MustCompile(`[.*:\s*]?(\d+)[,\.]?(\d+)?\s*$`)
	if !re.MatchString(t) {
		return nil, fmt.Errorf("string doesn't match format for parsing: %s (%s)", t, re)
	}
	groups := re.FindStringSubmatch(t)
	dec, frac := groups[1], groups[2]
	iDec, err := strconv.ParseInt(dec, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing decimal value (from %v): %v", dec, err)
	}
	iFrac, err := strconv.ParseInt(frac, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing fractional value (from %v): %v", frac, err)
	}
	return money.New(iDec*100+iFrac, currency), nil
}
