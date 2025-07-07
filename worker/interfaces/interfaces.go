package interfaces

type ScheduleTask interface {
	Run()
	Shutdown()
}
