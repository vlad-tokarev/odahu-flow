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
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	"github.com/odahu/odahu-flow/packages/operator/pkg/utils"
	"github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes"
	v1Routes "github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes/v1"
	"github.com/rakyll/statik/fs"
	"net/http"
)

const (
	URLPrefix = "/api/v1"
)

type Server struct {
	manager utils.ManagerCloser
	server  *http.Server
}

func NewAPIServer(config *config.Config) (*Server, error) {
	staticFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	server := gin.Default()

	mgr, err := utils.NewManager(config.API.Backend)
	if err != nil {
		return nil, err
	}

	routes.SetUpHealthCheck(server)
	routes.SetUpSwagger(server.Group(""), staticFS)
	routes.SetUpPrometheus(server)

	v1Group := server.Group(URLPrefix)
	err = v1Routes.SetupV1Routes(v1Group, mgr.GetClient(), mgr.GetConfig(), *config)

	return &Server{manager: mgr, server: &http.Server{
		Addr:    fmt.Sprintf(":%d", config.API.Port),
		Handler: server,
	}}, err
}

// @title API Gateway
// @version 1.0
// @description This is an API Gateway server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully
func (s *Server) Close(ctx context.Context) error {
	if err := s.manager.Close(); err != nil {
		return err
	}

	return s.server.Shutdown(ctx)
}
