package cloudamqp

import (
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

// Instance represents a CloudAMQP instance.
// Based on https://customer.cloudamqp.com/team/api.
type Instance struct {
	ID     int    `json:"id"`
	Plan   string `json:"plan"`
	Region string `json:"region"`
	Name   string `json:"name"`
	URL    string `json:"url,omitempty"`
	APIKey string `json:"apikey,omitempty"`
}

// InstanceService provides methods for accessing CloudAMQP instance API endpoints.
// https://customer.cloudamqp.com/team/api
type InstanceService struct {
	sling *sling.Sling
}

func newInstanceService(sling *sling.Sling) *InstanceService {
	return &InstanceService{
		sling: sling.Path("instances"),
	}
}

// List instances available to the authenticated session.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) List() ([]Instance, *http.Response, error) {
	instances := new([]Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").Receive(instances, apiError)
	return *instances, resp, err
}

// Get a CloudAMQP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Get(id int) (*Instance, *http.Response, error) {
	instance := new(Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Path("instances/").Get(strconv.Itoa(id)).Receive(instance, apiError)
	return instance, resp, relevantError(err, *apiError)
}

// CreateInstanceParams are the parameters for OrganizationService.Create.
type CreateInstanceParams struct {
	Name   string `url:"name"`
	Plan   string `url:"plan"`
	Region string `url:"region"`
}

// Create a new CloudAMP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Create(params *CreateInstanceParams) (*Instance, *http.Response, error) {
	instance := new(Instance)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("instances").BodyForm(params).Receive(instance, apiError)
	return instance, resp, relevantError(err, *apiError)
}

//
// // UpdateOrganizationParams are the parameters for OrganizationService.Update.
// type UpdateOrganizationParams struct {
// 	Name string `json:"name,omitempty"`
// 	Slug string `json:"slug,omitempty"`
// }
//
// // Update a Sentry organization.
// // https://docs.sentry.io/api/organizations/put-organization-details/
// func (s *OrganizationService) Update(slug string, params *UpdateOrganizationParams) (*Organization, *http.Response, error) {
// 	org := new(Organization)
// 	apiError := new(APIError)
// 	resp, err := s.sling.New().Put(slug+"/").BodyJSON(params).Receive(org, apiError)
// 	return org, resp, relevantError(err, *apiError)
// }

// Delete a CloudAMQP instance.
// https://customer.cloudamqp.com/team/api
func (s *InstanceService) Delete(id int) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Path("instances/").Delete(strconv.Itoa(id)).Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
