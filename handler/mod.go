package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gogo/status"
	"github.com/mxbikes/mxbikesclient.service.mod/models"
	"github.com/mxbikes/mxbikesclient.service.mod/repository"
	protobuffer "github.com/mxbikes/protobuf/mod"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type Mod struct {
	protobuffer.UnimplementedModServiceServer
	repository repository.ModRepository
	logger     logrus.Logger
	validate   *validator.Validate
}

// Return a new handler
func New(postgres repository.ModRepository, logger logrus.Logger) *Mod {
	return &Mod{repository: postgres, validate: validator.New(), logger: logger}
}

func (e *Mod) GetModByID(ctx context.Context, req *protobuffer.GetModByIDRequest) (*protobuffer.GetModByIDResponse, error) {
	mod := e.repository.GetModByID(req.ID)
	if mod != nil {
		return nil, status.Error(codes.NotFound, "Error requested ModID, is not found!")
	}

	e.logger.WithFields(logrus.Fields{"prefix": "SERVICE.Mod_GetModByID"}).Infof("mod with id: {%s} ", req.ID)

	return &protobuffer.GetModByIDResponse{Mod: models.ModToProto(mod)}, nil
}

func (e *Mod) SearchMod(ctx context.Context, req *protobuffer.SearchModRequest) (*protobuffer.SearchModResponse, error) {
	if req.SearchText == "" {
		req.SearchText = "%%"
	}

	pagination := &models.ListQuery{
		Size: int(req.Size),
		Page: int(req.Page),
	}

	result := e.repository.SearchMod(req.SearchText, pagination)

	// items to mods
	items := result.Items.(*[]models.Mod)
	var mods []*models.Mod
	for _, item := range *items {
		var mod = item
		mods = append(mods, &mod)
	}

	e.logger.WithFields(logrus.Fields{"prefix": "SERVICE.Mod_SearchMod"}).Infof("mod like: {%s} ", req.SearchText)

	return &protobuffer.SearchModResponse{Pagination: models.PaginationToProto(result), Mods: models.ModsToProto(mods)}, nil
}
