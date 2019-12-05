// licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// CloneAndAddLabel the given map and returns a new map with the given key and value added.
// Returns the given map, if labelKey is empty.
func CloneAndAddLabel(labels map[string]string, labelKey, labelValue string) map[string]string {
	if labelKey == "" {
		// Don't need to add a label.
		return labels
	}
	// Clone.
	newLabels := map[string]string{}
	for key, value := range labels {
		newLabels[key] = value
	}
	newLabels[labelKey] = labelValue
	return newLabels
}

//MatchLabelForLabelSelector match labels for labelselector, if labelSelecor is nil, select everything
func MatchLabelForLabelSelector(targetLabels map[string]string, labelSelector *metav1.LabelSelector) bool {
	selector, err := ConvertLabels(labelSelector)
	if err != nil {
		// this should not happen if the workset passed validation
		return false
	}
	if selector.Matches(labels.Set(targetLabels)) {
		return true
	}
	return false
}

func AddOwnersLabel(owners, resource, name, namespace string) string {
	if len(owners) == 0 {
		return resource + "." + namespace + "." + name
	}
	return owners + "," + resource + "." + namespace + "." + name
}