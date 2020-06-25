package go5e

import (
  "errors"
)
type StartingEquipment struct {
}

func GetStartingEquipment(index string) (StartingEquipment, error) {
  var ret StartingEquipment
  return ret, errors.New("Not implemented due to an API issue with indexing.")
}
