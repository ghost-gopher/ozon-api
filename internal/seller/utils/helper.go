package utils

import (
	"fmt"
)

const (
	EndpointPattern = "%s/%s/%s"

	APIVersionTwo = "v2"
)

func PathEndpointVersionTwo(entity string, path string) string {
	return fmt.Sprintf(EndpointPattern, APIVersionTwo, entity, path)
}
