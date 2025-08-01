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

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// NodePoolAutoscalingKind is the name of the type used to represent objects
// of type 'node_pool_autoscaling'.
const NodePoolAutoscalingKind = "NodePoolAutoscaling"

// NodePoolAutoscalingLinkKind is the name of the type used to represent links
// to objects of type 'node_pool_autoscaling'.
const NodePoolAutoscalingLinkKind = "NodePoolAutoscalingLink"

// NodePoolAutoscalingNilKind is the name of the type used to nil references
// to objects of type 'node_pool_autoscaling'.
const NodePoolAutoscalingNilKind = "NodePoolAutoscalingNil"

// NodePoolAutoscaling represents the values of the 'node_pool_autoscaling' type.
//
// Representation of a autoscaling in a node pool.
type NodePoolAutoscaling struct {
	fieldSet_  []bool
	id         string
	href       string
	maxReplica int
	minReplica int
}

// Kind returns the name of the type of the object.
func (o *NodePoolAutoscaling) Kind() string {
	if o == nil {
		return NodePoolAutoscalingNilKind
	}
	if len(o.fieldSet_) > 0 && o.fieldSet_[0] {
		return NodePoolAutoscalingLinkKind
	}
	return NodePoolAutoscalingKind
}

// Link returns true if this is a link.
func (o *NodePoolAutoscaling) Link() bool {
	return o != nil && len(o.fieldSet_) > 0 && o.fieldSet_[0]
}

// ID returns the identifier of the object.
func (o *NodePoolAutoscaling) ID() string {
	if o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1] {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *NodePoolAutoscaling) GetID() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 1 && o.fieldSet_[1]
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *NodePoolAutoscaling) HREF() string {
	if o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2] {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *NodePoolAutoscaling) GetHREF() (value string, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 2 && o.fieldSet_[2]
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *NodePoolAutoscaling) Empty() bool {
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

// MaxReplica returns the value of the 'max_replica' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The maximum number of replicas for the node pool.
func (o *NodePoolAutoscaling) MaxReplica() int {
	if o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3] {
		return o.maxReplica
	}
	return 0
}

// GetMaxReplica returns the value of the 'max_replica' attribute and
// a flag indicating if the attribute has a value.
//
// The maximum number of replicas for the node pool.
func (o *NodePoolAutoscaling) GetMaxReplica() (value int, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 3 && o.fieldSet_[3]
	if ok {
		value = o.maxReplica
	}
	return
}

// MinReplica returns the value of the 'min_replica' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The minimum number of replicas for the node pool.
func (o *NodePoolAutoscaling) MinReplica() int {
	if o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4] {
		return o.minReplica
	}
	return 0
}

// GetMinReplica returns the value of the 'min_replica' attribute and
// a flag indicating if the attribute has a value.
//
// The minimum number of replicas for the node pool.
func (o *NodePoolAutoscaling) GetMinReplica() (value int, ok bool) {
	ok = o != nil && len(o.fieldSet_) > 4 && o.fieldSet_[4]
	if ok {
		value = o.minReplica
	}
	return
}

// NodePoolAutoscalingListKind is the name of the type used to represent list of objects of
// type 'node_pool_autoscaling'.
const NodePoolAutoscalingListKind = "NodePoolAutoscalingList"

// NodePoolAutoscalingListLinkKind is the name of the type used to represent links to list
// of objects of type 'node_pool_autoscaling'.
const NodePoolAutoscalingListLinkKind = "NodePoolAutoscalingListLink"

// NodePoolAutoscalingNilKind is the name of the type used to nil lists of objects of
// type 'node_pool_autoscaling'.
const NodePoolAutoscalingListNilKind = "NodePoolAutoscalingListNil"

// NodePoolAutoscalingList is a list of values of the 'node_pool_autoscaling' type.
type NodePoolAutoscalingList struct {
	href  string
	link  bool
	items []*NodePoolAutoscaling
}

// Kind returns the name of the type of the object.
func (l *NodePoolAutoscalingList) Kind() string {
	if l == nil {
		return NodePoolAutoscalingListNilKind
	}
	if l.link {
		return NodePoolAutoscalingListLinkKind
	}
	return NodePoolAutoscalingListKind
}

// Link returns true iif this is a link.
func (l *NodePoolAutoscalingList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *NodePoolAutoscalingList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *NodePoolAutoscalingList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *NodePoolAutoscalingList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *NodePoolAutoscalingList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *NodePoolAutoscalingList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *NodePoolAutoscalingList) SetItems(items []*NodePoolAutoscaling) {
	l.items = items
}

// Items returns the items of the list.
func (l *NodePoolAutoscalingList) Items() []*NodePoolAutoscaling {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *NodePoolAutoscalingList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *NodePoolAutoscalingList) Get(i int) *NodePoolAutoscaling {
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
func (l *NodePoolAutoscalingList) Slice() []*NodePoolAutoscaling {
	var slice []*NodePoolAutoscaling
	if l == nil {
		slice = make([]*NodePoolAutoscaling, 0)
	} else {
		slice = make([]*NodePoolAutoscaling, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *NodePoolAutoscalingList) Each(f func(item *NodePoolAutoscaling) bool) {
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
func (l *NodePoolAutoscalingList) Range(f func(index int, item *NodePoolAutoscaling) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
