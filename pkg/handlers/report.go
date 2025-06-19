package handlers

import (
	"net/http"

	"BRSBackend/pkg/api"
)

func (h *Handler) ListOverdueRentals(w http.ResponseWriter, r *http.Request, params api.ListOverdueRentalsParams) {

	if params.StudentCardId == nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Student card ID is required")
		return
	}

	if params.Limit == nil || int(*params.Limit) < 0 {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter")
	}

	if params.Offset == nil || int(*params.Offset) < 0 {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid offset parameter")
	}

	overdueRentals, err := h.reportService.GetOverdueRentals(r.Context(), params.StudentCardId, int(*params.Limit), int(*params.Offset))
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, overdueRentals)
}

func (h *Handler) GetRentalReports(w http.ResponseWriter, r *http.Request, params api.GetRentalReportsParams) {
	limit := 0
	offset := 0
	if params.Limit == nil || int(*params.Limit) > 0 {
		limit = 20
	}

	if params.Offset == nil || int(*params.Offset) > 0 {
		offset = 0
	}

	report, err := h.reportService.GetRentalReport(r.Context(), limit, offset)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, report)
}
