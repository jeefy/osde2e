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
	"net/http"
)

func readAddOnDeleteRequest(request *AddOnDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeAddOnDeleteRequest(request *AddOnDeleteRequest, writer io.Writer) error {
	return nil
}
func readAddOnDeleteResponse(response *AddOnDeleteResponse, reader io.Reader) error {
	return nil
}
func writeAddOnDeleteResponse(response *AddOnDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readAddOnGetRequest(request *AddOnGetServerRequest, r *http.Request) error {
	return nil
}
func writeAddOnGetRequest(request *AddOnGetRequest, writer io.Writer) error {
	return nil
}
func readAddOnGetResponse(response *AddOnGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalAddOn(reader)
	return err
}
func writeAddOnGetResponse(response *AddOnGetServerResponse, w http.ResponseWriter) error {
	return MarshalAddOn(response.body, w)
}
func readAddOnUpdateRequest(request *AddOnUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalAddOn(r.Body)
	return err
}
func writeAddOnUpdateRequest(request *AddOnUpdateRequest, writer io.Writer) error {
	return MarshalAddOn(request.body, writer)
}
func readAddOnUpdateResponse(response *AddOnUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalAddOn(reader)
	return err
}
func writeAddOnUpdateResponse(response *AddOnUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalAddOn(response.body, w)
}
