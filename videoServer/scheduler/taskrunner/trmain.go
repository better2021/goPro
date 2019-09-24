package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval *time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for c = range w.ticker.C{

	}

	for {
		select {
			case <- w.ticker.C:
				go w.runner.StarAll()
		}
	}
}

func Start(){
	// Start video file cleaning
	r := NewRunner(3,true,VideoClearDispatcher)
}