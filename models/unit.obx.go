// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type unit_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var UnitBinding = unit_EntityInfo{
	Entity: objectbox.Entity{
		Id: 10,
	},
	Uid: 7333998516145050358,
}

// Unit_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Unit_ = struct {
	Id        *objectbox.PropertyUint64
	ExtId     *objectbox.PropertyString
	Name      *objectbox.PropertyString
	IsDeleted *objectbox.PropertyBool
	CreatedAt *objectbox.PropertyInt64
	UpdatedAt *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &UnitBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &UnitBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &UnitBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &UnitBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &UnitBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &UnitBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (unit_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (unit_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Unit", 10, 7333998516145050358)
	model.Property("Id", 6, 1, 3823507343141096680)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 7322082217165942378)
	model.Property("Name", 9, 3, 8324870919092323762)
	model.Property("IsDeleted", 1, 4, 2994449961417320882)
	model.Property("CreatedAt", 10, 5, 3184251330947109922)
	model.Property("UpdatedAt", 10, 6, 6454938316269374724)
	model.EntityLastPropertyId(6, 6454938316269374724)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (unit_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Unit).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (unit_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Unit).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (unit_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (unit_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Unit)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Unit.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Unit.UpdatedAt: " + err.Error())
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
func (unit_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Unit' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 12))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Unit.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 14))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Unit.UpdatedAt: " + err.Error())
	}

	return &Unit{
		Id:        propId,
		ExtId:     fbutils.GetStringSlot(table, 6),
		Name:      fbutils.GetStringSlot(table, 8),
		IsDeleted: fbutils.GetBoolSlot(table, 10),
		CreatedAt: propCreatedAt,
		UpdatedAt: propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (unit_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Unit, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (unit_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Unit), nil)
	}
	return append(slice.([]*Unit), object.(*Unit))
}

// Box provides CRUD access to Unit objects
type UnitBox struct {
	*objectbox.Box
}

// BoxForUnit opens a box of Unit objects
func BoxForUnit(ob *objectbox.ObjectBox) *UnitBox {
	return &UnitBox{
		Box: ob.InternalBox(10),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Unit.Id property on the passed object will be assigned the new ID as well.
func (box *UnitBox) Put(object *Unit) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Unit.Id property on the passed object will be assigned the new ID as well.
func (box *UnitBox) Insert(object *Unit) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *UnitBox) Update(object *Unit) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *UnitBox) PutAsync(object *Unit) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Unit.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Unit.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *UnitBox) PutMany(objects []*Unit) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *UnitBox) Get(id uint64) (*Unit, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Unit), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *UnitBox) GetMany(ids ...uint64) ([]*Unit, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Unit), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *UnitBox) GetManyExisting(ids ...uint64) ([]*Unit, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Unit), nil
}

// GetAll reads all stored objects
func (box *UnitBox) GetAll() ([]*Unit, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Unit), nil
}

// Remove deletes a single object
func (box *UnitBox) Remove(object *Unit) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *UnitBox) RemoveMany(objects ...*Unit) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Unit_ struct to create conditions.
// Keep the *UnitQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *UnitBox) Query(conditions ...objectbox.Condition) *UnitQuery {
	return &UnitQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Unit_ struct to create conditions.
// Keep the *UnitQuery if you intend to execute the query multiple times.
func (box *UnitBox) QueryOrError(conditions ...objectbox.Condition) (*UnitQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &UnitQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See UnitAsyncBox for more information.
func (box *UnitBox) Async() *UnitAsyncBox {
	return &UnitAsyncBox{AsyncBox: box.Box.Async()}
}

// UnitAsyncBox provides asynchronous operations on Unit objects.
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
type UnitAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForUnit creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use UnitBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForUnit(ob *objectbox.ObjectBox, timeoutMs uint64) *UnitAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 10, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 10: %s" + err.Error())
	}
	return &UnitAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *UnitAsyncBox) Put(object *Unit) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *UnitAsyncBox) Insert(object *Unit) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *UnitAsyncBox) Update(object *Unit) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *UnitAsyncBox) Remove(object *Unit) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Unit which Id is either 42 or 47:
//
//	box.Query(Unit_.Id.In(42, 47)).Find()
type UnitQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *UnitQuery) Find() ([]*Unit, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Unit), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *UnitQuery) Offset(offset uint64) *UnitQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *UnitQuery) Limit(limit uint64) *UnitQuery {
	query.Query.Limit(limit)
	return query
}
