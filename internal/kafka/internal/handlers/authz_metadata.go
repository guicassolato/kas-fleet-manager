package handlers

import (
	"net/http"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/services/account"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/presenters"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/services"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/handlers"
	"github.com/gorilla/mux"
)

type authzMetadataHandler struct {
	kafkaService   services.KafkaService
	accountService account.AccountService
	clusterService services.ClusterService
}

func NewAuthzMetadataHandler(kafkaService services.KafkaService, accountService account.AccountService, clusterService services.ClusterService) *authzMetadataHandler {
	return &authzMetadataHandler{
		kafkaService:   kafkaService,
		accountService: accountService,
		clusterService: clusterService,
	}
}

func (h authzMetadataHandler) GetKafkaRequestData(w http.ResponseWriter, r *http.Request) {
	cfg := &handlers.HandlerConfig{
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			kafkaRequest, err := h.kafkaService.GetById(id)
			if err != nil {
				return nil, err
			}
			if kafkaRequest == nil {
				return nil, errors.NotFound("Unable to find kafka with id '%s'", id)
			}
			return presenters.PresentKafkaRequestAuthzMetadata(kafkaRequest, h.accountService)
		},
	}
	handlers.HandleGet(w, r, cfg)
}

func (h *authzMetadataHandler) GetDataPlaneClusterData(w http.ResponseWriter, r *http.Request) {
	cfg := &handlers.HandlerConfig{
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			cluster, err := h.clusterService.FindClusterByID(id)
			if err != nil {
				return nil, err
			}
			if cluster == nil {
				return nil, errors.NotFound("Unable to find cluster with id '%s'", id)
			}
			return presenters.PresentDataPlaneClusterAuthzMetadata(cluster)
		},
	}
	handlers.HandleGet(w, r, cfg)
}
