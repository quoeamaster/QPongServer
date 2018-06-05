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

package datastore

func NewProjectEntity(id string) *Project {
	project := Project{}
	project.Id = id
	return &project
}

func (o *Project) AddDesign(design *Design) *Project {
	o.Design = design
	return o
}

func (o *Project) SetTemplates(t []*Template) *Project {
	o.Templates = t
	return o
}

func (o *Project) AddTemplate(t *Template) *Project {
	if len(o.Templates)==0 {
		o.Templates = make([]*Template, 0)
	}
	// force the slice to increment the capacity / length / size
	o.Templates[len(o.Templates)] = t

	return o
}