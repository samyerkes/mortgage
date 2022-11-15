package loan

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

var (
	AmortizationTableHeaders = table.Row{
		"Date", "Principal", "Interest", "Escrow", "Extra", "Total Payment", "Balance",
	}

	dn = DisplayNumber
)

type Loan struct {
	OriginalBalance float64 `mapstructure:"original_balance"`
	OriginalTerm    int     `mapstructure:"original_term"`
	Rate            float64 `mapstructure:"rate"`
	Escrow          float64 `mapstructure:"escrow"`
	Additional      float64 `mapstructure:"additional"`
	CurrentBalance  float64 `mapstructure:"current_balance"`
}

type ScheduleRow struct {
	month         string
	principal     float64
	interest      float64
	escrow        float64
	extra         float64
	total_payment float64
	balance       float64
}

type AmortizationLoanTotals struct {
	principle     float64
	interest      float64
	escrow        float64
	extra         float64
	total_payment float64
}

func (l *Loan) MonthlyPayment() (monthlyPayment float64) {
	term := float64(l.OriginalTerm)
	monthlyPayment = (l.OriginalBalance * l.MonthlyInterestRate() * math.Pow(1+l.MonthlyInterestRate(), term)) / (math.Pow(1+l.MonthlyInterestRate(), term) - 1)
	return
}

func (l *Loan) MonthlyInterestRate() (monthlyInterestRate float64) {
	monthlyInterestRate = l.Rate / 12
	return
}

func (l *Loan) AmortizationSchedule() (schedule []ScheduleRow) {
	balance := l.CurrentBalance
	for i := 1; i <= l.OriginalTerm; i++ {
		month := time.Now().AddDate(0, i, 0).Format("Jan 2006")
		month = fmt.Sprintf("%s (%d)", month, i)
		monthlyInterest := l.MonthlyInterestRate() * balance
		monthlyPrincipal := l.MonthlyPayment() - monthlyInterest
		extra := l.Additional
		principalExtra := monthlyPrincipal + extra
		if principalExtra > balance {
			if monthlyPrincipal > balance {
				monthlyPrincipal = balance
				extra = 0
				balance = 0
			} else {
				extra = extra - (principalExtra - balance)
				balance = 0
			}
		} else {
			balance = balance - (monthlyPrincipal + l.Additional)
		}
		totalMonthlyPayment := monthlyPrincipal + monthlyInterest + l.Escrow + extra
		s := ScheduleRow{
			month:         month,
			principal:     monthlyPrincipal,
			interest:      monthlyInterest,
			escrow:        l.Escrow,
			extra:         extra,
			total_payment: totalMonthlyPayment,
			balance:       balance,
		}
		schedule = append(schedule, s)
		if balance <= 0 {
			break
		}
	}
	return
}

func (l *Loan) AmortizationLoanTotals() (totals AmortizationLoanTotals) {
	for _, row := range l.AmortizationSchedule() {
		totals.principle += row.principal
		totals.interest += row.interest
		totals.escrow += row.escrow
		totals.extra += row.extra
		totals.total_payment += row.total_payment
	}
	return
}

func (l *Loan) PrintAmortizationSchedule() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(AmortizationTableHeaders)
	for _, row := range l.AmortizationSchedule() {
		t.AppendRow(table.Row{
			row.month,
			dn(row.principal),
			dn(row.interest),
			dn(row.escrow),
			dn(row.extra),
			dn(row.total_payment),
			dn(row.balance),
		})
	}
	totals := l.AmortizationLoanTotals()
	t.AppendFooter(table.Row{
		"",
		dn(totals.principle),
		dn(totals.interest),
		dn(totals.escrow),
		dn(totals.extra),
		dn(totals.total_payment),
		"",
	})
	t.Render()
}

func (l *Loan) PrintSummary() {
	fmt.Println("Loan:", dn(l.OriginalBalance))
	fmt.Println("Term:", l.OriginalTerm, "months")
	fmt.Println("Interest Rate: ", l.Rate)
	fmt.Println("Monthly Payment: ", dn(l.MonthlyPayment()))
	fmt.Println("Current Balance: ", dn(l.CurrentBalance))
}

func DisplayNumber(number float64) (displayNumber string) {
	n := toFixed(number, 2)
	displayNumber = fmt.Sprintf("$%.2f", n)
	return
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
