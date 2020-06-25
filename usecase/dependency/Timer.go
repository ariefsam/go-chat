package dependency

type Timer interface {
	CurrentTimestamp() int64
}
