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

// StoragePoolModifyReader is a Reader for the StoragePoolModify structure.
type StoragePoolModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoragePoolModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoragePoolModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStoragePoolModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStoragePoolModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoragePoolModifyOK creates a StoragePoolModifyOK with default headers values
func NewStoragePoolModifyOK() *StoragePoolModifyOK {
	return &StoragePoolModifyOK{}
}

/* StoragePoolModifyOK describes a response with status code 200, with default header values.

OK
*/
type StoragePoolModifyOK struct {
	Payload *models.StoragePoolPatch
}

func (o *StoragePoolModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/pools/{uuid}][%d] storagePoolModifyOK  %+v", 200, o.Payload)
}
func (o *StoragePoolModifyOK) GetPayload() *models.StoragePoolPatch {
	return o.Payload
}

func (o *StoragePoolModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePoolPatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoragePoolModifyAccepted creates a StoragePoolModifyAccepted with default headers values
func NewStoragePoolModifyAccepted() *StoragePoolModifyAccepted {
	return &StoragePoolModifyAccepted{}
}

/* StoragePoolModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type StoragePoolModifyAccepted struct {
	Payload *models.StoragePoolPatch
}

func (o *StoragePoolModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /storage/pools/{uuid}][%d] storagePoolModifyAccepted  %+v", 202, o.Payload)
}
func (o *StoragePoolModifyAccepted) GetPayload() *models.StoragePoolPatch {
	return o.Payload
}

func (o *StoragePoolModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePoolPatch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoragePoolModifyDefault creates a StoragePoolModifyDefault with default headers values
func NewStoragePoolModifyDefault(code int) *StoragePoolModifyDefault {
	return &StoragePoolModifyDefault{
		_statusCode: code,
	}
}

/* StoragePoolModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 11211658 | Node does not have enough spare capacity. |
| 11211659 | Valid allocation unit input is required. |
| 11211662 | Specified node is not part of the storage pool. |
| 11211663 | Failed to reassign available capacity in the storage pool. |
| 11211664 | Could not fix the broken allocation unit for the storage pool. |
| 11212673 | Could not grow one or more aggregates. |
| 11212679 | Adding specified number of disks will expand storage pool beyond maximum supported disk limit. |
| 11212680 | Incorrect node specified. |
| 11212681 | 0 is an invalid value for disk_count. |
| 11212682 | Adding the specified number of disks will result in the storage pool reaching the maximum disk limit reserved for RAID-TEC use only. At this limit, the storage pool can only allocate capacity to aggregates containing RAID-TEC RAID groups. Existing aggregates containing RAID groups other than RAID-TEC will not automatically grow to the new capacity. |
| 11212683 | Renaming storage pool to new name failed. |
| 11212763 | Storage pool add job failed. |
| 11215657 | Storage pool PATCH request have missing parameters. |
| 11215658 | Storage pool PATCH request for reassign is invalid. |
| 11215659 | Storage pool PATCH request for reassign have invalid allocation unit count. |
| 11215660 | Storage pool PATCH request for reassign have invalid node name. |
| 11215662 | Storage pool PATCH request have invalid disk count. |

*/
type StoragePoolModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the storage pool modify default response
func (o *StoragePoolModifyDefault) Code() int {
	return o._statusCode
}

func (o *StoragePoolModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/pools/{uuid}][%d] storage_pool_modify default  %+v", o._statusCode, o.Payload)
}
func (o *StoragePoolModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StoragePoolModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
