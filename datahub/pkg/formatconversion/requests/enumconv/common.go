package enumconv

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var ResourceBoundaryNameMap = map[ApiCommon.ResourceBoundary]schemas.ResourceBoundary{
	ApiCommon.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED: schemas.ResourceBoundaryUndefined,
	ApiCommon.ResourceBoundary_RESOURCE_RAW:                schemas.ResourceRaw,
	ApiCommon.ResourceBoundary_RESOURCE_UPPER_BOUND:        schemas.ResourceUpperBound,
	ApiCommon.ResourceBoundary_RESOURCE_LOWER_BOUND:        schemas.ResourceLowerBound,
}

var ResourceQuotaNameMap = map[ApiCommon.ResourceQuota]schemas.ResourceQuota{
	ApiCommon.ResourceQuota_RESOURCE_QUOTA_UNDEFINED: schemas.ResourceQuotaUndefined,
	ApiCommon.ResourceQuota_RESOURCE_LIMIT:           schemas.ResourceLimit,
	ApiCommon.ResourceQuota_RESOURCE_REQUEST:         schemas.ResourceRequest,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_LIMIT:   schemas.ResourceInitialLimit,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_REQUEST: schemas.ResourceInitialRequest,
}
