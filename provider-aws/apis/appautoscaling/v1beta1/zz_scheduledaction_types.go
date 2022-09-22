/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ScalableTargetActionObservation struct {
}

type ScalableTargetActionParameters struct {

	// The maximum capacity. At least one of max_capacity or min_capacity must be set.
	// +kubebuilder:validation:Optional
	MaxCapacity *string `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// The minimum capacity. At least one of min_capacity or max_capacity must be set.
	// +kubebuilder:validation:Optional
	MinCapacity *string `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type ScheduledActionObservation struct {

	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScheduledActionParameters struct {

	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of timezone.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The name of the scheduled action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: AWS Application Auto Scaling API Reference
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("resource_id",false)
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Target in appautoscaling to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// The scalable dimension. Documentation can be found in the parameter at: AWS Application Auto Scaling API Reference Example: ecs:service:DesiredCount
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("scalable_dimension",false)
	// +kubebuilder:validation:Optional
	ScalableDimension *string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`

	// Reference to a Target in appautoscaling to populate scalableDimension.
	// +kubebuilder:validation:Optional
	ScalableDimensionRef *v1.Reference `json:"scalableDimensionRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate scalableDimension.
	// +kubebuilder:validation:Optional
	ScalableDimensionSelector *v1.Selector `json:"scalableDimensionSelector,omitempty" tf:"-"`

	// The new minimum and maximum capacity. You can set both values or just one. See below
	// +kubebuilder:validation:Required
	ScalableTargetAction []ScalableTargetActionParameters `json:"scalableTargetAction" tf:"scalable_target_action,omitempty"`

	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in timezone. Documentation can be found in the parameter at: AWS Application Auto Scaling API Reference
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`

	// The namespace of the AWS service. Documentation can be found in the parameter at: AWS Application Auto Scaling API Reference Example: ecs
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/appautoscaling/v1beta1.Target
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("service_namespace",false)
	// +kubebuilder:validation:Optional
	ServiceNamespace *string `json:"serviceNamespace,omitempty" tf:"service_namespace,omitempty"`

	// Reference to a Target in appautoscaling to populate serviceNamespace.
	// +kubebuilder:validation:Optional
	ServiceNamespaceRef *v1.Reference `json:"serviceNamespaceRef,omitempty" tf:"-"`

	// Selector for a Target in appautoscaling to populate serviceNamespace.
	// +kubebuilder:validation:Optional
	ServiceNamespaceSelector *v1.Selector `json:"serviceNamespaceSelector,omitempty" tf:"-"`

	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of timezone.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for start_time and end_time. Valid values are the canonical names of the IANA time zones supported by Joda-Time, such as Etc/GMT+9 or Pacific/Tahiti. Default is UTC.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

// ScheduledActionSpec defines the desired state of ScheduledAction
type ScheduledActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduledActionParameters `json:"forProvider"`
}

// ScheduledActionStatus defines the observed state of ScheduledAction.
type ScheduledActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduledActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledAction is the Schema for the ScheduledActions API. Provides an Application AutoScaling ScheduledAction resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ScheduledAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduledActionSpec   `json:"spec"`
	Status            ScheduledActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledActionList contains a list of ScheduledActions
type ScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledAction `json:"items"`
}

// Repository type metadata.
var (
	ScheduledAction_Kind             = "ScheduledAction"
	ScheduledAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScheduledAction_Kind}.String()
	ScheduledAction_KindAPIVersion   = ScheduledAction_Kind + "." + CRDGroupVersion.String()
	ScheduledAction_GroupVersionKind = CRDGroupVersion.WithKind(ScheduledAction_Kind)
)

func init() {
	SchemeBuilder.Register(&ScheduledAction{}, &ScheduledActionList{})
}
