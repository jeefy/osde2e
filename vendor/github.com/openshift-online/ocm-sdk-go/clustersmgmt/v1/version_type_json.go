/*
Copyright (c) 2019 Red Hat, Inc.

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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalVersion writes a value of the 'version' type to the given writer.
func MarshalVersion(object *Version, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeVersion(object, stream)
	stream.Flush()
	return stream.Error
}

// writeVersion writes a value of the 'version' type to the given stream.
func writeVersion(object *Version, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if count > 0 {
		stream.WriteMore()
	}
	stream.WriteObjectField("kind")
	if object.link {
		stream.WriteString(VersionLinkKind)
	} else {
		stream.WriteString(VersionKind)
	}
	count++
	if object.id != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(*object.id)
		count++
	}
	if object.href != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(*object.href)
		count++
	}
	if object.default_ != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("default")
		stream.WriteBool(*object.default_)
		count++
	}
	if object.enabled != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("enabled")
		stream.WriteBool(*object.enabled)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalVersion reads a value of the 'version' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalVersion(source interface{}) (object *Version, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readVersion(iterator)
	err = iterator.Error
	return
}

// readVersion reads a value of the 'version' type from the given iterator.
func readVersion(iterator *jsoniter.Iterator) *Version {
	object := &Version{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			object.link = value == VersionLinkKind
		case "id":
			value := iterator.ReadString()
			object.id = &value
		case "href":
			value := iterator.ReadString()
			object.href = &value
		case "default":
			value := iterator.ReadBool()
			object.default_ = &value
		case "enabled":
			value := iterator.ReadBool()
			object.enabled = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
