# BelleseChallenge.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**itemGet**](DefaultApi.md#itemGet) | **GET** /item | 
[**itemIdDelete**](DefaultApi.md#itemIdDelete) | **DELETE** /item/{id} | 
[**itemIdGet**](DefaultApi.md#itemIdGet) | **GET** /item/{id} | 
[**itemIdPatch**](DefaultApi.md#itemIdPatch) | **PATCH** /item/{id} | 
[**itemPost**](DefaultApi.md#itemPost) | **POST** /item | 


<a name="itemGet"></a>
# **itemGet**
> [Item] itemGet()



### Example
```javascript
var BelleseChallenge = require('bellese_challenge');

var apiInstance = new BelleseChallenge.DefaultApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.itemGet(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Item]**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="itemIdDelete"></a>
# **itemIdDelete**
> itemIdDelete(id)



### Example
```javascript
var BelleseChallenge = require('bellese_challenge');

var apiInstance = new BelleseChallenge.DefaultApi();

var id = 8.14; // Number | id of the item to fetch


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.itemIdDelete(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of the item to fetch | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="itemIdGet"></a>
# **itemIdGet**
> Item itemIdGet(id)



### Example
```javascript
var BelleseChallenge = require('bellese_challenge');

var apiInstance = new BelleseChallenge.DefaultApi();

var id = 8.14; // Number | id of the item to fetch


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.itemIdGet(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of the item to fetch | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="itemIdPatch"></a>
# **itemIdPatch**
> Item itemIdPatch(id, opts)



### Example
```javascript
var BelleseChallenge = require('bellese_challenge');

var apiInstance = new BelleseChallenge.DefaultApi();

var id = 8.14; // Number | id of the item to fetch

var opts = { 
  'item': new BelleseChallenge.Item() // Item | the item to create
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.itemIdPatch(id, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| id of the item to fetch | 
 **item** | [**Item**](Item.md)| the item to create | [optional] 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="itemPost"></a>
# **itemPost**
> Item itemPost(opts)



### Example
```javascript
var BelleseChallenge = require('bellese_challenge');

var apiInstance = new BelleseChallenge.DefaultApi();

var opts = { 
  'item': new BelleseChallenge.Item() // Item | the item to create
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.itemPost(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**Item**](Item.md)| the item to create | [optional] 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

