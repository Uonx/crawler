package scheduler

import (
	"learn/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

//创建requestChan
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request) //创建workChan
	s.requestChan = make(chan engine.Request)     //创建requestChan
	go func() {
		var requestQ []engine.Request     //声明request队列
		var workerQ []chan engine.Request //声明requestChan队列
		for {
			var activeRequest engine.Request           //用来读取request队列内容
			var activeWorker chan engine.Request       //用来读取requestChan队列内容
			if len(requestQ) > 0 && len(workerQ) > 0 { //当request队列和requestChan队列有数据时将request队列第一条数据放到activeRequest,requestChan队列第一条数据放到activeWorker
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan: //接收request数据并加入到request队列
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: //接收requestChan数据并加入到requestChan队列
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //当activeRequest数据写入到activeWorker时去掉request队列第一条数据和requestChan队列第一条数据,就会继续执行获取数据
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

//将 requestChan写入workerChan
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//将request写入requestChan
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
