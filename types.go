package cnregion

import (
	_ "embed"

	"github.com/goravel/framework/support/json"
)

//go:embed data.json
var data string

type Region struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Children []Region `json:"children"`
}

var regions []Region

func init() {
	if len(regions) > 0 {
		return
	}
	if err := json.UnmarshalString(data, &regions); err != nil {
		panic(err)
	}
}
