// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migration

import (
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	v1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
)

// Converter converts a managed resource or a Composition's ComposedTemplate
// from the migration source provider's schema to the migration target
// provider's schema.
type Converter interface {
	// Resource takes a managed resource and returns zero or more managed
	// resources to be created.
	Resource(mg resource.Managed) ([]resource.Managed, error)

	// Composition receives a migration source v1.ComposedTemplate
	// that has been converted, by a resource converter, to the
	// v1.ComposedTemplates with new shapes specified in the
	// `convertedTemplates` argument.
	// Conversion of the v1.ComposedTemplate.Bases is handled
	// via Converter.Resources and Converter.Composition must only
	// convert the other fields (`Patches`, `ConnectionDetails`,
	// `PatchSet`s, etc.)
	// Returns the converted v1.PatchSet array or any errors encountered.
	Composition(sourcePatchSets []v1.PatchSet, sourceTemplate v1.ComposedTemplate, convertedTemplates ...*v1.ComposedTemplate) ([]v1.PatchSet, error)
}

// Source is a source for reading resource manifests
type Source interface {
	// HasNext returns `true` if the Source implementation has a next manifest
	// available to return with a call to Next. Any errors encountered while
	// determining whether a next manifest exists will also be reported.
	HasNext() (bool, error)
	// Next returns the next resource manifest available or
	// any errors encountered while reading the next resource manifest.
	Next() (UnstructuredWithMetadata, error)
}

// Target is a target where resource manifests can be manipulated
// (e.g., added, deleted, patched, etc.)
type Target interface {
	// Put writes a resource manifest to this Target
	Put(o UnstructuredWithMetadata) error
	// Delete deletes a resource manifest from this Target
	Delete(o UnstructuredWithMetadata) error
}
