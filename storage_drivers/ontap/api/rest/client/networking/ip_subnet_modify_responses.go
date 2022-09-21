// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IPSubnetModifyReader is a Reader for the IPSubnetModify structure.
type IPSubnetModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPSubnetModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPSubnetModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPSubnetModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPSubnetModifyOK creates a IPSubnetModifyOK with default headers values
func NewIPSubnetModifyOK() *IPSubnetModifyOK {
	return &IPSubnetModifyOK{}
}

/* IPSubnetModifyOK describes a response with status code 200, with default header values.

OK
*/
type IPSubnetModifyOK struct {
}

func (o *IPSubnetModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/subnets/{uuid}][%d] ipSubnetModifyOK ", 200)
}

func (o *IPSubnetModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIPSubnetModifyDefault creates a IPSubnetModifyDefault with default headers values
func NewIPSubnetModifyDefault(code int) *IPSubnetModifyDefault {
	return &IPSubnetModifyDefault{
		_statusCode: code,
	}
}

/* IPSubnetModifyDefault describes a response with status code -1, with default header values.

 Fill error codes below.
ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1377658 | Invalid gateway for subnet in IPspace. |
| 1377659 | Subnet would overlap with existing subnet named in IPspace. |
| 1377660 | A subnet with the name already exists in the IPspace. |
| 1377661 | Subnet in IPspace cannot use subnet address because that address is already used by subnet in the same IPspace. |
| 1377662 | The IP range address is not within the subnet in IPspace. |
| 1377663 | The specified IP address range of subnet in IPspace contains an address already in use by a LIF. |
| 1377664 | The specified IP address range of subnet in IPspace contains an address already in use by the Service Processor. |
| 1377673 | The addresses provided must have the same address family. |
| 53282568 | The subnet.address must be specified together with subnet.netmask. |
| 53282569 | The specified subnet.netmask is not valid. |
| 53282570 | Each pair of ranges must have ip_ranges.start less than or equal to ip_ranges.end. |
| 53282571 | The ip_ranges.start and ip_ranges.end fields must have the same number of items. |
| 53282572 | PATCH partially succeeded with error. |

*/
type IPSubnetModifyDefault struct {
	_statusCode int
}

// Code gets the status code for the ip subnet modify default response
func (o *IPSubnetModifyDefault) Code() int {
	return o._statusCode
}

func (o *IPSubnetModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/subnets/{uuid}][%d] ip_subnet_modify default ", o._statusCode)
}

func (o *IPSubnetModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
