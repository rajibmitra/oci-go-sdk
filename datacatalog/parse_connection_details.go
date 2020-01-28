// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ParseConnectionDetails Parse connections from the connection metadata and oracle wallet file.
type ParseConnectionDetails struct {
	ConnectionDetail *Connection `mandatory:"false" json:"connectionDetail"`

	// The information used to parse the connection from the wallet file payload.
	ConnectionPayload []byte `mandatory:"false" json:"connectionPayload"`
}

func (m ParseConnectionDetails) String() string {
	return common.PointerString(m)
}