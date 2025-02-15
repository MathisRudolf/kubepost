package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Role is the Schema for the roles API
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

type RoleSpec struct {
	InstanceRef InstanceRef `json:"instanceRef"`
	RoleName    string      `json:"roleName"`
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=true
	PreventDeletion bool `json:"preventDeletion"`
	//+kubebuilder:validation:Optional
	//+kubebuilder:default:=false
	CascadeDelete bool `json:"cascadeDelete"`
	//+kubebuilder:validation:Optional
	Grant Grant `json:"grant"`
}

type Grant struct {
	Database string `json:"database,omitempty"`
	Schema string `json:"schema,omitempty"`
	ObjectType string `json:"objectType"`
	Privileges []string `json:"privileges"`
}

type RoleStatus struct {
	Status string `json:"status"`
}

//+kubebuilder:object:root=true

// RoleList contains a list of Role
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}