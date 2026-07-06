package viewstate

const (
	ErrorDuplicate     = "duplicate"
	ErrorInvalidCreate = "invalid-create"
	ErrorInvalidEdit   = "invalid-edit"
	ErrorBrokenCreate  = "broken-create"
	ErrorBrokenEdit    = "broken-edit"
)

const (
	SuccessCreated = "created"
	SuccessUpdated = "updated"
	SuccessDeleted = "deleted"
)

func ErrorType(value string) string {
	switch value {
	case ErrorDuplicate, ErrorInvalidCreate, ErrorInvalidEdit, ErrorBrokenCreate, ErrorBrokenEdit:
		return value
	default:
		return ""
	}
}

func SuccessType(value string) string {
	switch value {
	case SuccessCreated, SuccessUpdated, SuccessDeleted:
		return value
	default:
		return ""
	}
}
