package selenium

type GsfSelenium struct {
	taskWatch *TaskWatch
}

func (gs *GsfSelenium) Run() {
	gs.taskWatch.Watch()
}

func (gs *GsfSelenium) Token() {
	gs.taskWatch.Add("token", GetToken)
}

func NewGsfSelenium() *GsfSelenium {
	return &GsfSelenium{
		taskWatch: NewTaskWatch(5),
	}
}
