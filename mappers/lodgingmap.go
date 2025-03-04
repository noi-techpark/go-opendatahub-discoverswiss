// SPDX-FileCopyrightText: 2024 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package mappers

import (
	"fmt"
	"strconv"

	"github.com/noi-techpark/go-opendatahub-discoverswiss/models"
)

// try to use regex, since something is not working ....
func MapAdditionalTypeToAccoTypeId(value string) string {   
    if value == "Hotel" || value == "Pension" {
        return "HotelPension"
    } else if value == "" {
        return "Notdefined"
    } else if value == "ServicedApartments" || value == "HolidayApartment" || value == "GroupAccommodation" {
        return "Apartment"
    } else if value == "BedAndBreakfast" || value == "HolidayHouse" || value == "GuestHouse" || value == "PrivateRoom" {
		return "BedBreakfast"  
	} else if value == "Hostel" {
		return "Youth"
	} else if value == "Campground" {
		return "Camping"
	} else if value == "Mountainhut" {
		return "Mountain"
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


	for _, room := range lb.NumberOfRooms {
		value,err := strconv.Atoi(room.Value)
		if err != nil {
			fmt.Println("Error converting room value to int")
			continue
		}

		switch room.PropertyID {
		case "total":
			acco.AccoOverview.TotalRooms = &value
		case "single":
			acco.AccoOverview.SingleRooms = &value
		case "double":
			acco.AccoOverview.DoubleRooms = &value
		}
	}

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
	
	acco.AccoTypeId = MapAdditionalTypeToAccoTypeId(lb.AdditionalType)

	return acco
}