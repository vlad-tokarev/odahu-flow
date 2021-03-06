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

package v1alpha1

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ToolchainIntegrationSpec defines the desired state of ToolchainIntegration
type ToolchainIntegrationSpec struct {
	// Path to binary which starts a training process
	Entrypoint string `json:"entrypoint"`
	// Default training Docker image
	DefaultImage string `json:"defaultImage"`
	// Additional environments for a training process
	AdditionalEnvironments map[string]string `json:"additionalEnvironments,omitempty"`
}

// ToolchainIntegrationStatus defines the observed state of ToolchainIntegration
type ToolchainIntegrationStatus struct {
	// Info about create and update
	//CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	//UpdatedAt *metav1.Time `json:"updatedAt,omitempty"`
	Modifiable `json:",inline"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainIntegration is the Schema for the toolchainintegrations API
// +k8s:openapi-gen=true
type ToolchainIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToolchainIntegrationSpec   `json:"spec,omitempty"`
	Status ToolchainIntegrationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToolchainIntegrationList contains a list of ToolchainIntegration
type ToolchainIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ToolchainIntegration `json:"items"`
}

func (tiSpec ToolchainIntegrationSpec) Value() (driver.Value, error) {
	return json.Marshal(tiSpec)
}

func (tiSpec *ToolchainIntegrationSpec) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	res := json.Unmarshal(b, &tiSpec)
	return res
}

func (tiStatus ToolchainIntegrationStatus) Value() (driver.Value, error) {
	return json.Marshal(tiStatus)
}

func (tiStatus *ToolchainIntegrationStatus) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	res := json.Unmarshal(b, &tiStatus)
	return res
}

func init() {
	SchemeBuilder.Register(&ToolchainIntegration{}, &ToolchainIntegrationList{})
}
