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
package modelpackaging

import (
	"context"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/connection"
	odahuflowv1alpha1 "github.com/odahu/odahu-flow/packages/operator/pkg/apis/odahuflow/v1alpha1"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/packaging"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	"github.com/odahu/odahu-flow/packages/operator/pkg/odahuflow"
	pi_kubernetes "github.com/odahu/odahu-flow/packages/operator/pkg/repository/packaging/kubernetes"
	"github.com/odahu/odahu-flow/packages/operator/pkg/repository/util/kubernetes"
	"github.com/odahu/odahu-flow/packages/operator/pkg/utils"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	tektonschema "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sync"
	"testing"
	"time"
)

const (
	mpName                     = "test-mp"
	timeout                    = time.Second * 5
	testNamespace              = "default"
	testPackagingIntegrationID = "ti"

	tolerationKey      = "key"
	tolerationValue    = "value"
	tolerationOperator = "operator"
	tolerationEffect   = "effect"

	modelPackageImage = "model-image:packager"
	integrationImage  = "model-image:packaging"
)

var (
	nodeSelector = map[string]string{"node-key": "node-value"}

	expectedRequest = reconcile.Request{
		NamespacedName: types.NamespacedName{Name: mpName, Namespace: testNamespace},
	}
	mpNamespacedName = types.NamespacedName{Name: mpName, Namespace: testNamespace}
	testResValue     = "5"
	toleration       = map[string]string{
		config.TolerationKey:      tolerationKey,
		config.TolerationValue:    tolerationValue,
		config.TolerationOperator: tolerationOperator,
		config.TolerationEffect:   tolerationEffect,
	}
)

type ModelPackagingControllerSuite struct {
	suite.Suite
	g               *GomegaWithT
	testEnvironment *envtest.Environment
	cfg             *rest.Config

	k8sClient  client.Client
	k8sManager manager.Manager
	stopMgr    chan struct{}
	mgrStopped *sync.WaitGroup
	requests   chan reconcile.Request
}

func (s *ModelPackagingControllerSuite) createPackagingIntegration() *odahuflowv1alpha1.PackagingIntegration {
	testPackagingIntegration, err := pi_kubernetes.TransformPiToK8s(&packaging.PackagingIntegration{
		ID: testPackagingIntegrationID,
		Spec: packaging.PackagingIntegrationSpec{
			DefaultImage: integrationImage,
			Schema: packaging.Schema{
				Targets: []odahuflowv1alpha1.TargetSchema{
					{
						Name: "target-1",
						ConnectionTypes: []string{
							string(connection.S3Type),
							string(connection.GcsType),
							string(connection.AzureBlobType),
						},
						Required: false,
					},
				},
				Arguments: packaging.JsonSchema{
					Properties: []packaging.Property{
						{
							Name: "argument-1",
							Parameters: []packaging.Parameter{
								{
									Name:  "type",
									Value: "string",
								},
							},
						},
					},
					Required: []string{"argument-1"},
				},
			},
		},
	}, testNamespace)

	if err != nil {
		s.T().Fatal(err)
	}

	return testPackagingIntegration
}

