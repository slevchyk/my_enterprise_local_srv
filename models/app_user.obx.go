// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package models

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type appUser_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var AppUserBinding = appUser_EntityInfo{
	Entity: objectbox.Entity{
		Id: 2,
	},
	Uid: 3772536781229214116,
}

// AppUser_ contains type-based Property helpers to facilitate some common operations such as Queries.
var AppUser_ = struct {
	Id                  *objectbox.PropertyUint64
	ExtId               *objectbox.PropertyString
	CreatedAt           *objectbox.PropertyInt64
	UpdatedAt           *objectbox.PropertyInt64
	FirstName           *objectbox.PropertyString
	LastName            *objectbox.PropertyString
	Email               *objectbox.PropertyString
	Phone               *objectbox.PropertyString
	Token               *objectbox.PropertyString
	IsBlocked           *objectbox.PropertyBool
	IsFarm              *objectbox.PropertyBool
	IsGasStation        *objectbox.PropertyBool
	IsHarvesting        *objectbox.PropertyBool
	IsPayDesk           *objectbox.PropertyBool
	IsWarehouse         *objectbox.PropertyBool
	Password            *objectbox.PropertyString
	IsAdministrator     *objectbox.PropertyBool
	IsDictionaries      *objectbox.PropertyBool
	TokenExpirationDate *objectbox.PropertyInt64
	PhotoPath           *objectbox.PropertyString
	IsManualSelecting   *objectbox.PropertyBool
	IsElevator          *objectbox.PropertyBool
	IsViewMode          *objectbox.PropertyBool
	IsReports           *objectbox.PropertyBool
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &AppUserBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &AppUserBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &AppUserBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &AppUserBinding.Entity,
		},
	},
	FirstName: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &AppUserBinding.Entity,
		},
	},
	LastName: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &AppUserBinding.Entity,
		},
	},
	Email: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &AppUserBinding.Entity,
		},
	},
	Phone: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &AppUserBinding.Entity,
		},
	},
	Token: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsBlocked: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsFarm: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsGasStation: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     14,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsHarvesting: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     15,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsPayDesk: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     16,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsWarehouse: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     17,
			Entity: &AppUserBinding.Entity,
		},
	},
	Password: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     18,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsAdministrator: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     19,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsDictionaries: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     20,
			Entity: &AppUserBinding.Entity,
		},
	},
	TokenExpirationDate: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     21,
			Entity: &AppUserBinding.Entity,
		},
	},
	PhotoPath: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     22,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsManualSelecting: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     24,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsElevator: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     25,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsViewMode: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     26,
			Entity: &AppUserBinding.Entity,
		},
	},
	IsReports: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     27,
			Entity: &AppUserBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (appUser_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (appUser_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("AppUser", 2, 3772536781229214116)
	model.Property("Id", 6, 1, 1727067599752868172)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 1751440294866873485)
	model.Property("CreatedAt", 10, 5, 2560859963279789654)
	model.Property("UpdatedAt", 10, 6, 6879137370614292715)
	model.Property("FirstName", 9, 7, 5371987638100000232)
	model.Property("LastName", 9, 8, 7064546187212738440)
	model.Property("Email", 9, 9, 4256892128440117593)
	model.Property("Phone", 9, 10, 8766869086574508498)
	model.Property("Token", 9, 11, 6198435271220425000)
	model.Property("IsBlocked", 1, 12, 5936673092836650890)
	model.Property("IsFarm", 1, 13, 6391471651868091211)
	model.Property("IsGasStation", 1, 14, 8380961054502993825)
	model.Property("IsHarvesting", 1, 15, 5881266140801353583)
	model.Property("IsPayDesk", 1, 16, 6970975710831344500)
	model.Property("IsWarehouse", 1, 17, 6958713036560849743)
	model.Property("Password", 9, 18, 7965139643554191174)
	model.Property("IsAdministrator", 1, 19, 3797028530307446237)
	model.Property("IsDictionaries", 1, 20, 6894104621440000829)
	model.Property("TokenExpirationDate", 10, 21, 6565581724974588785)
	model.Property("PhotoPath", 9, 22, 2572403520298632668)
	model.Property("IsManualSelecting", 1, 24, 8642972488482966165)
	model.Property("IsElevator", 1, 25, 8686396539319905421)
	model.Property("IsViewMode", 1, 26, 8979646598945838209)
	model.Property("IsReports", 1, 27, 8973871661349720378)
	model.EntityLastPropertyId(27, 8973871661349720378)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (appUser_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*AppUser).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (appUser_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*AppUser).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (appUser_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (appUser_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*AppUser)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on AppUser.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on AppUser.UpdatedAt: " + err.Error())
		}
	}

	var propTokenExpirationDate int64
	{
		var err error
		propTokenExpirationDate, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.TokenExpirationDate)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on AppUser.TokenExpirationDate: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)
	var offsetFirstName = fbutils.CreateStringOffset(fbb, obj.FirstName)
	var offsetLastName = fbutils.CreateStringOffset(fbb, obj.LastName)
	var offsetEmail = fbutils.CreateStringOffset(fbb, obj.Email)
	var offsetPhone = fbutils.CreateStringOffset(fbb, obj.Phone)
	var offsetToken = fbutils.CreateStringOffset(fbb, obj.Token)
	var offsetPassword = fbutils.CreateStringOffset(fbb, obj.Password)
	var offsetPhotoPath = fbutils.CreateStringOffset(fbb, obj.PhotoPath)

	// build the FlatBuffers object
	fbb.StartObject(27)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetFirstName)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetLastName)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetEmail)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetPhone)
	fbutils.SetUOffsetTSlot(fbb, 17, offsetPassword)
	fbutils.SetUOffsetTSlot(fbb, 21, offsetPhotoPath)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetToken)
	fbutils.SetInt64Slot(fbb, 20, propTokenExpirationDate)
	fbutils.SetBoolSlot(fbb, 18, obj.IsAdministrator)
	fbutils.SetBoolSlot(fbb, 23, obj.IsManualSelecting)
	fbutils.SetBoolSlot(fbb, 11, obj.IsBlocked)
	fbutils.SetBoolSlot(fbb, 12, obj.IsFarm)
	fbutils.SetBoolSlot(fbb, 13, obj.IsGasStation)
	fbutils.SetBoolSlot(fbb, 14, obj.IsHarvesting)
	fbutils.SetBoolSlot(fbb, 15, obj.IsPayDesk)
	fbutils.SetBoolSlot(fbb, 16, obj.IsWarehouse)
	fbutils.SetBoolSlot(fbb, 26, obj.IsReports)
	fbutils.SetBoolSlot(fbb, 19, obj.IsDictionaries)
	fbutils.SetBoolSlot(fbb, 24, obj.IsElevator)
	fbutils.SetBoolSlot(fbb, 25, obj.IsViewMode)
	fbutils.SetInt64Slot(fbb, 4, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 5, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (appUser_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'AppUser' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 12))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on AppUser.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 14))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on AppUser.UpdatedAt: " + err.Error())
	}

	propTokenExpirationDate, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 44))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on AppUser.TokenExpirationDate: " + err.Error())
	}

	return &AppUser{
		Id:                  propId,
		ExtId:               fbutils.GetStringSlot(table, 6),
		FirstName:           fbutils.GetStringSlot(table, 16),
		LastName:            fbutils.GetStringSlot(table, 18),
		Email:               fbutils.GetStringSlot(table, 20),
		Phone:               fbutils.GetStringSlot(table, 22),
		Password:            fbutils.GetStringSlot(table, 38),
		PhotoPath:           fbutils.GetStringSlot(table, 46),
		Token:               fbutils.GetStringSlot(table, 24),
		TokenExpirationDate: propTokenExpirationDate,
		IsAdministrator:     fbutils.GetBoolSlot(table, 40),
		IsManualSelecting:   fbutils.GetBoolSlot(table, 50),
		IsBlocked:           fbutils.GetBoolSlot(table, 26),
		IsFarm:              fbutils.GetBoolSlot(table, 28),
		IsGasStation:        fbutils.GetBoolSlot(table, 30),
		IsHarvesting:        fbutils.GetBoolSlot(table, 32),
		IsPayDesk:           fbutils.GetBoolSlot(table, 34),
		IsWarehouse:         fbutils.GetBoolSlot(table, 36),
		IsReports:           fbutils.GetBoolSlot(table, 56),
		IsDictionaries:      fbutils.GetBoolSlot(table, 42),
		IsElevator:          fbutils.GetBoolSlot(table, 52),
		IsViewMode:          fbutils.GetBoolSlot(table, 54),
		CreatedAt:           propCreatedAt,
		UpdatedAt:           propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (appUser_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*AppUser, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (appUser_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*AppUser), nil)
	}
	return append(slice.([]*AppUser), object.(*AppUser))
}

