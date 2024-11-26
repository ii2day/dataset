package dataset

import (
	"context"

	datasetv1alpha1 "github.com/BaizeAI/dataset/api/dataset/v1alpha1"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/BaizeAI/dataset/internal/pkg/constants"
)

func datasetConfigMapName(ds *datasetv1alpha1.Dataset) string {
	return "dataset-" + ds.Name + "-config"
}

func (r *DatasetReconciler) getConfigMap(ctx context.Context, ds *datasetv1alpha1.Dataset) (*corev1.ConfigMap, error) {
	cm, err := r.KubeClient.CoreV1().ConfigMaps(ds.Namespace).Get(ctx, datasetConfigMapName(ds), metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return cm, nil
}

type condaOptions struct {
	environmentYAML *string
	requirementsTxt *string
}

type condaOption func(*condaOptions)

func withCondaEnvironmentYAML(yaml string) condaOption {
	return func(o *condaOptions) {
		o.environmentYAML = &yaml
	}
}

func withPipRequirementsTxt(txt string) condaOption {
	return func(o *condaOptions) {
		o.requirementsTxt = &txt
	}
}

func (r *DatasetReconciler) createConfigMap(ctx context.Context, ds *datasetv1alpha1.Dataset, opts ...condaOption) (*corev1.ConfigMap, error) {
	defaultOpts := new(condaOptions)
	for _, opt := range opts {
		opt(defaultOpts)
	}

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      datasetConfigMapName(ds),
			Namespace: ds.Namespace,
			Labels: lo.Assign(ds.Labels, map[string]string{
				constants.DatasetNameLabel: ds.Name,
			}),
			OwnerReferences: datasetOwnerRef(ds),
		},
		Data: make(map[string]string),
	}
	if defaultOpts.environmentYAML != nil {
		cm.Data[constants.DatasetJobCondaCondaEnvironmentYAMLFilename] = *defaultOpts.environmentYAML
	}
	if defaultOpts.requirementsTxt != nil {
		cm.Data[constants.DatasetJobCondaPipRequirementsTxtFilename] = *defaultOpts.requirementsTxt
	}

	return r.KubeClient.CoreV1().ConfigMaps(ds.Namespace).Create(ctx, cm, metav1.CreateOptions{})
}

func (r *DatasetReconciler) updateConfigMap(ctx context.Context, cm *corev1.ConfigMap, opts ...condaOption) (*corev1.ConfigMap, error) {
	defaultOpts := new(condaOptions)
	for _, opt := range opts {
		opt(defaultOpts)
	}
	if cm.Data == nil {
		// NOTICE: .data is potentially nil when user deletes .data field in the manifest
		cm.Data = make(map[string]string)
	}

	if defaultOpts.environmentYAML != nil {
		cm.Data[constants.DatasetJobCondaCondaEnvironmentYAMLFilename] = *defaultOpts.environmentYAML
	}
	if defaultOpts.requirementsTxt != nil {
		cm.Data[constants.DatasetJobCondaPipRequirementsTxtFilename] = *defaultOpts.requirementsTxt
	}

	return r.KubeClient.CoreV1().ConfigMaps(cm.Namespace).Update(ctx, cm, metav1.UpdateOptions{})
}
