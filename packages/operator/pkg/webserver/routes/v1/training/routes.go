//
//    Copyright 2019 EPAM Systems
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package training

import (
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	conn_repository "github.com/odahu/odahu-flow/packages/operator/pkg/repository/connection"
	mt_repository "github.com/odahu/odahu-flow/packages/operator/pkg/repository/training"
)

func ConfigureRoutes(routeGroup *gin.RouterGroup, mtRepository mt_repository.Repository,
	tiRepository mt_repository.ToolchainRepository, connRepository conn_repository.Repository,
	config config.ModelTrainingConfig, gpuResourceName string) {

	mtController := ModelTrainingController{
		mtRepository: mtRepository,
		validator: NewMtValidator(
			tiRepository,
			connRepository,
			config.DefaultResources,
			config.OutputConnectionID,
			gpuResourceName,
		),
	}

	routeGroup.GET(GetModelTrainingURL, mtController.getMT)
	routeGroup.GET(GetAllModelTrainingURL, mtController.getAllMTs)
	routeGroup.GET(GetModelTrainingLogsURL, mtController.getModelTrainingLog)
	routeGroup.POST(CreateModelTrainingURL, mtController.createMT)
	routeGroup.PUT(UpdateModelTrainingURL, mtController.updateMT)
	routeGroup.PUT(SaveModelTrainingResultURL, mtController.saveMPResults)
	routeGroup.DELETE(DeleteModelTrainingURL, mtController.deleteMT)

}
