package gonetworkmanager

import (
	"github.com/godbus/dbus"
)

const (
	ActiveConnectionInterface             = NetworkManagerInterface + ".Connection.Active"
	ActiveConnectionProperyConnection     = ActiveConnectionInterface + ".Connection"
	ActiveConnectionProperySpecificObject = ActiveConnectionInterface + ".SpecificObject"
	ActiveConnectionProperyId             = ActiveConnectionInterface + ".Id"
	ActiveConnectionProperyUuid           = ActiveConnectionInterface + ".Uuid"
	ActiveConnectionProperyType           = ActiveConnectionInterface + ".Type"
	ActiveConnectionProperyDevices        = ActiveConnectionInterface + ".Devices"
	ActiveConnectionProperyState          = ActiveConnectionInterface + ".State"
	ActiveConnectionProperyStateFlags     = ActiveConnectionInterface + ".StateFlags"
	ActiveConnectionProperyDefault        = ActiveConnectionInterface + ".Default"
	ActiveConnectionProperyIP4Config      = ActiveConnectionInterface + ".Ip4Config"
	ActiveConnectionProperyDHCP4Config    = ActiveConnectionInterface + ".Dhcp4Config"
	ActiveConnectionProperyDefault6       = ActiveConnectionInterface + ".Default6"
	ActiveConnectionProperyVpn            = ActiveConnectionInterface + ".Vpn"
	ActiveConnectionProperyMaster         = ActiveConnectionInterface + ".Master"
)

type ActiveConnection interface {
	// GetConnection gets connection object of the connection.
	GetConnection() Connection

	// GetID gets the ID of the connection.
	GetID() string

	// GetUUID gets the UUID of the connection.
	GetUUID() string

	// GetType gets the type of the connection.
	GetType() string

	// GetDevices gets array of device objects which are part of this active connection.
	GetDevices() []Device

	// GetState gets the state of the connection.
	GetState() uint32

	// GetStateFlags gets the state flags of the connection.
	GetStateFlags() uint32

	// GetDefault gets the default IPv4 flag of the connection.
	GetDefault() uint32

	// GetIP4Config gets the IP4Config of the connection.
	GetIP4Config() IP4Config

	// GetDHCP4Config gets the DHCP4Config of the connection.
	GetDHCP4Config() DHCP4Config

	// GetVPN gets the VPN flag of the connection.
	GetVPN() uint32

	// GetMaster gets the master device of the connection.
	GetMaster() Device
}

func NewActiveConnection(objectPath dbus.ObjectPath) (ActiveConnection, error) {
	var a activeConnection
	return &a, a.init(NetworkManagerInterface, objectPath)
}

type activeConnection struct {
	dbusBase
}

func (a *activeConnection) GetConnection() Connection {
	path := a.getObjectProperty(ActiveConnectionProperyConnection)
	con, err := NewConnection(path)
	if err != nil {
		panic(err)
	}
	return con
}

func (a *activeConnection) GetID() string {
	return ""
}

func (a *activeConnection) GetUUID() string {
	return ""
}

func (a *activeConnection) GetType() string {
	return ""
}

func (a *activeConnection) GetDevices() []Device {
	return nil
}

func (a *activeConnection) GetState() uint32 {
	return 0
}

func (a *activeConnection) GetStateFlags() uint32 {
	return 0
}

func (a *activeConnection) GetDefault() uint32 {
	return 0
}

func (a *activeConnection) GetIP4Config() IP4Config {
	return nil
}

func (a *activeConnection) GetDHCP4Config() DHCP4Config {
	return nil
}

func (a *activeConnection) GetVPN() uint32 {
	return 0
}

func (a *activeConnection) GetMaster() Device {
	return nil
}
