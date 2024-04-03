package heartbeat

import "fmt"

// Структура для хранения информации о конкретной сущности
type InstanceHeath struct {
	ObjType string // LB = LoadBalancer || Server = Server || DBC = DataBaseCluster
	Address string // Network address of instance
	Status  string // OK || Killed || Ready || DOWN || Reboot || Service
}

// Returns a structure for storing information about every instance in system
func CreateHealthCheck() *[]InstanceHeath {
	HealthCheck := new([]InstanceHeath)
	return HealthCheck
}

// Отправить запрос на подтверждение функционирования узла
func RequestAvailability(sender, receiver string) error {
	// TODO
	return nil
}

// Отправить подтверждение функционирования узла
func ApproveAvailability(sender, receiver string) error {
	// TODO
	return nil
}

// Задать промежуток времени с которым будет обновляться статистика загрузки узла
func SetStatusLogDuration() error {
	// TODO
	return nil
}

func TestMe() {
	fmt.Println("Наконец-то научился работать с модулями")
}
