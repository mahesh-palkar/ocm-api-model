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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/statusboard/v1

import (
	time "time"
)

// ServiceDependencyKind is the name of the type used to represent objects
// of type 'service_dependency'.
const ServiceDependencyKind = "ServiceDependency"

// ServiceDependencyLinkKind is the name of the type used to represent links
// to objects of type 'service_dependency'.
const ServiceDependencyLinkKind = "ServiceDependencyLink"

// ServiceDependencyNilKind is the name of the type used to nil references
// to objects of type 'service_dependency'.
const ServiceDependencyNilKind = "ServiceDependencyNil"

// ServiceDependency represents the values of the 'service_dependency' type.
//
// Definition of a Status Board service dependency.
type ServiceDependency struct {
	fieldSet_     []bool
	id            string
	href          string
	childService  *Service
	createdAt     time.Time
	metadata      interface{}
	name          string
	owners        []*Owner
	parentService *Service
	type_         string
	updatedAt     time.Time
}

// Kind returns the name of the type of the object.
func (o *ServiceDependency) Kind() string {
	if o == nil {
		return ServiceDependencyNilKind
	}
	if len(o.fieldSet_) > 0 && o.fieldSet_[0] {
		return ServiceDependencyLinkKind
	}
	return ServiceDependencyKind
}

// Link returns true if this is a link.
func (o *ServiceDependency) Link() bool {
	return o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0]
}

// ID returns the identifier of the object.
func (o *ServiceDependency) ID() string {
	if o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1] {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ServiceDependency) GetID() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1]
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *ServiceDependency) HREF() string {
	if o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2] {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ServiceDependency) GetHREF() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2]
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ServiceDependency) Empty() bool {
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

// ChildService returns the value of the 'child_service' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The service dependency's child service
func (o *ServiceDependency) ChildService() *Service {
	if o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3] {
		return o.childService
	}
	return nil
}

// GetChildService returns the value of the 'child_service' attribute and
// a flag indicating if the attribute has a value.
//
// The service dependency's child service
func (o *ServiceDependency) GetChildService() (value *Service, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3]
	if ok {
		value = o.childService
	}
	return
}

// CreatedAt returns the value of the 'created_at' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Object creation timestamp.
func (o *ServiceDependency) CreatedAt() time.Time {
	if o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4] {
		return o.createdAt
	}
	return time.Time{}
}

// GetCreatedAt returns the value of the 'created_at' attribute and
// a flag indicating if the attribute has a value.
//
// Object creation timestamp.
func (o *ServiceDependency) GetCreatedAt() (value time.Time, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4]
	if ok {
		value = o.createdAt
	}
	return
}

// Metadata returns the value of the 'metadata' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Miscellaneous metadata about the service dependency.
func (o *ServiceDependency) Metadata() interface{} {
	if o != nil && len(o.fieldSet_) > 5 && o.fieldSet_[5] {
		return o.metadata
	}
	return nil
}

