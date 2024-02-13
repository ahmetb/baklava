package genericparser

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/Rhymond/go-money"
)

type GenericParser struct{}

func (_ GenericParser) FromURL(selector, url string) (*money.Money, error) {
	cl := http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   time.Second * 5,
			ResponseHeaderTimeout: time.Second * 5,
			Proxy:                 http.ProxyFromEnvironment,
		}}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "curl/7.64.1") // with Go UA, or a real ChromeUA, karakoygulluoglu returns price in USD (wtf)
	req.Header.Set("Accept-Language", "tr-TR, tr, en-US, en")
	resp, err := cl.Do(req)
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

	currency := "TRY" // assume default
	if regexp.MustCompile(`\b(tl|try|₺)\b`).MatchString(strings.ToLower(t)) {
		currency = "TRY"
	} else if regexp.MustCompile(`\b(usd|\$)\b`).MatchString(strings.ToLower(t)) {
		currency = "USD"
	}

	// Remove all characters except digits, period, and comma
	re := regexp.MustCompile(`[^0-9.,]`)
	t = re.ReplaceAllString(t, "")

	// If the length is less than 3, return error
	if len(t) < 3 {
		return nil, fmt.Errorf("price too low: %s", t)
	}

	thousand := false

	// Check price indicates thousands
	if len(t) >= 4 && (t[1] == '.' || t[1] == ',' || t[4] == ',') {
		thousand = true
	}

	// Remove all non-digits
	re = regexp.MustCompile(`[^0-9]`)
	t = re.ReplaceAllString(t, "")

	// Trim the string
	if thousand {
		t = t[:4]
	} else {
		t = t[:3]
	}

	// Append two zeros to represent the cents part
	t += "00"

	price, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error transforming string to integer (from %v): %v", t, err)
	}

	return money.New(price, currency), nil
}
