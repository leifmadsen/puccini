package v1_1

import (
	"github.com/tliron/puccini/tosca"
)

type GroupType struct {
	*Type `name:"group type"`

	PropertyDefinitions    PropertyDefinitions    `read:"properties,PropertyDefinition" inherit:"properties,Parent"`
	CapabilityDefinitions  CapabilityDefinitions  `read:"capabilities,CapabilityDefinition" inherit:"capabilities,Parent"`
	RequirementDefinitions RequirementDefinitions `read:"requirements,{}RequirementDefinition" inherit:"requirements,Parent"` // sequenced list, but we read it into map
	InterfaceDefinitions   InterfaceDefinitions   `read:"interfaces,InterfaceDefinition" inherit:"interfaces,Parent"`
	MemberNodeTypeNames    *[]string              `read:"members" inherit:"members,Parent"`

	Parent          *GroupType  `lookup:"derived_from,ParentName" inherit:"members,Parent" json:"-" yaml:"-"`
	MemberNodeTypes []*NodeType `lookup:"members,MemberNodeTypeNames" inherit:"members,Parent" json:"-" yaml:"-"`
}

func NewGroupType(context *tosca.Context) *GroupType {
	return &GroupType{
		Type:                   NewType(context),
		PropertyDefinitions:    make(PropertyDefinitions),
		CapabilityDefinitions:  make(CapabilityDefinitions),
		RequirementDefinitions: make(RequirementDefinitions),
		InterfaceDefinitions:   make(InterfaceDefinitions),
	}
}

// tosca.Reader signature
func ReadGroupType(context *tosca.Context) interface{} {
	self := NewGroupType(context)
	context.ValidateUnsupportedFields(context.ReadFields(self, Readers))
	return self
}

func init() {
	Readers["GroupType"] = ReadGroupType
}

// tosca.Hierarchical interface
func (self *GroupType) GetParent() interface{} {
	return self.Parent
}

// tosca.Inherits interface
func (self *GroupType) Inherit() {
	log.Infof("{inherit} group type: %s", self.Name)

	if self.Parent == nil {
		return
	}

	self.PropertyDefinitions.Inherit(self.Parent.PropertyDefinitions)
	self.CapabilityDefinitions.Inherit(self.Parent.CapabilityDefinitions)
	self.RequirementDefinitions.Inherit(self.Parent.RequirementDefinitions)
	self.InterfaceDefinitions.Inherit(self.Parent.InterfaceDefinitions)
}
