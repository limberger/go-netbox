// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// IPAMServicesListReader is a Reader for the IPAMServicesList structure.
type IPAMServicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMServicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMServicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMServicesListOK creates a IPAMServicesListOK with default headers values
func NewIPAMServicesListOK() *IPAMServicesListOK {
	return &IPAMServicesListOK{}
}

/*IPAMServicesListOK handles this case with default header values.

IPAMServicesListOK ipam services list o k
*/
type IPAMServicesListOK struct {
	Payload *models.IPAMServicesListOKBody
}

func (o *IPAMServicesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/services/][%d] ipamServicesListOK  %+v", 200, o.Payload)
}

func (o *IPAMServicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMServicesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