// Box provides CRUD access to AppUser objects
type AppUserBox struct {
	*objectbox.Box
}

// BoxForAppUser opens a box of AppUser objects
func BoxForAppUser(ob *objectbox.ObjectBox) *AppUserBox {
	return &AppUserBox{
		Box: ob.InternalBox(2),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the AppUser.Id property on the passed object will be assigned the new ID as well.
func (box *AppUserBox) Put(object *AppUser) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the AppUser.Id property on the passed object will be assigned the new ID as well.
func (box *AppUserBox) Insert(object *AppUser) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *AppUserBox) Update(object *AppUser) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *AppUserBox) PutAsync(object *AppUser) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the AppUser.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the AppUser.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *AppUserBox) PutMany(objects []*AppUser) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *AppUserBox) Get(id uint64) (*AppUser, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*AppUser), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *AppUserBox) GetMany(ids ...uint64) ([]*AppUser, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUser), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *AppUserBox) GetManyExisting(ids ...uint64) ([]*AppUser, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUser), nil
}

// GetAll reads all stored objects
func (box *AppUserBox) GetAll() ([]*AppUser, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUser), nil
}

// Remove deletes a single object
func (box *AppUserBox) Remove(object *AppUser) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *AppUserBox) RemoveMany(objects ...*AppUser) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the AppUser_ struct to create conditions.
// Keep the *AppUserQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *AppUserBox) Query(conditions ...objectbox.Condition) *AppUserQuery {
	return &AppUserQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the AppUser_ struct to create conditions.
