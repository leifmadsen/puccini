package v1_1

import (
	"github.com/tliron/puccini/tosca"
)

//
// PropertyDefinition
//

type PropertyDefinition struct {
	*AttributeDefinition `name:"property definition"`

	Required          *bool             `read:"required"`
	ConstraintClauses ConstraintClauses `read:"constraints,[]ConstraintClause"`
}

func NewPropertyDefinition(context *tosca.Context) *PropertyDefinition {
	return &PropertyDefinition{AttributeDefinition: NewAttributeDefinition(context)}
}

// tosca.Reader signature
func ReadPropertyDefinition(context *tosca.Context) interface{} {
	self := NewPropertyDefinition(context)
	context.ValidateUnsupportedFields(context.ReadFields(self, Readers))
	return self
}

func init() {
	Readers["PropertyDefinition"] = ReadPropertyDefinition
}

func (self *PropertyDefinition) Inherit(parentDefinition *PropertyDefinition) {
	if parentDefinition != nil {
		self.AttributeDefinition.Inherit(parentDefinition.AttributeDefinition)

		if (self.Required == nil) && (parentDefinition.Required != nil) {
			self.Required = parentDefinition.Required
		}
		if (self.ConstraintClauses == nil) && (parentDefinition.ConstraintClauses != nil) {
			self.ConstraintClauses = parentDefinition.ConstraintClauses
		}
	} else {
		self.AttributeDefinition.Inherit(nil)
	}
}

//
// PropertyDefinitions
//

type PropertyDefinitions map[string]*PropertyDefinition

func (self PropertyDefinitions) Inherit(parentDefinitions PropertyDefinitions) {
	for name, definition := range parentDefinitions {
		if _, ok := self[name]; !ok {
			self[name] = definition
		}
	}

	for name, definition := range self {
		if parentDefinitions != nil {
			if parentDefinition, ok := parentDefinitions[name]; ok {
				if definition != parentDefinition {
					definition.Inherit(parentDefinition)
				}
				continue
			}
		}

		definition.Inherit(nil)
	}
}
