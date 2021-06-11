// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type harvestType_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var HarvestTypeBinding = harvestType_EntityInfo{
	Entity: objectbox.Entity{
		Id: 7,
	},
	Uid: 995889323428318313,
}

// HarvestType_ contains type-based Property helpers to facilitate some common operations such as Queries.
var HarvestType_ = struct {
	Id        *objectbox.PropertyInt64
	ExtId     *objectbox.PropertyString
	Name      *objectbox.PropertyString
	IsDeleted *objectbox.PropertyBool
	CreatedAt *objectbox.PropertyInt64
	UpdatedAt *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &HarvestTypeBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (harvestType_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (harvestType_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("HarvestType", 7, 995889323428318313)
	model.Property("Id", 6, 1, 4992084047845681368)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 2055242425374455957)
	model.Property("Name", 9, 3, 722190376636770616)
	model.Property("IsDeleted", 1, 4, 7809446098984603525)
	model.Property("CreatedAt", 10, 5, 5210310893644494224)
	model.Property("UpdatedAt", 10, 6, 4233746714675610160)
	model.EntityLastPropertyId(6, 4233746714675610160)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (harvestType_EntityInfo) GetId(object interface{}) (uint64, error) {
	return uint64(object.(*HarvestType).Id), nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (harvestType_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*HarvestType).Id = int64(id)
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (harvestType_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (harvestType_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*HarvestType)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on HarvestType.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on HarvestType.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)

	// build the FlatBuffers object
	fbb.StartObject(6)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetBoolSlot(fbb, 3, obj.IsDeleted)
	fbutils.SetInt64Slot(fbb, 4, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 5, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (harvestType_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'HarvestType' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetInt64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 12))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on HarvestType.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 14))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on HarvestType.UpdatedAt: " + err.Error())
	}

	return &HarvestType{
		Id:        propId,
		ExtId:     fbutils.GetStringSlot(table, 6),
		Name:      fbutils.GetStringSlot(table, 8),
		IsDeleted: fbutils.GetBoolSlot(table, 10),
		CreatedAt: propCreatedAt,
		UpdatedAt: propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (harvestType_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*HarvestType, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (harvestType_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*HarvestType), nil)
	}
	return append(slice.([]*HarvestType), object.(*HarvestType))
}

// Box provides CRUD access to HarvestType objects
type HarvestTypeBox struct {
	*objectbox.Box
}

// BoxForHarvestType opens a box of HarvestType objects
func BoxForHarvestType(ob *objectbox.ObjectBox) *HarvestTypeBox {
	return &HarvestTypeBox{
		Box: ob.InternalBox(7),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the HarvestType.Id property on the passed object will be assigned the new ID as well.
func (box *HarvestTypeBox) Put(object *HarvestType) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the HarvestType.Id property on the passed object will be assigned the new ID as well.
func (box *HarvestTypeBox) Insert(object *HarvestType) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *HarvestTypeBox) Update(object *HarvestType) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *HarvestTypeBox) PutAsync(object *HarvestType) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the HarvestType.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the HarvestType.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *HarvestTypeBox) PutMany(objects []*HarvestType) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *HarvestTypeBox) Get(id uint64) (*HarvestType, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*HarvestType), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *HarvestTypeBox) GetMany(ids ...uint64) ([]*HarvestType, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*HarvestType), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *HarvestTypeBox) GetManyExisting(ids ...uint64) ([]*HarvestType, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*HarvestType), nil
}

// GetAll reads all stored objects
func (box *HarvestTypeBox) GetAll() ([]*HarvestType, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*HarvestType), nil
}

// Remove deletes a single object
func (box *HarvestTypeBox) Remove(object *HarvestType) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *HarvestTypeBox) RemoveMany(objects ...*HarvestType) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = uint64(object.Id)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the HarvestType_ struct to create conditions.
// Keep the *HarvestTypeQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *HarvestTypeBox) Query(conditions ...objectbox.Condition) *HarvestTypeQuery {
	return &HarvestTypeQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the HarvestType_ struct to create conditions.
// Keep the *HarvestTypeQuery if you intend to execute the query multiple times.
func (box *HarvestTypeBox) QueryOrError(conditions ...objectbox.Condition) (*HarvestTypeQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &HarvestTypeQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See HarvestTypeAsyncBox for more information.
func (box *HarvestTypeBox) Async() *HarvestTypeAsyncBox {
	return &HarvestTypeAsyncBox{AsyncBox: box.Box.Async()}
}

// HarvestTypeAsyncBox provides asynchronous operations on HarvestType objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type HarvestTypeAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForHarvestType creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use HarvestTypeBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForHarvestType(ob *objectbox.ObjectBox, timeoutMs uint64) *HarvestTypeAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 7, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 7: %s" + err.Error())
	}
	return &HarvestTypeAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *HarvestTypeAsyncBox) Put(object *HarvestType) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *HarvestTypeAsyncBox) Insert(object *HarvestType) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *HarvestTypeAsyncBox) Update(object *HarvestType) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *HarvestTypeAsyncBox) Remove(object *HarvestType) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all HarvestType which Id is either 42 or 47:
// 		box.Query(HarvestType_.Id.In(42, 47)).Find()
type HarvestTypeQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *HarvestTypeQuery) Find() ([]*HarvestType, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*HarvestType), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *HarvestTypeQuery) Offset(offset uint64) *HarvestTypeQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *HarvestTypeQuery) Limit(limit uint64) *HarvestTypeQuery {
	query.Query.Limit(limit)
	return query
}