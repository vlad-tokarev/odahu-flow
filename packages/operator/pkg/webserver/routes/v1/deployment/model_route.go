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

package deployment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/deployment"
	mr_repository "github.com/odahu/odahu-flow/packages/operator/pkg/repository/deployment"
	"github.com/odahu/odahu-flow/packages/operator/pkg/repository/util/kubernetes"
	"github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes"

	"net/http"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var logMR = logf.Log.WithName("mr-controller")

const (
	GetModelRouteURL    = "/model/route/:id"
	GetAllModelRouteURL = "/model/route"
	CreateModelRouteURL = "/model/route"
	UpdateModelRouteURL = "/model/route"
	DeleteModelRouteURL = "/model/route/:id"
	IDMrURLParam        = "id"
)

var (
	emptyCache = map[string]int{}
)

type ModelRouteController struct {
	mrRepository mr_repository.Repository
	validator    *MrValidator
}

// @Summary Get a Model route
// @Description Get a Model route by id
// @Tags Route
// @Name id
// @Accept  json
// @Produce  json
// @Param id path string true "Model route id"
// @Success 200 {object} deployment.ModelRoute
// @Failure 404 {object} routes.HTTPResult
// @Failure 400 {object} routes.HTTPResult
// @Router /api/v1/model/route/{id} [get]
func (mrc *ModelRouteController) getMR(c *gin.Context) {
	mrID := c.Param(IDMrURLParam)

	mr, err := mrc.mrRepository.GetModelRoute(mrID)
	if err != nil {
		logMR.Error(err, fmt.Sprintf("Retrieving %s model route", mrID))
		c.AbortWithStatusJSON(routes.CalculateHTTPStatusCode(err), routes.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, mr)
}

// @Summary Get list of Model routes
// @Description Get list of Model routes
// @Tags Route
// @Accept  json
// @Produce  json
// @Param size path int false "Number of entities in a response"
// @Param page path int false "Number of a page"
// @Success 200 {array} deployment.ModelRoute
// @Failure 400 {object} routes.HTTPResult
// @Router /api/v1/model/route [get]
func (mrc *ModelRouteController) getAllMRs(c *gin.Context) {
	size, page, err := routes.URLParamsToFilter(c, nil, emptyCache)
	if err != nil {
		logMR.Error(err, "Malformed url parameters of model route request")
		c.AbortWithStatusJSON(http.StatusBadRequest, routes.HTTPResult{Message: err.Error()})

		return
	}

	mrList, err := mrc.mrRepository.GetModelRouteList(
		kubernetes.Size(size),
		kubernetes.Page(page),
	)
	if err != nil {
		logMR.Error(err, "Retrieving list of model routes")
		c.AbortWithStatusJSON(routes.CalculateHTTPStatusCode(err), routes.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, mrList)
}

// @Summary Create a Model route
// @Description Create a Model route. Results is created Model route.
// @Param mr body deployment.ModelRoute true "Create a Model route"
// @Tags Route
// @Accept  json
// @Produce  json
// @Success 201 {object} deployment.ModelRoute
// @Failure 400 {object} routes.HTTPResult
// @Router /api/v1/model/route [post]
func (mrc *ModelRouteController) createMR(c *gin.Context) {
	var mr deployment.ModelRoute

	if err := c.ShouldBindJSON(&mr); err != nil {
		logMR.Error(err, "JSON binding of the model route is failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, routes.HTTPResult{Message: err.Error()})

		return
	}

	if err := mrc.validator.ValidatesAndSetDefaults(&mr); err != nil {
		logMR.Error(err, fmt.Sprintf("Validation of the model route is failed: %v", mr))
		c.AbortWithStatusJSON(http.StatusBadRequest, routes.HTTPResult{Message: err.Error()})

		return
	}

	if err := mrc.mrRepository.CreateModelRoute(&mr); err != nil {
		logMR.Error(err, fmt.Sprintf("Creation of the model route: %+v", mr))
		c.AbortWithStatusJSON(routes.CalculateHTTPStatusCode(err), routes.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusCreated, mr)
}

// @Summary Update a Model route
// @Description Update a Model route. Results is updated Model route.
// @Param mr body deployment.ModelRoute true "Update a Model route"
// @Tags Route
// @Accept  json
// @Produce  json
// @Success 200 {object} deployment.ModelRoute
// @Failure 404 {object} routes.HTTPResult
// @Failure 400 {object} routes.HTTPResult
// @Router /api/v1/model/route [put]
func (mrc *ModelRouteController) updateMR(c *gin.Context) {
	var mr deployment.ModelRoute

	if err := c.ShouldBindJSON(&mr); err != nil {
		logMR.Error(err, "JSON binding of the model route is failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, routes.HTTPResult{Message: err.Error()})

		return
	}

	if err := mrc.validator.ValidatesAndSetDefaults(&mr); err != nil {
		logMR.Error(err, fmt.Sprintf("Validation of the model route is failed: %v", mr))
		c.AbortWithStatusJSON(http.StatusBadRequest, routes.HTTPResult{Message: err.Error()})

		return
	}

	if err := mrc.mrRepository.UpdateModelRoute(&mr); err != nil {
		logMR.Error(err, fmt.Sprintf("Update of the model route: %+v", mr))
		c.AbortWithStatusJSON(routes.CalculateHTTPStatusCode(err), routes.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, mr)
}

// @Summary Delete a Model route
// @Description Delete a Model route by id
// @Tags Route
// @Name id
// @Accept  json
// @Produce  json
// @Param id path string true "Model route id"
// @Success 200 {object} routes.HTTPResult
// @Failure 404 {object} routes.HTTPResult
// @Failure 400 {object} routes.HTTPResult
// @Router /api/v1/model/route/{id} [delete]
func (mrc *ModelRouteController) deleteMR(c *gin.Context) {
	mrID := c.Param(IDMrURLParam)

	if err := mrc.mrRepository.DeleteModelRoute(mrID); err != nil {
		logMR.Error(err, fmt.Sprintf("Deletion of %s model route is failed", mrID))
		c.AbortWithStatusJSON(routes.CalculateHTTPStatusCode(err), routes.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, routes.HTTPResult{Message: fmt.Sprintf("Model route %s was deleted", mrID)})
}
