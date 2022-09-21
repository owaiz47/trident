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

// CounterRowCollectionGetReader is a Reader for the CounterRowCollectionGet structure.
type CounterRowCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CounterRowCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCounterRowCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCounterRowCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCounterRowCollectionGetOK creates a CounterRowCollectionGetOK with default headers values
func NewCounterRowCollectionGetOK() *CounterRowCollectionGetOK {
	return &CounterRowCollectionGetOK{}
}

/* CounterRowCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CounterRowCollectionGetOK struct {
	Payload *models.CounterRowResponse
}

func (o *CounterRowCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows][%d] counterRowCollectionGetOK  %+v", 200, o.Payload)
}
func (o *CounterRowCollectionGetOK) GetPayload() *models.CounterRowResponse {
	return o.Payload
}

func (o *CounterRowCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CounterRowResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCounterRowCollectionGetDefault creates a CounterRowCollectionGetDefault with default headers values
func NewCounterRowCollectionGetDefault(code int) *CounterRowCollectionGetDefault {
	return &CounterRowCollectionGetDefault{
		_statusCode: code,
	}
}

/* CounterRowCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 8585320 | Table requested is not found |
| 8586228 | Invalid counter name request. |
| 8586229 | Invalid counter property request. |

*/
type CounterRowCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the counter row collection get default response
func (o *CounterRowCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *CounterRowCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/counter/tables/{counter_table.name}/rows][%d] counter_row_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *CounterRowCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CounterRowCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
