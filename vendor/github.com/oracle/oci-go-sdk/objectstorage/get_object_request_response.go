// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetObjectRequest wrapper for the GetObject operation
type GetObjectRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// The entity tag to match. For creating and committing a multipart upload to an object, this is the entity tag of the target object.
	// For uploading a part, this is the entity tag of the target part.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag to avoid matching. The only valid value is ‘*’, which indicates that the request should fail if the object already exists.
	// For creating and committing a multipart upload, this is the entity tag of the target object. For uploading a part, this is the entity tag
	// of the target part.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Optional byte range to fetch, as described in [RFC 7233](https://tools.ietf.org/rfc/rfc7233), section 2.1.
	// Note, only a single range of bytes is supported.
	Range *string `mandatory:"false" contributesTo:"header" name:"range"`
}

func (request GetObjectRequest) String() string {
	return common.PointerString(request)
}

// GetObjectResponse wrapper for the GetObject operation
type GetObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, please provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The entity tag for the object.
	ETag *string `presentIn:"header" name:"etag"`

	// The user-defined metadata for the object.
	OpcMeta map[string]string `presentIn:"header-collection" prefix:"opc-meta-"`

	// The object size in bytes.
	ContentLength *int `presentIn:"header" name:"content-length"`

	// Content-Range header for range requests, per [RFC 7233](https://tools.ietf.org/rfc/rfc7233), section 4.2.
	ContentRange *string `presentIn:"header" name:"content-range"`

	// Content-MD5 header, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// Unavailable for objects uploaded using multipart upload.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// Only applicable to objects uploaded using multipart upload.
	// Base-64 representation of the multipart object hash.
	// The multipart object hash is calculated by taking the MD5 hashes of the parts,
	// concatenating the binary representation of those hashes in order of their part numbers,
	// and then calculating the MD5 hash of the concatenated values.
	OpcMultipartMd5 *string `presentIn:"header" name:"opc-multipart-md5"`

	// Content-Type header, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.17.
	ContentType *string `presentIn:"header" name:"content-type"`

	// Content-Language header, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.12.
	ContentLanguage *string `presentIn:"header" name:"content-language"`

	// Content-Encoding header, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.11.
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// The object modification time, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetObjectResponse) String() string {
	return common.PointerString(response)
}
