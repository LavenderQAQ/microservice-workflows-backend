package common

import (
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type NodeInterface interface {
	// Get node id
	GetId() string
	// Get the category of node
	GetCustom() int
	// Gets the ids of all the in nodes
	GetInNode() []string
	// Gets the ids of all the out nodes
	GetOutNode() []string
	// Append in nodes
	AppendInNode(string)
	// Append out nodes
	AppendOutNode(string)
	// Check if this node has in nodes
	HaveInNode() bool
	// Check if this node has out nodes
	HaveOutNode() bool
	// Generate template
	GenerateTemplate() v1alpha1.Template
}
