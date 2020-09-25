package application

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"prophetstor.com/alameda/datahub/pkg/entities"
	autoscalingv1alpha2 "prophetstor.com/alameda/operator/apis/autoscaling/v1alpha2"
	tenantv1alpha1 "prophetstor.com/alameda/operator/apis/tenant/v1alpha1"
	datahubscaler "prophetstor.com/alameda/operator/datahub/api/scaler"
	datahubtenant "prophetstor.com/alameda/operator/datahub/api/tenant"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func RemoveOutOfDate(client client.Client, datahubClient *datahubpkg.Client, enabledDA bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// sync alamedascaler
	scalersList := autoscalingv1alpha2.AlamedaScalerList{}
	if err := client.List(ctx, &scalersList); err != nil {
		return errors.Errorf(
			"Sync applications with datahub failed due to list applications from cluster failed: %s", err.Error())
	}

	resourceApps := []entities.ResourceClusterStatusApplication{}
	if err := datahubClient.List(&resourceApps); err != nil {
		return err
	}

	oodAppList := []entities.ResourceClusterStatusApplication{}
	for _, resourceApp := range resourceApps {
		isOOD := true
		for _, scaler := range scalersList.Items {
			if scaler.Spec.ClusterName == resourceApp.ClusterName &&
				scaler.GetNamespace() == resourceApp.Namespace &&
				scaler.GetName() == resourceApp.Name {
				isOOD = false
				break
			}
		}
		if isOOD {
			oodAppList = append(oodAppList, resourceApp)
		}
	}

	if len(oodAppList) > 0 {
		for _, oodApp := range oodAppList {
			if delErr := datahubscaler.DeleteV1Alpha2Scaler(datahubClient, client,
				oodApp.Namespace, oodApp.Name, enabledDA); delErr != nil {
				return delErr
			}
		}
	}

	// sync alamedaorganization
	orgList := tenantv1alpha1.AlamedaOrganizationList{}
	if err := client.List(ctx, &orgList); err != nil {
		return errors.Errorf(
			"Sync AlamedaOrganization with datahub failed due to list organization from cluster failed: %s", err.Error())
	}
	tenantOrgs := []entities.ConfigTenancyOrganization{}
	if err := datahubClient.List(&tenantOrgs); err != nil {
		return err
	}
	oodTenantOrgList := []entities.ConfigTenancyOrganization{}
	for _, tenantOrg := range tenantOrgs {
		isOOD := true
		for _, org := range orgList.Items {
			if org.GetName() == tenantOrg.Name {
				isOOD = false
				break
			}
		}
		if isOOD {
			oodTenantOrgList = append(oodTenantOrgList, tenantOrg)
		}
	}
	if len(oodTenantOrgList) > 0 {
		for _, oodTenantOrg := range oodTenantOrgList {
			if delErr := datahubtenant.DeleteOrganization(
				datahubClient, oodTenantOrg.Name); delErr != nil {
				return delErr
			}
		}
	}
	return nil
}
