// Package metadata contains components for accessing Kyma Application
package metadata

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/kyma-project/kyma/components/central-application-gateway/internal/metadata/applications"
	"github.com/kyma-project/kyma/components/central-application-gateway/internal/metadata/model"
	"github.com/kyma-project/kyma/components/central-application-gateway/internal/metadata/serviceapi"
	"github.com/kyma-project/kyma/components/central-application-gateway/pkg/apperrors"
)

// ServiceDefinitionService is a service that manages ServiceDefinition objects.
//
//go:generate mockery -name=ServiceDefinitionService
type ServiceDefinitionService interface {
	// GetAPI gets API of a service with given ID
	GetAPIByServiceName(appName, serviceName string) (*model.API, apperrors.AppError)
	GetAPIByEntryName(appName, serviceName, entryName string) (*model.API, apperrors.AppError)
}

type serviceDefinitionService struct {
	serviceAPIService     serviceapi.Service
	applicationRepository applications.ServiceRepository
}

// NewServiceDefinitionService creates new ServiceDefinitionService with provided dependencies.
func NewServiceDefinitionService(serviceAPIService serviceapi.Service, applicationRepository applications.ServiceRepository) ServiceDefinitionService {
	return &serviceDefinitionService{
		serviceAPIService:     serviceAPIService,
		applicationRepository: applicationRepository,
	}
}

// GetAPI gets API of a service with given name
func (sds *serviceDefinitionService) GetAPIByServiceName(appName, serviceName string) (*model.API, apperrors.AppError) {
	service, err := sds.applicationRepository.GetByServiceName(appName, serviceName)

	if err != nil {
		notFoundMessage := fmt.Sprintf("service with name %s not found", serviceName)
		internalErrMessage := fmt.Sprintf("failed to get service with name '%s': %s", serviceName, err.Error())

		return nil, handleError(err, notFoundMessage, internalErrMessage)
	}

	return sds.getAPI(service)
}

func (sds *serviceDefinitionService) GetAPIByEntryName(appName, serviceName, entryName string) (*model.API, apperrors.AppError) {
	service, err := sds.applicationRepository.GetByEntryName(appName, serviceName, entryName)

	if err != nil {
		notFoundMessage := fmt.Sprintf("service with name %s and entry name %s not found", serviceName, entryName)
		internalErrMessage := fmt.Sprintf("failed to get service with name '%s' and entry name '%s': %s", serviceName, entryName, err.Error())

		return nil, handleError(err, notFoundMessage, internalErrMessage)
	}

	return sds.getAPI(service)
}

func (sds *serviceDefinitionService) getAPI(service applications.Service) (*model.API, apperrors.AppError) {

	if service.API == nil {
		return nil, apperrors.WrongInputf("service '%s' has no API", service.Name)
	}

	api, err := sds.serviceAPIService.Read(service.API)
	if err != nil {
		zap.L().Error("failed to read api for serviceID",
			zap.String("serviceID", service.Name),
			zap.Error(err))
		return nil, apperrors.Internalf("failed to read API for %s service, %s", service.Name, err)
	}
	return api, nil
}

func handleError(err apperrors.AppError, notFoundMessage, internalErrorMEssage string) apperrors.AppError {
	if err.Code() == apperrors.CodeNotFound {
		return apperrors.NotFound(notFoundMessage)
	}
	zap.L().Error(internalErrorMEssage)

	if err.Code() == apperrors.CodeWrongInput {
		return apperrors.WrongInput(internalErrorMEssage)
	}
	return apperrors.Internal(internalErrorMEssage)
}
