package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// ========== STRUCT REQUEST ==========

type VirtualAccountDeleteRequest struct {
	PartnerServiceId string `json:"partnerServiceId"`
	CustomerNo       string `json:"customerNo"`
	VirtualAccountNo string `json:"virtualAccountNo"`
	TrxId            string `json:"trxId"`
}

type VirtualAccountRequest struct {
	PartnerServiceId    string         `json:"partnerServiceId"`
	CustomerNo          string         `json:"customerNo"`
	VirtualAccountNo    string         `json:"virtualAccountNo"`
	VirtualAccountName  string         `json:"virtualAccountName"`
	VirtualAccountEmail string         `json:"virtualAccountEmail"`
	TrxId               string         `json:"trxId"`
	TotalAmount         Amount         `json:"totalAmount"`
	BillDetails         []BillDetail   `json:"billDetails"`
	ExpiredDate         string         `json:"expiredDate"`
	AdditionalInfo      AdditionalInfo `json:"additionalInfo"`
}

// ========== STRUCT RESPONSE ==========
type VirtualAccountResponse struct {
	ResponseCode    string             `json:"responseCode"`
	ResponseMessage string             `json:"responseMessage"`
	VirtualAccount  VirtualAccountData `json:"virtualAccountData"`
}

type VirtualAccountData struct {
	PartnerServiceId    string         `json:"partnerServiceId"`
	CustomerNo          string         `json:"customerNo"`
	VirtualAccountNo    string         `json:"virtualAccountNo"`
	VirtualAccountName  string         `json:"virtualAccountName"`
	VirtualAccountEmail string         `json:"virtualAccountEmail"`
	TrxId               string         `json:"trxId"`
	TotalAmount         Amount         `json:"totalAmount"`
	BillDetails         []BillDetail   `json:"billDetails"`
	ExpiredDate         time.Time      `json:"expiredDate"`
	AdditionalInfo      AdditionalInfo `json:"additionalInfo"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type BillDetail struct {
	BillDescription Description `json:"billDescription"`
}

type Description struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

type AdditionalInfo struct {
	InsertId string `json:"insertId,omitempty"`
}

// ========== GENERIC ERROR RESPONSE ==========
func errorResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, VirtualAccountResponse{
		ResponseCode:    "5002700",
		ResponseMessage: message,
		VirtualAccount:  VirtualAccountData{},
	})
}

// ========== MAIN HANDLER ==========
// Note: The CreateVA function is a placeholder and does not implement actual creation logic.
func CreateVA(c echo.Context) error {
	// ===== Header Validation =====
	headers := []string{
		"X-TIMESTAMP", "X-SIGNATURE", "X-PARTNER-ID", "X-EXTERNAL-ID", "CHANNEL-ID",
	}

	for _, h := range headers {
		if c.Request().Header.Get(h) == "" {
			return errorResponse(c, "Missing required header: "+h)
		}
	}

	// ===== Bind Body =====
	var req VirtualAccountRequest
	if err := c.Bind(&req); err != nil {
		return errorResponse(c, "Invalid JSON format")
	}

	// ===== Field Validation =====
	if req.PartnerServiceId == "" ||
		req.CustomerNo == "" ||
		req.VirtualAccountNo == "" ||
		req.VirtualAccountName == "" ||
		req.VirtualAccountEmail == "" ||
		req.TrxId == "" ||
		req.TotalAmount.Value == "" ||
		req.TotalAmount.Currency == "" {
		return errorResponse(c, "Missing required field in request body")
	}

	// ===== Parse Expiration Date =====
	expiredDate, err := time.Parse(time.RFC3339, req.ExpiredDate)
	if err != nil {
		expiredDate = time.Now().Add(24 * time.Hour) // default expired in 1 day
	}

	// ===== Success Response =====
	resp := VirtualAccountResponse{
		ResponseCode:    "2002700",
		ResponseMessage: "Successful",
		VirtualAccount: VirtualAccountData{
			PartnerServiceId:    req.PartnerServiceId,
			CustomerNo:          req.CustomerNo,
			VirtualAccountNo:    req.VirtualAccountNo,
			VirtualAccountName:  req.VirtualAccountName,
			VirtualAccountEmail: req.VirtualAccountEmail,
			TrxId:               req.TrxId,
			TotalAmount:         req.TotalAmount,
			BillDetails:         req.BillDetails,
			ExpiredDate:         expiredDate,
			AdditionalInfo:      req.AdditionalInfo,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

// Note: The DeleteVA function is a placeholder and does not implement actual deletion logic.
func DeleteVA(c echo.Context) error {
	// ===== Header Validation =====
	headers := []string{
		"X-TIMESTAMP", "X-SIGNATURE", "X-PARTNER-ID", "X-EXTERNAL-ID", "CHANNEL-ID",
	}

	for _, h := range headers {
		if c.Request().Header.Get(h) == "" {
			return errorResponse(c, "Missing required header: "+h)
		}
	}

	// ===== Bind Body =====
	var req VirtualAccountDeleteRequest
	if err := c.Bind(&req); err != nil {
		return errorResponse(c, "Invalid JSON format")
	}

	// ===== Field Validation =====
	if req.VirtualAccountNo == "" {
		return errorResponse(c, "Missing required field: virtualAccountNo")
	}

	// ===== Success Response =====
	resp := VirtualAccountResponse{
		ResponseCode:    "2003100",
		ResponseMessage: "Success",
		VirtualAccount: VirtualAccountData{
			PartnerServiceId: req.PartnerServiceId,
			CustomerNo:       req.CustomerNo,
			VirtualAccountNo: req.VirtualAccountNo,
			TrxId:            req.TrxId,
		},
	}

	return c.JSON(http.StatusOK, resp)
}
