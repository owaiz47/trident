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

// FcZoneCollectionGetReader is a Reader for the FcZoneCollectionGet structure.
type FcZoneCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcZoneCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcZoneCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcZoneCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcZoneCollectionGetOK creates a FcZoneCollectionGetOK with default headers values
func NewFcZoneCollectionGetOK() *FcZoneCollectionGetOK {
	return &FcZoneCollectionGetOK{}
}

/* FcZoneCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FcZoneCollectionGetOK struct {
	Payload *models.FcZoneResponse
}

func (o *FcZoneCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/zones][%d] fcZoneCollectionGetOK  %+v", 200, o.Payload)
}
func (o *FcZoneCollectionGetOK) GetPayload() *models.FcZoneResponse {
	return o.Payload
}

func (o *FcZoneCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcZoneCollectionGetDefault creates a FcZoneCollectionGetDefault with default headers values
func NewFcZoneCollectionGetDefault(code int) *FcZoneCollectionGetDefault {
	return &FcZoneCollectionGetDefault{
		_statusCode: code,
	}
}

/* FcZoneCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5375053 | The Fibre Channel fabric specified by name in the request URI was not found in the FC fabric cache. |

*/
type FcZoneCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc zone collection get default response
func (o *FcZoneCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *FcZoneCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/fc/fabrics/{fabric.name}/zones][%d] fc_zone_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *FcZoneCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcZoneCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
