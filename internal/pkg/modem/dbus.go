package modem

import (
	"github.com/godbus/dbus/v5"
	"github.com/maltegrosse/go-modemmanager"
)

func (m *Modem) call(objectPath dbus.ObjectPath, method string, args ...any) error {
	dbusConn, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	obj := dbusConn.Object(modemmanager.ModemManagerInterface, objectPath)
	return obj.Call(method, 0, args...).Err
}

func (m *Modem) getProperty(objectPath dbus.ObjectPath, property string) (dbus.Variant, error) {
	dbusConn, err := dbus.SystemBus()
	if err != nil {
		return dbus.Variant{}, err
	}
	obj := dbusConn.Object(modemmanager.ModemManagerInterface, objectPath)
	return obj.GetProperty(property)
}

func (m *Modem) systemBusPrivate() (*dbus.Conn, error) {
	dbusConn, err := dbus.SystemBusPrivate()
	if err != nil {
		return nil, err
	}

	err = dbusConn.Auth(nil)
	if err != nil {
		dbusConn.Close()
		return nil, err
	}

	err = dbusConn.Hello()
	if err != nil {
		dbusConn.Close()
		return nil, err
	}
	return dbusConn, nil
}
