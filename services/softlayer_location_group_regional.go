package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/maximilien/softlayer-go/common"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Location_Group_Regional_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Location_Group_Regional_Service(client softlayer.Client) *softLayer_Location_Group_Regional_Service {
	return &softLayer_Location_Group_Regional_Service{
		client: client,
	}
}

func (sllgrs *softLayer_Location_Group_Regional_Service) GetName() string {
	return "SoftLayer_Location_Group_Regional"
}

// Get all regional groups
//
// See: http://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getAllObjects
func (sllgrs *softLayer_Location_Group_Regional_Service) GetAllObjects() ([]datatypes.SoftLayer_Location_Group, error) {
	path := fmt.Sprintf("%s/%s", sllgrs.GetName(), "getAllObjects.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location_Group{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getAllObjects, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return []datatypes.SoftLayer_Location_Group{}, errors.New(errorMessage)
	}

	var locationGroups []datatypes.SoftLayer_Location_Group

	err = json.Unmarshal(responseBytes, &locationGroups)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Location_Group{}, errors.New(errorMessage)
	}

	return locationGroups, nil
}

// Retrieve the datacenters in a SoftLayer_Location_Group_Regional
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getDatacenters
func (sllgrs *softLayer_Location_Group_Regional_Service) GetDatacenters(groupId int) ([]datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%d/%s", sllgrs.GetName(), groupId, "getDatacenters.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getDatacenters, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	var datacenters []datatypes.SoftLayer_Location

	err = json.Unmarshal(responseBytes, &datacenters)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	return datacenters, nil
}

// Retrieve the locations in a SoftLayer_Location_Group_Regional
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getLocations
func (sllgrs *softLayer_Location_Group_Regional_Service) GetLocations(groupId int) ([]datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%d/%s", sllgrs.GetName(), groupId, "getLocations.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getLocations, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	var locations []datatypes.SoftLayer_Location

	err = json.Unmarshal(responseBytes, &locations)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	return locations, nil
}

// Retrieve a SoftLayer_Location_Group_Regional object
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getObject
func (sllgrs *softLayer_Location_Group_Regional_Service) GetObject(groupId int) (datatypes.SoftLayer_Location_Group_Regional, error) {
	path := fmt.Sprintf("%s/%d/%s", sllgrs.GetName(), groupId, "getObject.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Location_Group_Regional{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getObject, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return datatypes.SoftLayer_Location_Group_Regional{}, errors.New(errorMessage)
	}

	var groupRegional datatypes.SoftLayer_Location_Group_Regional

	err = json.Unmarshal(responseBytes, &groupRegional)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Location_Group_Regional{}, errors.New(errorMessage)
	}

	return groupRegional, nil
}

// Retrieve the preferred datacenter for a SoftLayer_Location_Group_Regional object
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getPreferredDatacenter
func (sllgrs *softLayer_Location_Group_Regional_Service) GetPreferredDatacenter(groupId int) (datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%d/%s", sllgrs.GetName(), groupId, "getPreferredDatacenter.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Location{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getPreferredDatacenter, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	var datacenter datatypes.SoftLayer_Location

	err = json.Unmarshal(responseBytes, &datacenter)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	return datacenter, nil
}

// Retrieve the type for the SoftLayer_Location_Group_Regional
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location_Group_Regional/getLocationGroupType
func (sllgrs *softLayer_Location_Group_Regional_Service) GetLocationGroupType(groupId int) (datatypes.SoftLayer_Location_Group_Type, error) {
	path := fmt.Sprintf("%s/%d/%s", sllgrs.GetName(), groupId, "getLocationGroupType.json")

	responseBytes, errorCode, err := sllgrs.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Location_Group_Type{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getLocationGroupType, HTTP error code: '%d'", sllgrs.GetName(), errorCode)
		return datatypes.SoftLayer_Location_Group_Type{}, errors.New(errorMessage)
	}

	groupType := datatypes.SoftLayer_Location_Group_Type{}

	err = json.Unmarshal(responseBytes, &groupType)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Location_Group_Type{}, errors.New(errorMessage)
	}

	return groupType, nil
}
