package errores

const (
	ErrInvalidJSON           = "invalid JSON structure, check the documentation"
	ErrInvalidQueryParam     = "Invalid query params, review the documentation"
	ErrInvalidToken          = "the token is invalid"
	ErrTokenNull             = "token not found"
	ErrSigningTokenString    = "can't authenticate"
	ErrNoDefined             = "there was an error, unexpected"
	ErrDatabaseRequest       = "the operation could not be performed"
	ErrDatabaseInternal      = "something happened, the operation could not be performed"
	ErrRecordNotFound        = "record not found"
	ErrRecord                = "could not save the record"
	ErrUsernameExists        = "user already exists"
	ErrAuthorizationHeader   = "authorization header not found"
	ErrUserOrPasswordInvalid = "wrong username or password"
)
