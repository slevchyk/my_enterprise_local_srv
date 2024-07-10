// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type goodsConsignmentNoteIn_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var GoodsConsignmentNoteInBinding = goodsConsignmentNoteIn_EntityInfo{
	Entity: objectbox.Entity{
		Id: 4,
	},
	Uid: 8444046993855405783,
}

// GoodsConsignmentNoteIn_ contains type-based Property helpers to facilitate some common operations such as Queries.
var GoodsConsignmentNoteIn_ = struct {
	Id                *objectbox.PropertyUint64
	ExtId             *objectbox.PropertyString
	LoadingPercentage *objectbox.PropertyFloat32
	Quantity          *objectbox.PropertyFloat32
	CreatedAt         *objectbox.PropertyInt64
	UpdatedAt         *objectbox.PropertyInt64
	Subdivision       *objectbox.RelationToOne
	GoodsGroup        *objectbox.RelationToOne
	Goods             *objectbox.RelationToOne
	Unit              *objectbox.RelationToOne
	ConsignmentNoteIn *objectbox.RelationToOne
	AppId             *objectbox.PropertyString
	Locality          *objectbox.RelationToOne
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	LoadingPercentage: &objectbox.PropertyFloat32{
		BaseProperty: &objectbox.BaseProperty{
			Id:     24,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	Quantity: &objectbox.PropertyFloat32{
		BaseProperty: &objectbox.BaseProperty{
			Id:     25,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     62,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     63,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	Subdivision: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     64,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &SubdivisionBinding.Entity,
	},
	GoodsGroup: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     65,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &GoodsGroupBinding.Entity,
	},
	Goods: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     66,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &GoodsBinding.Entity,
	},
	Unit: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     67,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &UnitBinding.Entity,
	},
	ConsignmentNoteIn: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     68,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &ConsignmentNoteInBinding.Entity,
	},
	AppId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     69,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
	},
	Locality: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     70,
			Entity: &GoodsConsignmentNoteInBinding.Entity,
		},
		Target: &LocalityBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (goodsConsignmentNoteIn_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (goodsConsignmentNoteIn_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("GoodsConsignmentNoteIn", 4, 8444046993855405783)
	model.Property("Id", 6, 1, 7746942835379170791)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 2286777172924440151)
	model.PropertyFlags(2080)
	model.PropertyIndex(29, 8173145340567612110)
	model.Property("LoadingPercentage", 7, 24, 6002045979613560070)
	model.Property("Quantity", 7, 25, 7898911385562250666)
	model.Property("CreatedAt", 10, 62, 1137569487564150713)
	model.Property("UpdatedAt", 10, 63, 306223304406205015)
	model.Property("Subdivision", 11, 64, 1018232616173266192)
	model.PropertyFlags(520)
	model.PropertyRelation("Subdivision", 21, 5448937857734871962)
	model.Property("GoodsGroup", 11, 65, 4383679316513977820)
	model.PropertyFlags(520)
	model.PropertyRelation("GoodsGroup", 22, 6946803912572605339)
	model.Property("Goods", 11, 66, 8301365480842725703)
	model.PropertyFlags(520)
	model.PropertyRelation("Goods", 23, 3284539892381155022)
	model.Property("Unit", 11, 67, 2126803295336845819)
	model.PropertyFlags(520)
	model.PropertyRelation("Unit", 24, 8145676210367791504)
	model.Property("ConsignmentNoteIn", 11, 68, 7586328412098447940)
	model.PropertyFlags(520)
	model.PropertyRelation("ConsignmentNoteIn", 25, 8711340933066469198)
	model.Property("AppId", 9, 69, 6763306841751096049)
	model.PropertyFlags(2080)
	model.PropertyIndex(30, 3500087503737737130)
	model.Property("Locality", 11, 70, 7898631285694533807)
	model.PropertyFlags(520)
	model.PropertyRelation("Locality", 42, 6729944132186829092)
	model.EntityLastPropertyId(70, 7898631285694533807)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (goodsConsignmentNoteIn_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*GoodsConsignmentNoteIn).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (goodsConsignmentNoteIn_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*GoodsConsignmentNoteIn).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (goodsConsignmentNoteIn_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := object.(*GoodsConsignmentNoteIn).ConsignmentNoteIn; rel != nil {
		if rId, err := ConsignmentNoteInBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForConsignmentNoteIn(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*GoodsConsignmentNoteIn).Locality; rel != nil {
		if rId, err := LocalityBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForLocality(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*GoodsConsignmentNoteIn).Subdivision; rel != nil {
		if rId, err := SubdivisionBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForSubdivision(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*GoodsConsignmentNoteIn).GoodsGroup; rel != nil {
		if rId, err := GoodsGroupBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForGoodsGroup(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*GoodsConsignmentNoteIn).Goods; rel != nil {
		if rId, err := GoodsBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForGoods(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*GoodsConsignmentNoteIn).Unit; rel != nil {
		if rId, err := UnitBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForUnit(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (goodsConsignmentNoteIn_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*GoodsConsignmentNoteIn)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on GoodsConsignmentNoteIn.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on GoodsConsignmentNoteIn.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetAppId = fbutils.CreateStringOffset(fbb, obj.AppId)

	var rIdConsignmentNoteIn uint64
	if rel := obj.ConsignmentNoteIn; rel != nil {
		if rId, err := ConsignmentNoteInBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdConsignmentNoteIn = rId
		}
	}

	var rIdLocality uint64
	if rel := obj.Locality; rel != nil {
		if rId, err := LocalityBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdLocality = rId
		}
	}

	var rIdSubdivision uint64
	if rel := obj.Subdivision; rel != nil {
		if rId, err := SubdivisionBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdSubdivision = rId
		}
	}

	var rIdGoodsGroup uint64
	if rel := obj.GoodsGroup; rel != nil {
		if rId, err := GoodsGroupBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdGoodsGroup = rId
		}
	}

	var rIdGoods uint64
	if rel := obj.Goods; rel != nil {
		if rId, err := GoodsBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdGoods = rId
		}
	}

	var rIdUnit uint64
	if rel := obj.Unit; rel != nil {
		if rId, err := UnitBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdUnit = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(70)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 68, offsetAppId)
	if obj.ConsignmentNoteIn != nil {
		fbutils.SetUint64Slot(fbb, 67, rIdConsignmentNoteIn)
	}
	if obj.Locality != nil {
		fbutils.SetUint64Slot(fbb, 69, rIdLocality)
	}
	if obj.Subdivision != nil {
		fbutils.SetUint64Slot(fbb, 63, rIdSubdivision)
	}
	if obj.GoodsGroup != nil {
		fbutils.SetUint64Slot(fbb, 64, rIdGoodsGroup)
	}
	if obj.Goods != nil {
		fbutils.SetUint64Slot(fbb, 65, rIdGoods)
	}
	if obj.Unit != nil {
		fbutils.SetUint64Slot(fbb, 66, rIdUnit)
	}
	fbutils.SetFloat32Slot(fbb, 23, obj.LoadingPercentage)
	fbutils.SetFloat32Slot(fbb, 24, obj.Quantity)
	fbutils.SetInt64Slot(fbb, 61, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 62, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (goodsConsignmentNoteIn_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'GoodsConsignmentNoteIn' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 126))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on GoodsConsignmentNoteIn.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 128))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on GoodsConsignmentNoteIn.UpdatedAt: " + err.Error())
	}

	var relConsignmentNoteIn *ConsignmentNoteIn
	if rId := fbutils.GetUint64PtrSlot(table, 138); rId != nil && *rId > 0 {
		if rObject, err := BoxForConsignmentNoteIn(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relConsignmentNoteIn = rObject
		}
	}

	var relLocality *Locality
	if rId := fbutils.GetUint64PtrSlot(table, 142); rId != nil && *rId > 0 {
		if rObject, err := BoxForLocality(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relLocality = rObject
		}
	}

	var relSubdivision *Subdivision
	if rId := fbutils.GetUint64PtrSlot(table, 130); rId != nil && *rId > 0 {
		if rObject, err := BoxForSubdivision(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relSubdivision = rObject
		}
	}

	var relGoodsGroup *GoodsGroup
	if rId := fbutils.GetUint64PtrSlot(table, 132); rId != nil && *rId > 0 {
		if rObject, err := BoxForGoodsGroup(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relGoodsGroup = rObject
		}
	}

	var relGoods *Goods
	if rId := fbutils.GetUint64PtrSlot(table, 134); rId != nil && *rId > 0 {
		if rObject, err := BoxForGoods(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relGoods = rObject
		}
	}

	var relUnit *Unit
	if rId := fbutils.GetUint64PtrSlot(table, 136); rId != nil && *rId > 0 {
		if rObject, err := BoxForUnit(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relUnit = rObject
		}
	}

	return &GoodsConsignmentNoteIn{
		Id:                propId,
		ExtId:             fbutils.GetStringSlot(table, 6),
		AppId:             fbutils.GetStringSlot(table, 140),
		ConsignmentNoteIn: relConsignmentNoteIn,
		Locality:          relLocality,
		Subdivision:       relSubdivision,
		GoodsGroup:        relGoodsGroup,
		Goods:             relGoods,
		Unit:              relUnit,
		LoadingPercentage: fbutils.GetFloat32Slot(table, 50),
		Quantity:          fbutils.GetFloat32Slot(table, 52),
		CreatedAt:         propCreatedAt,
		UpdatedAt:         propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (goodsConsignmentNoteIn_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*GoodsConsignmentNoteIn, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (goodsConsignmentNoteIn_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*GoodsConsignmentNoteIn), nil)
	}
	return append(slice.([]*GoodsConsignmentNoteIn), object.(*GoodsConsignmentNoteIn))
}

