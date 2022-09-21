// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SwitchModifyReader is a Reader for the SwitchModify structure.
type SwitchModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwitchModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSwitchModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSwitchModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSwitchModifyAccepted creates a SwitchModifyAccepted with default headers values
func NewSwitchModifyAccepted() *SwitchModifyAccepted {
	return &SwitchModifyAccepted{}
}

/* SwitchModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SwitchModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SwitchModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switchModifyAccepted  %+v", 202, o.Payload)
}
func (o *SwitchModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SwitchModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchModifyDefault creates a SwitchModifyDefault with default headers values
func NewSwitchModifyDefault(code int) *SwitchModifyDefault {
	return &SwitchModifyDefault{
		_statusCode: code,
	}
}

/* SwitchModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 12517378 | Settings updated, but the IP address \"{address}\" is not reachable. Verify that the address is valid or check the network path. |
| 12517380 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid. |
| 12517382 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid (i.e., the SNMPv3 user exists in ONTAP and on the remote switch). If the \"snmp.user\" parameter is valid, verify that the SNMPv3 user's credentials are the same both in ONTAP as well as in the remote switch. If a custom engine-id was provided for the SNMPv3 user, ensure it is the same as that of the remote switch. |

*/
type SwitchModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the switch modify default response
func (o *SwitchModifyDefault) Code() int {
	return o._statusCode
}

func (o *SwitchModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/switches/{name}][%d] switch_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SwitchModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwitchModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
