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

func NewSpecEntity() *Spec {
	return &Spec{}
}

func (o *Spec) AddBackgroundImagePath(path string) *Spec {
	o.BackgroundImagePath = path
	return o
}

func (o *Spec) AddTitle(t *TextBlock) *Spec {
	o.Title = t
	return o
}

func (o *Spec) AddSubTitle(t *TextBlock) *Spec {
	o.SubTitle = t
	return o
}

func (o *Spec) AddDescription(t *TextBlock) *Spec {
	o.Description = t
	return o
}
