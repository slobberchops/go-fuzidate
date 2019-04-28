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
	"testing"
	"time"
)

func TestYmdo(t *testing.T) {
	fzd := Ymdo(1918, 7, 28, 4)
	checkFuzidate(t, fzd, 1918, 7, 28, 4, 19180728)
}

func TestYmo(t *testing.T) {
	fzd := Ymo(1918, 7,4)
	checkFuzidate(t, fzd, 1918, 7, 0, 4, 19180700)
}

func TestYo(t *testing.T) {
	fzd := Yo(1918, 4)
	checkFuzidate(t, fzd, 1918, 0, 0, 4, 19180000)
}

func TestFromNumberO(t *testing.T) {
	fzd := FromNumberO(19180728, 4)
	checkFuzidate(t, fzd, 1918, 7, 28, 4, 19180728)
}

func TestYmd(t *testing.T) {
	fzd := Ymd(1918, 7, 28)
	checkFuzidate(t, fzd, 1918, 7, 28, 0, 19180728)
}

func TestYm(t *testing.T) {
	fzd := Ym(1918, 7)
	checkFuzidate(t, fzd, 1918, 7, 0, 0, 19180700)
}

func TestY(t *testing.T) {
	fzd := Y(1918)
	checkFuzidate(t, fzd, 1918, 0, 0, 0, 19180000)
}

func TestFromNumber(t *testing.T) {
	fzd := FromNumber(19180728)
	checkFuzidate(t, fzd, 1918, 7, 28, 0, 19180728)
}

func TestFromTime(t *testing.T) {
	dt, _ := time.Parse("2006-1-2", "1918-7-28")
	fzd := FromTime(dt)
	checkFuzidate(t, fzd, 1918, 7, 28, 0, 19180728)
}
