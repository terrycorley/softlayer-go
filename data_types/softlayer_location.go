package data_types

type SoftLayer_Location struct {
	Id       int    `json:"id"`
	LongName string `json:"longName"`
	Name     string `json:"name"`
}

// SoftLayer_Location_Status models the state of any location. SoftLayer uses the following status codes:
// *ACTIVE: The location is currently active and available for public usage.
// *PLANNED: Used when a location is planned but not yet active.
// *RETIRED: Used when a location has been retired and no longer active.
// Locations in use should stay in the ACTIVE state.
type SoftLayer_Location_Status struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

// See http://sldn.softlayer.com/reference/datatypes/SoftLayer_Location_Group
type SoftLayer_Location_Group struct {
	Id                  int                            `json:"id"`
	Description         string                         `json:"description"`
	LocationGroupTypeId int                            `json:"locationGroupTypeId"`
	LocationGroupType   *SoftLayer_Location_Group_Type `json:"locationGroupType"`
}

// A region is made up of a keyname and a description of that region. A region keyname
// can be used as part of an order. Check the SoftLayer_Product_Order service for more details.
type SoftLayer_Location_Region struct {
	Description string `json:"description"`
	Keyname     string `json:"keyname"`
}

type SoftLayer_Location_Group_Type struct {
	Name string `json:"name"`
}

// See http://sldn.softlayer.com/reference/datatypes/SoftLayer_Location_Group_Regional
type SoftLayer_Location_Group_Regional struct {
	Id                  int    `json:"id"`
	Description         string `json:"description"`
	LocationGroupTypeId int    `json:"locationGroupTypeId"`
	Name                string `json:"name"`
	SecurityLevelId     int    `json:"securityLevelId,omitempty"`
}
