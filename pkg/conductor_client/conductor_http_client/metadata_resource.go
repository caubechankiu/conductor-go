package conductor_http_client

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/antihax/optional"
// )

// // Linger please
// var (
// 	_ context.Context
// )

// type MetadataResourceApiService service

// /*
// MetadataResourceApiService Create a new workflow definition
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *MetadataResourceApiService) Create(ctx context.Context, body WorkflowDef) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Post")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/workflow"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Retrieves workflow definition along with blueprint
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param name
//  * @param optional nil or *MetadataResourceApiGetOpts - Optional Parameters:
//      * @param "Version" (optional.Int32) -
// @return WorkflowDef
// */

// type MetadataResourceApiGetOpts struct {
// 	Version optional.Int32
// }

// func (a *MetadataResourceApiService) Get(ctx context.Context, name string, localVarOptionals *MetadataResourceApiGetOpts) (WorkflowDef, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue WorkflowDef
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/workflow/{name}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
// 		localVarQueryParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v WorkflowDef
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Retrieves all workflow definition along with blueprint
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return []WorkflowDef
// */
// func (a *MetadataResourceApiService) GetAll(ctx context.Context) ([]WorkflowDef, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue []WorkflowDef
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/workflow"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v []WorkflowDef
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Gets the task definition
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param tasktype
// @return TaskDef
// */
// func (a *MetadataResourceApiService) GetTaskDef(ctx context.Context, tasktype string) (TaskDef, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue TaskDef
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/taskdefs/{tasktype}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"tasktype"+"}", fmt.Sprintf("%v", tasktype), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v TaskDef
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Gets all task definition
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return []TaskDef
// */
// func (a *MetadataResourceApiService) GetTaskDefs(ctx context.Context) ([]TaskDef, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue []TaskDef
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/taskdefs"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v []TaskDef
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Update an existing task
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *MetadataResourceApiService) RegisterTaskDef(ctx context.Context, body TaskDef) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Put")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/taskdefs"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Create new task definition(s)
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *MetadataResourceApiService) RegisterTaskDef1(ctx context.Context, body []TaskDef) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Post")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/taskdefs"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Remove a task definition
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param tasktype

// */
// func (a *MetadataResourceApiService) UnregisterTaskDef(ctx context.Context, tasktype string) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Delete")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/taskdefs/{tasktype}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"tasktype"+"}", fmt.Sprintf("%v", tasktype), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Removes workflow definition. It does not remove workflows associated with the definition.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param name
//  * @param version

// */
// func (a *MetadataResourceApiService) UnregisterWorkflowDef(ctx context.Context, name string, version int32) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Delete")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/workflow/{name}/{version}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
// 	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// MetadataResourceApiService Create or update workflow definition
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *MetadataResourceApiService) Update(ctx context.Context, body []WorkflowDef) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Put")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/metadata/workflow"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }