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

// MetroclusterSvmCollectionGetReader is a Reader for the MetroclusterSvmCollectionGet structure.
type MetroclusterSvmCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterSvmCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterSvmCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterSvmCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterSvmCollectionGetOK creates a MetroclusterSvmCollectionGetOK with default headers values
func NewMetroclusterSvmCollectionGetOK() *MetroclusterSvmCollectionGetOK {
	return &MetroclusterSvmCollectionGetOK{}
}

/* MetroclusterSvmCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterSvmCollectionGetOK struct {
	Payload *models.MetroclusterSvmResponse
}

func (o *MetroclusterSvmCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/svms][%d] metroclusterSvmCollectionGetOK  %+v", 200, o.Payload)
}
func (o *MetroclusterSvmCollectionGetOK) GetPayload() *models.MetroclusterSvmResponse {
	return o.Payload
}

func (o *MetroclusterSvmCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterSvmResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterSvmCollectionGetDefault creates a MetroclusterSvmCollectionGetDefault with default headers values
func NewMetroclusterSvmCollectionGetDefault(code int) *MetroclusterSvmCollectionGetDefault {
	return &MetroclusterSvmCollectionGetDefault{
		_statusCode: code,
	}
}

/* MetroclusterSvmCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type MetroclusterSvmCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the metrocluster svm collection get default response
func (o *MetroclusterSvmCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MetroclusterSvmCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/svms][%d] metrocluster_svm_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *MetroclusterSvmCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterSvmCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
