// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
	"github.com/slevchyk/my_enterprise_local_srv/core"
)

type vehicle_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var VehicleBinding = vehicle_EntityInfo{
	Entity: objectbox.Entity{
		Id: 11,
	},
	Uid: 1924487076928445410,
}

// Vehicle_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Vehicle_ = struct {
	Id         *objectbox.PropertyUint64
	ExtId      *objectbox.PropertyString
	Name       *objectbox.PropertyString
	IsDeleted  *objectbox.PropertyBool
	CreatedAt  *objectbox.PropertyInt64
	UpdatedAt  *objectbox.PropertyInt64
	PhotoPath  *objectbox.PropertyString
	NfcId      *objectbox.PropertyString
	Length     *objectbox.PropertyFloat64
	Width      *objectbox.PropertyFloat64
	Height     *objectbox.PropertyFloat64
	MinWeight  *objectbox.PropertyFloat64
	Comment    *objectbox.PropertyString
	MaxWeight  *objectbox.PropertyFloat64
	DefTrailer *objectbox.RelationToOne
	DefDriver  *objectbox.RelationToOne
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &VehicleBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &VehicleBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &VehicleBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &VehicleBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &VehicleBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &VehicleBinding.Entity,
		},
	},
	PhotoPath: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &VehicleBinding.Entity,
		},
	},
	NfcId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &VehicleBinding.Entity,
		},
	},
	Length: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &VehicleBinding.Entity,
		},
	},
	Width: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &VehicleBinding.Entity,
		},
	},
	Height: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     14,
			Entity: &VehicleBinding.Entity,
		},
	},
	MinWeight: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     15,
			Entity: &VehicleBinding.Entity,
		},
	},
	Comment: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     16,
			Entity: &VehicleBinding.Entity,
		},
	},
	MaxWeight: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     17,
			Entity: &VehicleBinding.Entity,
		},
	},
	DefTrailer: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     18,
			Entity: &VehicleBinding.Entity,
		},
		Target: &TrailerBinding.Entity,
	},
	DefDriver: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     19,
			Entity: &VehicleBinding.Entity,
		},
		Target: &ServiceWorkerBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (vehicle_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (vehicle_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Vehicle", 11, 1924487076928445410)
	model.Property("Id", 6, 1, 6836724743619873469)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 5774404675760953993)
	model.Property("Name", 9, 3, 5464126720817488171)
	model.Property("IsDeleted", 1, 4, 2047802226670451377)
	model.Property("CreatedAt", 10, 5, 8294796580298202505)
	model.Property("UpdatedAt", 10, 6, 7281112080798607170)
	model.Property("PhotoPath", 9, 8, 6739210181211194073)
	model.Property("NfcId", 9, 9, 5502794683142699154)
	model.Property("Length", 8, 12, 523922608250770412)
	model.Property("Width", 8, 13, 2933999020946124024)
	model.Property("Height", 8, 14, 15511630527782703)
	model.Property("MinWeight", 8, 15, 6019667518297623819)
	model.Property("Comment", 9, 16, 5714631707887049789)
	model.Property("MaxWeight", 8, 17, 2475014504135017624)
	model.Property("DefTrailer", 11, 18, 3865780643844888825)
	model.PropertyFlags(520)
	model.PropertyRelation("Trailer", 48, 3222562739687006292)
	model.Property("DefDriver", 11, 19, 419646093088121380)
	model.PropertyFlags(520)
	model.PropertyRelation("ServiceWorker", 49, 547996278691202718)
	model.EntityLastPropertyId(19, 419646093088121380)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (vehicle_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Vehicle).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (vehicle_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Vehicle).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (vehicle_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := object.(*Vehicle).DefTrailer; rel != nil {
		if rId, err := TrailerBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForTrailer(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*Vehicle).DefDriver; rel != nil {
		if rId, err := ServiceWorkerBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForServiceWorker(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (vehicle_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Vehicle)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Vehicle.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Vehicle.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetPhotoPath = fbutils.CreateStringOffset(fbb, obj.PhotoPath)
	var offsetNfcId = fbutils.CreateStringOffset(fbb, obj.NfcId)
	var offsetComment = fbutils.CreateStringOffset(fbb, obj.Comment)

	var rIdDefTrailer uint64
	if rel := obj.DefTrailer; rel != nil {
		if rId, err := TrailerBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdDefTrailer = rId
		}
	}

	var rIdDefDriver uint64
	if rel := obj.DefDriver; rel != nil {
		if rId, err := ServiceWorkerBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdDefDriver = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(19)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetBoolSlot(fbb, 3, obj.IsDeleted)
	fbutils.SetFloat64Slot(fbb, 11, obj.Length)
	fbutils.SetFloat64Slot(fbb, 12, obj.Width)
	fbutils.SetFloat64Slot(fbb, 13, obj.Height)
	fbutils.SetFloat64Slot(fbb, 14, obj.MinWeight)
	fbutils.SetFloat64Slot(fbb, 16, obj.MaxWeight)
	fbutils.SetUOffsetTSlot(fbb, 15, offsetComment)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetPhotoPath)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetNfcId)
	if obj.DefTrailer != nil {
		fbutils.SetUint64Slot(fbb, 17, rIdDefTrailer)
	}
	if obj.DefDriver != nil {
		fbutils.SetUint64Slot(fbb, 18, rIdDefDriver)
	}
	fbutils.SetInt64Slot(fbb, 4, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 5, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (vehicle_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Vehicle' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 12))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Vehicle.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 14))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Vehicle.UpdatedAt: " + err.Error())
	}

	var relDefTrailer *Trailer
	if rId := fbutils.GetUint64PtrSlot(table, 38); rId != nil && *rId > 0 {
		if rObject, err := BoxForTrailer(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relDefTrailer = rObject
		}
	}

	var relDefDriver *ServiceWorker
	if rId := fbutils.GetUint64PtrSlot(table, 40); rId != nil && *rId > 0 {
		if rObject, err := BoxForServiceWorker(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relDefDriver = rObject
		}
	}

	return &Vehicle{
		Id:         propId,
		ExtId:      fbutils.GetStringSlot(table, 6),
		Name:       fbutils.GetStringSlot(table, 8),
		IsDeleted:  fbutils.GetBoolSlot(table, 10),
		Length:     fbutils.GetFloat64Slot(table, 26),
		Width:      fbutils.GetFloat64Slot(table, 28),
		Height:     fbutils.GetFloat64Slot(table, 30),
		MinWeight:  fbutils.GetFloat64Slot(table, 32),
		MaxWeight:  fbutils.GetFloat64Slot(table, 36),
		Comment:    fbutils.GetStringSlot(table, 34),
		PhotoPath:  fbutils.GetStringSlot(table, 18),
		NfcId:      fbutils.GetStringSlot(table, 20),
		DefTrailer: relDefTrailer,
		DefDriver:  relDefDriver,
		CreatedAt:  propCreatedAt,
		UpdatedAt:  propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (vehicle_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Vehicle, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (vehicle_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Vehicle), nil)
	}
	return append(slice.([]*Vehicle), object.(*Vehicle))
}

// Box provides CRUD access to Vehicle objects
type VehicleBox struct {
	*objectbox.Box
}

// BoxForVehicle opens a box of Vehicle objects
func BoxForVehicle(ob *objectbox.ObjectBox) *VehicleBox {
	return &VehicleBox{
		Box: ob.InternalBox(11),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Vehicle.Id property on the passed object will be assigned the new ID as well.
func (box *VehicleBox) Put(object *Vehicle) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Vehicle.Id property on the passed object will be assigned the new ID as well.
func (box *VehicleBox) Insert(object *Vehicle) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *VehicleBox) Update(object *Vehicle) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *VehicleBox) PutAsync(object *Vehicle) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Vehicle.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Vehicle.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *VehicleBox) PutMany(objects []*Vehicle) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *VehicleBox) Get(id uint64) (*Vehicle, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Vehicle), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *VehicleBox) GetMany(ids ...uint64) ([]*Vehicle, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Vehicle), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *VehicleBox) GetManyExisting(ids ...uint64) ([]*Vehicle, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Vehicle), nil
}

