package scheduler

import "learn/crawler/engine"

// Scheduler实现结构并定义结构体需要的变量
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(r chan engine.Request) {
}

//在main方法设置的workerCount数量全部用完时就会出现阻塞现象,所以需要新建一个goroutine用来接收request
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

//返回workerChan
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

//创建workerChan
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
