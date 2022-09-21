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

// NetgroupsSettingsModifyReader is a Reader for the NetgroupsSettingsModify structure.
type NetgroupsSettingsModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupsSettingsModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupsSettingsModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupsSettingsModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupsSettingsModifyOK creates a NetgroupsSettingsModifyOK with default headers values
func NewNetgroupsSettingsModifyOK() *NetgroupsSettingsModifyOK {
	return &NetgroupsSettingsModifyOK{}
}

/* NetgroupsSettingsModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupsSettingsModifyOK struct {
}

func (o *NetgroupsSettingsModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings/{svm.uuid}][%d] netgroupsSettingsModifyOK ", 200)
}

func (o *NetgroupsSettingsModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetgroupsSettingsModifyDefault creates a NetgroupsSettingsModifyDefault with default headers values
func NewNetgroupsSettingsModifyDefault(code int) *NetgroupsSettingsModifyDefault {
	return &NetgroupsSettingsModifyDefault{
		_statusCode: code,
	}
}

/* NetgroupsSettingsModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |

*/
type NetgroupsSettingsModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the netgroups settings modify default response
func (o *NetgroupsSettingsModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetgroupsSettingsModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings/{svm.uuid}][%d] netgroups_settings_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NetgroupsSettingsModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupsSettingsModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
