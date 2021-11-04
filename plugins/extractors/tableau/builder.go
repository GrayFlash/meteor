package tableau

import (
	"fmt"
	"time"

	"github.com/odpf/meteor/models/odpf/assets/common"
	"github.com/odpf/meteor/models/odpf/assets/facets"
	"github.com/pkg/errors"
)

func buildTimestamps(model BaseModel) (createdAt time.Time, updatedAt time.Time, err error) {
	createdAt, err = model.CreatedAt()
	if err != nil {
		err = errors.Wrap(err, "failed parsing createdAt")
		return
	}
	updatedAt, err = model.UpdatedAt()
	if err != nil {
		err = errors.Wrap(err, "failed parsing updatedAt")
		return
	}

	return
}

func buildLineage(wbConns []WorkbookConnection) (lineage *facets.Lineage) {
	upstreamLineages := []*common.Resource{}
	for _, wbc := range wbConns {
		upstreamLineages = append(upstreamLineages, buildLineageResource(wbc))
	}
	lineage = &facets.Lineage{Upstreams: upstreamLineages}
	return
}

func buildLineageResource(wbc WorkbookConnection) (resource *common.Resource) {
	// !TODO source name microsoft sql server -> sqlserver
	resource = &common.Resource{
		Urn:  fmt.Sprintf("%s::%s", wbc.Type, wbc.Datasource.ID),
		Name: wbc.Datasource.Name,
	}
	return
}
