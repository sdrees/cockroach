// Code generated by execgen; DO NOT EDIT.

// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecsel

import (
	"bytes"
	"regexp"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
)

type selPrefixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selPrefixBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selSuffixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selSuffixBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selContainsBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selContainsBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selRegexpBytesBytesConstOp struct {
	selConstOpBase
	constArg *regexp.Regexp
}

func (p *selRegexpBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selNotPrefixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selNotPrefixBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasPrefix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selNotSuffixBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selNotSuffixBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.HasSuffix(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selNotContainsBytesBytesConstOp struct {
	selConstOpBase
	constArg []byte
}

func (p *selNotContainsBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !bytes.Contains(arg, p.constArg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}

type selNotRegexpBytesBytesConstOp struct {
	selConstOpBase
	constArg *regexp.Regexp
}

func (p *selNotRegexpBytesBytesConstOp) Next() coldata.Batch {
	// In order to inline the templated code of overloads, we need to have a
	// `_overloadHelper` local variable of type `execgen.OverloadHelper`.
	_overloadHelper := p.overloadHelper
	// However, the scratch is not used in all of the selection operators, so
	// we add this to go around "unused" error.
	_ = _overloadHelper
	for {
		batch := p.Input.Next()
		if batch.Length() == 0 {
			return batch
		}

		vec := batch.ColVec(p.colIdx)
		col := vec.Bytes()
		var idx int
		n := batch.Length()
		if vec.MaybeHasNulls() {
			nulls := vec.Nulls()
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					if nulls.NullAt(i) {
						continue
					}
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		} else {
			if sel := batch.Selection(); sel != nil {
				sel = sel[:n]
				for _, i := range sel {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			} else {
				batch.SetSelection(true)
				sel := batch.Selection()
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					var cmp bool
					arg := col.Get(i)
					cmp = !p.constArg.Match(arg)
					if cmp {
						sel[idx] = i
						idx++
					}
				}
			}
		}
		if idx > 0 {
			batch.SetLength(idx)
			return batch
		}
	}
}
