package gremlin

import (
	"fmt"
)

type Property struct {
	Id   string
	Data interface{}
}

func (p *Property) String() (string, bool) {
	val, ok := p.Data.(string)
	return val, ok
}

func (p *Property) AssertString() string {
	val, ok := p.String()
	if ok {
		return val
	} else {
		panic(fmt.Sprintf("Expected a string, but got %#v", p.Data))
	}
}

func (p *Property) Float() (float64, bool) {
	val, ok := p.Data.(float64)
	return val, ok
}

func (p *Property) AssertFloat() float64 {
	val, ok := p.Float()
	if ok {
		return val
	} else {
		panic(fmt.Sprintf("Expected a float, but got %#v", p.Data))
	}
}

func (p *Property) Int() (int, bool) {
	val, ok := p.Data.(float64)
	return int(val), ok
}

func (p *Property) AssertInt() int {
	val, ok := p.Int()
	if ok {
		return val
	} else {
		panic(fmt.Sprintf("Expected a int, but got %#v", p.Data))
	}
}

func (p *Property) Int64() (int64, bool) {
	val, ok := p.Data.(float64)
	return int64(val), ok
}

func (p *Property) AssertInt64() int64 {
	val, ok := p.Int64()
	if ok {
		return val
	} else {
		panic(fmt.Sprintf("Expected a int, but got %#v", p.Data))
	}
}

func (p *Property) Bool() (bool, bool) {
	val, ok := p.Data.(bool)
	return val, ok
}

func (p *Property) AssertBool() bool {
	val, ok := p.Bool()
	if ok {
		return val
	} else {
		panic(fmt.Sprintf("Expected a bool, but got %#v", p.Data))
	}
}

func extractProperties(props map[string]interface{}) (map[string]Property, error) {
	properties := make(map[string]Property)
	for key, prop_val := range props {
		prop_val_maps, ok := prop_val.([]interface{})

		if !ok {
			return nil, fmt.Errorf("Property is not a list %#v", prop_val)
		}

		if len(prop_val_maps) != 1 {
			//TODO fix; read up on where/how this is possible.
			panic(fmt.Sprintf("should be exactly 1, but got %#v", prop_val_maps))
		}

		prop_val_map, ok := prop_val_maps[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Property value map is not an object %#v", prop_val)
		}

		prop_val, ok := prop_val_map["value"]
		if !ok {
			return nil, fmt.Errorf("no 'value' in property object")
		}

		prop_id_interface, ok := prop_val_map["id"]
		if !ok {
			return nil, fmt.Errorf("no 'id' in property object")
		}

		prop_id, ok := prop_id_interface.(string)
		if !ok {
			return nil, fmt.Errorf("'id' in property object is not a string")
		}

		property := Property{
			Id:   prop_id,
			Data: prop_val,
		}

		properties[key] = property
	}

	return properties, nil
}
