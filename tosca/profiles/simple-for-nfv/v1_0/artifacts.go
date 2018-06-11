// This file was auto-generated from YAML files

package v1_0

func init() {
	Profile["/tosca/simple-for-nfv/1.0/artifacts.yaml"] = `
# Modified from a file that was distributed with this NOTICE:
#
#   Apache AriaTosca
#   Copyright 2016-2017 The Apache Software Foundation
#
#   This product includes software developed at
#   The Apache Software Foundation (http://www.apache.org/).

tosca_definitions_version: tosca_simple_yaml_1_1

artifact_types:

  tosca.artifacts.nfv.SwImage:
    metadata:
      normative: 'true'
      specification: tosca-simple-nfv-1.0
      specification_section: 5.4.1
    derived_from: tosca.artifacts.Deployment.Image
    properties:
      name:
        description: >-
          Name of this software image.
        type: string
        required: true
      version:
        description: >-
          Version of this software image.
        type: string
        required: true
      checksum:
        description: >-
          Checksum of the software image file.
        type: string
      container_format:
        description: >-
          The container format describes the container file format in which software image is
          provided.
        type: string
        required: true
      disk_format:
        description: >-
          The disk format of a software image is the format of the underlying disk image.
        type: string
        required: true
      min_disk:
        description: >-
          The minimal disk size requirement for this software image.
        type: scalar-unit.size
        required: true
      min_ram:
        description: >-
          The minimal disk size requirement for this software image.
        type: scalar-unit.size
        required: false
      size: # ARIA NOTE: section [5.4.1.1 Properties] calls this field 'Size'
        description: >-
          The size of this software image
        type: scalar-unit.size
        required: true
      sw_image:
        description: >-
          A reference to the actual software image within VNF Package, or url.
        type: string
        required: true
      operating_system:
        description: >-
          Identifies the operating system used in the software image.
        type: string
        required: false
      supported _virtualization_enviroment:
        description: >-
          Identifies the virtualization environments (e.g. hypervisor) compatible with this software
          image.
        type: list
        entry_schema:
          type: string
        required: false
`
}