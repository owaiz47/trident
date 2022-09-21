// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// MultiAdminVerifyApprovalGroupDeleteReader is a Reader for the MultiAdminVerifyApprovalGroupDelete structure.
type MultiAdminVerifyApprovalGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyApprovalGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiAdminVerifyApprovalGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyApprovalGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyApprovalGroupDeleteOK creates a MultiAdminVerifyApprovalGroupDeleteOK with default headers values
func NewMultiAdminVerifyApprovalGroupDeleteOK() *MultiAdminVerifyApprovalGroupDeleteOK {
	return &MultiAdminVerifyApprovalGroupDeleteOK{}
}

/* MultiAdminVerifyApprovalGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type MultiAdminVerifyApprovalGroupDeleteOK struct {
}

func (o *MultiAdminVerifyApprovalGroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multiAdminVerifyApprovalGroupDeleteOK ", 200)
}

func (o *MultiAdminVerifyApprovalGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMultiAdminVerifyApprovalGroupDeleteDefault creates a MultiAdminVerifyApprovalGroupDeleteDefault with default headers values
func NewMultiAdminVerifyApprovalGroupDeleteDefault(code int) *MultiAdminVerifyApprovalGroupDeleteDefault {
	return &MultiAdminVerifyApprovalGroupDeleteDefault{
		_statusCode: code,
	}
}

/* MultiAdminVerifyApprovalGroupDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type MultiAdminVerifyApprovalGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the multi admin verify approval group delete default response
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/multi-admin-verify/approval-groups/{owner.uuid}/{name}][%d] multi_admin_verify_approval_group_delete default  %+v", o._statusCode, o.Payload)
}
func (o *MultiAdminVerifyApprovalGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyApprovalGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
