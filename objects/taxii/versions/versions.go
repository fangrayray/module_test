// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package versions

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Versions - This type implements the TAXII 2 Versions Resource and defines
all of the properties and methods needed to create and work with the TAXII
Versions Resource.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint retrieves a list of versions of an object in a Collection. This
list can be used to decide whether it's worth retrieving the actual objects, or
if new versions have been added.

If a client fails authentication then this endpoint MUST return either an HTTP
401 (Unauthorized) or an HTTP 404 (Not Found).

If the Collection specifies can_read as false for a particular client, this
Endpoint MUST return an HTTP HTTP 403 (Forbidden) or HTTP 404 (Not Found) error.
*/
type Versions struct {
	More     bool     `json:"more,omitempty"`
	Versions []string `json:"versions,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Versions object and return it as a
pointer.
*/
func New() *Versions {
	var obj Versions
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Versions
// ----------------------------------------------------------------------

/*
AddVersion - This method will take in a version and append it to the Versions
property.
*/
func (o *Versions) AddVersion(v string) error {
	o.Versions = append(o.Versions, v)
	return nil
}

/*
GetMore - This method will return the more property
*/
func (o *Versions) GetMore() bool {
	return o.More
}

/*
SetMore - This method will set the more property to true
*/
func (o *Versions) SetMore() error {
	o.More = true
	return nil
}
