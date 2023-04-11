package task

type worker struct {
	maxWorker  int
	queuedTask chan func()
}

func NewWorker(maxWorker int) *worker {
	return &worker{
		maxWorker:  maxWorker,
		queuedTask: make(chan func()),
	}
}

func (w *worker) Run() {
	for i := 0; i < w.maxWorker; i++ {
		go func(workerID int) {
			for task := range w.queuedTask {
				task()
			}
		}(i + 1)
	}
}

func (w *worker) AddTask(task func()) {
	w.queuedTask <- task
}
