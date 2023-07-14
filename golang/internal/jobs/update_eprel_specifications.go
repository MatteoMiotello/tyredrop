package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/null/v8"
	"pillowww/titw/internal/db"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
	"pillowww/titw/pkg/clients/eprel"
	"pillowww/titw/pkg/constants"
	"pillowww/titw/pkg/log"
	"strconv"
	"time"
)

type UpdateTyresSpecificationJob struct {
	MaxChildren int
}

func (r UpdateTyresSpecificationJob) Run() {
	UpdateTyresSpecifications()
}

func UpdateTyresSpecifications() {
	ctx := context.Background()

	err := db.WithTx(ctx, func(tx *sql.Tx) error {
		pDao := product.NewDao(tx)
		p, err := pDao.
			ForUpdate().
			FindNextRemainingEprelProduct(ctx, string(constants.PRODUCT_CATEGORY_TYRE))

		if err != nil {
			fmt.Println(err)
			return nil
		}

		vDao := product.NewSpecificationValueDao(tx)

		p.EprelUpdatedAt = null.TimeFrom(time.Now())
		err = vDao.Update(ctx, p)
		if err != nil {
			return err
		}

		value := p.EprelProductCode.String
		if len(value) == 0 {
			return nil
		}

		eprelSpec, err := eprel.GetEprelSpecifications(value)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		sDao := product.NewSpecificationDao(tx)
		err = createSpecificationValues(ctx, sDao, vDao, p, eprelSpec)
		if err != nil {
			return err
		}

		p.EprelUpdatedAt = null.TimeFrom(time.Now())
		err = pDao.Update(ctx, p)
		if err != nil {
			return err
		}
		return nil
	})

	ctx.Done()

	if err != nil {
		log.Warn("Unable to update eprel spec", err)
		return
	}
}

func createSpecificationValues(ctx context.Context, psDao *product.SpecificationDao, vDao *product.SpecificationValueDao, product *models.Product, response *eprel.TyreResponse) error {
	specs := map[constants.ProductSpecification]string{
		constants.TYRE_SPEC_FUEL_EFFICIENCY:              response.EnergyClass,
		constants.TYRE_SPEC_WET_GRIP_CLASS:               response.WetGripClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_CLASS: response.ExternalRollingNoiseClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_LEVEL: strconv.Itoa(response.ExternalRollingNoiseValue),
		constants.TYRE_SPEC_LOAD_VERSION:                 response.LoadCapacityIndicator,
	}

	for key, value := range specs {
		specification, err := psDao.FindOneByCode(ctx, string(key))
		if err != nil {
			return err
		}

		v, _ := vDao.FindByProductAndCode(ctx, product, string(key))

		if v != nil {
			continue
		}

		specificationValue, _ := vDao.FindBySpecificationAndValue(ctx, specification, value)

		if specificationValue == nil {
			specificationValue = &models.ProductSpecificationValue{
				ProductSpecificationID: specification.ID,
				SpecificationValue:     value,
			}

			err := vDao.Insert(ctx, specificationValue)

			if err != nil {
				return err
			}
		}

		relation := &models.ProductProductSpecificationValue{
			ProductSpecificationValueID: specificationValue.ID,
			ProductID:                   product.ID,
		}

		err = vDao.Insert(ctx, relation)

		if err != nil {
			log.Error("Error inserting ProductProductSpecificationValue", err)
			return err
		}
	}

	return nil
}