// Keep the *AppUserQuery if you intend to execute the query multiple times.
func (box *AppUserBox) QueryOrError(conditions ...objectbox.Condition) (*AppUserQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &AppUserQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See AppUserAsyncBox for more information.
func (box *AppUserBox) Async() *AppUserAsyncBox {
	return &AppUserAsyncBox{AsyncBox: box.Box.Async()}
}

// AppUserAsyncBox provides asynchronous operations on AppUser objects.
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
type AppUserAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForAppUser creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use AppUserBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForAppUser(ob *objectbox.ObjectBox, timeoutMs uint64) *AppUserAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 2, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 2: %s" + err.Error())
	}
	return &AppUserAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *AppUserAsyncBox) Put(object *AppUser) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *AppUserAsyncBox) Insert(object *AppUser) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *AppUserAsyncBox) Update(object *AppUser) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *AppUserAsyncBox) Remove(object *AppUser) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all AppUser which Id is either 42 or 47:
// 		box.Query(AppUser_.Id.In(42, 47)).Find()
type AppUserQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *AppUserQuery) Find() ([]*AppUser, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUser), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *AppUserQuery) Offset(offset uint64) *AppUserQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *AppUserQuery) Limit(limit uint64) *AppUserQuery {
	query.Query.Limit(limit)
	return query
}

type appUserCniRecipient_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var AppUserCniRecipientBinding = appUserCniRecipient_EntityInfo{
	Entity: objectbox.Entity{
		Id: 17,
	},
	Uid: 1720428749164872770,
}

