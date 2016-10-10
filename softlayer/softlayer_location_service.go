package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Location_Service interface {
	Service

	GetObject(locationId int) (datatypes.SoftLayer_Location, error)
	GetDatacenters() ([]datatypes.SoftLayer_Location, error)
	GetAvailableObjectStorageDatacenters() ([]datatypes.SoftLayer_Location, error)
	GetLocationStatus(locationId int) (datatypes.SoftLayer_Location_Status, error)
	GetPriceGroups(locationId int) ([]datatypes.SoftLayer_Location_Group, error)

	GetRegions(locationId int) ([]datatypes.SoftLayer_Location_Region, error)
}
