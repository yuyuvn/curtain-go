package lib

import (
	"github.com/go-ble/ble"
)

type Command struct {
	uuid	ble.UUID
	value	[2]byte
}

var AppBatteryStatusUuid, _ = ble.Parse("79ab00029dfa4ae2bd46ac69d9fdd743");
var ControlServiceUuid, _ = ble.Parse("79ab10009dfa4ae2bd46ac69d9fdd743");
var ControlServiceControlUuid, _ = ble.Parse("79ab10019dfa4ae2bd46ac69d9fdd743");
var ControlSettingSettingUuid, _ = ble.Parse("79ab10029dfa4ae2bd46ac69d9fdd743");

var BatteryStatusCommand = Command{AppBatteryStatusUuid, [2]byte {0, 0}};
var OpenCommand = Command{ControlServiceControlUuid, [2]byte {0, 0}};
var CloseCommand = Command{ControlServiceControlUuid, [2]byte {0, 1}};
var StopCommand = Command{ControlServiceControlUuid, [2]byte {0, 2}};
var HightSpeedOpenCommand = Command{ControlServiceControlUuid, [2]byte {3, 0}};
var HightSpeedCloseCommand = Command{ControlServiceControlUuid, [2]byte {3, 1}};
