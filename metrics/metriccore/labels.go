package metriccore

const (
	FSIDLabel = "fs_id"
)

func ErrorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