// AppUserCniRecipient_ contains type-based Property helpers to facilitate some common operations such as Queries.
var AppUserCniRecipient_ = struct {
	Id        *objectbox.PropertyUint64
	ExtId     *objectbox.PropertyString
	AppUser   *objectbox.RelationToOne
	Recipient *objectbox.RelationToOne
	IsActive  *objectbox.PropertyBool
	CreatedAt *objectbox.PropertyInt64
	UpdatedAt *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
	},
	ExtId: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
	},
	AppUser: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     4,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
		Target: &AppUserBinding.Entity,
	},
	Recipient: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     5,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
		Target: &StorageBinding.Entity,
	},
	IsActive: &objectbox.PropertyBool{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
	},
	CreatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
	},
	UpdatedAt: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &AppUserCniRecipientBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (appUserCniRecipient_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (appUserCniRecipient_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("AppUserCniRecipient", 17, 1720428749164872770)
	model.Property("Id", 6, 1, 566824801514946335)
	model.PropertyFlags(1)
	model.Property("ExtId", 9, 2, 7893636772228203944)
	model.PropertyFlags(2080)
	model.PropertyIndex(44, 6950181797625918299)
	model.Property("AppUser", 11, 4, 570692274894039251)
	model.PropertyFlags(520)
	model.PropertyRelation("AppUser", 46, 4711253548737289129)
	model.Property("Recipient", 11, 5, 2253899333105330226)
	model.PropertyFlags(520)
	model.PropertyRelation("Storage", 47, 8288459648892467887)
	model.Property("IsActive", 1, 6, 5353244190303935941)
	model.Property("CreatedAt", 10, 7, 5691837511673698860)
	model.Property("UpdatedAt", 10, 8, 5882832746778650044)
	model.EntityLastPropertyId(8, 5882832746778650044)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (appUserCniRecipient_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*AppUserCniRecipient).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (appUserCniRecipient_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*AppUserCniRecipient).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (appUserCniRecipient_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := object.(*AppUserCniRecipient).AppUser; rel != nil {
		if rId, err := AppUserBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForAppUser(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := object.(*AppUserCniRecipient).Recipient; rel != nil {
		if rId, err := StorageBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForStorage(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (appUserCniRecipient_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*AppUserCniRecipient)
	var propCreatedAt int64
	{
		var err error
		propCreatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.CreatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on AppUserCniRecipient.CreatedAt: " + err.Error())
		}
	}

	var propUpdatedAt int64
	{
		var err error
		propUpdatedAt, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.UpdatedAt)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on AppUserCniRecipient.UpdatedAt: " + err.Error())
		}
	}

	var offsetExtId = fbutils.CreateStringOffset(fbb, obj.ExtId)

	var rIdAppUser uint64
	if rel := obj.AppUser; rel != nil {
		if rId, err := AppUserBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdAppUser = rId
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

	// build the FlatBuffers object
	fbb.StartObject(8)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetExtId)
	if obj.AppUser != nil {
		fbutils.SetUint64Slot(fbb, 3, rIdAppUser)
	}
	if obj.Recipient != nil {
		fbutils.SetUint64Slot(fbb, 4, rIdRecipient)
	}
	fbutils.SetBoolSlot(fbb, 5, obj.IsActive)
	fbutils.SetInt64Slot(fbb, 6, propCreatedAt)
	fbutils.SetInt64Slot(fbb, 7, propUpdatedAt)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (appUserCniRecipient_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'AppUserCniRecipient' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propCreatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 16))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on AppUserCniRecipient.CreatedAt: " + err.Error())
	}

	propUpdatedAt, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 18))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on AppUserCniRecipient.UpdatedAt: " + err.Error())
	}

	var relAppUser *AppUser
	if rId := fbutils.GetUint64PtrSlot(table, 10); rId != nil && *rId > 0 {
		if rObject, err := BoxForAppUser(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relAppUser = rObject
		}
	}

	var relRecipient *Storage
	if rId := fbutils.GetUint64PtrSlot(table, 12); rId != nil && *rId > 0 {
		if rObject, err := BoxForStorage(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relRecipient = rObject
		}
	}

	return &AppUserCniRecipient{
		Id:        propId,
		ExtId:     fbutils.GetStringSlot(table, 6),
		AppUser:   relAppUser,
		Recipient: relRecipient,
		IsActive:  fbutils.GetBoolSlot(table, 14),
		CreatedAt: propCreatedAt,
		UpdatedAt: propUpdatedAt,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (appUserCniRecipient_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*AppUserCniRecipient, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (appUserCniRecipient_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*AppUserCniRecipient), nil)
	}
	return append(slice.([]*AppUserCniRecipient), object.(*AppUserCniRecipient))
}

