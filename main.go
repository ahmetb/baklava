package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"baklava/providers/farukgulluoglu"
	"baklava/providers/imamcagdas"
	"baklava/providers/karakoygulluoglu"
	"baklava/providers/kocakbaklava"
)

type BaklavaProvider interface {
	FistikliBaklava() (*money.Money, error)
	KuruBaklava() (*money.Money, error)
	FistikDolama() (*money.Money, error)
}

func main() {
	rates, err := GetRates()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		panic(err)
	}
	sheet, err := sheets.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		panic(err)
	}
	values := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         [][]interface{}{},
	}
	sheetID := "16ZsPZED1ovC_PreauREKA0gbhCKLAPq_giY_X3L7YdA"
	y, m, d := time.Now().Date()
	date := fmt.Sprintf("%d/%d/%d", m, d, y)
	for _, v := range []BaklavaProvider{
		karakoygulluoglu.KarakoyGulluogluProvider{},
		farukgulluoglu.FarukGulluoglu{},
		kocakbaklava.KocakProvider{},
		imamcagdas.ImamCagdasProvider{},
	} {
		cost, err := v.FistikliBaklava()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T fistikli baklava: %s\n", v, cost.Display())
		values.Values = append(values.Values, []interface{}{
			date,
			fmt.Sprintf("%T", v),
			"fistikli_baklava",
			float64(cost.Amount()) / 100.0,
			try2usd(rates, float64(cost.Amount())/100.0),
		})

		cost, err = v.KuruBaklava()
		if err != nil {
			panic(err)
		}
		values.Values = append(values.Values, []interface{}{
			date,
			fmt.Sprintf("%T", v),
			"kuru_baklava",
			float64(cost.Amount()) / 100.0,
			try2usd(rates, float64(cost.Amount())/100.0),
		})
		fmt.Printf("%T kuru baklava: %s\n", v, cost.Display())

		cost, err = v.FistikDolama()
		if err != nil {
			panic(err)
		}
		values.Values = append(values.Values, []interface{}{
			date,
			fmt.Sprintf("%T", v),
			"fistik_dolama",
			float64(cost.Amount()) / 100.0,
			try2usd(rates, float64(cost.Amount())/100.0),
		})
		fmt.Printf("%T fistik dolama: %s\n", v, cost.Display())
	}

	if _, err := sheet.Spreadsheets.Values.Append(sheetID, "Sheet1", values).
		ValueInputOption("USER_ENTERED").Do(); err != nil {
		panic(err)
	}
}
