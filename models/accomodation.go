// SPDX-FileCopyrightText: 2024 NOI Techpark <digital@noi.bz.it>
//
// SPDX-License-Identifier: MPL-2.0

package models

type discoverswissId struct{
	Id string	`json:"id"`
}

type Accommodation struct {
	Source     string `default:"discoverswiss"`
	Active     bool   `default:"true"`
	Shortname  string

	Mapping struct {
		DiscoverSwiss discoverswissId `json:"discoverswiss"`
	} `json:"Mapping"`

	AccoDetail struct {
		Language AccoDetailLanguage `json:"de"`
	} `json:"AccoDetail"`

	GpsInfo []struct {
		Gpstype   string  `json:"Gpstype"`
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
		Altitude  float64 `json:"Altitude"`
		AltitudeUnitofMeasure string `json:"AltitudeUnitofMeasure"`
	} `json:"GpsInfo"`

	AccoType struct {
		Id string `json:"Id"`
	} `json:"AccoType"`

	AccoOverview struct {
		TotalRooms   int    `json:"TotalRooms"`
		SingleRooms  int    `json:"SingleRooms"`
		DoubleRooms  int    `json:"DoubleRooms"`
		CheckInFrom  string `json:"CheckInFrom"`
		CheckInTo    string `json:"CheckInTo"`
		CheckOutFrom string `json:"CheckOutFrom"`
		CheckOutTo   string `json:"CheckOutTo"`
		MaxPersons   int    `json:"MaxPersons"`
	} `json:"AccoOverview"`

	LicenseInfo struct {
		Author string `json:"Author"`
		License string `json:"License"`
		ClosedData bool `json:"ClosedData"`
		LicenseHolder string `json:"LicenseHolder"`
	} `json:"LicenseInfo"`

	ImageGallery []ImageGalleryItem `json:"ImageGallery"`
}

type ImageGalleryItem struct {
    ImageUrl    string      `json:"ImageUrl"`     // From ContentUrl
    CopyRight   string      `json:"CopyRight"`    // From CopyrightNotice
    ImageDesc   LanguageMap `json:"ImageDesc"`    // From Name
    ImageName   *string     `json:"ImageName,omitempty"`  // From Identifier
    ImageSource *string     `json:"ImageSource,omitempty"` // From DataGovernance.Source.Name
}

type LanguageMap struct {
    DE string `json:"de,omitempty"`
    EN string `json:"en,omitempty"`
    IT string `json:"it,omitempty"`
}

type AccoDetailLanguage struct {
	Name        string `json:"Name"`
	Street      string `json:"Street"`
	Zip         string `json:"Zip"`
	City        string `json:"City"`
	CountryCode string `json:"CountryCode"`
	Email       string `json:"Email"`
	Phone       string `json:"Phone"`
}


type DiscoverSwissResponse struct {
	Count         int               `json:"count"`
	HasNextPage   bool              `json:"hasNextPage"`
	NextPageToken string            `json:"nextPageToken"`
	Data          []LodgingBusiness `json:"data"`
}
type LodgingBusiness struct {
	Name string `json:"name"`

	Address struct {
		AddressCountry  string `json:"addressCountry"`
		AddressLocality string `json:"addressLocality"`
		PostalCode      string `json:"postalCode"`
		StreetAddress   string `json:"streetAddress"`
		Email           string `json:"email"`
		Telephone       string `json:"telephone"`
	} `json:"address"`

	Geo struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geo"`

	NumberOfRooms []struct {
		PropertyID string `json:"propertyId"`
		Value      string `json:"value"`
	} `json:"numberOfRooms"`

	StarRating StarRating `json:"starRating"`

	NumberOfBeds int `json:"numberOfBeds"`

	Identifier string `json:"identifier"`

	CheckinTime      string `json:"checkinTime"`
	CheckinTimeTo    string `json:"checkinTimeTo"`
	CheckoutTimeFrom string `json:"checkoutTimeFrom"`
	CheckoutTime     string `json:"checkoutTime"`

	License string `json:"license"`

	Photo []Photo `json:"photo"`
}

type Photo struct {
    ContentUrl      string        `json:"contentUrl"`     // Maps to ImageUrl
    CopyrightNotice string        `json:"copyrightNotice"` // Maps to CopyRight
    DataGovernance  DataGovernance `json:"dataGovernance"` // For extracting ImageSource
    Identifier      string        `json:"identifier"`     // Could map to ImageName
    Name            string        `json:"name"`           // Could map to ImageDesc
}

type DataGovernance struct {
    Source Source `json:"source"`
}

type Source struct {
    Name string `json:"name"` // Maps to ImageSource
}
	

type StarRating struct {
	RatingValue    float64 `json:"ratingValue"`
	AdditionalType string  `json:"additionalType"`
	Name           string  `json:"name"`
}

