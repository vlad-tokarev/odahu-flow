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

package controller

import (
	istioschema "github.com/aspenmesh/istio-client-go/pkg/client/clientset/versioned/scheme"
	knservingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	tektonschema "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("controller-setup")

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager, odahuConfig *config.Config) error {
	if odahuConfig.Training.Enabled {
		log.Info("Setting up the training controller")

		if err := tektonschema.AddToScheme(m.GetScheme()); err != nil {
			log.Error(err, "unable to add tekton APIs to scheme")

			return err
		}

		if err := AddTrainingToManager(m, odahuConfig.Training, odahuConfig.Operator, odahuConfig.Common,
			odahuConfig.Common.ResourceGPUName); err != nil {
			return err
		}
	}

	if odahuConfig.Packaging.Enabled {
		log.Info("Setting up the packaging controller")

		if err := tektonschema.AddToScheme(m.GetScheme()); err != nil {
			log.Error(err, "unable to add tekton APIs to scheme")

			return err
		}

		if err := AddPackagingToManager(m, odahuConfig.Packaging, odahuConfig.Operator, odahuConfig.Common,
			odahuConfig.Common.ResourceGPUName); err != nil {
			return err
		}
	}

	if odahuConfig.Deployment.Enabled {
		log.Info("Setting up the deployment controller")

		log.Info("Setting up Istio scheme")
		istioschema.AddToScheme(m.GetScheme())

		log.Info("Setting up Knative scheme")
		if err := knservingv1alpha1.AddToScheme(m.GetScheme()); err != nil {
			log.Error(err, "unable add Knative APIs to scheme")

			return err
		}

		if err := AddDeploymentToManager(
			m, odahuConfig.Deployment, odahuConfig.Operator, odahuConfig.Common.ResourceGPUName,
		); err != nil {
			return err
		}
	}

	return nil
}
