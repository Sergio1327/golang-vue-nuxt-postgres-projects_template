package bimport

import (
	"projects_template/internal/bridge"
)

type BridgeImports struct {
	Bridge Bridge
}

func (b *BridgeImports) InitBridge(template bridge.Template) {
	b.Bridge = Bridge{
		Template: template,
	}
}

func NewEmptyBridge() *BridgeImports {
	return &BridgeImports{}
}
