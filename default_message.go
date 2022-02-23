package answer

import "github.com/samuskitchen/go-message-tool/errors"

const (
	OperationSuccess   = "successful operation"
	CreateSuccess      = "record saved successfully"
	OperationError     = "the operation could not be completed"
	ForbiddenOperation = "operation not allowed"
	InternalError      = "error, something happened"
)

var (
	ErrorDefaultForbidden = errores.NewForbidden(nil, ForbiddenOperation)
)
