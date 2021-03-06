package v1_1

import (
	"github.com/tliron/puccini/tosca"
)

//
// OperationDefinition
//

type OperationDefinition struct {
	*Entity `name:"operation definition"`
	Name    string

	Description      *string                  `read:"description"`
	Implementation   *OperationImplementation `read:"implementation,OperationImplementation"`
	InputDefinitions PropertyDefinitions      `read:"inputs,PropertyDefinition"`
}

func NewOperationDefinition(context *tosca.Context) *OperationDefinition {
	return &OperationDefinition{
		Entity:           NewEntity(context),
		Name:             context.Name,
		InputDefinitions: make(PropertyDefinitions),
	}
}

// tosca.Reader signature
func ReadOperationDefinition(context *tosca.Context) interface{} {
	self := NewOperationDefinition(context)
	if context.Is("map") {
		context.ValidateUnsupportedFields(context.ReadFields(self, Readers))
	} else if context.ValidateType("map", "string") {
		self.Implementation = ReadOperationImplementation(context).(*OperationImplementation)
	}
	return self
}

func init() {
	Readers["OperationDefinition"] = ReadOperationDefinition
}

// tosca.Mappable interface
func (self *OperationDefinition) GetKey() string {
	return self.Name
}

func (self *OperationDefinition) Inherit(parentDefinition *OperationDefinition) {
	if parentDefinition != nil {
		if (self.Description == nil) && (parentDefinition.Description != nil) {
			self.Description = parentDefinition.Description
		}

		self.InputDefinitions.Inherit(parentDefinition.InputDefinitions)
	} else {
		self.InputDefinitions.Inherit(nil)
	}
}

//
// OperationDefinitions
//

type OperationDefinitions map[string]*OperationDefinition

func (self OperationDefinitions) Inherit(parentDefinitions OperationDefinitions) {
	for name, definition := range parentDefinitions {
		if _, ok := self[name]; !ok {
			self[name] = definition
		}
	}

	for name, definition := range self {
		if parentDefinition, ok := parentDefinitions[name]; ok {
			if definition != parentDefinition {
				definition.Inherit(parentDefinition)
			}
		} else {
			definition.Inherit(nil)
		}
	}
}
