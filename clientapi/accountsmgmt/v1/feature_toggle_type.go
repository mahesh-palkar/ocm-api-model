/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-api-model/clientapi/accountsmgmt/v1

// FeatureToggleKind is the name of the type used to represent objects
// of type 'feature_toggle'.
const FeatureToggleKind = "FeatureToggle"

// FeatureToggleLinkKind is the name of the type used to represent links
// to objects of type 'feature_toggle'.
const FeatureToggleLinkKind = "FeatureToggleLink"

// FeatureToggleNilKind is the name of the type used to nil references
// to objects of type 'feature_toggle'.
const FeatureToggleNilKind = "FeatureToggleNil"

// FeatureToggle represents the values of the 'feature_toggle' type.
type FeatureToggle struct {
	fieldSet_ []bool
	id        string
	href      string
	enabled   bool
}

// Kind returns the name of the type of the object.
func (o *FeatureToggle) Kind() string {
	if o == nil {
		return FeatureToggleNilKind
	}
	if len(o.fieldSet_) > 0 && o.fieldSet_[0] {
		return FeatureToggleLinkKind
	}
	return FeatureToggleKind
}

// Link returns true if this is a link.
func (o *FeatureToggle) Link() bool {
	return o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0]
}

// ID returns the identifier of the object.
func (o *FeatureToggle) ID() string {
	if o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1] {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *FeatureToggle) GetID() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1]
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *FeatureToggle) HREF() string {
	if o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2] {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *FeatureToggle) GetHREF() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2]
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *FeatureToggle) Empty() bool {
	if o == nil || len(o.fieldSet_) == 0 {
		return true
	}

	// Check all fields except the link flag (index 0)
	for i := 1; i < len(o.fieldSet_); i++ {
		if o.fieldSet_[i] {
			return false
		}
	}
	return true
}

// Enabled returns the value of the 'enabled' attribute, or
// the zero value of the type if the attribute doesn't have a value.
func (o *FeatureToggle) Enabled() bool {
	if o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3] {
		return o.enabled
	}
	return false
}

// GetEnabled returns the value of the 'enabled' attribute and
// a flag indicating if the attribute has a value.
func (o *FeatureToggle) GetEnabled() (value bool, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3]
	if ok {
		value = o.enabled
	}
	return
}

// FeatureToggleListKind is the name of the type used to represent list of objects of
// type 'feature_toggle'.
const FeatureToggleListKind = "FeatureToggleList"

// FeatureToggleListLinkKind is the name of the type used to represent links to list
// of objects of type 'feature_toggle'.
const FeatureToggleListLinkKind = "FeatureToggleListLink"

// FeatureToggleNilKind is the name of the type used to nil lists of objects of
// type 'feature_toggle'.
const FeatureToggleListNilKind = "FeatureToggleListNil"

// FeatureToggleList is a list of values of the 'feature_toggle' type.
type FeatureToggleList struct {
	href  string
	link  bool
	items []*FeatureToggle
}

// Kind returns the name of the type of the object.
func (l *FeatureToggleList) Kind() string {
	if l == nil {
		return FeatureToggleListNilKind
	}
	if l.link {
		return FeatureToggleListLinkKind
	}
	return FeatureToggleListKind
}

// Link returns true iif this is a link.
func (l *FeatureToggleList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *FeatureToggleList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *FeatureToggleList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *FeatureToggleList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *FeatureToggleList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *FeatureToggleList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *FeatureToggleList) SetItems(items []*FeatureToggle) {
	l.items = items
}

// Items returns the items of the list.
func (l *FeatureToggleList) Items() []*FeatureToggle {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *FeatureToggleList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *FeatureToggleList) Get(i int) *FeatureToggle {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *FeatureToggleList) Slice() []*FeatureToggle {
	var slice []*FeatureToggle
	if l == nil {
		slice = make([]*FeatureToggle, 0)
	} else {
		slice = make([]*FeatureToggle, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *FeatureToggleList) Each(f func(item *FeatureToggle) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *FeatureToggleList) Range(f func(index int, item *FeatureToggle) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
