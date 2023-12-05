package icp

import "time"

type IcpAccount struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 用户 id
	UserId uint `json:"userId"`
	// user name
	UserName string `json:"userName"`
	// principal
	Principal string `json:"principal"`
	// accountid
	AccountId string `json:"accountId"`
	// wallet
	Wallet string `json:"wallet"`
	// cycles
	Cycles uint `json:"cycles"`
	// balance
	Balance uint `json:"balance"`
	// project count
	Projects uint `json:"projects"`
	// canister count
	Canisters uint `json:"canisters"`
}

type IcpProject struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// project name
	ProjectName string `json:"projectName"`
	// user id
	UserId uint `json:"userId"`
	// canister count
	Canisters uint `json:"canisters"`
}

type CanisterStatus string

const (
	Running CanisterStatus = "Running"
	Stopped CanisterStatus = "Stopped"
)

type ControllerRole string

const (
	Administrator ControllerRole = "Administrator"
	ReadWrite     ControllerRole = "Read-Write"
)

type IcpCannister struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// canister id
	CanisterId string `json:"canisterId"`
	// canister name
	CanisterName string `json:"canisterName"`
	// status
	Status CanisterStatus `json:"status"`
	// controllers
	Controllers string `json:"controllers"`
	// cycles
	Cycles uint `json:"cycles"`
	// update at
	UpdatedAt time.Time `json:"updatedAt"`
	// project name
	ProjectName string `json:"projectName"`
}
