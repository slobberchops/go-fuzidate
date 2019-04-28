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

func (f Fuzidate) Compare(o Fuzidate) int {
	fi := f.Uint()
	oi := o.Uint()
	if fi > oi {
		return 1
	}
	if fi < oi {
		return -1
	}
	return 0
}

func (f Fuzidate) Lt(o Fuzidate) bool {
	return f.Compare(o) < 0
}

func (f Fuzidate) Le(o Fuzidate) bool {
	return f.Compare(o) <= 0
}

func (f Fuzidate) Gt(o Fuzidate) bool {
	return f.Compare(o) > 0
}

func (f Fuzidate) Ge(o Fuzidate) bool {
	return f.Compare(o) >= 0
}

func (f Sequence) Len() int {
	return len(f)
}

func (f Sequence) Less(i, j int) bool {
	return f[i].Compare(f[j]) < 0
}

func (f Sequence) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
