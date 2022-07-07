// Code generated by ObjectBox; DO NOT EDIT.

package models

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(6)

	model.RegisterBinding(PersonBinding)
	model.RegisterBinding(AppUserBinding)
	model.RegisterBinding(ConsignmentNoteInBinding)
	model.RegisterBinding(GoodsConsignmentNoteInBinding)
	model.RegisterBinding(GoodsBinding)
	model.RegisterBinding(GoodsGroupBinding)
	model.RegisterBinding(HarvestTypeBinding)
	model.RegisterBinding(StorageBinding)
	model.RegisterBinding(SubdivisionBinding)
	model.RegisterBinding(UnitBinding)
	model.RegisterBinding(VehicleBinding)
	model.RegisterBinding(TrailerBinding)
	model.RegisterBinding(LocalityBinding)
	model.RegisterBinding(ServiceWorkerBinding)
	model.RegisterBinding(AppUserCniRecipientBinding)
	model.LastEntityId(17, 1720428749164872770)
	model.LastIndexId(47, 8288459648892467887)

	return model
}