// GetMetadata returns the value of the 'metadata' attribute and
// a flag indicating if the attribute has a value.
//
// Miscellaneous metadata about the service dependency.
func (o *ServiceDependency) GetMetadata() (value interface{}, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 5 && o.fieldSet_[5]
	if ok {
		value = o.metadata
	}
	return
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The name of the service dependency.
func (o *ServiceDependency) Name() string {
	if o != nil && len(o.fieldSet_) > 6 && o.fieldSet_[6] {
		return o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
// The name of the service dependency.
func (o *ServiceDependency) GetName() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 6 && o.fieldSet_[6]
	if ok {
		value = o.name
	}
	return
}

// Owners returns the value of the 'owners' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The service dependency owners (name and email)
func (o *ServiceDependency) Owners() []*Owner {
	if o != nil && len(o.fieldSet_) > 7 && o.fieldSet_[7] {
		return o.owners
	}
	return nil
}

// GetOwners returns the value of the 'owners' attribute and
// a flag indicating if the attribute has a value.
//
// The service dependency owners (name and email)
func (o *ServiceDependency) GetOwners() (value []*Owner, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 7 && o.fieldSet_[7]
	if ok {
		value = o.owners
	}
	return
}

// ParentService returns the value of the 'parent_service' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The service dependency's parent service
func (o *ServiceDependency) ParentService() *Service {
	if o != nil && len(o.fieldSet_) > 8 && o.fieldSet_[8] {
		return o.parentService
	}
	return nil
}

// GetParentService returns the value of the 'parent_service' attribute and
// a flag indicating if the attribute has a value.
//
// The service dependency's parent service
func (o *ServiceDependency) GetParentService() (value *Service, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 8 && o.fieldSet_[8]
	if ok {
		value = o.parentService
	}
	return
}

// Type returns the value of the 'type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The type of service dependency, e.g. soft or hard.
func (o *ServiceDependency) Type() string {
	if o != nil && len(o.fieldSet_) > 9 && o.fieldSet_[9] {
		return o.type_
	}
	return ""
}

// GetType returns the value of the 'type' attribute and
// a flag indicating if the attribute has a value.
//
// The type of service dependency, e.g. soft or hard.
func (o *ServiceDependency) GetType() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 9 && o.fieldSet_[9]
	if ok {
		value = o.type_
	}
	return
}

// UpdatedAt returns the value of the 'updated_at' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Object modification timestamp.
func (o *ServiceDependency) UpdatedAt() time.Time {
	if o != nil && len(o.fieldSet_) > 10 && o.fieldSet_[10] {
		return o.updatedAt
	}
	return time.Time{}
}

// GetUpdatedAt returns the value of the 'updated_at' attribute and
// a flag indicating if the attribute has a value.
//
// Object modification timestamp.
func (o *ServiceDependency) GetUpdatedAt() (value time.Time, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 10 && o.fieldSet_[10]
	if ok {
		value = o.updatedAt
	}
	return
}

// ServiceDependencyListKind is the name of the type used to represent list of objects of
// type 'service_dependency'.
const ServiceDependencyListKind = "ServiceDependencyList"

// ServiceDependencyListLinkKind is the name of the type used to represent links to list
// of objects of type 'service_dependency'.
const ServiceDependencyListLinkKind = "ServiceDependencyListLink"

// ServiceDependencyNilKind is the name of the type used to nil lists of objects of
// type 'service_dependency'.
const ServiceDependencyListNilKind = "ServiceDependencyListNil"

// ServiceDependencyList is a list of values of the 'service_dependency' type.
type ServiceDependencyList struct {
	href  string
	link  bool
	items []*ServiceDependency
}

// Kind returns the name of the type of the object.
func (l *ServiceDependencyList) Kind() string {
	if l == nil {
		return ServiceDependencyListNilKind
	}
	if l.link {
		return ServiceDependencyListLinkKind
	}
	return ServiceDependencyListKind
}

// Link returns true iif this is a link.
func (l *ServiceDependencyList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ServiceDependencyList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ServiceDependencyList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *ServiceDependencyList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *ServiceDependencyList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *ServiceDependencyList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *ServiceDependencyList) SetItems(items []*ServiceDependency) {
	l.items = items
}

// Items returns the items of the list.
func (l *ServiceDependencyList) Items() []*ServiceDependency {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *ServiceDependencyList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ServiceDependencyList) Get(i int) *ServiceDependency {
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
func (l *ServiceDependencyList) Slice() []*ServiceDependency {
	var slice []*ServiceDependency
	if l == nil {
		slice = make([]*ServiceDependency, 0)
	} else {
		slice = make([]*ServiceDependency, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ServiceDependencyList) Each(f func(item *ServiceDependency) bool) {
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
func (l *ServiceDependencyList) Range(f func(index int, item *ServiceDependency) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
