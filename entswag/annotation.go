// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entswag

import "github.com/facebook/ent/schema"

// Annotation annotates fields and edges with metadata for templates.
type Annotation struct {
	// Model is a swagger annotation on the struct
	Model bool
}

// Name implements ent.Annotation interface.
func (Annotation) Name() string {
	return "EntSwag"
}

// Model returns a swagger:model annotation.
func Model(name bool) Annotation {
	return Annotation{Model: true}
}

// Merge implements the schema.Merger interface.
func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}
	if ant.Model != false {
		a.Model = ant.Model
	}

	return a
}

var (
	_ schema.Annotation = (*Annotation)(nil)
	_ schema.Merger     = (*Annotation)(nil)
)
