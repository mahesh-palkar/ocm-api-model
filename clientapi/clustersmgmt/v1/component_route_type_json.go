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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalComponentRoute writes a value of the 'component_route' type to the given writer.
func MarshalComponentRoute(object *ComponentRoute, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteComponentRoute(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteComponentRoute writes a value of the 'component_route' type to the given stream.
func WriteComponentRoute(object *ComponentRoute, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if len(object.fieldSet_) > 0 && object.fieldSet_[0] {
		stream.WriteString(ComponentRouteLinkKind)
	} else {
		stream.WriteString(ComponentRouteKind)
	}
	count++
	if len(object.fieldSet_) > 1 && object.fieldSet_[1] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if len(object.fieldSet_) > 2 && object.fieldSet_[2] {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = len(object.fieldSet_) > 3 && object.fieldSet_[3]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("hostname")
		stream.WriteString(object.hostname)
		count++
	}
	present_ = len(object.fieldSet_) > 4 && object.fieldSet_[4]
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("tls_secret_ref")
		stream.WriteString(object.tlsSecretRef)
	}
	stream.WriteObjectEnd()
}

// UnmarshalComponentRoute reads a value of the 'component_route' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalComponentRoute(source interface{}) (object *ComponentRoute, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadComponentRoute(iterator)
	err = iterator.Error
	return
}

// ReadComponentRoute reads a value of the 'component_route' type from the given iterator.
func ReadComponentRoute(iterator *jsoniter.Iterator) *ComponentRoute {
	object := &ComponentRoute{
		fieldSet_: make([]bool, 5),
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == ComponentRouteLinkKind {
				object.fieldSet_[0] = true
			}
		case "id":
			object.id = iterator.ReadString()
			object.fieldSet_[1] = true
		case "href":
			object.href = iterator.ReadString()
			object.fieldSet_[2] = true
		case "hostname":
			value := iterator.ReadString()
			object.hostname = value
			object.fieldSet_[3] = true
		case "tls_secret_ref":
			value := iterator.ReadString()
			object.tlsSecretRef = value
			object.fieldSet_[4] = true
		default:
			iterator.ReadAny()
		}
	}
	return object
}
