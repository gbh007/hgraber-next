package metriccore

const (
	ActionLabel      = "action"
	TypeLabel        = "type"
	SuccessLabel     = "success"
	ServiceNameLabel = "service_name"
	ServiceTypeLabel = "service_type"
)

const (
	ServiceTypeLabelValueServer  = "server"
	ServiceTypeLabelValueAgent   = "agent"
	ServiceTypeLabelValueUnknown = "unknown"
)

func ErrorLabel(ok bool) string {
	if ok {
		return OkLabelValue
	}

	return ErrorLabelValue
}
