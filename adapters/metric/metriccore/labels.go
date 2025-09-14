package metriccore

const (
	FSIDLabel        = "fs_id"
	ActionLabel      = "action"
	SuccessLabel     = "success"
	ServiceNameLabel = "service_name"
)

func ErrorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
