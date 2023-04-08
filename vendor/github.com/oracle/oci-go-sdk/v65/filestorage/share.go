// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Share A share makes a file system directory accessible to SMB clients on
// the network. Each share resource belongs to exactly one share set.
// Each share is identified by a share name. The share name is unique to a
// SMB server. A SMB client sees only the share name. Share names are not
// case-sensitive. Share name cannot be an empty string. The maximum share name
// length is 80 unicode characters (or 240 bytes). Unicode share names are supported.
// The following special characters are not supported for share name:
// / \ : * ? < > " |
// For example, the following are acceptable:
//   * example and path
//   * example1 and example2
//   * example and example1
// The following examples are not acceptable:
//   * example and example/path
//   * / and /example
//   * my@ and my*
// Each share has a share comment. This is the description of the share
// when SMB client to list the SMB server's shares.
// Use `shareOptions` to control access to a share.
type Share struct {

	// Policies that apply to SMB requests made through this
	// share. `shareOptions` contains a sequential list of
	// `ClientShareOptions`. Each `ClientShareOptions` item defines the
	// share options that are applied to a specified
	// set of clients.
	// For each SMB request, the first `ShareOptions` option
	// in the list whose `source` attribute matches the source
	// IP address of the request is applied.
	// If a client source IP address does not match the `source`
	// property of any `ClientShareOptions` in the list, then the
	// share will be invisible to that client. This share will
	// not be returned by list shares calls made by the client
	// and any attempt to access the file system through
	// this share will result in an error.
	// **Shares without defined `ClientShareOptions` are invisible to all clients.**
	// If one share is invisible to a particular client, associated file
	// systems may still be accessible through other shares on the same
	// mount target.
	// To completely deny client access to a file system, be sure that the client
	// source IP address is not included in any share for any mount target
	// associated with the file system.
	ShareOptions []ClientShareOptions `mandatory:"true" json:"shareOptions"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the share set of this share is in.
	ShareSetId *string `mandatory:"true" json:"shareSetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this share's file system.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this share.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this share.
	LifecycleState ShareLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Share name used to access the associated file system.
	// Avoid entering confidential information.
	// Example: `accounting`
	ShareName *string `mandatory:"true" json:"shareName"`

	// A short comment description of the Share.
	// Avoid entering confidential information.
	// Example: `accounting`
	Comment *string `mandatory:"true" json:"comment"`

	// The date and time the share was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Share) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Share) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShareLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShareLifecycleStateEnum Enum with underlying type: string
type ShareLifecycleStateEnum string

// Set of constants representing the allowable values for ShareLifecycleStateEnum
const (
	ShareLifecycleStateCreating ShareLifecycleStateEnum = "CREATING"
	ShareLifecycleStateActive   ShareLifecycleStateEnum = "ACTIVE"
	ShareLifecycleStateDeleting ShareLifecycleStateEnum = "DELETING"
	ShareLifecycleStateDeleted  ShareLifecycleStateEnum = "DELETED"
)

var mappingShareLifecycleStateEnum = map[string]ShareLifecycleStateEnum{
	"CREATING": ShareLifecycleStateCreating,
	"ACTIVE":   ShareLifecycleStateActive,
	"DELETING": ShareLifecycleStateDeleting,
	"DELETED":  ShareLifecycleStateDeleted,
}

var mappingShareLifecycleStateEnumLowerCase = map[string]ShareLifecycleStateEnum{
	"creating": ShareLifecycleStateCreating,
	"active":   ShareLifecycleStateActive,
	"deleting": ShareLifecycleStateDeleting,
	"deleted":  ShareLifecycleStateDeleted,
}

// GetShareLifecycleStateEnumValues Enumerates the set of values for ShareLifecycleStateEnum
func GetShareLifecycleStateEnumValues() []ShareLifecycleStateEnum {
	values := make([]ShareLifecycleStateEnum, 0)
	for _, v := range mappingShareLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetShareLifecycleStateEnumStringValues Enumerates the set of values in String for ShareLifecycleStateEnum
func GetShareLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingShareLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShareLifecycleStateEnum(val string) (ShareLifecycleStateEnum, bool) {
	enum, ok := mappingShareLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
