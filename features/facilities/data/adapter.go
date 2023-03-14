package data

import (
	"alta-airbnb-be/features/facilities"
	_facilityModel "alta-airbnb-be/features/facilities/models"
)

func ConvertToGorm(facilityEntity *facilities.FacilityEntity) _facilityModel.Facility {
	facilityModel := _facilityModel.Facility{
		RoomID: facilityEntity.RoomID,
		Name:   facilityEntity.Name,
	}
	return facilityModel
}

func ConvertToEntity(facilityModel *_facilityModel.Facility) facilities.FacilityEntity {
	facilityEntity := facilities.FacilityEntity{
		RoomID: facilityModel.RoomID,
		Name:   facilityModel.Name,
	}
	return facilityEntity
}