// GetAll reads all stored objects
func (box *VehicleBox) GetAll() ([]*Vehicle, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Vehicle), nil
}

// Remove deletes a single object
func (box *VehicleBox) Remove(object *Vehicle) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *VehicleBox) RemoveMany(objects ...*Vehicle) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Vehicle_ struct to create conditions.
// Keep the *VehicleQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *VehicleBox) Query(conditions ...objectbox.Condition) *VehicleQuery {
	return &VehicleQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Vehicle_ struct to create conditions.
// Keep the *VehicleQuery if you intend to execute the query multiple times.
func (box *VehicleBox) QueryOrError(conditions ...objectbox.Condition) (*VehicleQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &VehicleQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See VehicleAsyncBox for more information.
func (box *VehicleBox) Async() *VehicleAsyncBox {
	return &VehicleAsyncBox{AsyncBox: box.Box.Async()}
}

// VehicleAsyncBox provides asynchronous operations on Vehicle objects.
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
type VehicleAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForVehicle creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use VehicleBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForVehicle(ob *objectbox.ObjectBox, timeoutMs uint64) *VehicleAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 11, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 11: %s" + err.Error())
	}
	return &VehicleAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *VehicleAsyncBox) Put(object *Vehicle) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *VehicleAsyncBox) Insert(object *Vehicle) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *VehicleAsyncBox) Update(object *Vehicle) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *VehicleAsyncBox) Remove(object *Vehicle) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Vehicle which Id is either 42 or 47:
