// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimInterfacesListParams creates a new DcimInterfacesListParams object
// with the default values initialized.
func NewDcimInterfacesListParams() *DcimInterfacesListParams {
	var ()
	return &DcimInterfacesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesListParamsWithTimeout creates a new DcimInterfacesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfacesListParamsWithTimeout(timeout time.Duration) *DcimInterfacesListParams {
	var ()
	return &DcimInterfacesListParams{

		timeout: timeout,
	}
}

// NewDcimInterfacesListParamsWithContext creates a new DcimInterfacesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfacesListParamsWithContext(ctx context.Context) *DcimInterfacesListParams {
	var ()
	return &DcimInterfacesListParams{

		Context: ctx,
	}
}

// NewDcimInterfacesListParamsWithHTTPClient creates a new DcimInterfacesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfacesListParamsWithHTTPClient(client *http.Client) *DcimInterfacesListParams {
	var ()
	return &DcimInterfacesListParams{
		HTTPClient: client,
	}
}

/*DcimInterfacesListParams contains all the parameters to send to the API endpoint
for the dcim interfaces list operation typically these are written to a http.Request
*/
type DcimInterfacesListParams struct {

	/*Cabled*/
	Cabled *string
	/*ConnectionStatus*/
	ConnectionStatus *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *float64
	/*Enabled*/
	Enabled *string
	/*FormFactor*/
	FormFactor *string
	/*LagID*/
	LagID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MacAddress*/
	MacAddress *string
	/*MgmtOnly*/
	MgmtOnly *string
	/*Mtu*/
	Mtu *float64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Tag*/
	Tag *string
	/*Type*/
	Type *string
	/*Vlan*/
	Vlan *string
	/*VlanID*/
	VlanID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTimeout(timeout time.Duration) *DcimInterfacesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithContext(ctx context.Context) *DcimInterfacesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithHTTPClient(client *http.Client) *DcimInterfacesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithCabled(cabled *string) *DcimInterfacesListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnectionStatus adds the connectionStatus to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithConnectionStatus(connectionStatus *string) *DcimInterfacesListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithDevice adds the device to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDevice(device *string) *DcimInterfacesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDeviceID(deviceID *float64) *DcimInterfacesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDeviceID(deviceID *float64) {
	o.DeviceID = deviceID
}

// WithEnabled adds the enabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithEnabled(enabled *string) *DcimInterfacesListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithFormFactor adds the formFactor to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithFormFactor(formFactor *string) *DcimInterfacesListParams {
	o.SetFormFactor(formFactor)
	return o
}

// SetFormFactor adds the formFactor to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetFormFactor(formFactor *string) {
	o.FormFactor = formFactor
}

// WithLagID adds the lagID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithLagID(lagID *string) *DcimInterfacesListParams {
	o.SetLagID(lagID)
	return o
}

// SetLagID adds the lagId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetLagID(lagID *string) {
	o.LagID = lagID
}

// WithLimit adds the limit to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithLimit(limit *int64) *DcimInterfacesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddress(macAddress *string) *DcimInterfacesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMgmtOnly adds the mgmtOnly to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMgmtOnly(mgmtOnly *string) *DcimInterfacesListParams {
	o.SetMgmtOnly(mgmtOnly)
	return o
}

// SetMgmtOnly adds the mgmtOnly to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMgmtOnly(mgmtOnly *string) {
	o.MgmtOnly = mgmtOnly
}

// WithMtu adds the mtu to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtu(mtu *float64) *DcimInterfacesListParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtu(mtu *float64) {
	o.Mtu = mtu
}

// WithName adds the name to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithName(name *string) *DcimInterfacesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithOffset(offset *int64) *DcimInterfacesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithQ(q *string) *DcimInterfacesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTag(tag *string) *DcimInterfacesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithType adds the typeVar to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithType(typeVar *string) *DcimInterfacesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithVlan adds the vlan to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithVlan(vlan *string) *DcimInterfacesListParams {
	o.SetVlan(vlan)
	return o
}

// SetVlan adds the vlan to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetVlan(vlan *string) {
	o.Vlan = vlan
}

// WithVlanID adds the vlanID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithVlanID(vlanID *string) *DcimInterfacesListParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetVlanID(vlanID *string) {
	o.VlanID = vlanID
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID float64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatFloat64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}

	}

	if o.FormFactor != nil {

		// query param form_factor
		var qrFormFactor string
		if o.FormFactor != nil {
			qrFormFactor = *o.FormFactor
		}
		qFormFactor := qrFormFactor
		if qFormFactor != "" {
			if err := r.SetQueryParam("form_factor", qFormFactor); err != nil {
				return err
			}
		}

	}

	if o.LagID != nil {

		// query param lag_id
		var qrLagID string
		if o.LagID != nil {
			qrLagID = *o.LagID
		}
		qLagID := qrLagID
		if qLagID != "" {
			if err := r.SetQueryParam("lag_id", qLagID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}

	}

	if o.MgmtOnly != nil {

		// query param mgmt_only
		var qrMgmtOnly string
		if o.MgmtOnly != nil {
			qrMgmtOnly = *o.MgmtOnly
		}
		qMgmtOnly := qrMgmtOnly
		if qMgmtOnly != "" {
			if err := r.SetQueryParam("mgmt_only", qMgmtOnly); err != nil {
				return err
			}
		}

	}

	if o.Mtu != nil {

		// query param mtu
		var qrMtu float64
		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := swag.FormatFloat64(qrMtu)
		if qMtu != "" {
			if err := r.SetQueryParam("mtu", qMtu); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.Vlan != nil {

		// query param vlan
		var qrVlan string
		if o.Vlan != nil {
			qrVlan = *o.Vlan
		}
		qVlan := qrVlan
		if qVlan != "" {
			if err := r.SetQueryParam("vlan", qVlan); err != nil {
				return err
			}
		}

	}

	if o.VlanID != nil {

		// query param vlan_id
		var qrVlanID string
		if o.VlanID != nil {
			qrVlanID = *o.VlanID
		}
		qVlanID := qrVlanID
		if qVlanID != "" {
			if err := r.SetQueryParam("vlan_id", qVlanID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}