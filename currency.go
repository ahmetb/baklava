package main

import (
	"encoding/xml"
	"net/http"
)

type envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Gesmes  string   `xml:"gesmes,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Subject string   `xml:"subject"`
	Sender  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube struct {
		Text string `xml:",chardata"`
		Cube struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
			Cube []struct {
				Text     string  `xml:",chardata"`
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

var rates map[string]float64

func GetRates() (map[string]float64, error) {
	if len(rates) > 0 {
		return rates, nil
	}
	resp, err := http.Get(`https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml`)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var p envelope
	if err := xml.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}
	out := make(map[string]float64)
	for _, v := range p.Cube.Cube.Cube {
		out[v.Currency] = v.Rate
	}
	return out, nil
}