// Box provides CRUD access to GoodsConsignmentNoteIn objects
type GoodsConsignmentNoteInBox struct {
	*objectbox.Box
}

// BoxForGoodsConsignmentNoteIn opens a box of GoodsConsignmentNoteIn objects
func BoxForGoodsConsignmentNoteIn(ob *objectbox.ObjectBox) *GoodsConsignmentNoteInBox {
	return &GoodsConsignmentNoteInBox{
		Box: ob.InternalBox(4),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the GoodsConsignmentNoteIn.Id property on the passed object will be assigned the new ID as well.
func (box *GoodsConsignmentNoteInBox) Put(object *GoodsConsignmentNoteIn) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the GoodsConsignmentNoteIn.Id property on the passed object will be assigned the new ID as well.
func (box *GoodsConsignmentNoteInBox) Insert(object *GoodsConsignmentNoteIn) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *GoodsConsignmentNoteInBox) Update(object *GoodsConsignmentNoteIn) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *GoodsConsignmentNoteInBox) PutAsync(object *GoodsConsignmentNoteIn) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the GoodsConsignmentNoteIn.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the GoodsConsignmentNoteIn.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *GoodsConsignmentNoteInBox) PutMany(objects []*GoodsConsignmentNoteIn) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *GoodsConsignmentNoteInBox) Get(id uint64) (*GoodsConsignmentNoteIn, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*GoodsConsignmentNoteIn), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *GoodsConsignmentNoteInBox) GetMany(ids ...uint64) ([]*GoodsConsignmentNoteIn, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*GoodsConsignmentNoteIn), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *GoodsConsignmentNoteInBox) GetManyExisting(ids ...uint64) ([]*GoodsConsignmentNoteIn, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*GoodsConsignmentNoteIn), nil
}

