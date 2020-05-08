package models

type CustomError interface {
	Error() string
	ServeCustomErr() *ErrResponseBody
}

type ErrResponseBody struct {
	Err string				`json:"error"`
	UUID string				`json:"request_id,omitempty"`
	Keys []string			`json:"keys,omitempty"`
	Details []string		`json:"details,omitempty"`
}

// ======== ErrNotFound =====================

type ErrNotFound struct {
	Err error
}

func NewErrNotFound(err error) CustomError {
	return &ErrNotFound{Err: err}
}

func (err *ErrNotFound) Error() string {
	return err.Error()
}

func (err *ErrNotFound) ServeCustomErr() *ErrResponseBody {
	return &ErrResponseBody{Err: "not_found"}
}


// ======== ErrBadParam =====================

type ErrBadParam struct {
	Err error
	Keys []string
	Details []string
}

func NewErrBadParam(err error, keys []string, values []string) CustomError {
	return &ErrBadParam{Err: err, Keys: keys, Details: values}
}

func (err *ErrBadParam) Error() string {
	return err.Error()
}

func (err *ErrBadParam) ServeCustomErr() *ErrResponseBody {
	return &ErrResponseBody{Err: "bad_parameters", Keys: err.Keys, Details: err.Details}
}

// ======== ErrUnauthorized =====================

type ErrUnauthorized struct {
	Err error
	Keys []string
	Details []string
}

func NewErrUnauthorized(err error, keys []string, values []string) CustomError {
	return &ErrUnauthorized{Err: err, Keys: keys, Details: values}
}

func (err *ErrUnauthorized) Error() string {
	return err.Error()
}

func (err *ErrUnauthorized) ServeCustomErr() *ErrResponseBody {
	return &ErrResponseBody{Err: "bad_parameters", Keys: err.Keys, Details: err.Details}
}

// ======== ErrInternal =====================

type ErrInternal struct {
	Err error
	UUID string
}

func NewErrInternal(uuid string, err error) CustomError {
	return &ErrInternal{Err: err, UUID: uuid}
}

func (err *ErrInternal) Error() string {
	return err.Error()
}

func (err *ErrInternal) ServeCustomErr() *ErrResponseBody {
	return &ErrResponseBody{Err: "internal_server_error", UUID: err.UUID}
}