// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package descs

import (
	"github.com/cockroachdb/cockroach/pkg/sql/catalog"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descriptortree"
)

type syntheticDescriptors struct {
	descs descriptortree.Tree
}

func makeSyntheticDescriptors() syntheticDescriptors {
	return syntheticDescriptors{descs: descriptortree.Make()}
}

func (sd *syntheticDescriptors) set(descs []catalog.Descriptor) {
	sd.descs.Clear()
	for _, desc := range descs {
		if mut, ok := desc.(catalog.MutableDescriptor); ok {
			desc = mut.ImmutableCopy()
		}
		sd.descs.Upsert(desc)
	}
}

func (sd *syntheticDescriptors) reset() {
	sd.descs.Clear()
}

func (sd *syntheticDescriptors) getByName(
	dbID descpb.ID, schemaID descpb.ID, name string,
) (found bool, desc catalog.Descriptor) {
	desc, found = sd.descs.GetByName(dbID, schemaID, name)
	return found, desc
}

func (sd *syntheticDescriptors) getByID(id descpb.ID) (found bool, desc catalog.Descriptor) {
	desc, found = sd.descs.GetByID(id)
	return found, desc
}
