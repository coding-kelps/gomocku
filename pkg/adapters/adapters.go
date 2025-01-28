package adapters

import (
	"github.com/coding-kelps/gomocku/pkg/adapters/stdio"
	"github.com/coding-kelps/gomocku/pkg/adapters/tcp"
)

var NewStdioManagerInterface 	= stdio.NewStdioManagerInterface
var NewTCPManagerInterface		= tcp.NewTCPManagerInterface
