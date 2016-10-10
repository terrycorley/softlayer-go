package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Product_Package_Server_Service interface {
	Service

	GetAllObjects() ([]datatypes.SoftLayer_Product_Package_Server, error)
	GetObject(packageServerId int) (datatypes.SoftLayer_Product_Package_Server, error)
}
