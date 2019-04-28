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

func TestPrecision_String(t *testing.T) {
	precisions := []struct {
		precision Precision
		expected string
	}{
		{None, "None"},
		{AtDay, "AtDay"},
		{AtMonth, "AtMonth"},
		{AtYear, "AtYear"},
	}
	for _, tc := range precisions {
		t.Run(tc.expected, func(t *testing.T) {
			if tc.precision.String() != tc.expected {
				t.Errorf("%s != %s", tc.precision.String(), tc.expected)
			}
		})
	}
}
