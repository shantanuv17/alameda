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
