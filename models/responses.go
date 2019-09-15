package models


type BalanceResponse struct {
	Data Account `json:"data"`
	Code int `json:"code"`
	Error Error `json:"error"`
}

type UserResponse struct {
	Data User `json:"data"`
	Code int `json:"code"`
	Error Error `json:"error"`
}

type PackageResponse struct {
	Data Package `json:"data"`
	Code int `json:"code"`
	Error Error `json:"error"`
}

type UserPackageResponse struct {
	Data UserPackage `json:"data"`
	Code int `json:"code"`
	Error Error `json:"error"`
}

type Error struct{
	Message string `json:"message"`
}