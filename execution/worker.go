package execution

import kiteconnect "github.com/zerodhatech/gokiteconnect"

type Worker struct {
	WorkerJob Job
	Client    *kiteconnect.Client
}

func (w *Worker) Work() (kiteconnect.OrderResponse, error) {

	return w.Client.PlaceOrder("regular", kiteconnect.OrderParams{
		Exchange:        w.WorkerJob.Exchange,
		Tradingsymbol:   w.WorkerJob.Instrument,
		OrderType:       "LIMIT",
		Product:         "MIS",
		Validity:        "DAY",
		TransactionType: w.WorkerJob.Type,
		Quantity:        w.WorkerJob.Quantity,
		Price:           w.WorkerJob.Price,
	})
}
