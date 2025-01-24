package adapters

import (
	"github.com/coding-kelps/gomocku/pkg/adapters/stdio"
	"github.com/coding-kelps/gomocku/pkg/adapters/tcp"
)

var NewStdio 	= stdio.NewStdio
var NewTCP		= tcp.NewTCP
