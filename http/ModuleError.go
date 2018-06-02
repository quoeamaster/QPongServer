package http

import "reflect"

/* ***************************** */
/*    structs for the modules    */
/* ***************************** */

type ModuleError struct {
	ErrorMsg string
	ErrorType string
	Meta map[string]interface{}
}

func PopulateModuleError(e *error, meta... map[string]interface{}) ModuleError {
	m := ModuleError{}
	errType := reflect.TypeOf(*e)

	m.ErrorMsg = (*e).Error()
	m.ErrorType = errType.String()
	if meta != nil && len(meta) > 0 {
		m.Meta = meta[0]
	}
	// add additional meta based on error type...

	return m
}
