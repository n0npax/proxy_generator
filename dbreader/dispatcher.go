package dbreader

import (
	"github.com/n0npax/proxy_generator/parser"
	"strings"
)

// ReadNginxRedirection dispatcher right functions for different dbs
func ReadNginxRedirection(dnEndpoint string) []parser.NginxRedirection {

	dbEndpointSubparts := strings.Split(dnEndpoint, "//")
	if len(dbEndpointSubparts) != 2 {
		panic("endpoint should contain \"//\" just once")
	}
	dbPath := dbEndpointSubparts[1]
	dbType := dbEndpointSubparts[0]
	switch {

	case strings.HasPrefix(dbType, "sqlite"):
		return SqliteReadNginxRedirection(dbPath)
	default:
		panic("Not supported yet")
	}
}
