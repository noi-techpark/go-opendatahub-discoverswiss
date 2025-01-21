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
}

	

type StarRating struct {
	RatingValue    float64 `json:"ratingValue"`
	AdditionalType string  `json:"additionalType"`
	Name           string  `json:"name"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
