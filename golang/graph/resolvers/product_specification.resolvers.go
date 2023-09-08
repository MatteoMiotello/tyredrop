package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"pillowww/titw/graph"
	"pillowww/titw/graph/converters"
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

// ProductCategory is the resolver for the productCategory field.
func (r *productSpecificationResolver) ProductCategory(ctx context.Context, obj *model.ProductSpecification) (*model.ProductCategory, error) {
	c, err := r.ProductCategoryDao.
		Load(models.ProductCategoryRels.ProductCategoryLanguages).
		FindCategoryById(ctx, obj.ProductCategoryID)

	if err != nil {
		return nil, err
	}

	return converters.ProductCategoryToGraphQL(c), err
}

// Values is the resolver for the values field.
func (r *productSpecificationResolver) Values(ctx context.Context, obj *model.ProductSpecification) ([]*model.ProductSpecificationValue, error) {
	values, err := r.ProductSpecificationValueDao.FindBySpecificationId(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	var graphValues []*model.ProductSpecificationValue

	for _, value := range values {
		graphValues = append(graphValues, converters.ProductSpecificationValueToGraphQL(value))
	}

	return graphValues, nil
}

// Specification is the resolver for the specification field.
func (r *productSpecificationValueResolver) Specification(ctx context.Context, obj *model.ProductSpecificationValue) (*model.ProductSpecification, error) {
	spec, err := r.ProductSpecificationDao.
		Load(models.ProductSpecificationRels.ProductSpecificationLanguages).
		FindById(ctx, obj.SpecificationID)

	if err != nil {
		return nil, err
	}

	return converters.ProductSpecificationToGraphQL(spec), nil
}

// ProductSpecification returns graph.ProductSpecificationResolver implementation.
func (r *Resolver) ProductSpecification() graph.ProductSpecificationResolver {
	return &productSpecificationResolver{r}
}

// ProductSpecificationValue returns graph.ProductSpecificationValueResolver implementation.
func (r *Resolver) ProductSpecificationValue() graph.ProductSpecificationValueResolver {
	return &productSpecificationValueResolver{r}
}

type productSpecificationResolver struct{ *Resolver }
type productSpecificationValueResolver struct{ *Resolver }
