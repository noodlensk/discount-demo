package ports

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/noodlensk/discount-demo/internal/common/auth"
	"github.com/noodlensk/discount-demo/internal/common/server/httperr"
	"github.com/noodlensk/discount-demo/internal/discount/app"
	"github.com/noodlensk/discount-demo/internal/discount/app/command"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(application app.Application) HTTPServer {
	return HTTPServer{application}
}

func (h HTTPServer) GenerateCode(w http.ResponseWriter, r *http.Request, brandID string) {
	postGenerateCodes := GenerateDiscountCode{}
	if err := render.Decode(r, &postGenerateCodes); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)

		return
	}

	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	if !canGenerateCodes(user, brandID) {
		httperr.Unauthorized("invalid-permissions", nil, w, r)

		return
	}

	cmd := command.GenerateDiscountCodes{
		BrandID:       brandID,
		DesiredAmount: postGenerateCodes.Amount,
	}

	err = h.app.Commands.GenerateDiscountCodes.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) FetchCode(w http.ResponseWriter, r *http.Request, brandID string) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	cmd := command.FetchDiscountCode{
		BrandID: brandID,
		UserID:  user.ID,
	}

	dc, err := h.app.Commands.FetchDiscount.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)

		return
	}

	render.Respond(w, r, DiscountCode{
		Brand: dc.BrandID(),
		Code:  dc.Code(),
	})
}

func canGenerateCodes(u auth.User, brandID string) bool {
	if u.Role != "brand" {
		return false
	}

	for _, id := range u.AllowedBrandIDs {
		if id == brandID {
			return true
		}
	}

	return false
}
