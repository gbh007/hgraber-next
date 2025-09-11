package metriccore

const (
	FSIDLabel    = "fs_id"
	ActionLabel  = "action"
	SuccessLabel = "success"
)

func ErrorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
