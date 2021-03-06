package compiler

import (
	"github.com/tliron/puccini/ard"
	"github.com/tliron/puccini/clout"
	"github.com/tliron/puccini/js"
	"github.com/tliron/puccini/tosca/problems"
)

type Transformer func(interface{}, interface{}, interface{}, interface{}, *clout.Clout) (interface{}, bool, error)

func Transform(transformer Transformer, c *clout.Clout, p *problems.Problems) {
	if tosca, ok := GetMap(c.Properties, "tosca"); ok {
		TransformValues(transformer, tosca, "inputs", nil, nil, nil, c, p)
		TransformValues(transformer, tosca, "outputs", nil, nil, nil, c, p)
	}

	for _, vertex := range c.Vertexes {
		if nodeTemplate, ok := GetToscaVertexProperties(vertex); ok {
			TransformValues(transformer, nodeTemplate, "properties", vertex, nil, nil, c, p)
			TransformValues(transformer, nodeTemplate, "attributes", vertex, nil, nil, c, p)
			TransformInterfaces(transformer, nodeTemplate, vertex, nil, nil, c, p)

			if capabilities, ok := GetMap(nodeTemplate, "capabilities"); ok {
				for _, value := range capabilities {
					if capability, ok := value.(ard.Map); ok {
						TransformValues(transformer, capability, "properties", vertex, nil, nil, c, p)
						TransformValues(transformer, capability, "attributes", vertex, nil, nil, c, p)
					}
				}
			}

			if artifacts, ok := GetMap(nodeTemplate, "artifacts"); ok {
				for _, value := range artifacts {
					if artifact, ok := value.(ard.Map); ok {
						TransformValues(transformer, artifact, "properties", vertex, nil, nil, c, p)
					}
				}
			}
		}

		for _, edge := range vertex.EdgesOut {
			if relationship, ok := GetToscaEdgeProperties(edge); ok {
				TransformValues(transformer, relationship, "properties", edge, vertex, edge.Target, c, p)
				TransformValues(transformer, relationship, "attributes", edge, vertex, edge.Target, c, p)
				TransformInterfaces(transformer, relationship, edge, vertex, edge.Target, c, p)
			}
		}
	}
}

func TransformValues(transformer Transformer, o ard.Map, fieldName string, site interface{}, source interface{}, target interface{}, c *clout.Clout, p *problems.Problems) {
	value, ok := o[fieldName]
	if !ok {
		return
	}

	map_, ok := value.(ard.Map)
	if !ok {
		return
	}

	for k, v := range map_ {
		var err error
		v, ok, err = transformer(v, site, source, target, c)
		if !ok {
			continue
		}
		if err != nil {
			if jsError, ok := err.(*js.Error); ok {
				p.Report(jsError.ColorError())
			} else {
				p.ReportError(err)
			}
		} else {
			map_[k] = v
		}
	}
}

func TransformInterfaces(transformer Transformer, o ard.Map, site interface{}, source interface{}, target interface{}, c *clout.Clout, p *problems.Problems) {
	if interfaces, ok := GetMap(o, "interfaces"); ok {
		for _, value := range interfaces {
			if intr, ok := value.(ard.Map); ok {
				TransformValues(transformer, intr, "inputs", site, source, target, c, p)
				if operations, ok := GetMap(intr, "operations"); ok {
					for _, value := range operations {
						if operation, ok := value.(ard.Map); ok {
							TransformValues(transformer, operation, "inputs", site, source, target, c, p)
						}
					}
				}
			}
		}
	}
}

func GetToscaVertexProperties(vertex *clout.Vertex) (ard.Map, bool) {
	if p, ok := vertex.Metadata["puccini-tosca"]; ok {
		if map_, ok := p.(ard.Map); ok {
			if v, ok := map_["version"]; ok {
				if version, ok := v.(string); ok {
					if version == "1.0" {
						return vertex.Properties, true
					}
				}
			}
		}
	}
	return nil, false
}

func GetToscaEdgeProperties(edge *clout.Edge) (ard.Map, bool) {
	if p, ok := edge.Metadata["puccini-tosca"]; ok {
		if map_, ok := p.(ard.Map); ok {
			if v, ok := map_["version"]; ok {
				if version, ok := v.(string); ok {
					if version == "1.0" {
						return edge.Properties, true
					}
				}
			}
		}
	}
	return nil, false
}

func GetMap(map_ ard.Map, key string) (ard.Map, bool) {
	v, ok := map_[key]
	if !ok {
		return nil, false
	}
	map_, ok = v.(ard.Map)
	return map_, ok
}
