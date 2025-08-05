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

// ========== STRUCTS FOR STATUS VA RESPONSE ==========
type PaidAmount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type StatusVaVirtualAccountData struct {
	PartnerServiceId string     `json:"partnerServiceId"`
	CustomerNo       string     `json:"customerNo"`
	VirtualAccountNo string     `json:"virtualAccountNo"`
	InquiryRequestId string     `json:"inquiryRequestId"`
	PaymentRequestId string     `json:"paymentRequestId"`
	PaidAmount       PaidAmount `json:"paidAmount"`
	TotalAmount      PaidAmount `json:"totalAmount"`
	ReferenceNo      string     `json:"referenceNo"`
	TrxDateTime      string     `json:"trxDateTime"`
	TransactionDate  string     `json:"transactionDate"`
}

type StatusVaAdditionalInfo struct {
	InsertId   string `json:"insertId"`
	TagId      string `json:"tagId"`
	TrxStatus  string `json:"trxStatus"`
	TrxMessage string `json:"trxMessage"`
}

type StatusVaResponse struct {
	ResponseCode       string                     `json:"responseCode"`
	ResponseMessage    string                     `json:"responseMessage"`
	VirtualAccountData StatusVaVirtualAccountData `json:"virtualAccountData"`
	AdditionalInfo     StatusVaAdditionalInfo     `json:"additionalInfo"`
}

type StatusVaErrorResponse struct {
	ResponseCode       string                 `json:"responseCode"`
	ResponseMessage    string                 `json:"responseMessage"`
	VirtualAccountData map[string]interface{} `json:"virtualAccountData"`
}

// ========== STATUS VA HANDLER ==========

func StatusVa(c echo.Context) error {
	// Set X-TIMESTAMP header
	c.Response().Header().Set("X-TIMESTAMP", time.Now().Format(time.RFC3339))
	c.Response().Header().Set("Content-Type", "application/json")

	// Example: check for required query param (simulate error)
	// vaNo := c.QueryParam("virtualAccountNo")
	// if vaNo == "" {
	// 	errResp := StatusVaErrorResponse{
	// 		ResponseCode:       "5002600",
	// 		ResponseMessage:    "General Error",
	// 		VirtualAccountData: map[string]interface{}{},
	// 	}
	// 	return c.JSON(http.StatusBadRequest, errResp)
	// }

	// Success response (static example, replace with real data as needed)
	resp := StatusVaResponse{
		ResponseCode:    "2002600",
		ResponseMessage: "Successful",
		VirtualAccountData: StatusVaVirtualAccountData{
			PartnerServiceId: "7421",
			CustomerNo:       "000587616",
			VirtualAccountNo: "7421000587616",
			InquiryRequestId: "13432232",
			PaymentRequestId: "",
			PaidAmount: PaidAmount{
				Currency: "IDR",
				Value:    "0.00",
			},
			TotalAmount: PaidAmount{
				Currency: "IDR",
				Value:    "26900.00",
			},
			ReferenceNo:     "",
			TrxDateTime:     "2024-02-28T16:56:05+07:00",
			TransactionDate: "",
		},
		AdditionalInfo: StatusVaAdditionalInfo{
			InsertId:   "293126",
			TagId:      "293126",
			TrxStatus:  "03",
			TrxMessage: "Pending",
		},
	}
	return c.JSON(http.StatusOK, resp)
}
