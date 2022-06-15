// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type trailer_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var TrailerBinding = trailer_EntityInfo{
	Entity: objectbox.Entity{
		Id: 13,
	},
	Uid: 3574534838454594203,
}

// Trailer_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Trailer_ = struct {
	Id        *objectbox.PropertyUint64
	ExtId     *objectbox.PropertyString
	Name      *objectbox.PropertyString
	IsDeleted *objectbox.PropertyBool
	MaxWeight *objectbox.PropertyFloat32
	CreatedAt *objectbox.PropertyInt64
	UpdatedAt *objectbox.PropertyInt64
	PhotoPath *objectbox.PropertyString
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &TrailerBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &TrailerBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &TrailerBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &TrailerBinding.Entity,
		},
	},
	MaxWeight: &objectbox.PropertyFloat32{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &TrailerBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &TrailerBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &TrailerBinding.Entity,
		},
	},
	PhotoPath: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &TrailerBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (trailer_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (trailer_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Trailer", 13, 3574534838454594203)
	model.Property("Id", 6, 1, 2784396226860608106)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 3924205386727081595)
	model.Property("Name", 9, 3, 7439380208901930246)
	model.Property("IsDeleted", 1, 4, 1198123329569085660)
	model.Property("MaxWeight", 7, 5, 1863085959529203247)
	model.Property("CreatedAt", 10, 7, 889475617324735826)
	model.Property("UpdatedAt", 10, 8, 4577958236363959362)
	model.Property("PhotoPath", 9, 9, 9185563778196944025)
	model.EntityLastPropertyId(9, 9185563778196944025)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (trailer_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Trailer).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (trailer_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Trailer).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (trailer_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (trailer_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Trailer)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Trailer.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Trailer.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetPhotoPath = fbutils.CreateStringOffset(fbb, obj.PhotoPath)

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetBoolSlot(fbb, 3, obj.IsDeleted)
	fbutils.SetFloat32Slot(fbb, 4, obj.MaxWeight)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetPhotoPath)
	fbutils.SetInt64Slot(fbb, 6, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 7, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (trailer_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Trailer' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 16))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Trailer.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 18))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Trailer.UpdatedAt: " + err.Error())
	}

	return &Trailer{
		Id:        propId,
		ExtId:     fbutils.GetStringSlot(table, 6),
		Name:      fbutils.GetStringSlot(table, 8),
		IsDeleted: fbutils.GetBoolSlot(table, 10),
		MaxWeight: fbutils.GetFloat32Slot(table, 12),
		PhotoPath: fbutils.GetStringSlot(table, 20),
		CreatedAt: propCreatedAt,
		UpdatedAt: propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (trailer_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Trailer, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (trailer_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Trailer), nil)
	}
	return append(slice.([]*Trailer), object.(*Trailer))
}

// Box provides CRUD access to Trailer objects
type TrailerBox struct {
	*objectbox.Box
}

// BoxForTrailer opens a box of Trailer objects
func BoxForTrailer(ob *objectbox.ObjectBox) *TrailerBox {
	return &TrailerBox{
		Box: ob.InternalBox(13),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Trailer.Id property on the passed object will be assigned the new ID as well.
func (box *TrailerBox) Put(object *Trailer) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Trailer.Id property on the passed object will be assigned the new ID as well.
func (box *TrailerBox) Insert(object *Trailer) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *TrailerBox) Update(object *Trailer) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *TrailerBox) PutAsync(object *Trailer) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Trailer.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Trailer.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *TrailerBox) PutMany(objects []*Trailer) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *TrailerBox) Get(id uint64) (*Trailer, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Trailer), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *TrailerBox) GetMany(ids ...uint64) ([]*Trailer, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Trailer), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *TrailerBox) GetManyExisting(ids ...uint64) ([]*Trailer, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Trailer), nil
}

// GetAll reads all stored objects
func (box *TrailerBox) GetAll() ([]*Trailer, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Trailer), nil
}

// Remove deletes a single object
func (box *TrailerBox) Remove(object *Trailer) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *TrailerBox) RemoveMany(objects ...*Trailer) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Trailer_ struct to create conditions.
// Keep the *TrailerQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *TrailerBox) Query(conditions ...objectbox.Condition) *TrailerQuery {
	return &TrailerQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Trailer_ struct to create conditions.
// Keep the *TrailerQuery if you intend to execute the query multiple times.
func (box *TrailerBox) QueryOrError(conditions ...objectbox.Condition) (*TrailerQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &TrailerQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See TrailerAsyncBox for more information.
func (box *TrailerBox) Async() *TrailerAsyncBox {
	return &TrailerAsyncBox{AsyncBox: box.Box.Async()}
}

// TrailerAsyncBox provides asynchronous operations on Trailer objects.
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
type TrailerAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForTrailer creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use TrailerBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForTrailer(ob *objectbox.ObjectBox, timeoutMs uint64) *TrailerAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 13, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 13: %s" + err.Error())
	}
	return &TrailerAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *TrailerAsyncBox) Put(object *Trailer) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *TrailerAsyncBox) Insert(object *Trailer) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *TrailerAsyncBox) Update(object *Trailer) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *TrailerAsyncBox) Remove(object *Trailer) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Trailer which Id is either 42 or 47:
// 		box.Query(Trailer_.Id.In(42, 47)).Find()
type TrailerQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *TrailerQuery) Find() ([]*Trailer, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Trailer), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *TrailerQuery) Offset(offset uint64) *TrailerQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *TrailerQuery) Limit(limit uint64) *TrailerQuery {
	query.Query.Limit(limit)
	return query
}