func (s *ModelPackagingControllerSuite) SetupSuite() {
	s.testEnvironment = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "..", "config", "crds"),
			filepath.Join("..", "..", "..", "hack", "tests", "thirdparty_crds"),
		},
	}
	err := apis.AddToScheme(scheme.Scheme)
	if err != nil {
		s.T().Fatal(err)
	}

	if err := tektonschema.AddToScheme(scheme.Scheme); err != nil {
		s.T().Fatal(err)
	}

	if s.cfg, err = s.testEnvironment.Start(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *ModelPackagingControllerSuite) TearDownSuite() {
	if err := s.testEnvironment.Stop(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *ModelPackagingControllerSuite) SetupTest() {
	s.g = NewGomegaWithT(s.T())

	mgr, err := manager.New(s.cfg, manager.Options{})
	s.g.Expect(err).NotTo(HaveOccurred())
	s.k8sClient = mgr.GetClient()
	s.k8sManager = mgr

	s.requests = make(chan reconcile.Request)

	s.stopMgr = make(chan struct{})
	s.mgrStopped = &sync.WaitGroup{}
	s.mgrStopped.Add(1)
	go func() {
		defer s.mgrStopped.Done()
		s.g.Expect(mgr.Start(s.stopMgr)).NotTo(HaveOccurred())
	}()

	if err := s.k8sClient.Create(context.TODO(), s.createPackagingIntegration()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *ModelPackagingControllerSuite) initReconciler(packagingConfig config.ModelPackagingConfig) {
	packagingConfig.PackagingIntegrationNamespace = testNamespace
	packagingConfig.Namespace = testNamespace

	mpReconciler := newReconciler(
		s.k8sManager, packagingConfig, config.NewDefaultOperatorConfig(),
		config.NewDefaultCommonConfig(), config.NvidiaResourceName,
	)
	recFn := reconcile.Func(func(req reconcile.Request) (reconcile.Result, error) {
		result, err := mpReconciler.Reconcile(req)
		s.requests <- req
		return result, err
	})

	s.g.Expect(add(s.k8sManager, recFn)).NotTo(HaveOccurred())
}

func (s *ModelPackagingControllerSuite) TearDownTest() {
	if err := s.k8sClient.Delete(context.TODO(), s.createPackagingIntegration()); err != nil {
		s.T().Fatal(err)
	}

	close(s.stopMgr)
	s.mgrStopped.Wait()
}

func TestModelPackagingControllerSuite(t *testing.T) {
	suite.Run(t, new(ModelPackagingControllerSuite))
}

func (s *ModelPackagingControllerSuite) templateNodeSelectorTest(
	mpResources *odahuflowv1alpha1.ResourceRequirements,
	expectedNodeSelector map[string]string,
	expectedToleration []v1.Toleration,
) {
	mp := &odahuflowv1alpha1.ModelPackaging{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mpName,
			Namespace: testNamespace,
		},
		Spec: odahuflowv1alpha1.ModelPackagingSpec{
			Resources: mpResources,
			Type:      testPackagingIntegrationID,
		},
	}

	err := s.k8sClient.Create(context.TODO(), mp)
	s.g.Expect(err).NotTo(HaveOccurred())
	defer s.k8sClient.Delete(context.TODO(), mp)

	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))
	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))

	s.g.Expect(s.k8sClient.Get(context.TODO(), mpNamespacedName, mp)).ToNot(HaveOccurred())

	tr := &tektonv1alpha1.TaskRun{}
	trKey := types.NamespacedName{Name: mp.Name, Namespace: testNamespace}
	s.g.Expect(s.k8sClient.Get(context.TODO(), trKey, tr)).ToNot(HaveOccurred())

	s.g.Expect(tr.Spec.PodTemplate.NodeSelector).Should(Equal(expectedNodeSelector))
	s.g.Expect(tr.Spec.PodTemplate.Tolerations).Should(Equal(expectedToleration))
}

func (s *ModelPackagingControllerSuite) TestEmptyNodePools() {
	packagingConfig := config.NewDefaultModelPackagingConfig()
	s.initReconciler(packagingConfig)

	s.templateNodeSelectorTest(
		&odahuflowv1alpha1.ResourceRequirements{
			Limits: &odahuflowv1alpha1.ResourceList{
				CPU: &testResValue,
			},
		},
		nil,
		nil,
	)
}

func (s *ModelPackagingControllerSuite) TestNodePools() {
	packagingConfig := config.NewDefaultModelPackagingConfig()
	packagingConfig.NodeSelector = nodeSelector
	packagingConfig.Toleration = toleration
	s.initReconciler(packagingConfig)

	s.templateNodeSelectorTest(
		&odahuflowv1alpha1.ResourceRequirements{
			Limits: &odahuflowv1alpha1.ResourceList{
				CPU: &testResValue,
			},
		},
		nodeSelector,
		[]v1.Toleration{{
			Key:      tolerationKey,
			Operator: tolerationOperator,
			Value:    tolerationValue,
			Effect:   tolerationEffect,
		}},
	)
}

func (s *ModelPackagingControllerSuite) TestOnlyNodeSelector() {
	packagingConfig := config.NewDefaultModelPackagingConfig()
	packagingConfig.NodeSelector = nodeSelector
	s.initReconciler(packagingConfig)

	s.templateNodeSelectorTest(
		&odahuflowv1alpha1.ResourceRequirements{
			Limits: &odahuflowv1alpha1.ResourceList{
				CPU: &testResValue,
			},
		},
		nodeSelector,
		nil,
	)
}

