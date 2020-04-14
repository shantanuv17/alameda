package dispatcher

import (
	"fmt"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
)

func GetUnitScalerNSName(unit *config.Unit, rowValues []string, columns []string) (string, string, error) {
	errMsg := ""
	scaleNS, nsErr := utils.GetRowValue(rowValues,
		columns, unit.UnitValueKeys.ScaleNamespace)
	if nsErr != nil {
		errMsg = nsErr.Error()
	}
	scaleName, nameErr := utils.GetRowValue(rowValues,
		columns, unit.UnitValueKeys.ScaleName)
	if nameErr != nil {
		errMsg = fmt.Sprintf("%s.%s", errMsg, nameErr.Error())
	}

	if nsErr == nil && nameErr == nil {
		return scaleNS, scaleName, nil
	}
	return scaleNS, scaleName, fmt.Errorf(errMsg)
}

func GetUnitResourceK8SNSName(unit *config.Unit, rowValues []string, columns []string) (*string, *string, error) {
	errMsg := ""
	var resourceK8SNS *string = nil
	var resourceK8SName *string = nil
	var nsErr, nameErr error
	if unit.UnitValueKeys.ResourceK8SNamespace != nil {
		nsVal, nsErr := utils.GetRowValue(rowValues,
			columns, *unit.UnitValueKeys.ResourceK8SNamespace)
		if nsErr != nil {
			errMsg = nsErr.Error()
		} else {
			resourceK8SNS = &nsVal
		}
	}
	if unit.UnitValueKeys.ResourceK8SName != nil {
		nameVal, nameErr := utils.GetRowValue(rowValues,
			columns, *unit.UnitValueKeys.ResourceK8SName)
		if nameErr != nil {
			errMsg = fmt.Sprintf("%s.%s", errMsg, nameErr.Error())
		} else {
			resourceK8SName = &nameVal
		}
	}
	if nsErr == nil && nameErr == nil {
		return resourceK8SNS, resourceK8SName, nil
	}
	return resourceK8SNS, resourceK8SName, fmt.Errorf(errMsg)
}
