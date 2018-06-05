/*
 *  Copyright Project - CFactor, Author - quoeamaster, (C) 2018
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package http

import (
	"reflect"
)

/* ***************************** */
/*    structs for the modules    */
/* ***************************** */

type ModuleError struct {
	ErrorMsg string
	ErrorType string
	Meta map[string]interface{}
}

func NewModuleError(e *error, meta... map[string]interface{}) ModuleError {
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
