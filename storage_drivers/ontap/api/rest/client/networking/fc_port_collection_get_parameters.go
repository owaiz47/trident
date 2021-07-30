// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewFcPortCollectionGetParams creates a new FcPortCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcPortCollectionGetParams() *FcPortCollectionGetParams {
	return &FcPortCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcPortCollectionGetParamsWithTimeout creates a new FcPortCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFcPortCollectionGetParamsWithTimeout(timeout time.Duration) *FcPortCollectionGetParams {
	return &FcPortCollectionGetParams{
		timeout: timeout,
	}
}

// NewFcPortCollectionGetParamsWithContext creates a new FcPortCollectionGetParams object
// with the ability to set a context for a request.
func NewFcPortCollectionGetParamsWithContext(ctx context.Context) *FcPortCollectionGetParams {
	return &FcPortCollectionGetParams{
		Context: ctx,
	}
}

// NewFcPortCollectionGetParamsWithHTTPClient creates a new FcPortCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcPortCollectionGetParamsWithHTTPClient(client *http.Client) *FcPortCollectionGetParams {
	return &FcPortCollectionGetParams{
		HTTPClient: client,
	}
}

/* FcPortCollectionGetParams contains all the parameters to send to the API endpoint
   for the fc port collection get operation.

   Typically these are written to a http.Request.
*/
type FcPortCollectionGetParams struct {

	/* Description.

	   Filter by description
	*/
	DescriptionQueryParameter *string

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* FabricConnected.

	   Filter by fabric.connected
	*/
	FabricConnectedQueryParameter *bool

	/* FabricConnectedSpeed.

	   Filter by fabric.connected_speed
	*/
	FabricConnectedSpeedQueryParameter *int64

	/* FabricName.

	   Filter by fabric.name
	*/
	FabricNameQueryParameter *string

	/* FabricPortAddress.

	   Filter by fabric.port_address
	*/
	FabricPortAddressQueryParameter *string

	/* FabricSwitchPort.

	   Filter by fabric.switch_port
	*/
	FabricSwitchPortQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* PhysicalProtocol.

	   Filter by physical_protocol
	*/
	PhysicalProtocolQueryParameter *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SpeedConfigured.

	   Filter by speed.configured
	*/
	SpeedConfiguredQueryParameter *string

	/* SpeedMaximum.

	   Filter by speed.maximum
	*/
	SpeedMaximumQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* SupportedProtocols.

	   Filter by supported_protocols
	*/
	SupportedProtocolsQueryParameter *string

	/* TransceiverCapabilities.

	   Filter by transceiver.capabilities
	*/
	TransceiverCapabilitiesQueryParameter *int64

	/* TransceiverFormFactor.

	   Filter by transceiver.form-factor
	*/
	TransceiverFormFactorQueryParameter *string

	/* TransceiverManufacturer.

	   Filter by transceiver.manufacturer
	*/
	TransceiverManufacturerQueryParameter *string

	/* TransceiverPartNumber.

	   Filter by transceiver.part_number
	*/
	TransceiverPartNumberQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	/* Wwnn.

	   Filter by wwnn
	*/
	WwnnQueryParameter *string

	/* Wwpn.

	   Filter by wwpn
	*/
	WwpnQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortCollectionGetParams) WithDefaults() *FcPortCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := FcPortCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fc port collection get params
