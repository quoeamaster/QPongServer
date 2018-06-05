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

/* ********************************************* */
/*     data structs for the "project" entity     */
/* ********************************************* */

/**
 *  PROJECT is the starting point of everything:
 *  1. Design (Specifications for this project's main design)
 *  2. Template(s) (templates generated and picked eventually)
 */
type Project struct {
	Id        string
	Design    *Design
	Templates []*Template
}
/**
 *  involve Specifications information
 *  1. Background image path
 *  2. Title
 *  3. Sub Title
 *  4. Description
 */
type Design struct {
	Spec *Spec
}
/**
 *  same contents with a "Design"
 */
type Template struct {
	Spec *Spec
}

/**
 *  encapsulation of a Design / Template specification
 */
type Spec struct {
	BackgroundImagePath string
	Title               *TextBlock
	SubTitle            *TextBlock
	Description         *TextBlock
}

/**
 *  a block of Text with:
 *  1. wordings
 *  2. Dimension information (x, y, w, h)
 */
type TextBlock struct {
	Text  string
	Dimen *Dimension
}
/**
 *  encapsulation of "dimension"
 */
type Dimension struct {
	X, Y float32
	w, h float32
}
