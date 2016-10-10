package data_types

type SoftLayer_Product_Package_Order_Configuration struct {
	Id             int `json:"id"`
	IsRequired     int `json:"isRequired"`
	ItemCategoryId int `json:"itemCategoryId"`
	OrderStepId    int `json:"orderStepId"`
	PackageId      int `json:"packageId"`
}
