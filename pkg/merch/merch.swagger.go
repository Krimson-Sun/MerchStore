package merch

import _ "embed"

//go:embed merch.swagger.json
var swaggerFile []byte

// New creates a new instance of the workouts service
func GetSwaggerDescription() []byte {
	return swaggerFile
}