// GetAll reads all stored objects
func (box *GoodsConsignmentNoteInBox) GetAll() ([]*GoodsConsignmentNoteIn, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*GoodsConsignmentNoteIn), nil
}

// Remove deletes a single object
func (box *GoodsConsignmentNoteInBox) Remove(object *GoodsConsignmentNoteIn) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *GoodsConsignmentNoteInBox) RemoveMany(objects ...*GoodsConsignmentNoteIn) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the GoodsConsignmentNoteIn_ struct to create conditions.
// Keep the *GoodsConsignmentNoteInQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *GoodsConsignmentNoteInBox) Query(conditions ...objectbox.Condition) *GoodsConsignmentNoteInQuery {
	return &GoodsConsignmentNoteInQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the GoodsConsignmentNoteIn_ struct to create conditions.
// Keep the *GoodsConsignmentNoteInQuery if you intend to execute the query multiple times.
func (box *GoodsConsignmentNoteInBox) QueryOrError(conditions ...objectbox.Condition) (*GoodsConsignmentNoteInQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &GoodsConsignmentNoteInQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See GoodsConsignmentNoteInAsyncBox for more information.
func (box *GoodsConsignmentNoteInBox) Async() *GoodsConsignmentNoteInAsyncBox {
	return &GoodsConsignmentNoteInAsyncBox{AsyncBox: box.Box.Async()}
}

// GoodsConsignmentNoteInAsyncBox provides asynchronous operations on GoodsConsignmentNoteIn objects.
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
type GoodsConsignmentNoteInAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForGoodsConsignmentNoteIn creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use GoodsConsignmentNoteInBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForGoodsConsignmentNoteIn(ob *objectbox.ObjectBox, timeoutMs uint64) *GoodsConsignmentNoteInAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 4, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 4: %s" + err.Error())
	}
	return &GoodsConsignmentNoteInAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *GoodsConsignmentNoteInAsyncBox) Put(object *GoodsConsignmentNoteIn) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *GoodsConsignmentNoteInAsyncBox) Insert(object *GoodsConsignmentNoteIn) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *GoodsConsignmentNoteInAsyncBox) Update(object *GoodsConsignmentNoteIn) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *GoodsConsignmentNoteInAsyncBox) Remove(object *GoodsConsignmentNoteIn) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all GoodsConsignmentNoteIn which Id is either 42 or 47:
//
//	box.Query(GoodsConsignmentNoteIn_.Id.In(42, 47)).Find()
type GoodsConsignmentNoteInQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *GoodsConsignmentNoteInQuery) Find() ([]*GoodsConsignmentNoteIn, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*GoodsConsignmentNoteIn), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *GoodsConsignmentNoteInQuery) Offset(offset uint64) *GoodsConsignmentNoteInQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *GoodsConsignmentNoteInQuery) Limit(limit uint64) *GoodsConsignmentNoteInQuery {
	query.Query.Limit(limit)
	return query
}
