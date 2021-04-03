package strategy

import (
	"autobot/execution"
	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

func StrategyOne(Client *kiteconnect.Client) {
	JobChannel := make(chan execution.Job, 2)
	ResultsChannel := make(chan kiteconnect.OrderResponse, 2)

	ex := execution.Executor{
		Client:         Client,
		JobsChannel:    JobChannel,
		ResultsChannel: ResultsChannel,
	}

	//Two workers.
	go ex.ExecuteTrades()
	go ex.ExecuteTrades()

	job1 := execution.Job{
		Instrument: "TCS",
		Price:      3167,
		Quantity:   1,
		Exchange:   "NSE",
	}

	job2 := execution.Job{
		Instrument: "TCS",
		Price:      3167,
		Quantity:   1,
		Exchange:   "NSE",
	}

	ex.JobsChannel <- job1
	ex.JobsChannel <- job2

	for i := 0; i < 2; i++ {
		<-ex.ResultsChannel
	}

}
