package services

import (
	"../lib"
	"github.com/go-ble/ble"
)

const (
	AppServiceUuid = ble.Parse('79ab00009dfa4ae2bd46ac69d9fdd743')
	AppServiceStatusUuid = ble.Parse('79ab00019dfa4ae2bd46ac69d9fdd743')
)

func runCommand(command lib.Command) error {
	defer lib.runDisconnect()
	lib.runExplore()
	auth()
	lib.runWrite(command)
	lib.runDisconnect()
}

func readStatus(command lib.Command) ([]byte, error) {
	defer lib.runDisconnect()
	lib.runExplore()
	auth()
	return lib.runRead(command)
}

func auth() error {
	data, err := lib.runRead(lib.Command{AppServiceStatusUuid})
	if err != nil {
		return fmt.Errorf("Failed to get seed")
	}

	token := data[11:15]

}
