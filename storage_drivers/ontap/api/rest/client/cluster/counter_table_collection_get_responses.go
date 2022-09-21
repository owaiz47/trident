// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CounterTableCollectionGetReader is a Reader for the CounterTableCollectionGet structure.
type CounterTableCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CounterTableCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCounterTableCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCounterTableCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCounterTableCollectionGetOK creates a CounterTableCollectionGetOK with default headers values
func NewCounterTableCollectionGetOK() *CounterTableCollectionGetOK {
	return &CounterTableCollectionGetOK{}
}

/* CounterTableCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CounterTableCollectionGetOK struct {
	Payload *models.CounterTableResponse
}

func (o *CounterTableCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counterTableCollectionGetOK  %+v", 200, o.Payload)
}
func (o *CounterTableCollectionGetOK) GetPayload() *models.CounterTableResponse {
	return o.Payload
}

func (o *CounterTableCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CounterTableResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCounterTableCollectionGetDefault creates a CounterTableCollectionGetDefault with default headers values
func NewCounterTableCollectionGetDefault(code int) *CounterTableCollectionGetDefault {
	return &CounterTableCollectionGetDefault{
		_statusCode: code,
	}
}

/* CounterTableCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 8585368 | The system has not completed it's initialization |

*/
type CounterTableCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the counter table collection get default response
func (o *CounterTableCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CounterTableCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/counter/tables][%d] counter_table_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *CounterTableCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CounterTableCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
