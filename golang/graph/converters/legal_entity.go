package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func LegalEntityTypeToGraphQL(entityType models.LegalEntityType) *model.LegalEntityType {
	return &model.LegalEntityType{
		ID:   entityType.ID,
		Name: entityType.Name,
	}
}
