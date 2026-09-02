package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/urbaniakmichal/data-generator/internal/account"
	"github.com/urbaniakmichal/data-generator/internal/fraud"
	"github.com/urbaniakmichal/data-generator/internal/transaction"
)

func main() {
	// Common flags
	countFlag := flag.Int("count", 1, "number of events to generate")
	intervalFlag := flag.Int("interval", 0, "generate event every X milliseconds (continuous mode)")

	// Transaction flags
	transactionFlag := flag.Bool("transaction", true, "generate transaction events")
	amountFlag := flag.Float64("amount", 0, "force transaction amount (0 for random)")
	currencyFlag := flag.String("currency", "", "force currency (PLN, USD, EUR)")
	statusFlag := flag.String("status", "", "force status (PENDING, SUCCESS, etc.)")
	paymentTypeFlag := flag.String("payment-type", "", "force payment type (CARD, CASH, etc.)")

	// Account flags
	accountFlag := flag.Bool("account", true, "generate account events")
	eventTypeFlag := flag.String("event-type", "", "force event type (ACCOUNT_CREATED, etc.)")

	// Fraud flags
	fraudFlag := flag.Bool("fraud", true, "generate fraud events")
	fraudScoreFlag := flag.Int("fraud-score", -1, "force fraud score 0-100 (-1 for random)")
	countryCodeFlag := flag.String("country", "", "force country code")
	deviceTypeFlag := flag.String("device", "", "force device type")

	flag.Parse()

	// Streaming mode
	if *intervalFlag > 0 {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("Starting continuous data generation...", slog.Int("interval_ms", *intervalFlag))

		for {
			generateAndLogEvents(accountFlag, fraudFlag, transactionFlag, amountFlag, currencyFlag, statusFlag, paymentTypeFlag, eventTypeFlag, fraudScoreFlag, countryCodeFlag, deviceTypeFlag)
			time.Sleep(time.Duration(*intervalFlag) * time.Millisecond)
		}
	}

	// Batch mode
	for i := 0; i < *countFlag; i++ {
		generateAndLogEvents(accountFlag, fraudFlag, transactionFlag, amountFlag, currencyFlag, statusFlag, paymentTypeFlag, eventTypeFlag, fraudScoreFlag, countryCodeFlag, deviceTypeFlag)
	}
}

func generateAndLogEvents(
	accountFlag, fraudFlag, transactionFlag *bool,
	amountFlag *float64, currencyFlag, statusFlag, paymentTypeFlag, eventTypeFlag *string,
	fraudScoreFlag *int, countryCodeFlag, deviceTypeFlag *string,
) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if *accountFlag {
		acc := account.GenerateNewEAccountData(eventTypeFlag)
		logger.Info("account event generated", slog.Any("data", acc))
	}

	if *fraudFlag {
		var scorePtr *int
		if *fraudScoreFlag >= 0 {
			scorePtr = fraudScoreFlag
		}
		frd := fraud.GenerateNewFraudData(scorePtr, countryCodeFlag, deviceTypeFlag)
		logger.Info("fraud event generated", slog.Any("data", frd))
	}

	if *transactionFlag {
		var amountPtr *float64
		if *amountFlag > 0 {
			amountPtr = amountFlag
		}

		var currPtr *string
		if *currencyFlag != "" {
			currPtr = currencyFlag
		}

		var statusPtr *string
		if *statusFlag != "" {
			statusPtr = statusFlag
		}

		var paymentPtr *string
		if *paymentTypeFlag != "" {
			paymentPtr = paymentTypeFlag
		}

		tx := transaction.GenerateNewTransactionData(amountPtr, currPtr, statusPtr, paymentPtr)
		logger.Info("transaction event generated", slog.Any("data", tx))
	}
}
