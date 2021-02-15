package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueTypeScreenSchemeService_Append(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		mappings                *[]IssueTypeScreenSchemeMappingPayloadScheme
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheIssueTypeScreenSchemeIDParamIsEmpty",
			issueTypeScreenSchemeID: "",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheMappingsParamIsNil",
			issueTypeScreenSchemeID: "10000",
			mappings:                nil,
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenschemes/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			mappings: &[]IssueTypeScreenSchemeMappingPayloadScheme{
				{
					IssueTypeID:    "10000",
					ScreenSchemeID: "10001",
				},
				{
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				},
				{
					IssueTypeID:    "10002",
					ScreenSchemeID: "10003",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Append(testCase.context, testCase.issueTypeScreenSchemeID, testCase.mappings)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Assign(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		projectID               string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheProjectIDIsNotSet",
			issueTypeScreenSchemeID: "10000",
			projectID:               "",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/projects",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Assign(testCase.context, testCase.issueTypeScreenSchemeID, testCase.projectID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *IssueTypeScreenSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueTypeSchemeWhenTheParametersAreCorrect",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateIssueTypeSchemeWhenTheIssueTypeScreenSchemePayloadSchemeParamIsNil",
			payload:            nil,
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheIssueTypeScreenSchemePayloadSchemeNameIsNotSet",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheIssueTypeScreenSchemePayloadSchemeIssueTypeMappingsIsNotSet",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name:              "Scrum issue type screen scheme",
				IssueTypeMappings: nil,
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheEndpointIsIncorrect",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/apis/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheContextIsNil",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheResponseBodyHasADifferentFormat",
			payload: &IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []IssueTypeScreenSchemeMappingPayloadScheme{{
					IssueTypeID:    "default",
					ScreenSchemeID: "10001",
				}, {
					IssueTypeID:    "10001",
					ScreenSchemeID: "10002",
				}},
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "DeleteIssueTypeSchemeWhenTheIssueTypeScreenSchemeIDIsValid",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10001",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueTypeScreenSchemeID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		ids                []int
		startAt            int
		maxResults         int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTypeSchemesWhenTheParametersAreCorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheIdsAreNotSet",
			ids:                nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheEndpointIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenschemes?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheRequestMethodIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheStatusCodeIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheContextIsNil",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheResponseBodyHasADifferentFormat",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.ids, testCase.startAt, testCase.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Update(t *testing.T) {

	testCases := []struct {
		name                      string
		issueTypeScreenSchemeID   string
		issueTypeScreenSchemeName string
		description               string
		mockFile                  string
		wantHTTPMethod            string
		endpoint                  string
		context                   context.Context
		wantHTTPCodeReturn        int
		wantErr                   bool
	}{
		{
			name:                      "UpdateIssueTypeSchemeWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   false,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID:   "",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10001",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPost,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusBadRequest,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   nil,
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Update(testCase.context, testCase.issueTypeScreenSchemeID,
				testCase.issueTypeScreenSchemeName, testCase.description)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}
}
