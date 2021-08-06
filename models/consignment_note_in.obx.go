// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type consignmentNoteIn_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ConsignmentNoteInBinding = consignmentNoteIn_EntityInfo{
	Entity: objectbox.Entity{
		Id: 3,
	},
	Uid: 8603667634404417786,
}

// ConsignmentNoteIn_ contains type-based Property helpers to facilitate some common operations such as Queries.
var ConsignmentNoteIn_ = struct {
	Id            *objectbox.PropertyUint64
	ExtId         *objectbox.PropertyString
	Date          *objectbox.PropertyInt64
	Number        *objectbox.PropertyString
	DepartureDate *objectbox.PropertyInt64
	IsDeleted     *objectbox.PropertyBool
	CreatedAt     *objectbox.PropertyInt64
	UpdatedAt     *objectbox.PropertyInt64
	Driver        *objectbox.RelationToOne
	Recipient     *objectbox.RelationToOne
	Sender        *objectbox.RelationToOne
	HarvestType   *objectbox.RelationToOne
	Vehicle       *objectbox.RelationToOne
	AppId         *objectbox.PropertyString
	Gross         *objectbox.PropertyFloat64
	Tare          *objectbox.PropertyFloat64
	Net           *objectbox.PropertyFloat64
	Humidity      *objectbox.PropertyFloat64
	Weediness     *objectbox.PropertyFloat64
	AppUser       *objectbox.RelationToOne
	ChangedByAcc  *objectbox.PropertyBool
	ChangedByApp  *objectbox.PropertyBool
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Date: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Number: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	DepartureDate: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     14,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	IsDeleted: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     35,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     36,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     37,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Driver: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     45,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &PersonBinding.Entity,
	},
	Recipient: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     46,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &StorageBinding.Entity,
	},
	Sender: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     47,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &StorageBinding.Entity,
	},
	HarvestType: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     58,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &HarvestTypeBinding.Entity,
	},
	Vehicle: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     59,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &VehicleBinding.Entity,
	},
	AppId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     60,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Gross: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     61,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Tare: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     62,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Net: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     63,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Humidity: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     64,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	Weediness: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     65,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	AppUser: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     66,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
		Target: &AppUserBinding.Entity,
	},
	ChangedByAcc: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     67,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
	ChangedByApp: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     68,
			Entity: &ConsignmentNoteInBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (consignmentNoteIn_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (consignmentNoteIn_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("ConsignmentNoteIn", 3, 8603667634404417786)
	model.Property("Id", 6, 1, 7381584349873625452)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 6555419099548205599)
	model.PropertyFlags(2080)
	model.PropertyIndex(27, 8472944323867149148)
	model.Property("Date", 10, 4, 3681808398951988176)
	model.Property("Number", 9, 5, 3024305087246974805)
	model.Property("DepartureDate", 10, 14, 214961481867191822)
	model.Property("IsDeleted", 1, 35, 638093717419802601)
	model.Property("CreatedAt", 10, 36, 8899633992386505214)
	model.Property("UpdatedAt", 10, 37, 8437818475573608947)
	model.Property("Driver", 11, 45, 6637254081577110712)
	model.PropertyFlags(520)
	model.PropertyRelation("Person", 5, 2579239408896235758)
	model.Property("Recipient", 11, 46, 1634970022263309275)
	model.PropertyFlags(520)
	model.PropertyRelation("Storage", 6, 2583116323191521888)
	model.Property("Sender", 11, 47, 5548493982286844419)
	model.PropertyFlags(520)
	model.PropertyRelation("Storage", 7, 8776324619587302163)
	model.Property("HarvestType", 11, 58, 3732019176814765919)
	model.PropertyFlags(520)
	model.PropertyRelation("HarvestType", 17, 5703090548301462582)
	model.Property("Vehicle", 11, 59, 5453828081365110257)
	model.PropertyFlags(520)
	model.PropertyRelation("Vehicle", 18, 2240474060564286524)
	model.Property("AppId", 9, 60, 6042246242714324374)
	model.PropertyFlags(2080)
	model.PropertyIndex(28, 529896894866044814)
	model.Property("Gross", 8, 61, 6424679353718418180)
	model.Property("Tare", 8, 62, 4959522081134873497)
	model.Property("Net", 8, 63, 51868586889488965)
	model.Property("Humidity", 8, 64, 7455670039623708115)
	model.Property("Weediness", 8, 65, 8126686263197659917)
	model.Property("AppUser", 11, 66, 5908904309199154956)
	model.PropertyFlags(520)
	model.PropertyRelation("AppUser", 26, 2796045731298863334)
	model.Property("ChangedByAcc", 1, 67, 5779134778575088849)
	model.Property("ChangedByApp", 1, 68, 3935944489437497465)
	model.EntityLastPropertyId(68, 3935944489437497465)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (consignmentNoteIn_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*ConsignmentNoteIn).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (consignmentNoteIn_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*ConsignmentNoteIn).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (consignmentNoteIn_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := object.(*ConsignmentNoteIn).HarvestType; rel != nil {
		if rId, err := HarvestTypeBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForHarvestType(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*ConsignmentNoteIn).Vehicle; rel != nil {
		if rId, err := VehicleBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForVehicle(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*ConsignmentNoteIn).Driver; rel != nil {
		if rId, err := PersonBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForPerson(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*ConsignmentNoteIn).Recipient; rel != nil {
		if rId, err := StorageBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForStorage(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*ConsignmentNoteIn).Sender; rel != nil {
		if rId, err := StorageBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForStorage(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*ConsignmentNoteIn).AppUser; rel != nil {
		if rId, err := AppUserBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForAppUser(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (consignmentNoteIn_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*ConsignmentNoteIn)
	var propDate int64
	{
		var err error
		propDate, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.Date)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on ConsignmentNoteIn.Date: " + err.Error())
		}
	}

	var propDepartureDate int64
	{
		var err error
		propDepartureDate, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.DepartureDate)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on ConsignmentNoteIn.DepartureDate: " + err.Error())
		}
	}

	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on ConsignmentNoteIn.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on ConsignmentNoteIn.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetNumber = fbutils.CreateStringOffset(fbb, obj.Number)
	var offsetAppId = fbutils.CreateStringOffset(fbb, obj.AppId)

	var rIdHarvestType uint64
	if rel := obj.HarvestType; rel != nil {
		if rId, err := HarvestTypeBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdHarvestType = rId
		}
	}

	var rIdVehicle uint64
	if rel := obj.Vehicle; rel != nil {
		if rId, err := VehicleBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdVehicle = rId
		}
	}

	var rIdDriver uint64
	if rel := obj.Driver; rel != nil {
		if rId, err := PersonBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdDriver = rId
		}
	}

	var rIdRecipient uint64
	if rel := obj.Recipient; rel != nil {
		if rId, err := StorageBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdRecipient = rId
		}
	}

	var rIdSender uint64
	if rel := obj.Sender; rel != nil {
		if rId, err := StorageBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdSender = rId
		}
	}

	var rIdAppUser uint64
	if rel := obj.AppUser; rel != nil {
		if rId, err := AppUserBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdAppUser = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(68)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 59, offsetAppId)
	fbutils.SetInt64Slot(fbb, 3, propDate)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetNumber)
	if obj.HarvestType != nil {
		fbutils.SetUint64Slot(fbb, 57, rIdHarvestType)
	}
	if obj.Vehicle != nil {
		fbutils.SetUint64Slot(fbb, 58, rIdVehicle)
	}
	fbutils.SetInt64Slot(fbb, 13, propDepartureDate)
	if obj.Driver != nil {
		fbutils.SetUint64Slot(fbb, 44, rIdDriver)
	}
	if obj.Recipient != nil {
		fbutils.SetUint64Slot(fbb, 45, rIdRecipient)
	}
	if obj.Sender != nil {
		fbutils.SetUint64Slot(fbb, 46, rIdSender)
	}
	if obj.AppUser != nil {
		fbutils.SetUint64Slot(fbb, 65, rIdAppUser)
	}
	fbutils.SetFloat64Slot(fbb, 60, obj.Gross)
	fbutils.SetFloat64Slot(fbb, 61, obj.Tare)
	fbutils.SetFloat64Slot(fbb, 62, obj.Net)
	fbutils.SetFloat64Slot(fbb, 63, obj.Humidity)
	fbutils.SetFloat64Slot(fbb, 64, obj.Weediness)
	fbutils.SetBoolSlot(fbb, 34, obj.IsDeleted)
	fbutils.SetInt64Slot(fbb, 35, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 36, propUpdatedAt)
	fbutils.SetBoolSlot(fbb, 67, obj.ChangedByApp)
	fbutils.SetBoolSlot(fbb, 66, obj.ChangedByAcc)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (consignmentNoteIn_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'ConsignmentNoteIn' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propDate, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 10))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on ConsignmentNoteIn.Date: " + err.Error())
	}

	propDepartureDate, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 30))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on ConsignmentNoteIn.DepartureDate: " + err.Error())
	}

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 74))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on ConsignmentNoteIn.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 76))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on ConsignmentNoteIn.UpdatedAt: " + err.Error())
	}

	var relHarvestType *HarvestType
	if rId := fbutils.GetUint64PtrSlot(table, 118); rId != nil && *rId > 0 {
		if rObject, err := BoxForHarvestType(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relHarvestType = rObject
		}
	}

	var relVehicle *Vehicle
	if rId := fbutils.GetUint64PtrSlot(table, 120); rId != nil && *rId > 0 {
		if rObject, err := BoxForVehicle(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relVehicle = rObject
		}
	}

	var relDriver *Person
	if rId := fbutils.GetUint64PtrSlot(table, 92); rId != nil && *rId > 0 {
		if rObject, err := BoxForPerson(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relDriver = rObject
		}
	}

	var relRecipient *Storage
	if rId := fbutils.GetUint64PtrSlot(table, 94); rId != nil && *rId > 0 {
		if rObject, err := BoxForStorage(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relRecipient = rObject
		}
	}

	var relSender *Storage
	if rId := fbutils.GetUint64PtrSlot(table, 96); rId != nil && *rId > 0 {
		if rObject, err := BoxForStorage(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relSender = rObject
		}
	}

	var relAppUser *AppUser
	if rId := fbutils.GetUint64PtrSlot(table, 134); rId != nil && *rId > 0 {
		if rObject, err := BoxForAppUser(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relAppUser = rObject
		}
	}

	return &ConsignmentNoteIn{
		Id:            propId,
		ExtId:         fbutils.GetStringSlot(table, 6),
		AppId:         fbutils.GetStringSlot(table, 122),
		Date:          propDate,
		Number:        fbutils.GetStringSlot(table, 12),
		HarvestType:   relHarvestType,
		Vehicle:       relVehicle,
		DepartureDate: propDepartureDate,
		Driver:        relDriver,
		Recipient:     relRecipient,
		Sender:        relSender,
		AppUser:       relAppUser,
		Gross:         fbutils.GetFloat64Slot(table, 124),
		Tare:          fbutils.GetFloat64Slot(table, 126),
		Net:           fbutils.GetFloat64Slot(table, 128),
		Humidity:      fbutils.GetFloat64Slot(table, 130),
		Weediness:     fbutils.GetFloat64Slot(table, 132),
		IsDeleted:     fbutils.GetBoolSlot(table, 72),
		CreatedAt:     propCreatedAt,
		UpdatedAt:     propUpdatedAt,
		ChangedByApp:  fbutils.GetBoolSlot(table, 138),
		ChangedByAcc:  fbutils.GetBoolSlot(table, 136),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (consignmentNoteIn_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*ConsignmentNoteIn, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (consignmentNoteIn_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*ConsignmentNoteIn), nil)
	}
	return append(slice.([]*ConsignmentNoteIn), object.(*ConsignmentNoteIn))
}

