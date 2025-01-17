// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPITroubleshootingAndMonitoringTraceflow.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_PHYSICAL = "PHYSICAL"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_LR = "LR"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_LS = "LS"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_DFW = "DFW"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_BRIDGE = "BRIDGE"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_EDGE_TUNNEL = "EDGE_TUNNEL"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_EDGE_HOSTSWITCH = "EDGE_HOSTSWITCH"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_FW_BRIDGE = "FW_BRIDGE"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_LOAD_BALANCER = "LOAD_BALANCER"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_NAT = "NAT"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_IPSEC = "IPSEC"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_SERVICE_INSERTION = "SERVICE_INSERTION"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_VMC = "VMC"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_SPOOFGUARD = "SPOOFGUARD"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_EDGE_FW = "EDGE_FW"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_DLB = "DLB"

// Possible value for ``componentType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_COMPONENT_TYPE_UNKNOWN = "UNKNOWN"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONFORWARDED = "TraceflowObservationForwarded"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDROPPED = "TraceflowObservationDropped"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDELIVERED = "TraceflowObservationDelivered"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRECEIVED = "TraceflowObservationReceived"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONFORWARDEDLOGICAL = "TraceflowObservationForwardedLogical"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDROPPEDLOGICAL = "TraceflowObservationDroppedLogical"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRECEIVEDLOGICAL = "TraceflowObservationReceivedLogical"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONREPLICATIONLOGICAL = "TraceflowObservationReplicationLogical"

// Possible value for ``resourceType`` of method ManagementPlaneAPITroubleshootingAndMonitoringTraceflow#getTraceflowObservations.
const ManagementPlaneAPITroubleshootingAndMonitoringTraceflow_GET_TRACEFLOW_OBSERVATIONS_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRELAYEDLOGICAL = "TraceflowObservationRelayedLogical"

func managementPlaneAPITroubleshootingAndMonitoringTraceflowCreateTraceflowInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_request"] = bindings.NewReferenceType(model.TraceflowRequestBindingType)
	fieldNameMap["traceflow_request"] = "TraceflowRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowCreateTraceflowOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TraceflowBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowCreateTraceflowRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traceflow_request"] = bindings.NewReferenceType(model.TraceflowRequestBindingType)
	fieldNameMap["traceflow_request"] = "TraceflowRequest"
	paramsTypeMap["traceflow_request"] = bindings.NewReferenceType(model.TraceflowRequestBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"traceflow_request",
		"POST",
		"/api/v1/traceflows",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowDeleteTraceflowInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_id"] = bindings.NewStringType()
	fieldNameMap["traceflow_id"] = "TraceflowId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowDeleteTraceflowOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowDeleteTraceflowRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traceflow_id"] = bindings.NewStringType()
	fieldNameMap["traceflow_id"] = "TraceflowId"
	paramsTypeMap["traceflow_id"] = bindings.NewStringType()
	paramsTypeMap["traceflowId"] = bindings.NewStringType()
	pathParams["traceflow_id"] = "traceflowId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/traceflows/{traceflowId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_id"] = bindings.NewStringType()
	fieldNameMap["traceflow_id"] = "TraceflowId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TraceflowBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traceflow_id"] = bindings.NewStringType()
	fieldNameMap["traceflow_id"] = "TraceflowId"
	paramsTypeMap["traceflow_id"] = bindings.NewStringType()
	paramsTypeMap["traceflowId"] = bindings.NewStringType()
	pathParams["traceflow_id"] = "traceflowId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/traceflows/{traceflowId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowObservationsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_id"] = bindings.NewStringType()
	fields["component_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["component_name"] = "ComponentName"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_node_name"] = "TransportNodeName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowObservationsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TraceflowObservationListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowGetTraceflowObservationsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traceflow_id"] = bindings.NewStringType()
	fields["component_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["component_name"] = "ComponentName"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_node_name"] = "TransportNodeName"
	paramsTypeMap["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["traceflow_id"] = bindings.NewStringType()
	paramsTypeMap["component_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_node_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["traceflowId"] = bindings.NewStringType()
	pathParams["traceflow_id"] = "traceflowId"
	queryParams["cursor"] = "cursor"
	queryParams["component_type"] = "component_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["component_name"] = "component_name"
	queryParams["resource_type"] = "resource_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["transport_node_name"] = "transport_node_name"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/traceflows/{traceflowId}/observations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowListTraceflowsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["lport_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowListTraceflowsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TraceflowListResultBindingType)
}

func managementPlaneAPITroubleshootingAndMonitoringTraceflowListTraceflowsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["lport_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lport_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["lport_id"] = "lport_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/traceflows",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
