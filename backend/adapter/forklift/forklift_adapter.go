// internal/adapter/forklift/forklift_adapter.go
package forklift

import (
	"corporation-site/service"
	"encoding/json"
	"net/http"

	api "corporation-site/infra/api/forklift"

	"github.com/gin-gonic/gin"
)

type ForkliftAdapter struct {
	forkliftService service.ForkliftService
}

func NewForkliftAdapter(forkliftService service.ForkliftService) *ForkliftAdapter {
	return &ForkliftAdapter{
		forkliftService: forkliftService,
	}
}

// GetForkliftsTypeEnginetype handles requests for forklifts by engine type
func (a *ForkliftAdapter) GetForkliftsTypeEnginetype(w http.ResponseWriter, r *http.Request, enginetype string) {
	forklifts, err := a.forkliftService.GetForkliftsByType(r.Context(), enginetype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(forklifts)
}

// GetForkliftsTypeEnginetypeModelSerial handles requests for a specific forklift
func (a *ForkliftAdapter) GetForkliftsTypeEnginetypeModelSerial(w http.ResponseWriter, r *http.Request, enginetype string, model string, serial string) {
	forklift, err := a.forkliftService.GetForkliftByEngineTypeModelSerial(r.Context(), enginetype, model, serial)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(forklift)
}

func (a *ForkliftAdapter) RegisterRoutes(router interface{}) {
	switch r := router.(type) {
	case *gin.Engine:
		a.registerGinRoutes(r)
	default:
		panic("Unsupported router type")
	}
}

func (a *ForkliftAdapter) registerGinRoutes(router *gin.Engine) {
	handler := api.HandlerWithOptions(a, api.StdHTTPServerOptions{})
	router.Any("/forklifts/type/:enginetype", gin.WrapH(handler))
	router.Any("/forklifts/type/:enginetype/:model/:serial", gin.WrapH(handler))
}
