package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	common "github.com/maximilien/softlayer-go/common"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

type softLayer_Location_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Location_Service(client softlayer.Client) *softLayer_Location_Service {
	return &softLayer_Location_Service{
		client: client,
	}
}

func (slls *softLayer_Location_Service) GetName() string {
	return "SoftLayer_Location"
}

func (slls *softLayer_Location_Service) GetObject(locationId int) (datatypes.SoftLayer_Location, error) {
	objectMask := []string{
		"id",
		"longName",
		"name",
	}

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequestWithObjectMask(fmt.Sprintf("%s/%d/getObject.json", slls.GetName(), locationId), objectMask, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Location{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getObject, HTTP error code: '%d'", slls.GetName(), errorCode)
		return datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	location := datatypes.SoftLayer_Location{}
	err = json.Unmarshal(responseBytes, &location)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	return location, nil

}

func (slls *softLayer_Location_Service) GetDatacenters() ([]datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%s", slls.GetName(), "getDatacenters.json")

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	var locations []datatypes.SoftLayer_Location

	if err != nil {
		return locations, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getDatacenters, HTTP error code: '%d'", slls.GetName(), errorCode)
		return locations, errors.New(errorMessage)
	}

	err = json.Unmarshal(responseBytes, &locations)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return locations, errors.New(errorMessage)
	}

	return locations, nil

}

// Retrieve the location status
func (slls *softLayer_Location_Service) GetLocationStatus(locationId int) (datatypes.SoftLayer_Location_Status, error) {
	path := fmt.Sprintf("%s/%d/%s", slls.GetName(), locationId, "getLocationStatus.json")

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Location_Status{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getLocationStatus, HTTP error code: '%d'", slls.GetName(), errorCode)
		return datatypes.SoftLayer_Location_Status{}, errors.New(errorMessage)
	}

	status := datatypes.SoftLayer_Location_Status{}

	err = json.Unmarshal(responseBytes, &status)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Location_Status{}, errors.New(errorMessage)
	}

	return status, nil
}

// Object Storage is only available in select datacenters. This method will return all the datacenters where object storage is available.
func (slls *softLayer_Location_Service) GetAvailableObjectStorageDatacenters() ([]datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%s", slls.GetName(), "getObjectStorageDatacenters.json")

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getObjectStorageDatacenters, HTTP error code: '%d'", slls.GetName(), errorCode)
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

// A location can be a member of 1 or more Price Groups. This will show which groups to which a location belongs.
// http://sldn.softlayer.com/reference/services/SoftLayer_Location/getPriceGroups
func (slls *softLayer_Location_Service) GetPriceGroups(locationId int) ([]datatypes.SoftLayer_Location_Group, error) {
	path := fmt.Sprintf("%s/%d/getPriceGroups.json", slls.GetName(), locationId)

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location_Group{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getPriceGroups, HTTP error code: '%d'", slls.GetName(), errorCode)
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

// Retrieve regions that a location belongs to
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Location/getRegions
func (slls *softLayer_Location_Service) GetRegions(locationId int) ([]datatypes.SoftLayer_Location_Region, error) {
	path := fmt.Sprintf("%s/%d/%s", slls.GetName(), locationId, "getRegions.json")

	responseBytes, errorCode, err := slls.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Location_Region{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getRegions, HTTP error code: '%d'", slls.GetName(), errorCode)
		return []datatypes.SoftLayer_Location_Region{}, errors.New(errorMessage)
	}

	var regions []datatypes.SoftLayer_Location_Region

	err = json.Unmarshal(responseBytes, &regions)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Location_Region{}, errors.New(errorMessage)
	}

	return regions, nil
}
