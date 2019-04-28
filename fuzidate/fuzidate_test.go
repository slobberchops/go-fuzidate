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
	"fmt"
	"testing"
)

func TestUnknown(t *testing.T) {
	checkFuzidate(t, Unknown, 0, 0, 0, 0, 0)
}

func TestPrecision_None(t *testing.T) {
	if Unknown.Precision() != None {
		t.Errorf("%d != None", Unknown.Precision())
	}
}

func TestPrecision_Year(t *testing.T) {
	fzd := Y(1928)
	if fzd.Precision() != AtYear {
		t.Errorf("%d != AtYear", fzd.Precision())
	}
}

func TestPrecision_Month(t *testing.T) {
	fzd := Ym(1928, 7)
	if fzd.Precision() != AtMonth {
		t.Errorf("%d != AtMonth", fzd.Precision())
	}
}

func TestPrecision_Day(t *testing.T) {
	fzd := Ymd(1928, 7, 28)
	if fzd.Precision() != AtDay {
		t.Errorf("%d != AtDay", fzd.Precision())
	}
}

func TestFuzidate_String(t *testing.T) {
	cases := []struct {
		fzd Fuzidate
		expected string
	}{
		{Ymdo(1918, 7, 28, 4), "1918-07-28+4"},
		{Ymdo(1918, 7, 28, 0), "1918-07-28"},
		{Ymdo(1918, 7, 0, 4), "1918-07+4"},
		{Ymdo(1918, 7, 0, 0), "1918-07"},
		{Ymdo(1918, 0, 0, 4), "1918+4"},
		{Ymdo(1918, 0, 0, 0), "1918"},
		{Unknown, "0"},
	}
	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			if tc.fzd.String() != tc.expected {
				t.Errorf("%s != %s", tc.fzd, tc.expected)
			}
		})
	}
}

func TestFuzidate_Uint(t *testing.T) {
	cases := []struct {
		fzd Fuzidate
		expected uint64
	}{
		{Ymdo(1918, 7, 28, 4), 82380599473471492},
		{Ymdo(1918, 7, 28, 0), 82380599473471488},
		{Ymdo(1918, 7, 0, 4), 82380479214387204},
		{Ymdo(1918, 7, 0, 0), 82380479214387200},
		{Ymdo(1918, 0, 0, 4), 82377472737280004},
		{Ymdo(1918, 0, 0, 0), 82377472737280000},
		{Unknown, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.expected), func(t *testing.T) {
			if tc.fzd.Uint() != tc.expected {
				t.Errorf("%d != %d", tc.fzd.Uint(), tc.expected)
			}
		})
	}
}
