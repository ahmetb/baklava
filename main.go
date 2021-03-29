package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rhymond/go-money"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"baklava/providers/farukgulluoglu"
	"baklava/providers/gaziantepgulluoglu"
	"baklava/providers/imamcagdas"
	"baklava/providers/karakoygulluoglu"
	"baklava/providers/kocakbaklava"
	"baklava/util"
)

type BaklavaProvider interface {
	Name() string
	FistikliBaklava() (*money.Money, error)
	KuruBaklava() (*money.Money, error)
	FistikDolama() (*money.Money, error)
}

var (
	flCLI       bool
	flNoUpdate  bool
	flSheetID   string
	flSheetName string
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	flag.BoolVar(&flCLI, "cli", false, "run as a one of tool")
	flag.BoolVar(&flNoUpdate, "no_update", false, "do not update google sheets")
	flag.StringVar(&flSheetID, "sheet_id", "", "google sheets ID (also $SHEET_ID)")
	flag.StringVar(&flSheetName, "sheet_name", "Sheet1", "google sheet spreadsheet name")
}

func main() {
	flag.Parse()
	if flCLI {
		log.Println("running one-off as cli tool")
		if err := run(); err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	if v := os.Getenv("SHEET_ID"); v != "" {
		flSheetID = v
	}
	if !flNoUpdate && flSheetID == "" {
		log.Fatal("SHEET_ID is empty")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT environment variable required if not running as -cli")
	}
	http.HandleFunc("/health", func(_ http.ResponseWriter, _ *http.Request) {})
	http.HandleFunc("/run", func(rw http.ResponseWriter, _ *http.Request) {
		if err := run(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(io.MultiWriter(os.Stderr, rw), "error: %v\n", err)
			return
		}
		fmt.Fprintf(rw, "ok")
	})
	log.Printf("server starting at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func run() error {
	rates, err := GetRates()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		return fmt.Errorf("token initialization error: %w", err)
	}
	sheet, err := sheets.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return fmt.Errorf("failed to init sheets client: %w", err)
	}
	values := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         [][]interface{}{},
	}
	y, m, d := time.Now().Date()
	date := fmt.Sprintf("%d/%d/%d", m, d, y)

	exchangeRate := rates["TRY"] / rates["USD"]
	log.Printf("USDTRY exchange rate is: %.2f", exchangeRate)
	addRow := func(provider BaklavaProvider, cost *money.Money, product string) {
		costTRY := toTRY(cost, rates)
		costUSD := toUSD(cost, rates)
		v := []interface{}{
			date, provider.Name(), product,
			fmt.Sprintf("%.2f", float64(costTRY.Amount())/100.0),
			fmt.Sprintf("%.2f", float64(costUSD.Amount())/100.0),
			fmt.Sprintf("%.2f", exchangeRate),
		}
		values.Values = append(values.Values, v)
		log.Printf("%#v", v)
	}

	var errs util.ErrGroup
	for _, v := range []BaklavaProvider{
		karakoygulluoglu.KarakoyGulluogluProvider{},
		farukgulluoglu.FarukGulluogluProvider{},
		kocakbaklava.KocakProvider{},
		imamcagdas.ImamCagdasProvider{},
		gaziantepgulluoglu.GaziantepGulluogluProvider{},
	} {
		cost, err := v.FistikliBaklava()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get price (%T): %w", v, err))
		} else {
			log.Printf("%T fistikli baklava: %s\n", v, cost.Display())
			addRow(v, cost, "fistikli_baklava")
		}

		cost, err = v.KuruBaklava()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get price (%T): %w", v, err))
		} else {
			log.Printf("%T kuru_baklava: %s", v, cost.Display())
			addRow(v, cost, "kuru_baklava")
		}

		cost, err = v.FistikDolama()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get price (%T): %w", v, err))
		} else {
			log.Printf("%T fistik dolama: %s", v, cost.Display())
			addRow(v, cost, "fistik_dolama")
		}
	}

	if !flNoUpdate {
		if _, err := sheet.Spreadsheets.Values.Append(flSheetID, flSheetName, values).
			ValueInputOption("USER_ENTERED").Do(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
