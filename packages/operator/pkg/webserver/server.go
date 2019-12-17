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

package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/utils"
	"github.com/odahu/odahu-flow/packages/operator/pkg/validation"
	"github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes"
	v1Routes "github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes/v1"
	"github.com/rakyll/statik/fs"
)

// @title API Gateway
// @version 1.0
// @description This is an API Gateway server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func SetUPMainServer() (*gin.Engine, error) {
	staticFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	server := gin.Default()

	mgr, err := utils.NewManager()
	if err != nil {
		return nil, err
	}

	validation.InitValidator()
	routes.SetUpHealthCheck(server)
	routes.SetUpSwagger(server.Group(""), staticFS)
	routes.SetUpPrometheus(server)
	err = routes.SetUpIndexPage(server.Group(""), staticFS)
	if err != nil {
		return nil, err
	}

	v1Group := server.Group("/api/v1")
	err = v1Routes.SetupV1Routes(v1Group, mgr.GetClient(), mgr.GetConfig())

	if err != nil {
		return nil, err
	}

	return server, nil
}
