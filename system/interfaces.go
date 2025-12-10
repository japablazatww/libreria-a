package system

type System interface {
	GetSystemStatus(code string) (string, error)
}