//
//	box.Query(Vehicle_.Id.In(42, 47)).Find()
type VehicleQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *VehicleQuery) Find() ([]*Vehicle, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Vehicle), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *VehicleQuery) Offset(offset uint64) *VehicleQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *VehicleQuery) Limit(limit uint64) *VehicleQuery {
	query.Query.Limit(limit)
	return query
}

type vehicleImport_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var VehicleImportBinding = vehicleImport_EntityInfo{
	Entity: objectbox.Entity{
		Id: 18,
	},
	Uid: 3928159787491915519,
}

// VehicleImport_ contains type-based Property helpers to facilitate some common operations such as Queries.
var VehicleImport_ = struct {
	Id              *objectbox.PropertyUint64
	ExtId           *objectbox.PropertyString
	Name            *objectbox.PropertyString
	IsDeleted       *objectbox.PropertyBool
	Length          *objectbox.PropertyFloat64
	Width           *objectbox.PropertyFloat64
	Height          *objectbox.PropertyFloat64
	MinWeight       *objectbox.PropertyFloat64
	MaxWeight       *objectbox.PropertyFloat64
	Comment         *objectbox.PropertyString
	PhotoPath       *objectbox.PropertyString
	NfcId           *objectbox.PropertyString
	DefTrailerExtId *objectbox.PropertyString
	DefDriverExtId  *objectbox.PropertyString
	CreatedAt       *objectbox.PropertyInt64
	UpdatedAt       *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	Length: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	Width: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	Height: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	MinWeight: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	MaxWeight: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	Comment: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	PhotoPath: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	NfcId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	DefTrailerExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	DefDriverExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     14,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     15,
			Entity: &VehicleImportBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     16,
			Entity: &VehicleImportBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (vehicleImport_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (vehicleImport_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("VehicleImport", 18, 3928159787491915519)
	model.Property("Id", 6, 1, 3305164315838620365)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 1959607379369983453)
	model.Property("Name", 9, 3, 4495523189833224085)
	model.Property("IsDeleted", 1, 4, 6477995439407178031)
	model.Property("Length", 8, 5, 884073393702277082)
	model.Property("Width", 8, 6, 6783753808289818300)
	model.Property("Height", 8, 7, 5390617234041942943)
	model.Property("MinWeight", 8, 8, 2983923145552962984)
	model.Property("MaxWeight", 8, 9, 5779199848780285083)
	model.Property("Comment", 9, 10, 3682145699734791637)
	model.Property("PhotoPath", 9, 11, 1846027210865495343)
	model.Property("NfcId", 9, 12, 5881614396633403948)
	model.Property("DefTrailerExtId", 9, 13, 7664158350002246278)
	model.Property("DefDriverExtId", 9, 14, 4594866031181129181)
	model.Property("CreatedAt", 10, 15, 8170612890188281368)
	model.Property("UpdatedAt", 10, 16, 8768853457139016146)
	model.EntityLastPropertyId(16, 8768853457139016146)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (vehicleImport_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*VehicleImport).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (vehicleImport_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*VehicleImport).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (vehicleImport_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (vehicleImport_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*VehicleImport)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on VehicleImport.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on VehicleImport.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetComment = fbutils.CreateStringOffset(fbb, obj.Comment)
	var offsetPhotoPath = fbutils.CreateStringOffset(fbb, obj.PhotoPath)
	var offsetNfcId = fbutils.CreateStringOffset(fbb, obj.NfcId)
	var offsetDefTrailerExtId = fbutils.CreateStringOffset(fbb, obj.DefTrailerExtId)
	var offsetDefDriverExtId = fbutils.CreateStringOffset(fbb, obj.DefDriverExtId)

	// build the FlatBuffers object
	fbb.StartObject(16)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetName)
	fbutils.SetBoolSlot(fbb, 3, obj.IsDeleted)
	fbutils.SetFloat64Slot(fbb, 4, float64(obj.Length))
	fbutils.SetFloat64Slot(fbb, 5, float64(obj.Width))
	fbutils.SetFloat64Slot(fbb, 6, float64(obj.Height))
	fbutils.SetFloat64Slot(fbb, 7, float64(obj.MinWeight))
	fbutils.SetFloat64Slot(fbb, 8, float64(obj.MaxWeight))
	fbutils.SetUOffsetTSlot(fbb, 9, offsetComment)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetPhotoPath)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetNfcId)
	fbutils.SetUOffsetTSlot(fbb, 12, offsetDefTrailerExtId)
	fbutils.SetUOffsetTSlot(fbb, 13, offsetDefDriverExtId)
	fbutils.SetInt64Slot(fbb, 14, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 15, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (vehicleImport_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'VehicleImport' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 32))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on VehicleImport.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 34))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on VehicleImport.UpdatedAt: " + err.Error())
	}

	return &VehicleImport{
		Id:              propId,
		ExtId:           fbutils.GetStringSlot(table, 6),
		Name:            fbutils.GetStringSlot(table, 8),
		IsDeleted:       fbutils.GetBoolSlot(table, 10),
		Length:          core.Float(fbutils.GetFloat64Slot(table, 12)),
		Width:           core.Float(fbutils.GetFloat64Slot(table, 14)),
		Height:          core.Float(fbutils.GetFloat64Slot(table, 16)),
		MinWeight:       core.Float(fbutils.GetFloat64Slot(table, 18)),
		MaxWeight:       core.Float(fbutils.GetFloat64Slot(table, 20)),
		Comment:         fbutils.GetStringSlot(table, 22),
		PhotoPath:       fbutils.GetStringSlot(table, 24),
		NfcId:           fbutils.GetStringSlot(table, 26),
		DefTrailerExtId: fbutils.GetStringSlot(table, 28),
		DefDriverExtId:  fbutils.GetStringSlot(table, 30),
		CreatedAt:       propCreatedAt,
		UpdatedAt:       propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (vehicleImport_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*VehicleImport, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (vehicleImport_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*VehicleImport), nil)
	}
	return append(slice.([]*VehicleImport), object.(*VehicleImport))
}

