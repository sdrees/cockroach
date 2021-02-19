// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
)

// limitOp is an operator that implements limit, returning only the first n
// tuples from its input.
type limitOp struct {
	oneInputCloserHelper

	limit uint64

	// seen is the number of tuples seen so far.
	seen uint64
	// done is true if the limit has been reached.
	done bool
}

var _ colexecbase.Operator = &limitOp{}
var _ closableOperator = &limitOp{}

// NewLimitOp returns a new limit operator with the given limit.
func NewLimitOp(input colexecbase.Operator, limit uint64) colexecbase.Operator {
	c := &limitOp{
		oneInputCloserHelper: makeOneInputCloserHelper(input),
		limit:                limit,
	}
	return c
}

func (c *limitOp) Init() {
	c.Input.Init()
}

func (c *limitOp) Next(ctx context.Context) coldata.Batch {
	if c.done {
		return coldata.ZeroBatch
	}
	bat := c.Input.Next(ctx)
	length := bat.Length()
	if length == 0 {
		return bat
	}
	newSeen := c.seen + uint64(length)
	if newSeen >= c.limit {
		c.done = true
		bat.SetLength(int(c.limit - c.seen))
		return bat
	}
	c.seen = newSeen
	return bat
}