// Box provides CRUD access to AppUserCniRecipient objects
type AppUserCniRecipientBox struct {
	*objectbox.Box
}

// BoxForAppUserCniRecipient opens a box of AppUserCniRecipient objects
func BoxForAppUserCniRecipient(ob *objectbox.ObjectBox) *AppUserCniRecipientBox {
	return &AppUserCniRecipientBox{
		Box: ob.InternalBox(17),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the AppUserCniRecipient.Id property on the passed object will be assigned the new ID as well.
func (box *AppUserCniRecipientBox) Put(object *AppUserCniRecipient) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the AppUserCniRecipient.Id property on the passed object will be assigned the new ID as well.
func (box *AppUserCniRecipientBox) Insert(object *AppUserCniRecipient) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *AppUserCniRecipientBox) Update(object *AppUserCniRecipient) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *AppUserCniRecipientBox) PutAsync(object *AppUserCniRecipient) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the AppUserCniRecipient.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the AppUserCniRecipient.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *AppUserCniRecipientBox) PutMany(objects []*AppUserCniRecipient) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *AppUserCniRecipientBox) Get(id uint64) (*AppUserCniRecipient, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*AppUserCniRecipient), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *AppUserCniRecipientBox) GetMany(ids ...uint64) ([]*AppUserCniRecipient, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUserCniRecipient), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *AppUserCniRecipientBox) GetManyExisting(ids ...uint64) ([]*AppUserCniRecipient, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUserCniRecipient), nil
}

// GetAll reads all stored objects
func (box *AppUserCniRecipientBox) GetAll() ([]*AppUserCniRecipient, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUserCniRecipient), nil
}

// Remove deletes a single object
func (box *AppUserCniRecipientBox) Remove(object *AppUserCniRecipient) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *AppUserCniRecipientBox) RemoveMany(objects ...*AppUserCniRecipient) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the AppUserCniRecipient_ struct to create conditions.
// Keep the *AppUserCniRecipientQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *AppUserCniRecipientBox) Query(conditions ...objectbox.Condition) *AppUserCniRecipientQuery {
	return &AppUserCniRecipientQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the AppUserCniRecipient_ struct to create conditions.
// Keep the *AppUserCniRecipientQuery if you intend to execute the query multiple times.
func (box *AppUserCniRecipientBox) QueryOrError(conditions ...objectbox.Condition) (*AppUserCniRecipientQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &AppUserCniRecipientQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See AppUserCniRecipientAsyncBox for more information.
func (box *AppUserCniRecipientBox) Async() *AppUserCniRecipientAsyncBox {
	return &AppUserCniRecipientAsyncBox{AsyncBox: box.Box.Async()}
}

// AppUserCniRecipientAsyncBox provides asynchronous operations on AppUserCniRecipient objects.
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
type AppUserCniRecipientAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForAppUserCniRecipient creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use AppUserCniRecipientBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForAppUserCniRecipient(ob *objectbox.ObjectBox, timeoutMs uint64) *AppUserCniRecipientAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 17, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 17: %s" + err.Error())
	}
	return &AppUserCniRecipientAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *AppUserCniRecipientAsyncBox) Put(object *AppUserCniRecipient) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *AppUserCniRecipientAsyncBox) Insert(object *AppUserCniRecipient) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *AppUserCniRecipientAsyncBox) Update(object *AppUserCniRecipient) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *AppUserCniRecipientAsyncBox) Remove(object *AppUserCniRecipient) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all AppUserCniRecipient which Id is either 42 or 47:
// 		box.Query(AppUserCniRecipient_.Id.In(42, 47)).Find()
type AppUserCniRecipientQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *AppUserCniRecipientQuery) Find() ([]*AppUserCniRecipient, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*AppUserCniRecipient), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *AppUserCniRecipientQuery) Offset(offset uint64) *AppUserCniRecipientQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *AppUserCniRecipientQuery) Limit(limit uint64) *AppUserCniRecipientQuery {
	query.Query.Limit(limit)
	return query
}
