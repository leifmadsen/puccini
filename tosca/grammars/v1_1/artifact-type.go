package v1_1

import (
	"github.com/tliron/puccini/tosca"
)

//
// ArtifactType
//

type ArtifactType struct {
	*Type `name:"artifact type"`

	PropertyDefinitions PropertyDefinitions `read:"properties,PropertyDefinition" inherit:"properties,Parent"`
	MimeType            *string             `read:"mime_type" inherit:"mime_type,Parent"`
	FileExt             *[]string           `read:"file_ext" inherit:"file_ext,Parent"`

	Parent *ArtifactType `lookup:"derived_from,ParentName" json:"-" yaml:"-"`
}

func NewArtifactType(context *tosca.Context) *ArtifactType {
	return &ArtifactType{
		Type:                NewType(context),
		PropertyDefinitions: make(PropertyDefinitions),
	}
}

// tosca.Reader signature
func ReadArtifactType(context *tosca.Context) interface{} {
	self := NewArtifactType(context)
	context.ValidateUnsupportedFields(context.ReadFields(self, Readers))
	return self
}

func init() {
	Readers["ArtifactType"] = ReadArtifactType
}

// tosca.Hierarchical interface
func (self *ArtifactType) GetParent() interface{} {
	return self.Parent
}

// tosca.Inherits interface
func (self *ArtifactType) Inherit() {
	log.Infof("{inherit} artifact type: %s", self.Name)

	if self.Parent == nil {
		return
	}

	self.PropertyDefinitions.Inherit(self.Parent.PropertyDefinitions)
}
