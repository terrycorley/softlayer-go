package data_types

// The SoftLayer_Product_Package_Server data type contains summarized information for bare
// metal servers regarding pricing, processor stats, and feature sets.
//
// See https://sldn.softlayer.com/reference/datatypes/SoftLayer_Product_Package_Server
type SoftLayer_Product_Package_Server struct {
	Id                     int    `json:"id"`
	CatalogId              int    `json:"catalogId"`
	Datacenters            string `json:"datacenters"`
	DefaultRamCapacity     string `json:"defaultRamCapacity"`
	DualPathNetwork        bool   `json:"dualPathNetworkFlag"`
	Gpu                    bool   `json:"gpuFlag"`
	HourlyBilling          bool   `json:"hourlyBillingFlag"`
	ItemId                 int    `json:"itemId"`
	ItemPriceId            int    `json:"itemPriceId"`
	MaximumDriveCount      int    `json:"maximumDriveCount"`
	MaximumPortSpeed       string `json:"maximumPortSpeed"`
	MaximumRamCapacity     string `json:"maximumRamCapacity"`
	MinimumPortSpeed       string `json:"minimumPortSpeed"`
	Outlet                 bool   `json:"outletFlag"`
	PackageId              int    `json:"packageId"`
	PackageType            string `json:"packageType"`
	PowerServer            bool   `json:"powerServerFlag"`
	PresetId               int    `json:"presetId,omitempty"`
	PrivateNetworkOnly     bool   `json:"privateNetworkOnlyFlag"`
	ProcessorBusSpeed      string `json:"processorBusSpeed"`
	ProcessorCache         string `json:"processorCache"`
	ProcessorCores         int    `json:"processorCores"`
	ProcessorCount         int    `json:"processorCount"`
	ProcessorManufacturer  string `json:"processorManufacturer"`
	ProcessorModel         string `json:"processorModel"`
	ProcessorName          string `json:"processorName"`
	ProcessorPhysicalCores int    `json:"processorPhysicalCores"`
	ProcessorSpeed         string `json:"processorSpeed"`
	ProductName            string `json:"productName"`
	RedundantPower         bool   `json:"redundantPowerFlag"`
	SapCertifiedServer     bool   `json:"sapCertifiedServerFlag"`
	StartingHourlyPrice    string `json:"startingHourlyPrice,omitempty"`
	StartingMonthlyPrice   string `json:"startingMonthlyPrice,omitempty"`
	TotalCoreCount         int    `json:"totalCoreCount"`
	TxtTpm                 bool   `json:"txtTpmFlag"`
	UnitSize               int    `json:"unitSize"`
}
