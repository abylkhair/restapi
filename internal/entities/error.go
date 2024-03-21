package entities

import "errors"

var (
	ErrInternal    = errors.New("internal error")
	ErrNotFound    = errors.New("not found")
	ErrInitFail    = errors.New("errors of initialize of new entity")
	ErrStorageRead = errors.New("errors in storage")
)
