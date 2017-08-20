package prototype

type Id uint16

type clonable interface {
  clone() clonable
  Id() Id
}

type Prototype struct{
  id Id
}

func (p *Prototype) clone() clonable {
  return &Prototype{id: p.id}
}

func (p *Prototype) SetId(id Id) {
  p.id = id
}

func (p *Prototype) Id() Id {
  return p.id
}

type Client struct {
  prototype clonable
}

func (c *Client) Init(clone clonable) {
  c.prototype = clone
}

func (c *Client) Duplicate() clonable {
  return c.prototype.clone()
}
