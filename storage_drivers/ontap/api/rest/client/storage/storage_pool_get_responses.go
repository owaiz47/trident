// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// StoragePoolGetReader is a Reader for the StoragePoolGet structure.
type StoragePoolGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoragePoolGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoragePoolGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStoragePoolGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoragePoolGetOK creates a StoragePoolGetOK with default headers values
func NewStoragePoolGetOK() *StoragePoolGetOK {
	return &StoragePoolGetOK{}
}

/* StoragePoolGetOK describes a response with status code 200, with default header values.

OK
*/
type StoragePoolGetOK struct {
	Payload *models.StoragePool
}

func (o *StoragePoolGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storagePoolGetOK  %+v", 200, o.Payload)
}
func (o *StoragePoolGetOK) GetPayload() *models.StoragePool {
	return o.Payload
}

func (o *StoragePoolGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoragePoolGetDefault creates a StoragePoolGetDefault with default headers values
func NewStoragePoolGetDefault(code int) *StoragePoolGetDefault {
	return &StoragePoolGetDefault{
		_statusCode: code,
	}
}

/* StoragePoolGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 11206662 | There is no storage pool matching the specified UUID or name. |

*/
type StoragePoolGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the storage pool get default response
func (o *StoragePoolGetDefault) Code() int {
	return o._statusCode
}

func (o *StoragePoolGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/pools/{uuid}][%d] storage_pool_get default  %+v", o._statusCode, o.Payload)
}
func (o *StoragePoolGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StoragePoolGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
