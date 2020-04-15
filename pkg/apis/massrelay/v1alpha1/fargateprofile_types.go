package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TagKey is a value used to satisfy the CRD generator
type TagKey string

// TagValue is a value used to satisfy the CRD generator
type TagValue string

// FargateProfileSpec defines the desired state of FargateProfile
type FargateProfileSpec struct {
	// Name of the Fargate profile. Maps to the
	// https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html#AmazonEKS-Type-FargateProfile-fargateProfileName
	// field in the AWS API.
	Name string `json:"name"`
	// The Amazon Resource Name (ARN) of the pod execution role to use for pods
	// that match the selectors in the Fargate profile. The pod execution role allows
	// Fargate infrastructure to register with your cluster as a node, and it provides
	// read access to Amazon ECR image repositories. For more information, see Pod
	// Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html)
	// in the Amazon EKS User Guide.
	PodExecutionRoleArn string `json:"podExecutionRoleArn"`
	// The selectors to match for pods to use this Fargate profile. Each selector
	// must have an associated namespace. Optionally, you can also specify labels
	// for a namespace. You may specify up to five selectors in a Fargate profile.
	Selectors []*FargateProfileSelector `json:"selectors"`
	// The IDs of subnets to launch your pods into. At this time, pods running on
	// Fargate are not assigned public IP addresses, so only private subnets (with
	// no direct route to an Internet Gateway) are accepted for this parameter.
	Subnets []*string `json:"subnets"`
	// The metadata to apply to the Fargate profile to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both
	// of which you define. Fargate profile tags do not propagate to any other resources
	// associated with the Fargate profile, such as the pods that are scheduled
	// with it.
	Tags map[TagKey]TagValue `json:"tags"`
}

// LabelKey is a type used to to satisfy CRD generator
type LabelKey string

// LabelValue is a type used to satisfy CRD generator
type LabelValue string

// Namespace is a type used to satisfy CRD generator
type Namespace string

// FargateProfileSelector mirrors the EKS Fargate API's profile selector.
type FargateProfileSelector struct {
	// The Kubernetes labels that the selector should match. A pod must contain
	// all of the labels that are specified in the selector for it to be considered
	// a match.
	Labels map[LabelKey]LabelValue `json:"labels"`
	// The Kubernetes namespace that the selector should match.
	Namespace Namespace `json:"namespace"`
}

// FargateProfileStatus defines the observed state of FargateProfile
type FargateProfileStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FargateProfile is the Schema for the fargateprofiles API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=fargateprofiles,scope=Namespaced
type FargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FargateProfileSpec   `json:"spec,omitempty"`
	Status FargateProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FargateProfileList contains a list of FargateProfile
type FargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FargateProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FargateProfile{}, &FargateProfileList{})
}