func (o *FcPortCollectionGetParams) WithTimeout(timeout time.Duration) *FcPortCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc port collection get params
func (o *FcPortCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc port collection get params
func (o *FcPortCollectionGetParams) WithContext(ctx context.Context) *FcPortCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc port collection get params
func (o *FcPortCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc port collection get params
func (o *FcPortCollectionGetParams) WithHTTPClient(client *http.Client) *FcPortCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc port collection get params
func (o *FcPortCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescriptionQueryParameter adds the description to the fc port collection get params
func (o *FcPortCollectionGetParams) WithDescriptionQueryParameter(description *string) *FcPortCollectionGetParams {
	o.SetDescriptionQueryParameter(description)
	return o
}

// SetDescriptionQueryParameter adds the description to the fc port collection get params
func (o *FcPortCollectionGetParams) SetDescriptionQueryParameter(description *string) {
	o.DescriptionQueryParameter = description
}

// WithEnabledQueryParameter adds the enabled to the fc port collection get params
func (o *FcPortCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *FcPortCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the fc port collection get params
func (o *FcPortCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFabricConnectedQueryParameter adds the fabricConnected to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFabricConnectedQueryParameter(fabricConnected *bool) *FcPortCollectionGetParams {
	o.SetFabricConnectedQueryParameter(fabricConnected)
	return o
}

// SetFabricConnectedQueryParameter adds the fabricConnected to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFabricConnectedQueryParameter(fabricConnected *bool) {
	o.FabricConnectedQueryParameter = fabricConnected
}

// WithFabricConnectedSpeedQueryParameter adds the fabricConnectedSpeed to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFabricConnectedSpeedQueryParameter(fabricConnectedSpeed *int64) *FcPortCollectionGetParams {
	o.SetFabricConnectedSpeedQueryParameter(fabricConnectedSpeed)
	return o
}

// SetFabricConnectedSpeedQueryParameter adds the fabricConnectedSpeed to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFabricConnectedSpeedQueryParameter(fabricConnectedSpeed *int64) {
	o.FabricConnectedSpeedQueryParameter = fabricConnectedSpeed
}

// WithFabricNameQueryParameter adds the fabricName to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFabricNameQueryParameter(fabricName *string) *FcPortCollectionGetParams {
	o.SetFabricNameQueryParameter(fabricName)
	return o
}

// SetFabricNameQueryParameter adds the fabricName to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFabricNameQueryParameter(fabricName *string) {
	o.FabricNameQueryParameter = fabricName
}

// WithFabricPortAddressQueryParameter adds the fabricPortAddress to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFabricPortAddressQueryParameter(fabricPortAddress *string) *FcPortCollectionGetParams {
	o.SetFabricPortAddressQueryParameter(fabricPortAddress)
	return o
}

// SetFabricPortAddressQueryParameter adds the fabricPortAddress to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFabricPortAddressQueryParameter(fabricPortAddress *string) {
	o.FabricPortAddressQueryParameter = fabricPortAddress
}

// WithFabricSwitchPortQueryParameter adds the fabricSwitchPort to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFabricSwitchPortQueryParameter(fabricSwitchPort *string) *FcPortCollectionGetParams {
	o.SetFabricSwitchPortQueryParameter(fabricSwitchPort)
	return o
}

// SetFabricSwitchPortQueryParameter adds the fabricSwitchPort to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFabricSwitchPortQueryParameter(fabricSwitchPort *string) {
	o.FabricSwitchPortQueryParameter = fabricSwitchPort
}