// Box provides CRUD access to ConsignmentNoteIn objects
type ConsignmentNoteInBox struct {
	*objectbox.Box
}

// BoxForConsignmentNoteIn opens a box of ConsignmentNoteIn objects
func BoxForConsignmentNoteIn(ob *objectbox.ObjectBox) *ConsignmentNoteInBox {
	return &ConsignmentNoteInBox{
		Box: ob.InternalBox(3),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ConsignmentNoteIn.Id property on the passed object will be assigned the new ID as well.
func (box *ConsignmentNoteInBox) Put(object *ConsignmentNoteIn) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ConsignmentNoteIn.Id property on the passed object will be assigned the new ID as well.
func (box *ConsignmentNoteInBox) Insert(object *ConsignmentNoteIn) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *ConsignmentNoteInBox) Update(object *ConsignmentNoteIn) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *ConsignmentNoteInBox) PutAsync(object *ConsignmentNoteIn) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the ConsignmentNoteIn.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the ConsignmentNoteIn.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ConsignmentNoteInBox) PutMany(objects []*ConsignmentNoteIn) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ConsignmentNoteInBox) Get(id uint64) (*ConsignmentNoteIn, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*ConsignmentNoteIn), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *ConsignmentNoteInBox) GetMany(ids ...uint64) ([]*ConsignmentNoteIn, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*ConsignmentNoteIn), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *ConsignmentNoteInBox) GetManyExisting(ids ...uint64) ([]*ConsignmentNoteIn, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*ConsignmentNoteIn), nil
}

