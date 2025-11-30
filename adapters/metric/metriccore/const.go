package metriccore

const (
	SystemName    = "hgraber_next"
	SubSystemName = "server" //nolint:goconst // ложно-положительное

	OkLabelValue    = "ok"
	ErrorLabelValue = "error"
)

const (
	Kilobyte = 1 << 10
	Megabyte = 1 << 20
	Gigabyte = 1 << 30
)
