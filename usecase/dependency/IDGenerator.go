package dependency

type IDGenerator interface {
	Generate() string
	GenerateNumberCode(length int) string
}