// WithFields adds the fields to the fc port collection get params
func (o *FcPortCollectionGetParams) WithFields(fields []string) *FcPortCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the fc port collection get params
func (o *FcPortCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the fc port collection get params
func (o *FcPortCollectionGetParams) WithMaxRecords(maxRecords *int64) *FcPortCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the fc port collection get params
func (o *FcPortCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the fc port collection get params
func (o *FcPortCollectionGetParams) WithNameQueryParameter(name *string) *FcPortCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the fc port collection get params
func (o *FcPortCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNodeNameQueryParameter adds the nodeName to the fc port collection get params
func (o *FcPortCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *FcPortCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the fc port collection get params
func (o *FcPortCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the fc port collection get params
func (o *FcPortCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *FcPortCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the fc port collection get params
func (o *FcPortCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the fc port collection get params
func (o *FcPortCollectionGetParams) WithOrderBy(orderBy []string) *FcPortCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the fc port collection get params
func (o *FcPortCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPhysicalProtocolQueryParameter adds the physicalProtocol to the fc port collection get params
func (o *FcPortCollectionGetParams) WithPhysicalProtocolQueryParameter(physicalProtocol *string) *FcPortCollectionGetParams {
	o.SetPhysicalProtocolQueryParameter(physicalProtocol)
	return o
}

// SetPhysicalProtocolQueryParameter adds the physicalProtocol to the fc port collection get params
func (o *FcPortCollectionGetParams) SetPhysicalProtocolQueryParameter(physicalProtocol *string) {
	o.PhysicalProtocolQueryParameter = physicalProtocol
}

// WithReturnRecords adds the returnRecords to the fc port collection get params
func (o *FcPortCollectionGetParams) WithReturnRecords(returnRecords *bool) *FcPortCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fc port collection get params
func (o *FcPortCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the fc port collection get params
func (o *FcPortCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *FcPortCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the fc port collection get params
func (o *FcPortCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSpeedConfiguredQueryParameter adds the speedConfigured to the fc port collection get params
func (o *FcPortCollectionGetParams) WithSpeedConfiguredQueryParameter(speedConfigured *string) *FcPortCollectionGetParams {
	o.SetSpeedConfiguredQueryParameter(speedConfigured)
	return o
}

// SetSpeedConfiguredQueryParameter adds the speedConfigured to the fc port collection get params
func (o *FcPortCollectionGetParams) SetSpeedConfiguredQueryParameter(speedConfigured *string) {
	o.SpeedConfiguredQueryParameter = speedConfigured
}

// WithSpeedMaximumQueryParameter adds the speedMaximum to the fc port collection get params
func (o *FcPortCollectionGetParams) WithSpeedMaximumQueryParameter(speedMaximum *string) *FcPortCollectionGetParams {
	o.SetSpeedMaximumQueryParameter(speedMaximum)
	return o
}

// SetSpeedMaximumQueryParameter adds the speedMaximum to the fc port collection get params
func (o *FcPortCollectionGetParams) SetSpeedMaximumQueryParameter(speedMaximum *string) {
	o.SpeedMaximumQueryParameter = speedMaximum
}

// WithStateQueryParameter adds the state to the fc port collection get params
func (o *FcPortCollectionGetParams) WithStateQueryParameter(state *string) *FcPortCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the fc port collection get params
func (o *FcPortCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithSupportedProtocolsQueryParameter adds the supportedProtocols to the fc port collection get params
func (o *FcPortCollectionGetParams) WithSupportedProtocolsQueryParameter(supportedProtocols *string) *FcPortCollectionGetParams {
	o.SetSupportedProtocolsQueryParameter(supportedProtocols)
	return o
}

// SetSupportedProtocolsQueryParameter adds the supportedProtocols to the fc port collection get params
func (o *FcPortCollectionGetParams) SetSupportedProtocolsQueryParameter(supportedProtocols *string) {
	o.SupportedProtocolsQueryParameter = supportedProtocols
}

// WithTransceiverCapabilitiesQueryParameter adds the transceiverCapabilities to the fc port collection get params
func (o *FcPortCollectionGetParams) WithTransceiverCapabilitiesQueryParameter(transceiverCapabilities *int64) *FcPortCollectionGetParams {
	o.SetTransceiverCapabilitiesQueryParameter(transceiverCapabilities)
	return o
}

// SetTransceiverCapabilitiesQueryParameter adds the transceiverCapabilities to the fc port collection get params
func (o *FcPortCollectionGetParams) SetTransceiverCapabilitiesQueryParameter(transceiverCapabilities *int64) {
	o.TransceiverCapabilitiesQueryParameter = transceiverCapabilities
}

// WithTransceiverFormFactorQueryParameter adds the transceiverFormFactor to the fc port collection get params
func (o *FcPortCollectionGetParams) WithTransceiverFormFactorQueryParameter(transceiverFormFactor *string) *FcPortCollectionGetParams {
	o.SetTransceiverFormFactorQueryParameter(transceiverFormFactor)
	return o
}

// SetTransceiverFormFactorQueryParameter adds the transceiverFormFactor to the fc port collection get params
func (o *FcPortCollectionGetParams) SetTransceiverFormFactorQueryParameter(transceiverFormFactor *string) {
	o.TransceiverFormFactorQueryParameter = transceiverFormFactor
}

// WithTransceiverManufacturerQueryParameter adds the transceiverManufacturer to the fc port collection get params
func (o *FcPortCollectionGetParams) WithTransceiverManufacturerQueryParameter(transceiverManufacturer *string) *FcPortCollectionGetParams {
	o.SetTransceiverManufacturerQueryParameter(transceiverManufacturer)
	return o
}

// SetTransceiverManufacturerQueryParameter adds the transceiverManufacturer to the fc port collection get params
func (o *FcPortCollectionGetParams) SetTransceiverManufacturerQueryParameter(transceiverManufacturer *string) {
	o.TransceiverManufacturerQueryParameter = transceiverManufacturer
}

// WithTransceiverPartNumberQueryParameter adds the transceiverPartNumber to the fc port collection get params
func (o *FcPortCollectionGetParams) WithTransceiverPartNumberQueryParameter(transceiverPartNumber *string) *FcPortCollectionGetParams {
	o.SetTransceiverPartNumberQueryParameter(transceiverPartNumber)
	return o
}

// SetTransceiverPartNumberQueryParameter adds the transceiverPartNumber to the fc port collection get params
func (o *FcPortCollectionGetParams) SetTransceiverPartNumberQueryParameter(transceiverPartNumber *string) {
	o.TransceiverPartNumberQueryParameter = transceiverPartNumber
}

// WithUUIDQueryParameter adds the uuid to the fc port collection get params
func (o *FcPortCollectionGetParams) WithUUIDQueryParameter(uuid *string) *FcPortCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the fc port collection get params
func (o *FcPortCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WithWwnnQueryParameter adds the wwnn to the fc port collection get params
func (o *FcPortCollectionGetParams) WithWwnnQueryParameter(wwnn *string) *FcPortCollectionGetParams {
	o.SetWwnnQueryParameter(wwnn)
	return o
}

// SetWwnnQueryParameter adds the wwnn to the fc port collection get params
func (o *FcPortCollectionGetParams) SetWwnnQueryParameter(wwnn *string) {
	o.WwnnQueryParameter = wwnn
}

// WithWwpnQueryParameter adds the wwpn to the fc port collection get params
func (o *FcPortCollectionGetParams) WithWwpnQueryParameter(wwpn *string) *FcPortCollectionGetParams {
	o.SetWwpnQueryParameter(wwpn)
	return o
}

// SetWwpnQueryParameter adds the wwpn to the fc port collection get params
func (o *FcPortCollectionGetParams) SetWwpnQueryParameter(wwpn *string) {
	o.WwpnQueryParameter = wwpn
}

// WriteToRequest writes these params to a swagger request
func (o *FcPortCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DescriptionQueryParameter != nil {

		// query param description
		var qrDescription string

		if o.DescriptionQueryParameter != nil {
			qrDescription = *o.DescriptionQueryParameter
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.FabricConnectedQueryParameter != nil {

		// query param fabric.connected
		var qrFabricConnected bool

		if o.FabricConnectedQueryParameter != nil {
			qrFabricConnected = *o.FabricConnectedQueryParameter
		}
		qFabricConnected := swag.FormatBool(qrFabricConnected)
		if qFabricConnected != "" {

			if err := r.SetQueryParam("fabric.connected", qFabricConnected); err != nil {
				return err
			}
		}
	}

	if o.FabricConnectedSpeedQueryParameter != nil {

		// query param fabric.connected_speed
		var qrFabricConnectedSpeed int64

		if o.FabricConnectedSpeedQueryParameter != nil {
			qrFabricConnectedSpeed = *o.FabricConnectedSpeedQueryParameter
		}
		qFabricConnectedSpeed := swag.FormatInt64(qrFabricConnectedSpeed)
		if qFabricConnectedSpeed != "" {

			if err := r.SetQueryParam("fabric.connected_speed", qFabricConnectedSpeed); err != nil {
				return err
			}
		}
	}

	if o.FabricNameQueryParameter != nil {

		// query param fabric.name
		var qrFabricName string

		if o.FabricNameQueryParameter != nil {
			qrFabricName = *o.FabricNameQueryParameter
		}
		qFabricName := qrFabricName
		if qFabricName != "" {

			if err := r.SetQueryParam("fabric.name", qFabricName); err != nil {
				return err
			}
		}
	}

	if o.FabricPortAddressQueryParameter != nil {

		// query param fabric.port_address
		var qrFabricPortAddress string

		if o.FabricPortAddressQueryParameter != nil {
			qrFabricPortAddress = *o.FabricPortAddressQueryParameter
		}
		qFabricPortAddress := qrFabricPortAddress
		if qFabricPortAddress != "" {

			if err := r.SetQueryParam("fabric.port_address", qFabricPortAddress); err != nil {
				return err
			}
		}
	}

	if o.FabricSwitchPortQueryParameter != nil {

		// query param fabric.switch_port
		var qrFabricSwitchPort string

		if o.FabricSwitchPortQueryParameter != nil {
			qrFabricSwitchPort = *o.FabricSwitchPortQueryParameter
		}
		qFabricSwitchPort := qrFabricSwitchPort
		if qFabricSwitchPort != "" {

			if err := r.SetQueryParam("fabric.switch_port", qFabricSwitchPort); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PhysicalProtocolQueryParameter != nil {

		// query param physical_protocol
		var qrPhysicalProtocol string

		if o.PhysicalProtocolQueryParameter != nil {
			qrPhysicalProtocol = *o.PhysicalProtocolQueryParameter
		}
		qPhysicalProtocol := qrPhysicalProtocol
		if qPhysicalProtocol != "" {

			if err := r.SetQueryParam("physical_protocol", qPhysicalProtocol); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SpeedConfiguredQueryParameter != nil {

		// query param speed.configured
		var qrSpeedConfigured string

		if o.SpeedConfiguredQueryParameter != nil {
			qrSpeedConfigured = *o.SpeedConfiguredQueryParameter
		}
		qSpeedConfigured := qrSpeedConfigured
		if qSpeedConfigured != "" {

			if err := r.SetQueryParam("speed.configured", qSpeedConfigured); err != nil {
				return err
			}
		}
	}

	if o.SpeedMaximumQueryParameter != nil {

		// query param speed.maximum
		var qrSpeedMaximum string

		if o.SpeedMaximumQueryParameter != nil {
			qrSpeedMaximum = *o.SpeedMaximumQueryParameter
		}
		qSpeedMaximum := qrSpeedMaximum
		if qSpeedMaximum != "" {

			if err := r.SetQueryParam("speed.maximum", qSpeedMaximum); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.SupportedProtocolsQueryParameter != nil {

		// query param supported_protocols
		var qrSupportedProtocols string

		if o.SupportedProtocolsQueryParameter != nil {
			qrSupportedProtocols = *o.SupportedProtocolsQueryParameter
		}
		qSupportedProtocols := qrSupportedProtocols
		if qSupportedProtocols != "" {

			if err := r.SetQueryParam("supported_protocols", qSupportedProtocols); err != nil {
				return err
			}
		}
	}

	if o.TransceiverCapabilitiesQueryParameter != nil {

		// query param transceiver.capabilities
		var qrTransceiverCapabilities int64

		if o.TransceiverCapabilitiesQueryParameter != nil {
			qrTransceiverCapabilities = *o.TransceiverCapabilitiesQueryParameter
		}
		qTransceiverCapabilities := swag.FormatInt64(qrTransceiverCapabilities)
		if qTransceiverCapabilities != "" {

			if err := r.SetQueryParam("transceiver.capabilities", qTransceiverCapabilities); err != nil {
				return err
			}
		}
	}

	if o.TransceiverFormFactorQueryParameter != nil {

		// query param transceiver.form-factor
		var qrTransceiverFormFactor string

		if o.TransceiverFormFactorQueryParameter != nil {
			qrTransceiverFormFactor = *o.TransceiverFormFactorQueryParameter
		}
		qTransceiverFormFactor := qrTransceiverFormFactor
		if qTransceiverFormFactor != "" {

			if err := r.SetQueryParam("transceiver.form-factor", qTransceiverFormFactor); err != nil {
				return err
			}
		}
	}

	if o.TransceiverManufacturerQueryParameter != nil {

		// query param transceiver.manufacturer
		var qrTransceiverManufacturer string

		if o.TransceiverManufacturerQueryParameter != nil {
			qrTransceiverManufacturer = *o.TransceiverManufacturerQueryParameter
		}
		qTransceiverManufacturer := qrTransceiverManufacturer
		if qTransceiverManufacturer != "" {

			if err := r.SetQueryParam("transceiver.manufacturer", qTransceiverManufacturer); err != nil {
				return err
			}
		}
	}

	if o.TransceiverPartNumberQueryParameter != nil {

		// query param transceiver.part_number
		var qrTransceiverPartNumber string

		if o.TransceiverPartNumberQueryParameter != nil {
			qrTransceiverPartNumber = *o.TransceiverPartNumberQueryParameter
		}
		qTransceiverPartNumber := qrTransceiverPartNumber
		if qTransceiverPartNumber != "" {

			if err := r.SetQueryParam("transceiver.part_number", qTransceiverPartNumber); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if o.WwnnQueryParameter != nil {

		// query param wwnn
		var qrWwnn string

		if o.WwnnQueryParameter != nil {
			qrWwnn = *o.WwnnQueryParameter
		}
		qWwnn := qrWwnn
		if qWwnn != "" {

			if err := r.SetQueryParam("wwnn", qWwnn); err != nil {
				return err
			}
		}
	}

	if o.WwpnQueryParameter != nil {

		// query param wwpn
		var qrWwpn string

		if o.WwpnQueryParameter != nil {
			qrWwpn = *o.WwpnQueryParameter
		}
		qWwpn := qrWwpn
		if qWwpn != "" {

			if err := r.SetQueryParam("wwpn", qWwpn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFcPortCollectionGet binds the parameter fields
func (o *FcPortCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamFcPortCollectionGet binds the parameter order_by
func (o *FcPortCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}