// GetAll reads all stored objects
func (box *ConsignmentNoteInBox) GetAll() ([]*ConsignmentNoteIn, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*ConsignmentNoteIn), nil
}

// Remove deletes a single object
func (box *ConsignmentNoteInBox) Remove(object *ConsignmentNoteIn) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *ConsignmentNoteInBox) RemoveMany(objects ...*ConsignmentNoteIn) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the ConsignmentNoteIn_ struct to create conditions.
// Keep the *ConsignmentNoteInQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ConsignmentNoteInBox) Query(conditions ...objectbox.Condition) *ConsignmentNoteInQuery {
	return &ConsignmentNoteInQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the ConsignmentNoteIn_ struct to create conditions.
// Keep the *ConsignmentNoteInQuery if you intend to execute the query multiple times.
func (box *ConsignmentNoteInBox) QueryOrError(conditions ...objectbox.Condition) (*ConsignmentNoteInQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ConsignmentNoteInQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See ConsignmentNoteInAsyncBox for more information.
func (box *ConsignmentNoteInBox) Async() *ConsignmentNoteInAsyncBox {
	return &ConsignmentNoteInAsyncBox{AsyncBox: box.Box.Async()}
}

// ConsignmentNoteInAsyncBox provides asynchronous operations on ConsignmentNoteIn objects.
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
type ConsignmentNoteInAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForConsignmentNoteIn creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use ConsignmentNoteInBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForConsignmentNoteIn(ob *objectbox.ObjectBox, timeoutMs uint64) *ConsignmentNoteInAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 3, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 3: %s" + err.Error())
	}
	return &ConsignmentNoteInAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *ConsignmentNoteInAsyncBox) Put(object *ConsignmentNoteIn) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *ConsignmentNoteInAsyncBox) Insert(object *ConsignmentNoteIn) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *ConsignmentNoteInAsyncBox) Update(object *ConsignmentNoteIn) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *ConsignmentNoteInAsyncBox) Remove(object *ConsignmentNoteIn) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all ConsignmentNoteIn which Id is either 42 or 47:
// 		box.Query(ConsignmentNoteIn_.Id.In(42, 47)).Find()
type ConsignmentNoteInQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ConsignmentNoteInQuery) Find() ([]*ConsignmentNoteIn, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*ConsignmentNoteIn), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ConsignmentNoteInQuery) Offset(offset uint64) *ConsignmentNoteInQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ConsignmentNoteInQuery) Limit(limit uint64) *ConsignmentNoteInQuery {
	query.Query.Limit(limit)
	return query
}
