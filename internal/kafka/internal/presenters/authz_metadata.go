package presenters

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/admin/private"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/services/account"
)

func PresentKafkaRequestAuthzMetadata(kafkaRequest *dbapi.KafkaRequest, accountService account.AccountService) (*private.Kafka, *errors.ServiceError) {
	reference := PresentReference(kafkaRequest.ID, kafkaRequest)

	var accountNumber string
	if org, _ := accountService.GetOrganization(fmt.Sprintf("external_id='%s'", kafkaRequest.OrganisationId)); org != nil {
		accountNumber = org.AccountNumber
	}

	// convert kafka storage size to bytes
	maxDataRetentionSizeQuantity := config.Quantity(kafkaRequest.KafkaStorageSize)
	maxDataRetentionSizeBytes, conversionErr := maxDataRetentionSizeQuantity.ToInt64()
	if conversionErr != nil {
		return nil, errors.NewWithCause(errors.ErrorGeneral, conversionErr, "failed to get bytes value for max_data_retention_size")
	}

	return &private.Kafka{
		Id:                         reference.Id,
		Kind:                       reference.Kind,
		Href:                       reference.Href,
		Status:                     kafkaRequest.Status,
		CloudProvider:              kafkaRequest.CloudProvider,
		MultiAz:                    kafkaRequest.MultiAZ,
		Region:                     kafkaRequest.Region,
		Owner:                      kafkaRequest.Owner,
		Name:                       kafkaRequest.Name,
		BootstrapServerHost:        setBootstrapServerHost(kafkaRequest.BootstrapServerHost),
		CreatedAt:                  kafkaRequest.CreatedAt,
		UpdatedAt:                  kafkaRequest.UpdatedAt,
		FailedReason:               kafkaRequest.FailedReason,
		DesiredKafkaVersion:        kafkaRequest.DesiredKafkaVersion,
		ActualKafkaVersion:         kafkaRequest.ActualKafkaVersion,
		DesiredStrimziVersion:      kafkaRequest.DesiredStrimziVersion,
		ActualStrimziVersion:       kafkaRequest.ActualStrimziVersion,
		DesiredKafkaIbpVersion:     kafkaRequest.DesiredKafkaIBPVersion,
		ActualKafkaIbpVersion:      kafkaRequest.ActualKafkaIBPVersion,
		KafkaUpgrading:             kafkaRequest.KafkaUpgrading,
		StrimziUpgrading:           kafkaRequest.StrimziUpgrading,
		KafkaIbpUpgrading:          kafkaRequest.KafkaIBPUpgrading,
		DeprecatedKafkaStorageSize: kafkaRequest.KafkaStorageSize,
		OrganisationId:             kafkaRequest.OrganisationId,
		SubscriptionId:             kafkaRequest.SubscriptionId,
		OwnerAccountId:             kafkaRequest.OwnerAccountId,
		AccountNumber:              accountNumber,
		QuotaType:                  kafkaRequest.QuotaType,
		Routes:                     GetRoutesFromKafkaRequest(kafkaRequest),
		RoutesCreated:              kafkaRequest.RoutesCreated,
		ClusterId:                  kafkaRequest.ClusterID,
		InstanceType:               kafkaRequest.InstanceType,
		Namespace:                  kafkaRequest.Namespace,
		SizeId:                     kafkaRequest.SizeId,
		MaxDataRetentionSize: private.SupportedKafkaSizeBytesValueItem{
			Bytes: maxDataRetentionSizeBytes,
		},
	}, nil
}

func PresentDataPlaneClusterAuthzMetadata(cluster *api.Cluster) (*api.Cluster, *errors.ServiceError) {
	// TODO: generate a proper (private) model using OpenAPI Generator (https://openapi-generator.tech)
	return &api.Cluster{
		CloudProvider:            cluster.CloudProvider,
		ClusterID:                cluster.ClusterID,
		ExternalID:               cluster.ExternalID,
		MultiAZ:                  cluster.MultiAZ,
		Region:                   cluster.Region,
		Status:                   cluster.Status,
		StatusDetails:            cluster.StatusDetails,
		IdentityProviderID:       cluster.IdentityProviderID,
		ClusterDNS:               cluster.ClusterDNS,
		ClientID:                 cluster.ClientID,
		ProviderType:             cluster.ProviderType,
		ProviderSpec:             cluster.ProviderSpec,
		ClusterSpec:              cluster.ClusterSpec,
		AvailableStrimziVersions: cluster.AvailableStrimziVersions,
		SupportedInstanceType:    cluster.SupportedInstanceType,
		DynamicCapacityInfo:      cluster.DynamicCapacityInfo,
	},
	nil
}
