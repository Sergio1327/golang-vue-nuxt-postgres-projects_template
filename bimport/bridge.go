package bimport

import "projects_template/internal/bridge"

type Bridge struct {
	Template bridge.Template
}

type TestBridge struct {
	Template *bridge.MockTemplate
}
