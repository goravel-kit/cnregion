package facades

import (
	"log"

	"github.com/goravel-kit/cnregion"
	"github.com/goravel-kit/cnregion/contracts"
)

func CnRegion() contracts.CnRegion {
	instance, err := cnregion.App.Make(cnregion.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.CnRegion)
}
