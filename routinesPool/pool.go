package routinesPool

// 任务类型
type Task struct {
	fun func() error
}

func NewTask(argcFun func() error) *Task {
	return &Task{
		fun: argcFun,
	}
}

func (t *Task) Execute() error {
	return t.fun()
}

// 协程池
type Pool struct {
	// 对外Task入口
	EntryChannel chan *Task

	// 内部任务队列
	JobChannel chan *Task

	// 协程池中worker的数量
	WorkNum int
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		JobChannel:   make(chan *Task),
		WorkNum:      cap,
	}
}

// 创建一个worker去执行Task
func (p *Pool) work() {
	for task := range p.JobChannel {
		_ = task.Execute()
	}
}

func (p *Pool) AddTask(task *Task) {
	p.EntryChannel <- task
}

// 开启协程池工作
func (p *Pool) Run() {
	for i := 0; i < p.WorkNum; i++ {
		go p.work()
	}
	for task := range p.EntryChannel {
		p.JobChannel <- task
	}
}
