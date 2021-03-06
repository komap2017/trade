package main

import (
	"io"

	"github.com/corpix/loggers/logger/logrus"
	"github.com/davecgh/go-spew/spew"
	logrusLogger "github.com/sirupsen/logrus"

	"github.com/cryptounicorns/trade/assets"
	"github.com/cryptounicorns/trade/currencies"
	"github.com/cryptounicorns/trade/markets"
	"github.com/cryptounicorns/trade/markets/market/bitfinex"
	"github.com/cryptounicorns/trade/markets/ticker"
	// XXX: Import any other market
)

func configuredLogrus() *logrusLogger.Logger {
	var (
		l = logrusLogger.New()
	)

	l.Level = logrusLogger.DebugLevel

	return l
}

func main() {
	var (
		connection io.ReadWriteCloser
		consumer   ticker.Consumer
		tickers    <-chan ticker.Result

		loader           = currencies.NewAssetLoader(assets.Asset)
		commonCurrencies currencies.Currencies
		marketCurrencies currencies.Currencies
		mapper           currencies.Mapper
		market           markets.Market

		bitcoin currencies.Currency
		dollar  currencies.Currency

		log = logrus.New(configuredLogrus())

		err error
	)

	commonCurrencies, err = loader.Common()
	if err != nil {
		panic(err)
	}
	marketCurrencies, err = loader.Market(bitfinex.Name)
	if err != nil {
		panic(err)
	}

	mapper = currencies.NewMapper(
		commonCurrencies,
		marketCurrencies,
	)

	market = bitfinex.New(
		bitfinex.DefaultConfig,
		mapper,
		log,
	)

	bitcoin, err = mapper.CommonByName("bitcoin")
	if err != nil {
		panic(err)
	}

	dollar, err = mapper.CommonByName("united-states-dollar")
	if err != nil {
		panic(err)
	}

	connection, err = market.Connect()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	consumer, err = market.NewTickerConsumer(connection)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	tickers, err = consumer.Consume(
		[]currencies.CurrencyPair{
			currencies.NewCurrencyPair(
				bitcoin,
				dollar,
			),
		},
	)
	if err != nil {
		panic(err)
	}

	for ticker := range tickers {
		spew.Dump(ticker)

		if ticker.Err != nil {
			panic(ticker.Err)
		}
	}
}
