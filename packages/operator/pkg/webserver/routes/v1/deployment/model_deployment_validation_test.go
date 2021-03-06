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

package deployment_test

import (
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/deployment"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/odahuflow/v1alpha1"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	"github.com/odahu/odahu-flow/packages/operator/pkg/validation"
	md_routes "github.com/odahu/odahu-flow/packages/operator/pkg/webserver/routes/v1/deployment"
	"github.com/stretchr/testify/suite"
	"testing"

	. "github.com/onsi/gomega"
)

var (
	mdRoleName = "test-tole"
)

type ModelDeploymentValidationSuite struct {
	suite.Suite
	g                     *GomegaWithT
	defaultModelValidator *md_routes.ModelDeploymentValidator
}

func (s *ModelDeploymentValidationSuite) SetupTest() {
	s.g = NewGomegaWithT(s.T())
	s.defaultModelValidator = md_routes.NewModelDeploymentValidator(
		config.NewDefaultModelDeploymentConfig(),
		config.NvidiaResourceName,
	)
}

func TestModelDeploymentValidationSuite(t *testing.T) {
	suite.Run(t, new(ModelDeploymentValidationSuite))
}

func (s *ModelDeploymentValidationSuite) TestMDMinReplicasDefaultValue() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}
	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)

	s.g.Expect(*md.Spec.MinReplicas).To(Equal(md_routes.MdDefaultMinimumReplicas))
}

func (s *ModelDeploymentValidationSuite) TestMDMaxReplicasDefaultValue() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}
	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)

	s.g.Expect(*md.Spec.MaxReplicas).To(Equal(md_routes.MdDefaultMaximumReplicas))
}

func (s *ModelDeploymentValidationSuite) TestMDResourcesDefaultValues() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}
	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)

	s.g.Expect(*md.Spec.Resources).To(Equal(config.NewDefaultModelTrainingConfig().DefaultResources))
}

func (s *ModelDeploymentValidationSuite) TestMDReadinessProbeDefaultValue() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}
	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)

	s.g.Expect(*md.Spec.ReadinessProbeInitialDelay).To(Equal(md_routes.MdDefaultReadinessProbeInitialDelay))
}

func (s *ModelDeploymentValidationSuite) TestMDLivenessProbeDefaultValue() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}
	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)

	s.g.Expect(*md.Spec.LivenessProbeInitialDelay).To(Equal(md_routes.MdDefaultLivenessProbeInitialDelay))
}

func (s *ModelDeploymentValidationSuite) TestValidateNegativeMinReplicas() {
	minReplicas := int32(-1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			MinReplicas: &minReplicas,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.NegativeMinReplicasErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateNegativeMaxReplicas() {
	maxReplicas := int32(-1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			MaxReplicas: &maxReplicas,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.NegativeMaxReplicasErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateMaximumReplicas() {
	maxReplicas := int32(-1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			MaxReplicas: &maxReplicas,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.NegativeMaxReplicasErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateMinLessMaxReplicas() {
	minReplicas := int32(2)
	maxReplicas := int32(1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			MinReplicas: &minReplicas,
			MaxReplicas: &maxReplicas,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.MaxMoreThanMinReplicasErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateMinModelThanDefaultMax() {
	minReplicas := int32(3)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			MinReplicas: &minReplicas,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).ToNot(ContainSubstring(md_routes.MaxMoreThanMinReplicasErrorMessage))
	s.g.Expect(*md.Spec.MinReplicas).To(Equal(minReplicas))
	s.g.Expect(*md.Spec.MaxReplicas).To(Equal(minReplicas))
}

func (s *ModelDeploymentValidationSuite) TestValidateImage() {
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.EmptyImageErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateReadinessProbe() {
	readinessProbe := int32(-1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			ReadinessProbeInitialDelay: &readinessProbe,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.ReadinessProbeErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestValidateLivenessProbe() {
	livenessProbe := int32(-1)
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			LivenessProbeInitialDelay: &livenessProbe,
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).To(HaveOccurred())
	s.g.Expect(err.Error()).To(ContainSubstring(md_routes.LivenessProbeErrorMessage))
}

func (s *ModelDeploymentValidationSuite) TestMdResourcesValidation() {
	wrongResourceValue := "wrong res"
	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			Resources: &v1alpha1.ResourceRequirements{
				Limits: &v1alpha1.ResourceList{
					Memory: &wrongResourceValue,
					GPU:    &wrongResourceValue,
					CPU:    &wrongResourceValue,
				},
				Requests: &v1alpha1.ResourceList{
					Memory: &wrongResourceValue,
					GPU:    &wrongResourceValue,
					CPU:    &wrongResourceValue,
				},
			},
		},
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).Should(HaveOccurred())

	errorMessage := err.Error()
	s.g.Expect(errorMessage).Should(ContainSubstring(
		"validation of memory request is failed: quantities must match the regular expression"))
	s.g.Expect(errorMessage).Should(ContainSubstring(
		"validation of cpu request is failed: quantities must match the regular expression"))
	s.g.Expect(errorMessage).Should(ContainSubstring(
		"validation of memory limit is failed: quantities must match the regular expression"))
	s.g.Expect(errorMessage).Should(ContainSubstring(
		"validation of cpu limit is failed: quantities must match the regular expression"))
	s.g.Expect(errorMessage).Should(ContainSubstring(
		"validation of gpu limit is failed: quantities must match the regular expression"))
}

func (s *ModelDeploymentValidationSuite) TestValidateDefaultDockerPullConnectionName() {
	newDefaultDockerPullConnectionName := "default-docker-pull-conn"
	mdConfig := config.NewDefaultModelDeploymentConfig()
	mdConfig.DefaultDockerPullConnName = newDefaultDockerPullConnectionName

	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{},
	}

	_ = md_routes.NewModelDeploymentValidator(mdConfig, config.NvidiaResourceName).ValidatesMDAndSetDefaults(md)
	s.g.Expect(md.Spec.ImagePullConnectionID).ShouldNot(BeNil())
	s.g.Expect(*md.Spec.ImagePullConnectionID).Should(Equal(newDefaultDockerPullConnectionName))
}

func (s *ModelDeploymentValidationSuite) TestValidateDockerPullConnectionName() {
	dockerPullConnectionName := "default-docker-pull-conn"

	md := &deployment.ModelDeployment{
		Spec: v1alpha1.ModelDeploymentSpec{
			ImagePullConnectionID: &dockerPullConnectionName,
		},
	}

	_ = s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(md.Spec.ImagePullConnectionID).ShouldNot(BeNil())
	s.g.Expect(*md.Spec.ImagePullConnectionID).Should(Equal(dockerPullConnectionName))
}

func (s *ModelDeploymentValidationSuite) TestValidateID() {
	md := &deployment.ModelDeployment{
		ID: "not-VALID-id-",
	}

	err := s.defaultModelValidator.ValidatesMDAndSetDefaults(md)
	s.g.Expect(err).Should(HaveOccurred())
	s.g.Expect(err.Error()).Should(ContainSubstring(validation.ErrIDValidation.Error()))
}
