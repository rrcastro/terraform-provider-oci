// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDhcpOptionsRequest wrapper for the ListDhcpOptions operation
type ListDhcpOptionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" contributesTo:"query" name:"vcnId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListDhcpOptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListDhcpOptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState DhcpOptionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
}

func (request ListDhcpOptionsRequest) String() string {
	return common.PointerString(request)
}

// ListDhcpOptionsResponse wrapper for the ListDhcpOptions operation
type ListDhcpOptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []DhcpOptions instance
	Items []DhcpOptions `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDhcpOptionsResponse) String() string {
	return common.PointerString(response)
}

// ListDhcpOptionsSortByEnum Enum with underlying type: string
type ListDhcpOptionsSortByEnum string

// Set of constants representing the allowable values for ListDhcpOptionsSortBy
const (
	ListDhcpOptionsSortByTimecreated ListDhcpOptionsSortByEnum = "TIMECREATED"
	ListDhcpOptionsSortByDisplayname ListDhcpOptionsSortByEnum = "DISPLAYNAME"
	ListDhcpOptionsSortByUnknown     ListDhcpOptionsSortByEnum = "UNKNOWN"
)

var mappingListDhcpOptionsSortBy = map[string]ListDhcpOptionsSortByEnum{
	"TIMECREATED": ListDhcpOptionsSortByTimecreated,
	"DISPLAYNAME": ListDhcpOptionsSortByDisplayname,
	"UNKNOWN":     ListDhcpOptionsSortByUnknown,
}

// GetListDhcpOptionsSortByEnumValues Enumerates the set of values for ListDhcpOptionsSortBy
func GetListDhcpOptionsSortByEnumValues() []ListDhcpOptionsSortByEnum {
	values := make([]ListDhcpOptionsSortByEnum, 0)
	for _, v := range mappingListDhcpOptionsSortBy {
		if v != ListDhcpOptionsSortByUnknown {
			values = append(values, v)
		}
	}
	return values
}

// ListDhcpOptionsSortOrderEnum Enum with underlying type: string
type ListDhcpOptionsSortOrderEnum string

// Set of constants representing the allowable values for ListDhcpOptionsSortOrder
const (
	ListDhcpOptionsSortOrderAsc     ListDhcpOptionsSortOrderEnum = "ASC"
	ListDhcpOptionsSortOrderDesc    ListDhcpOptionsSortOrderEnum = "DESC"
	ListDhcpOptionsSortOrderUnknown ListDhcpOptionsSortOrderEnum = "UNKNOWN"
)

var mappingListDhcpOptionsSortOrder = map[string]ListDhcpOptionsSortOrderEnum{
	"ASC":     ListDhcpOptionsSortOrderAsc,
	"DESC":    ListDhcpOptionsSortOrderDesc,
	"UNKNOWN": ListDhcpOptionsSortOrderUnknown,
}

// GetListDhcpOptionsSortOrderEnumValues Enumerates the set of values for ListDhcpOptionsSortOrder
func GetListDhcpOptionsSortOrderEnumValues() []ListDhcpOptionsSortOrderEnum {
	values := make([]ListDhcpOptionsSortOrderEnum, 0)
	for _, v := range mappingListDhcpOptionsSortOrder {
		if v != ListDhcpOptionsSortOrderUnknown {
			values = append(values, v)
		}
	}
	return values
}
