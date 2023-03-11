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
	"pillowww/titw/pkg/task"
	"strconv"
	"time"
)

type UpdateTyresSpecificationJob struct {
}

func (r UpdateTyresSpecificationJob) Run() {
	worker := task.NewWorker(5)
	worker.Run()

	worker.AddTask(UpdateTyresSpecifications)
	worker.AddTask(UpdateTyresSpecifications)
	worker.AddTask(UpdateTyresSpecifications)
	worker.AddTask(UpdateTyresSpecifications)
}

func UpdateTyresSpecifications() {
	ctx := context.Background()
	pDao := product.NewDao(db.DB)

	p, _ := pDao.FindNextRemainingEprelProduct(ctx, string(constants.PRODUCT_CATEGORY_TYRE))
	if p == nil {
		return
	}
	err := db.WithTx(ctx, func(tx *sql.Tx) error {
		pDao := product.NewDao(tx)

		p.EprelUpdatedAt = null.TimeFrom(time.Now())
		err := pDao.Update(ctx, p)
		if err != nil {
			return err
		}

		value, _ := pDao.FindProductSpecificationValue(ctx, p, string(constants.TYRE_SPEC_EPREL_ID))
		if value == nil {
			return nil
		}

		eprelSpec, err := eprel.GetEprelSpecifications(value.SpecificationValue)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = createSpecificationValues(ctx, pDao, p, eprelSpec)
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

	if err != nil {
		log.Warn("Unable to update eprel spec", err)
		return

	}
}

func createSpecificationValues(ctx context.Context, pDao *product.Dao, product *models.Product, response *eprel.TyreResponse) error {
	specs := map[constants.ProductSpecification]string{
		constants.TYRE_SPEC_FUEL_EFFICIENCY:              response.EnergyClass,
		constants.TYRE_SPEC_WET_GRIP_CLASS:               response.WetGripClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_CLASS: response.ExternalRollingNoiseClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_LEVEL: strconv.Itoa(response.ExternalRollingNoiseValue),
		constants.TYRE_SPEC_LOAD_VERSION:                 response.LoadCapacityIndicator,
	}

	for key, value := range specs {
		specification, err := pDao.FindOneProductSpecificationByCode(ctx, string(key))
		if err != nil {
			return err
		}

		value := &models.ProductSpecificationValue{
			SpecificationValue:     value,
			ProductID:              product.ID,
			ProductSpecificationID: specification.ID,
		}

		err = pDao.Upsert(ctx, value, false, []string{
			models.ProductSpecificationValueColumns.ProductID,
			models.ProductSpecificationValueColumns.ProductSpecificationID,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
