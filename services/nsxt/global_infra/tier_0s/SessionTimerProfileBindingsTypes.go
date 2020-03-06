/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: SessionTimerProfileBindings.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tier_0s

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func sessionTimerProfileBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sessionTimerProfileBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func sessionTimerProfileBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
}

func sessionTimerProfileBindingsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func sessionTimerProfileBindingsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sessionTimerProfileBindingsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"session_timer_profile_binding_map",
		"PATCH",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func sessionTimerProfileBindingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
}

func sessionTimerProfileBindingsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"session_timer_profile_binding_map",
		"PUT",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

