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

import "testing"

func checkFuzidate(
	t *testing.T,
	fzd Fuzidate,
	year Year,
	month Month,
	day Day,
	offset Offset,
	number Number) {
	if fzd.Year() != year {
		t.Errorf("%d != %d", fzd.Year(), year)
	}
	if fzd.Month() != month {
		t.Errorf("%d != %d", fzd.Month(), month)
	}
	if fzd.Day() != day {
		t.Errorf("%d != %d", fzd.Day(), day)
	}
	if fzd.Offset() != offset {
		t.Errorf("%d != %d", fzd.Offset(), offset)
	}
	if fzd.Number() != number {
		t.Errorf("%d != %d", fzd.Number(), number)
	}
}
