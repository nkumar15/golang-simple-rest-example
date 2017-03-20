package main

import "errors"

var (
	ErrNoMoreRows               = errors.New(`location: no more rows in this result set`)
	ErrNotConnected             = errors.New(`location: you're currently not connected`)
	ErrMissingDatabaseName      = errors.New(`location: missing database name`)
	ErrMissingCollectionName    = errors.New(`location: missing collection name`)
	ErrCollectionDoesNotExist   = errors.New(`location: collection does not exist`)
	ErrSockerOrHost             = errors.New(`location: you may connect either to a unix socket or a tcp address, but not both`)
	ErrQueryLimitParam          = errors.New(`location: a query can accept only one limit parameter`)
	ErrQuerySortParam           = errors.New(`location: a query can accept only one order by parameter`)
	ErrQueryOffsetParam         = errors.New(`location: a query can accept only one offset parameter`)
	ErrMissingConditions        = errors.New(`location: missing selector conditions`)
	ErrUnsupported              = errors.New(`location: this action is currently unsupported on this database`)
	ErrUndefined                = errors.New(`location: this value is undefined`)
	ErrQueryIsPending           = errors.New(`location: can't execute this instruction while the result set is still open`)
	ErrUnsupportedDestination   = errors.New(`location: unsupported destination type`)
	ErrUnsupportedType          = errors.New(`location: this type does not support marshaling`)
	ErrUnsupportedValue         = errors.New(`location: this value does not support unmarshaling`)
	ErrUnknownConditionType     = errors.New(`location: arguments of type %T can't be used as constraints`)
	ErrTooManyClients           = errors.New(`location: can't connect to database server: too many clients`)
	ErrGivingUpTryingToConnect  = errors.New(`location: giving up trying to connect: too many clients`)
	ErrMissingConnURL           = errors.New(`location: missing DSN`)
	ErrNotImplemented           = errors.New(`location: call not implemented`)
	ErrAlreadyWithinTransaction = errors.New(`location: already within a transaction`)
)
