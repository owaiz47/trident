// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetgroupsSettingsGetReader is a Reader for the NetgroupsSettingsGet structure.
type NetgroupsSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupsSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupsSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupsSettingsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupsSettingsGetOK creates a NetgroupsSettingsGetOK with default headers values
func NewNetgroupsSettingsGetOK() *NetgroupsSettingsGetOK {
	return &NetgroupsSettingsGetOK{}
}

/* NetgroupsSettingsGetOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupsSettingsGetOK struct {
	Payload *models.NetgroupsSettings
}

func (o *NetgroupsSettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings/{svm.uuid}][%d] netgroupsSettingsGetOK  %+v", 200, o.Payload)
}
func (o *NetgroupsSettingsGetOK) GetPayload() *models.NetgroupsSettings {
	return o.Payload
}

func (o *NetgroupsSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetgroupsSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetgroupsSettingsGetDefault creates a NetgroupsSettingsGetDefault with default headers values
func NewNetgroupsSettingsGetDefault(code int) *NetgroupsSettingsGetDefault {
	return &NetgroupsSettingsGetDefault{
		_statusCode: code,
	}
}

/* NetgroupsSettingsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetgroupsSettingsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the netgroups settings get default response
func (o *NetgroupsSettingsGetDefault) Code() int {
	return o._statusCode
}

func (o *NetgroupsSettingsGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings/{svm.uuid}][%d] netgroups_settings_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetgroupsSettingsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupsSettingsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
