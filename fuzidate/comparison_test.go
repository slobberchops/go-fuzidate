/*
Copyright 2019 Rafe Kaplan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzidate

import (
	"reflect"
	"sort"
	"testing"
)

func TestOrder(t *testing.T) {
	precisions := []struct {
		greater Fuzidate
		lesser Fuzidate
		descr string
	}{
		{Unknown, Y(1914), "Unknown is the least"},
		{ Y(1914), Yo(1914, 1), "Year < Year offset"},
		{ Yo(1914, 1),Y(1915), "Year < Year"},
		{ Yo(1914, 1),Ym(1914, 1),  "Offset < Month"},
		{ Ym(1914, 7), Ymo(1914, 7, 1), "Month < Month offset"},
		{ Ymo(1914, 7, 1),Ym(1914, 8), "Month < Month"},
		{ Ymo(1914, 7, 1),Ymd(1914, 7, 1),  "Offset < Day"},
		{ Ymd(1914, 7, 28), Ymdo(1914, 7, 28, 1), "Day < Day offset"},
		{ Ymdo(1914, 7, 28, 1),Ymd(1914, 7, 29), "Month < Month"},
	}
	for _, tc := range precisions {
		t.Run(tc.descr, func(t *testing.T) {
			if tc.greater.Ge(tc.lesser) {
				t.Errorf("%s <= %s", tc.greater.String(), tc.lesser.String())
			}
		})
	}
}

func TestEq(t *testing.T) {
	precisions := []struct {
		left  Fuzidate
		right Fuzidate
		descr string
	}{
		{Unknown, Y(0), "Unknown"},
		{Y(1914), Y(1914), "Year"},
		{Yo(1914, 4), Yo(1914, 4), "Year offset"},
		{Ym(1914, 7), Ym(1914, 7), "Month"},
		{Ymo(1914, 7, 4), Ymo(1914, 7, 4), "Month offset"},
		{Ymd(1914, 7, 28), Ymd(1914, 7, 28), "Day"},
		{Ymdo(1914, 7, 28, 4), Ymdo(1914, 7, 28, 4), "Day offset"},
	}
	for _, tc := range precisions {
		t.Run(tc.descr, func(t *testing.T) {
			if tc.left.Compare(tc.right) != 0 {
				t.Errorf("%s != %s", tc.left.String(), tc.right.String())
			}
		})
	}
}

func TestCompareLt(t *testing.T) {
	left := Y(1914)
	right := Y(1915)
	if left.Compare(right) >= 0 {
		t.Errorf("Y(%s).Compare(Y(%s)) >= 0", left, right)
	}
}

func TestCompareEq(t *testing.T) {
	left := Y(1914)
	right := Y(1914)
	if left.Compare(right) != 0 {
		t.Errorf("Y(%s).Compare(Y(%s)) != 0", left, right)
	}
}

func TestCompareGt(t *testing.T) {
	left := Y(1915)
	right := Y(1914)
	if left.Compare(right) <= 0 {
		t.Errorf("Y(%s).Compare(Y(%s)) <= 0", left, right)
	}
}

func TestLt(t *testing.T) {
	left := Y(1914)
	right := Y(1915)
	if !left.Lt(right) {
		t.Errorf("!Y(%s).Le(Y(%s))", left, right)
	}
}

func TestLe(t *testing.T) {
	left := Y(1914)
	right := Y(1915)
	if !left.Le(right) {
		t.Errorf("!Y(%s).Le(Y(%s))", left, right)
	}
	if !left.Le(left) {
		t.Errorf("!Y(%s).Le(Y(%s))", left, left)
	}
}

func TestGt(t *testing.T) {
	left := Y(1915)
	right := Y(1914)
	if !left.Gt(right) {
		t.Errorf("!Y(%s).Ge(Y(%s))", left, right)
	}
}

func TestGe(t *testing.T) {
	left := Y(1915)
	right := Y(1914)
	if !left.Ge(right) {
		t.Errorf("!Y(%s).Ge(Y(%s))", left, right)
	}
	if !left.Ge(left) {
		t.Errorf("!Y(%s).Ge(Y(%s))", left, left)
	}
}

func TestSort(t *testing.T) {
	fzds := Sequence {
		Y(1917),
		Y(1916),
		Y(1914),
		Y(1918),
		Y(1915),
	}
	sort.Sort(fzds)

	expected :=  Sequence {
		Y(1914),
		Y(1915),
		Y(1916),
		Y(1917),
		Y(1918),
	}
	if !reflect.DeepEqual(fzds, expected) {
		t.Errorf("Unexpected sort order: %s", fzds)
	}
}
