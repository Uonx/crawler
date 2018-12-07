package engine

import (
	"time"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run() //simpleSchedule用来创建workerChan  queuedSchedule用来创建workerChan和requestChan，并用select监听进行调度

	for i := 0; i < e.WorkerCount; i++ { //根据main配置的workerCount进行worker方法的创建，主要是创建多个goroutine去工作
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) //此处第三个参数虽然传的是scheduler整体，但是接收方只接收了workerReady，关注点是接口组合
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r) //将request赋给submit
	}

	for { //goroutine赋值后就必须接收,所以需要一个死循环一直接收parseResult里面的值,并将parseResult里面的requests再次提交给submit达到循环获取chan里面的item
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }() //将item传给itemChan main方法设置了itemSaver方法用来接收
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request) //获取到的request提交到submit
		}
	}
}

// 根据requestChan从worker方法里面拿到ParseResult并赋给parseResultChan（worker读取失败后进行5次的重试,主要是为了爬取时出现503后再次访问可以获取数据）
func (e *ConcurrentEngine)createWorker(in chan Request,
	out chan ParseResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in) //（只适用queuedScheduler）goroutine如果全部在工作时会阻塞，直到有空余goroutine才会进行传输，传进去后通过select进行调度
			request := <-in      //in 在simpleScheduler时是workerChan，在queuedScheduler时是workerChan
			//result, err := Retry(5, 1*time.Microsecond, func() (parseResult interface{}, err error) {
			//	return e.RequestProcessor(request)
			//})                   //请求失败后再次调用worker方法，如果五次依然没有获取到数据就直接跳出接收下一个requestChan
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// 根据Request地址获取内容并返回拿到的数据和需要下次访问的url以及获取url对应的方法

// 请求失败后重试, attempts 重试次数, sleep 重试时间, fn 重试时需要调用的方法
func Retry(attempts int, sleep time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	result, err := fn()
	if err != nil {
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return nil, err
	}
	return result, nil
}
