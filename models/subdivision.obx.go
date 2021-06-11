// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type subdivision_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var SubdivisionBinding = subdivision_EntityInfo{
	Entity: objectbox.Entity{
		Id: 9,
	},
	Uid: 2902826944044842194,
}

// Subdivision_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Subdivision_ = struct {
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
			Entity: &SubdivisionBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &SubdivisionBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &SubdivisionBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &SubdivisionBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &SubdivisionBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &SubdivisionBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (subdivision_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (subdivision_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Subdivision", 9, 2902826944044842194)
	model.Property("Id", 6, 1, 7852798639316692896)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 6958493076190342342)
	model.Property("Name", 9, 3, 2762885926623008400)
	model.Property("IsDeleted", 1, 4, 6519561072805446948)
	model.Property("CreatedAt", 10, 5, 5635390990865168775)
	model.Property("UpdatedAt", 10, 6, 448129442438902162)
	model.EntityLastPropertyId(6, 448129442438902162)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (subdivision_EntityInfo) GetId(object interface{}) (uint64, error) {
	return uint64(object.(*Subdivision).Id), nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (subdivision_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Subdivision).Id = int64(id)
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (subdivision_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (subdivision_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Subdivision)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Subdivision.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Subdivision.UpdatedAt: " + err.Error())
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
func (subdivision_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Subdivision' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetInt64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 12))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Subdivision.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 14))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Subdivision.UpdatedAt: " + err.Error())
	}

	return &Subdivision{
		Id:        propId,
		ExtId:     fbutils.GetStringSlot(table, 6),
		Name:      fbutils.GetStringSlot(table, 8),
		IsDeleted: fbutils.GetBoolSlot(table, 10),
		CreatedAt: propCreatedAt,
		UpdatedAt: propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (subdivision_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Subdivision, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (subdivision_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Subdivision), nil)
	}
	return append(slice.([]*Subdivision), object.(*Subdivision))
}

// Box provides CRUD access to Subdivision objects
type SubdivisionBox struct {
	*objectbox.Box
}

// BoxForSubdivision opens a box of Subdivision objects
func BoxForSubdivision(ob *objectbox.ObjectBox) *SubdivisionBox {
	return &SubdivisionBox{
		Box: ob.InternalBox(9),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Subdivision.Id property on the passed object will be assigned the new ID as well.
func (box *SubdivisionBox) Put(object *Subdivision) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Subdivision.Id property on the passed object will be assigned the new ID as well.
func (box *SubdivisionBox) Insert(object *Subdivision) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *SubdivisionBox) Update(object *Subdivision) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *SubdivisionBox) PutAsync(object *Subdivision) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Subdivision.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Subdivision.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *SubdivisionBox) PutMany(objects []*Subdivision) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *SubdivisionBox) Get(id uint64) (*Subdivision, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Subdivision), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *SubdivisionBox) GetMany(ids ...uint64) ([]*Subdivision, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Subdivision), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *SubdivisionBox) GetManyExisting(ids ...uint64) ([]*Subdivision, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Subdivision), nil
}

// GetAll reads all stored objects
func (box *SubdivisionBox) GetAll() ([]*Subdivision, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Subdivision), nil
}

// Remove deletes a single object
func (box *SubdivisionBox) Remove(object *Subdivision) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *SubdivisionBox) RemoveMany(objects ...*Subdivision) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = uint64(object.Id)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Subdivision_ struct to create conditions.
// Keep the *SubdivisionQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *SubdivisionBox) Query(conditions ...objectbox.Condition) *SubdivisionQuery {
	return &SubdivisionQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Subdivision_ struct to create conditions.
// Keep the *SubdivisionQuery if you intend to execute the query multiple times.
func (box *SubdivisionBox) QueryOrError(conditions ...objectbox.Condition) (*SubdivisionQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &SubdivisionQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See SubdivisionAsyncBox for more information.
func (box *SubdivisionBox) Async() *SubdivisionAsyncBox {
	return &SubdivisionAsyncBox{AsyncBox: box.Box.Async()}
}

// SubdivisionAsyncBox provides asynchronous operations on Subdivision objects.
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
type SubdivisionAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForSubdivision creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use SubdivisionBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForSubdivision(ob *objectbox.ObjectBox, timeoutMs uint64) *SubdivisionAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 9, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 9: %s" + err.Error())
	}
	return &SubdivisionAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *SubdivisionAsyncBox) Put(object *Subdivision) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *SubdivisionAsyncBox) Insert(object *Subdivision) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *SubdivisionAsyncBox) Update(object *Subdivision) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *SubdivisionAsyncBox) Remove(object *Subdivision) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Subdivision which Id is either 42 or 47:
// 		box.Query(Subdivision_.Id.In(42, 47)).Find()
type SubdivisionQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *SubdivisionQuery) Find() ([]*Subdivision, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Subdivision), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *SubdivisionQuery) Offset(offset uint64) *SubdivisionQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *SubdivisionQuery) Limit(limit uint64) *SubdivisionQuery {
	query.Query.Limit(limit)
	return query
}
