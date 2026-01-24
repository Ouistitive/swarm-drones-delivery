package world

import (
	"encoding/json"
	"os"
)

type Address struct {
    Street string `json:"street"`
}

func ReadAddresses(addressesPath string) ([]Address, error) {
    plan, _ := os.ReadFile(addressesPath)
    var adresses []Address
    err := json.Unmarshal(plan, &adresses)

    if err != nil {
        return []Address{}, err
    }

    return adresses, nil
}