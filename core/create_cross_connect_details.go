// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateCrossConnectDetails struct {

	// The OCID of the compartment to contain the cross-connect.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name of the FastConnect location where this cross-connect will be installed.
	// To get a list of the available locations, see
	// ListCrossConnectLocations.
	// Example: `CyrusOne, Chandler, AZ`
	LocationName *string `mandatory:"true" json:"locationName,omitempty"`

	// The port speed for this cross-connect. To get a list of the available port speeds, see
	// ListCrossconnectPortSpeedShapes.
	// Example: `10 Gbps`
	PortSpeedShapeName *string `mandatory:"true" json:"portSpeedShapeName,omitempty"`

	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupID *string `mandatory:"false" json:"crossConnectGroupId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// If you already have an existing cross-connect or cross-connect group at this FastConnect
	// location, and you want this new cross-connect to be on a different router (for the
	// purposes of redundancy), provide the OCID of that existing cross-connect or
	// cross-connect group.
	FarCrossConnectOrCrossConnectGroupID *string `mandatory:"false" json:"farCrossConnectOrCrossConnectGroupId,omitempty"`

	// If you already have an existing cross-connect or cross-connect group at this FastConnect
	// location, and you want this new cross-connect to be on the same router, provide the
	// OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupID *string `mandatory:"false" json:"nearCrossConnectOrCrossConnectGroupId,omitempty"`
}

func (model CreateCrossConnectDetails) String() string {
	return common.PointerString(model)
}