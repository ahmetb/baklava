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

	t := ""
	if selector[0] != '"' {
		t = d.Find(selector).Text()
	} else {
		re := regexp.MustCompile(selector)
		html, err := d.Html()
		if err != nil {
			return nil, fmt.Errorf("failed to get HTML: %w", err)
		}
		matches := re.FindStringSubmatch(html)
		if len(matches) > 0 {
			t = matches[0]
		} else {
			return nil, fmt.Errorf("selector didn't match anything")
		}
	}

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

	if strings.Contains(t, ",") {
		if len(strings.Split(t, ",")[1]) == 3 {
			t = strings.ReplaceAll(t, ",", "") + ".00"
		} else {
			t = strings.ReplaceAll(strings.ReplaceAll(t, ".", ""), ",", ".")
		}
	} else if strings.Contains(t, ".") && len(t)-strings.LastIndex(t, ".")-1 == 3 {
		t = strings.ReplaceAll(t, ".", "")
	}

	if !strings.Contains(t, ".") {
		t += ".00"
	}

	price, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse float: %w", err)
	}

	return money.New(int64(price*100), currency), nil
}