func (s *ModelPackagingControllerSuite) TestPackagingStepConfiguration() {
	packagingConfig := config.NewDefaultModelPackagingConfig()
	packagingConfig.ModelPackagerImage = modelPackageImage
	s.initReconciler(packagingConfig)

	packResources := &odahuflowv1alpha1.ResourceRequirements{
		Limits: &odahuflowv1alpha1.ResourceList{
			CPU: &testResValue,
			GPU: &testResValue,
		},
		Requests: &odahuflowv1alpha1.ResourceList{
			CPU: &testResValue,
			GPU: &testResValue,
		},
	}
	k8sPackagingResources, err := kubernetes.ConvertOdahuflowResourcesToK8s(packResources, config.NvidiaResourceName)
	s.g.Expect(err).Should(BeNil())

	mp := &odahuflowv1alpha1.ModelPackaging{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mpName,
			Namespace: testNamespace,
		},
		Spec: odahuflowv1alpha1.ModelPackagingSpec{
			Image:     integrationImage,
			Resources: packResources,
			Type:      testPackagingIntegrationID,
		},
	}

	err = s.k8sClient.Create(context.TODO(), mp)
	s.g.Expect(err).NotTo(HaveOccurred())
	defer s.k8sClient.Delete(context.TODO(), mp)

	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))
	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))

	s.g.Expect(s.k8sClient.Get(context.TODO(), mpNamespacedName, mp)).ToNot(HaveOccurred())

	tr := &tektonv1alpha1.TaskRun{}
	trKey := types.NamespacedName{Name: mp.Name, Namespace: testNamespace}
	s.g.Expect(s.k8sClient.Get(context.TODO(), trKey, tr)).ToNot(HaveOccurred())

	expectedHelperContainerResources := v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    *k8sPackagingResources.Limits.Cpu(),
			v1.ResourceMemory: *utils.DefaultHelperLimits.Memory(),
		},
		Requests: v1.ResourceList{
			v1.ResourceMemory: resource.MustParse("0"),
			v1.ResourceCPU:    resource.MustParse("0"),
		},
	}

	for _, step := range tr.Spec.TaskSpec.Steps {
		switch step.Name {
		case odahuflow.PackagerSetupStep:
			s.g.Expect(step.Image).Should(Equal(modelPackageImage))
			s.g.Expect(step.Resources).Should(Equal(expectedHelperContainerResources))
		case odahuflow.PackagerPackageStep:
			s.g.Expect(step.Image).Should(Equal(integrationImage))
			s.g.Expect(step.Resources).Should(Equal(k8sPackagingResources))
		case odahuflow.PackagerResultStep:
			s.g.Expect(step.Image).Should(Equal(modelPackageImage))
			s.g.Expect(step.Resources).Should(Equal(expectedHelperContainerResources))
		default:
			s.T().Errorf("Unexpected spep name: %s", step.Name)
		}
	}
}

func (s *ModelPackagingControllerSuite) TestPackagingTimeout() {
	packagingConfig := config.NewDefaultModelPackagingConfig()
	packagingConfig.Timeout = 3 * time.Hour
	s.initReconciler(packagingConfig)

	mp := &odahuflowv1alpha1.ModelPackaging{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mpName,
			Namespace: testNamespace,
		},
		Spec: odahuflowv1alpha1.ModelPackagingSpec{
			Type: testPackagingIntegrationID,
		},
	}

	err := s.k8sClient.Create(context.TODO(), mp)
	s.g.Expect(err).NotTo(HaveOccurred())
	defer s.k8sClient.Delete(context.TODO(), mp)

	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))
	s.g.Eventually(s.requests, timeout).Should(Receive(Equal(expectedRequest)))

	s.g.Expect(s.k8sClient.Get(context.TODO(), mpNamespacedName, mp)).ToNot(HaveOccurred())

	tr := &tektonv1alpha1.TaskRun{}
	trKey := types.NamespacedName{Name: mp.Name, Namespace: testNamespace}
	s.g.Expect(s.k8sClient.Get(context.TODO(), trKey, tr)).ToNot(HaveOccurred())

	s.g.Expect(tr.Spec.Timeout.Duration).Should(Equal(time.Hour * 3))
}
