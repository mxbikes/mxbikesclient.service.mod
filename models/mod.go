package models

import (
	protobuffer "github.com/mxbikes/protobuf/mod"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Mod struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" validate:"omitempty,uuid4"`
	Name        string `gorm:"type:varchar(50);not null;default:null;unique" validate:"min=1,max=50"`
	Description string `gorm:"type:varchar(250);not null;default:null" validate:"min=1,max=250"`
	//ModTypeCategoryID string 	`gorm:"type:uuid;default:uuid_generate_v4()" validate:"omitempty,uuid4"`

	ReleaseYear int16 `gorm:"type:smallint;not null;default:null" validate:"min=1950"`
}

func ModToProto(mod *Mod) *protobuffer.Mod {
	return &protobuffer.Mod{
		ID:          mod.ID,
		Name:        mod.Name,
		Description: mod.Description,
		//ModTypeCategoryID: mod.ModTypeCategoryID.String(),
		ReleaseYear: int32(mod.ReleaseYear),
		Create_At:   timestamppb.New(mod.CreatedAt),
	}
}

func ModsToProto(mods []*Mod) []*protobuffer.Mod {
	projections := make([]*protobuffer.Mod, 0, len(mods))
	for _, projection := range mods {
		projections = append(projections, ModToProto(projection))
	}
	return projections
}
