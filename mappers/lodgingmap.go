// SPDX-FileCopyrightText: 2024 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package mappers

import (
	"fmt"
	"strings"

	"github.com/noi-techpark/go-opendatahub-discoverswiss/models"
)


func MapAdditionalTypeToAccoTypeId(additionalType *string) string {
    // Check if additionalType is nil (null)
    if additionalType == nil {
        return "Notdefined"
    }
    
    // Now we can safely dereference and check the value
    value := *additionalType
    
    if strings.EqualFold(value, "Hotel") {
        return "HotelPension"
    } else if strings.EqualFold(value, "SwissLodge") {
        return "SwissLodge"
    } else if strings.EqualFold(value, "") {
        return "Notdefined"
    } else if strings.EqualFold(value, "ServicedApartments") {
        return "Apartment"
    }
    
    return value
}

func MapLodgingBusinessToAccommodation(lb models.LodgingBusiness) models.Accommodation {
	acco := models.Accommodation{
		Source:    "discoverswiss",
		Active:    true,
		Shortname: lb.Name,
	}

	acco.Mapping.DiscoverSwiss.Id = lb.Identifier
	acco.LicenseInfo.Author = ""
	acco.LicenseInfo.License = lb.License	
	acco.LicenseInfo.ClosedData = false
	acco.LicenseInfo.LicenseHolder = "www.discover.swiss"

	acco.GpsInfo = []struct {
		Gpstype              string  `json:"Gpstype"`
		Latitude             float64 `json:"Latitude"`
		Longitude            float64 `json:"Longitude"`
		Altitude             float64 `json:"Altitude"`
		AltitudeUnitofMeasure string `json:"AltitudeUnitofMeasure"`
	}{
		{
			Gpstype:              "position",
			Latitude:             lb.Geo.Latitude,
			Longitude:            lb.Geo.Longitude,
			Altitude:             lb.Geo.Elevation,
			AltitudeUnitofMeasure: "m",
		},
	}

	acco.AccoDetail.Language = models.AccoDetailLanguage{
		Name:        lb.Name,
		Street:      lb.Address.StreetAddress,
		Zip:         lb.Address.PostalCode,
		City:        lb.Address.AddressLocality,
		CountryCode: lb.Address.AddressCountry,
		Email:       lb.Address.Email,
		Phone:       lb.Address.Telephone,
	}

	var totalRooms, singleRooms, doubleRooms int
	for _, room := range lb.NumberOfRooms {
		value := 0
		fmt.Sscanf(room.Value, "%d", &value)

		switch room.PropertyID {
		case "total":
			totalRooms = value
		case "single":
			singleRooms = value
		case "double":
			doubleRooms = value
		}
	}

	acco.AccoOverview.TotalRooms = totalRooms
	acco.AccoOverview.SingleRooms = singleRooms
	acco.AccoOverview.DoubleRooms = doubleRooms
	acco.AccoOverview.CheckInFrom = lb.CheckinTime
	acco.AccoOverview.CheckInTo = lb.CheckinTimeTo
	acco.AccoOverview.CheckOutFrom = lb.CheckoutTimeFrom
	acco.AccoOverview.CheckOutTo = lb.CheckoutTime
	acco.AccoOverview.MaxPersons = lb.NumberOfBeds

	for _,photo := range lb.Photo {
		acco.ImageGallery = append(acco.ImageGallery, models.ImageGalleryItem{
			ImageUrl: photo.ContentUrl, CopyRight: photo.CopyrightNotice,
			ImageDesc: models.LanguageMap{DE: photo.Name, EN: photo.Name, IT: photo.Name},
			ImageName: &photo.Identifier,
			ImageSource: &photo.DataGovernance.Source.Name,
		})
	}
	
	acco.AccoTypeId = MapAdditionalTypeToAccoTypeId(&lb.AdditionalType)

	return acco
}