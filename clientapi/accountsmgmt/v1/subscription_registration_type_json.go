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

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalSubscriptionRegistration writes a value of the 'subscription_registration' type to the given writer.
func MarshalSubscriptionRegistration(object *SubscriptionRegistration, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteSubscriptionRegistration(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteSubscriptionRegistration writes a value of the 'subscription_registration' type to the given stream.
func WriteSubscriptionRegistration(object *SubscriptionRegistration, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_uuid")
		stream.WriteString(object.clusterUUID)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("console_url")
		stream.WriteString(object.consoleURL)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("display_name")
		stream.WriteString(object.displayName)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("plan_id")
		stream.WriteString(string(object.planID))
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("status")
		stream.WriteString(object.status)
	}
	stream.WriteObjectEnd()
}

// UnmarshalSubscriptionRegistration reads a value of the 'subscription_registration' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalSubscriptionRegistration(source interface{}) (object *SubscriptionRegistration, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadSubscriptionRegistration(iterator)
	err = iterator.Error
	return
}

// ReadSubscriptionRegistration reads a value of the 'subscription_registration' type from the given iterator.
func ReadSubscriptionRegistration(iterator *jsoniter.Iterator) *SubscriptionRegistration {
	object := &SubscriptionRegistration{
		fieldSet_: make([]bool, 5),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "cluster_uuid":
			value := iterator.ReadString()
			object.clusterUUID = value
			object.fieldSet_[0] = true
		case "console_url":
			value := iterator.ReadString()
			object.consoleURL = value
			object.fieldSet_[1] = true
		case "display_name":
			value := iterator.ReadString()
			object.displayName = value
			object.fieldSet_[2] = true
		case "plan_id":
			text := iterator.ReadString()
			value := PlanID(text)
			object.planID = value
			object.fieldSet_[3] = true
		case "status":
			value := iterator.ReadString()
			object.status = value
			object.fieldSet_[4] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
