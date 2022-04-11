package handler

import (
	"net/http"

	"gorm.io/gorm"
)

// HTTPDatabaseErrorCodes is http codes map of database errors
var HTTPDatabaseErrorCodes = map[error]int{
	gorm.ErrRecordNotFound:       http.StatusNotFound,
	gorm.ErrUnsupportedRelation:  http.StatusMethodNotAllowed,
	gorm.ErrPrimaryKeyRequired:   http.StatusBadRequest,
	gorm.ErrNotImplemented:       http.StatusNotFound,
	gorm.ErrModelValueRequired:   http.StatusBadRequest,
	gorm.ErrMissingWhereClause:   http.StatusMethodNotAllowed,
	gorm.ErrInvalidValueOfLength: http.StatusBadRequest,
	gorm.ErrInvalidValue:         http.StatusBadRequest,
	gorm.ErrInvalidTransaction:   http.StatusBadRequest,
	gorm.ErrInvalidField:         http.StatusNotFound,
	gorm.ErrInvalidData:          http.StatusBadRequest,
}
