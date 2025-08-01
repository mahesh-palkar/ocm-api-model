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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/authorizations/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalAccessReviewRequest writes a value of the 'access_review_request' type to the given writer.
func MarshalAccessReviewRequest(object *AccessReviewRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteAccessReviewRequest(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteAccessReviewRequest writes a value of the 'access_review_request' type to the given stream.
func WriteAccessReviewRequest(object *AccessReviewRequest, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = len(object.fieldSet_) > 0 && object.fieldSet_[0]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("account_username")
		stream.WriteString(object.accountUsername)
		count++
	}
	present_ = len(object.fieldSet_) > 1 && object.fieldSet_[1]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("action")
		stream.WriteString(object.action)
		count++
	}
	present_ = len(object.fieldSet_) > 2 && object.fieldSet_[2]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_id")
		stream.WriteString(object.clusterID)
		count++
	}
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster_uuid")
		stream.WriteString(object.clusterUUID)
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("organization_id")
		stream.WriteString(object.organizationID)
		count++
	}
	present_ = len(object.fieldSet_) > 5 && object.fieldSet_[5]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("resource_type")
		stream.WriteString(object.resourceType)
		count++
	}
	present_ = len(object.fieldSet_) > 6 && object.fieldSet_[6]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("subscription_id")
		stream.WriteString(object.subscriptionID)
	}
	stream.WriteObjectEnd()
}

// UnmarshalAccessReviewRequest reads a value of the 'access_review_request' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAccessReviewRequest(source interface{}) (object *AccessReviewRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadAccessReviewRequest(iterator)
	err = iterator.Error
	return
}

// ReadAccessReviewRequest reads a value of the 'access_review_request' type from the given iterator.
func ReadAccessReviewRequest(iterator *jsoniter.Iterator) *AccessReviewRequest {
	object := &AccessReviewRequest{
		fieldSet_: make([]bool, 7),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "account_username":
			value := iterator.ReadString()
			object.accountUsername = value
			object.fieldSet_[0] = true
		case "action":
			value := iterator.ReadString()
			object.action = value
			object.fieldSet_[1] = true
		case "cluster_id":
			value := iterator.ReadString()
			object.clusterID = value
			object.fieldSet_[2] = true
		case "cluster_uuid":
			value := iterator.ReadString()
			object.clusterUUID = value
			object.fieldSet_[3] = true
		case "organization_id":
			value := iterator.ReadString()
			object.organizationID = value
			object.fieldSet_[4] = true
		case "resource_type":
			value := iterator.ReadString()
			object.resourceType = value
			object.fieldSet_[5] = true
		case "subscription_id":
			value := iterator.ReadString()
			object.subscriptionID = value
			object.fieldSet_[6] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
