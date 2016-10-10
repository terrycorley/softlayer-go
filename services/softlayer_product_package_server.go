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

type softLayer_Product_Package_Server_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Product_Package_Server_Service(client softlayer.Client) *softLayer_Product_Package_Server_Service {
	return &softLayer_Product_Package_Server_Service{
		client: client,
	}
}

func (slppss *softLayer_Product_Package_Server_Service) GetName() string {
	return "SoftLayer_Product_Package_Server"
}

// Get all the package servers available
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Product_Package_Server/getAllObjects
func (slppss *softLayer_Product_Package_Server_Service) GetAllObjects() ([]datatypes.SoftLayer_Product_Package_Server, error) {

	path := fmt.Sprintf("%s/%s", slppss.GetName(), "getAllObjects.json")

	responseBytes, errorCode, err := slppss.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return []datatypes.SoftLayer_Product_Package_Server{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getAllObjects, HTTP error code: '%d'", slppss.GetName(), errorCode)
		return []datatypes.SoftLayer_Product_Package_Server{}, errors.New(errorMessage)
	}

	var packages []datatypes.SoftLayer_Product_Package_Server

	err = json.Unmarshal(responseBytes, &packages)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Product_Package_Server{}, errors.New(errorMessage)
	}

	return packages, nil
}

// Get the package server object
//
// See https://sldn.softlayer.com/reference/services/SoftLayer_Product_Package_Server/getObject
func (slppss *softLayer_Product_Package_Server_Service) GetObject(packageServerId int) (datatypes.SoftLayer_Product_Package_Server, error) {
	path := fmt.Sprintf("%s/%d/%s", slppss.GetName(), packageServerId, "getObject.json")

	responseBytes, errorCode, err := slppss.client.GetHttpClient().DoRawHttpRequest(path, "GET", new(bytes.Buffer))

	if err != nil {
		return datatypes.SoftLayer_Product_Package_Server{}, err
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not %s#getObject, HTTP error code: '%d'", slppss.GetName(), errorCode)
		return datatypes.SoftLayer_Product_Package_Server{}, errors.New(errorMessage)
	}

	pps := datatypes.SoftLayer_Product_Package_Server{}

	err = json.Unmarshal(responseBytes, &pps)

	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, error message '%s'", err.Error())
		return datatypes.SoftLayer_Product_Package_Server{}, errors.New(errorMessage)
	}

	return pps, nil
}
