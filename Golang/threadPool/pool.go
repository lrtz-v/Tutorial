package pool

import "errors"

// Task 代表执行任务
type Task struct {
	id int
	// task result
	result int

	ans chan bool

	// todo lock
}

func (task Task) run() {
	// go run task
	// set result
	task.result = 1 // example
	task.ans <- true
}

// Pool 结构定义
type Pool struct {
	MaxWorker   int    // 最大worker数量
	CurWorkerNum int   // 当前Worker数量
	Queue       []Task // 工作队列
	TaskNum     int    // 当前任务数量
	QueueLength int    // 工作队列长度
	// todo lock
}

// NewPool create new pool
func NewPool(num, length int) *Pool {
	pool := Pool{}
	pool.MaxWorker = num
	pool.Queue = make([]Task, length)
	pool.TaskNum = 0
	pool.QueueLength = length
	pool.CurWorkerNum = 0

	go pool.WorkerRun()
	return &pool
}

// Submit 提交任务
func (pool *Pool) Submit(task Task) (error) {
	// 1.判断是否可以新增worker
	if pool.CurWorkerNum < pool.MaxWorker {
		pool.AddWorker(task)
	} else if pool.TaskNum < pool.QueueLength {
		// 2.提交到任务队列
		pool.Queue = append(pool.Queue, task)
		pool.TaskNum++
	} else {
		// 3.rejec
		return errors.New("Submit task failed")
	}
	return nil
}

func (pool *Pool) runTask(task Task) {
	// go run task
	task.run()
	pool.TaskNum--
	pool.CurWorkerNum--
}

// AddWorker 增加任务处理协程
func (pool *Pool) AddWorker(task Task) {
	pool.CurWorkerNum++
	go pool.runTask(task)
}

// WorkerRun 异步处理工作队列
func (pool *Pool) WorkerRun() {
	// get task
	for {
		if pool.TaskNum == 0 {
			continue
		} else {
			task := pool.getNewTask()
			// 执行任务
			pool.AddWorker(task)
		}
	}
	// run task
}

// getNewTask 获取待执行任务
func (pool *Pool) getNewTask() Task {
	// get lock
	// return 任务队列的一个任务
	// release lock
}
