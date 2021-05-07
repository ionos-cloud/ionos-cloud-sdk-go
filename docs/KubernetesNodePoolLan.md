# KubernetesNodePoolLan

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **int32** | The LAN ID of an existing LAN at the related datacenter | [optional] |

## Methods

### NewKubernetesNodePoolLan

`func NewKubernetesNodePoolLan() *KubernetesNodePoolLan`

NewKubernetesNodePoolLan instantiates a new KubernetesNodePoolLan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolLanWithDefaults

`func NewKubernetesNodePoolLanWithDefaults() *KubernetesNodePoolLan`

NewKubernetesNodePoolLanWithDefaults instantiates a new KubernetesNodePoolLan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePoolLan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePoolLan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePoolLan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNodePoolLan) HasId() bool`

HasId returns a boolean if a field has been set.


