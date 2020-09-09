package tenant

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
)

func DeleteOrganization(
	datahubClient *datahubpkg.Client, orgName string) error {
	err := datahubClient.DeleteByOpts(&entities.ConfigTenancyOrganization{}, datahubpkg.Option{
		Entity: entities.ConfigTenancyOrganization{
			Name: orgName,
		},
		Fields: []string{"Name"},
	})
	if err != nil {
		return err
	}
	err = datahubClient.DeleteByOpts(&entities.ConfigTenancyCluster{}, datahubpkg.Option{
		Entity: entities.ConfigTenancyCluster{
			Organization: orgName,
		},
		Fields: []string{"Organization"},
	})
	return nil
}
