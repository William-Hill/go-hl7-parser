package hl7

import (
	"bytes"
	"fmt"
)

// Component is an HL7 component
type Component struct {
	SubComponents []SubComponent
	Value         []byte
}

func (c *Component) String() string {
	var str string
	for _, s := range c.SubComponents {
		str += "Component SubComponent: " + string(s.Value) + "\n"
	}
	return str
}

func (c *Component) Parse(seps *Delimeters) error {
	r := bytes.NewReader(c.Value)
	i := 0
	ii := 0
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case ch == eof || (ch == endMsg && seps.LFTermMsg):
			if ii > i {
				scmp := SubComponent{Value: c.Value[i : ii-1]}
				c.SubComponents = append(c.SubComponents, scmp)
			}
			return nil
		case ch == seps.SubComponent:
			scmp := SubComponent{Value: c.Value[i : ii-1]}
			c.SubComponents = append(c.SubComponents, scmp)
			i = ii
		case ch == seps.Escape:
			ii++
			r.ReadRune()
		}
	}
}

// SubComponent returns the subcomponent i (1-based index)
func (c *Component) SubComponent(i int) (*SubComponent, error) {
	// Convert 1-based index to 0-based for internal array access
	idx := i - 1
	if idx < 0 || idx >= len(c.SubComponents) {
		return nil, fmt.Errorf("SubComponent out of range")
	}
	return &c.SubComponents[idx], nil
}

func (c *Component) encode(seps *Delimeters) []byte {
	buf := [][]byte{}
	for _, sc := range c.SubComponents {
		buf = append(buf, sc.Value)
	}
	return bytes.Join(buf, []byte(string(seps.SubComponent)))
}

// Get returns the value specified by the Location
func (c *Component) Get(l *Location) (string, error) {
	if l.SubComp == -1 {
		return string(c.Value), nil
	}
	sc, err := c.SubComponent(l.SubComp)
	if err != nil {
		return "", err
	}
	return string(sc.Value), nil
}

// Set will insert a value into a message at Location
func (c *Component) Set(l *Location, val string, seps *Delimeters) error {
	// Convert 1-based index to 0-based
	subloc := l.SubComp
	if subloc < 0 {
		subloc = 0
	} else {
		subloc-- // Convert to 0-based
	}

	if x := subloc - len(c.SubComponents) + 1; x > 0 {
		c.SubComponents = append(c.SubComponents, make([]SubComponent, x)...)
	}
	c.SubComponents[subloc].Value = []byte(val)
	c.Value = c.encode(seps)
	return nil
}
