package prototype

import "testing"

func TestDuplicate(t *testing.T) {
  var expected_id Id = 12
  client := new(Client)
  proto := new(Prototype)
  proto.SetId(expected_id)
  client.Init(proto)
  cloned := client.Duplicate()

  if cloned == nil {
    t.Error("cloned object is nil")
  }
  if cloned.Id() != expected_id {
    t.Error("ids are not the same!")
  }
}
