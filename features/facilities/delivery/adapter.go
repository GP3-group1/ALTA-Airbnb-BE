package delivery

import (
	"alta-airbnb-be/features/facilities"
	"alta-airbnb-be/features/rooms"
	"strings"
)

func ConvertToEntities(roomRequest *rooms.RoomRequest) (facilityEntities []facilities.FacilityEntity) {
	for _, val := range strings.Split(roomRequest.Facilities, ", ") {
		facilityEntities = append(facilityEntities, facilities.FacilityEntity{
			Name: val,
		})
	}
	return facilityEntities
}
