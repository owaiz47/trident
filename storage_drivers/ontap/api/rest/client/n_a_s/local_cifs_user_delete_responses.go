// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LocalCifsUserDeleteReader is a Reader for the LocalCifsUserDelete structure.
type LocalCifsUserDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsUserDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsUserDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsUserDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsUserDeleteOK creates a LocalCifsUserDeleteOK with default headers values
func NewLocalCifsUserDeleteOK() *LocalCifsUserDeleteOK {
	return &LocalCifsUserDeleteOK{}
}

/* LocalCifsUserDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsUserDeleteOK struct {
}

func (o *LocalCifsUserDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] localCifsUserDeleteOK ", 200)
}

func (o *LocalCifsUserDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsUserDeleteDefault creates a LocalCifsUserDeleteDefault with default headers values
func NewLocalCifsUserDeleteDefault(code int) *LocalCifsUserDeleteDefault {
	return &LocalCifsUserDeleteDefault{
		_statusCode: code,
	}
}

/* LocalCifsUserDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 655735     | The local Administrator account cannot be deleted. |

*/
type LocalCifsUserDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs user delete default response
func (o *LocalCifsUserDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsUserDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-users/{svm.uuid}/{sid}][%d] local_cifs_user_delete default  %+v", o._statusCode, o.Payload)
}
func (o *LocalCifsUserDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsUserDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}