package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Location_Group_Regional_Service interface {
	Service

	GetAllObjects() ([]datatypes.SoftLayer_Location_Group, error)
	GetDatacenters(groupId int) ([]datatypes.SoftLayer_Location, error)
	GetLocations(groupId int) ([]datatypes.SoftLayer_Location, error)
	GetObject(groupId int) (datatypes.SoftLayer_Location_Group_Regional, error)
	GetPreferredDatacenter(groupId int) (datatypes.SoftLayer_Location, error)
	GetLocationGroupType(groupId int) (datatypes.SoftLayer_Location_Group_Type, error)
}
