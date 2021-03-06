/*
 * Copyright 2019 EPAM Systems
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package servicecatalog

import (
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	"github.com/odahu/odahu-flow/packages/operator/pkg/servicecatalog/catalog"
	"github.com/rakyll/statik/fs"
)

func SetUPMainServer(
	mrc *catalog.ModelRouteCatalog,
	config config.ServiceCatalog,
) (*gin.Engine, error) {
	staticFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	server := gin.Default()
	rootRouteGroup := server.Group(config.BaseURL)

	SetUpSwagger(rootRouteGroup, staticFS, mrc.ProcessSwaggerJSON)
	SetUpHealthCheck(server)

	return server, nil
}
