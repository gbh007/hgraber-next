package metriccore

const (
	FSIDLabel   = "fs_id"
	ActionLabel = "action"
)

func ErrorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