// Box provides CRUD access to VehicleImport objects
type VehicleImportBox struct {
	*objectbox.Box
}

// BoxForVehicleImport opens a box of VehicleImport objects
func BoxForVehicleImport(ob *objectbox.ObjectBox) *VehicleImportBox {
	return &VehicleImportBox{
		Box: ob.InternalBox(18),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the VehicleImport.Id property on the passed object will be assigned the new ID as well.
func (box *VehicleImportBox) Put(object *VehicleImport) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the VehicleImport.Id property on the passed object will be assigned the new ID as well.
func (box *VehicleImportBox) Insert(object *VehicleImport) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *VehicleImportBox) Update(object *VehicleImport) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *VehicleImportBox) PutAsync(object *VehicleImport) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the VehicleImport.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the VehicleImport.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *VehicleImportBox) PutMany(objects []*VehicleImport) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *VehicleImportBox) Get(id uint64) (*VehicleImport, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*VehicleImport), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *VehicleImportBox) GetMany(ids ...uint64) ([]*VehicleImport, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*VehicleImport), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *VehicleImportBox) GetManyExisting(ids ...uint64) ([]*VehicleImport, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*VehicleImport), nil
}

// GetAll reads all stored objects
func (box *VehicleImportBox) GetAll() ([]*VehicleImport, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*VehicleImport), nil
}

// Remove deletes a single object
func (box *VehicleImportBox) Remove(object *VehicleImport) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *VehicleImportBox) RemoveMany(objects ...*VehicleImport) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the VehicleImport_ struct to create conditions.
// Keep the *VehicleImportQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *VehicleImportBox) Query(conditions ...objectbox.Condition) *VehicleImportQuery {
	return &VehicleImportQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the VehicleImport_ struct to create conditions.
// Keep the *VehicleImportQuery if you intend to execute the query multiple times.
func (box *VehicleImportBox) QueryOrError(conditions ...objectbox.Condition) (*VehicleImportQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &VehicleImportQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See VehicleImportAsyncBox for more information.
func (box *VehicleImportBox) Async() *VehicleImportAsyncBox {
	return &VehicleImportAsyncBox{AsyncBox: box.Box.Async()}
}

// VehicleImportAsyncBox provides asynchronous operations on VehicleImport objects.
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
type VehicleImportAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForVehicleImport creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use VehicleImportBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForVehicleImport(ob *objectbox.ObjectBox, timeoutMs uint64) *VehicleImportAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 18, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 18: %s" + err.Error())
	}
	return &VehicleImportAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *VehicleImportAsyncBox) Put(object *VehicleImport) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *VehicleImportAsyncBox) Insert(object *VehicleImport) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *VehicleImportAsyncBox) Update(object *VehicleImport) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *VehicleImportAsyncBox) Remove(object *VehicleImport) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all VehicleImport which Id is either 42 or 47:
//
//	box.Query(VehicleImport_.Id.In(42, 47)).Find()
type VehicleImportQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *VehicleImportQuery) Find() ([]*VehicleImport, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*VehicleImport), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *VehicleImportQuery) Offset(offset uint64) *VehicleImportQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *VehicleImportQuery) Limit(limit uint64) *VehicleImportQuery {
	query.Query.Limit(limit)
	return query
}
