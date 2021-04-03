package execution

import (
	kiteconnect "github.com/zerodhatech/gokiteconnect"
	"log"
)

type Order string
type Type string
type Exchange string

const (
	BUY  Order = "BUY"
	SELL Order = "SELL"
)

const (
	CNC Type = "CNC"
	MIS Type = "MIS"
)

const (
	NSE Exchange = "NSE"
	BSE Exchange = "BSE"
)

type Job struct {
	Instrument string
	Price      float64
	Quantity   int
	Type       string
	Exchange   string
}

type Executor struct {
	Client         *kiteconnect.Client
	JobsChannel    chan Job
	ResultsChannel chan kiteconnect.OrderResponse
}

func (e *Executor) ExecuteTrades() {

	for Trade := range e.JobsChannel {
		worker := Worker{
			WorkerJob: Trade,
			Client:    e.Client,
		}

		response, err := worker.Work()
		if err != nil {
			log.Print(err)
		}

		e.ResultsChannel <- response
	}
}
