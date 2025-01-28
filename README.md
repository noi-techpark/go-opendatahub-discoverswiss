<!--
SPDX-FileCopyrightText: 2024 NOI Techpark <digital@noi.bz.it>

SPDX-License-Identifier: CC0-1.0
-->
# OpenDataHub DiscoverSwiss Mapper

## Overview

This repository contains the mapping structures and functions for transforming **DiscoverSwiss** data into the **OpenDataHub** content API ontology. 

## Purpose

The primary purpose of this repository is to:
- Maintain mapping functions between DiscoverSwiss and OpenDataHub data structures
- Define Go structs that represent the data models
- Provide reusable mapping packages for the collector and transformers implementation

## Structure

```
.
├── models/            # Data structures for both source and target schemas
│   |     
│   └── accomodation.go   
│       
├── mappers/          # Mapping functions
│   ├── lodgingmap.go
│   └── new_mapper.go
│
├── utilities/          # Utility functions to interact with content api
│   |
│   └── contentapi.go
├── go.mod
├── go.sum
└── README.md

```
## Usage

Import the required mapper packages in your collector implementation:

```go
import (
    "github.com/noi-techpark/go-opendatahub-discoverwiss/mappers"
    "github.com/noi-techpark/go-opendatahub-discoverwiss/models"
)
```

Example mapping usage:

```go
// Create a model
mapper := models.Accommodation{}

// Map DiscoverSwiss data to ODH format
odhData, err := mapper.MapAccommodation(discoverSwissData)
if err != nil {
    log.Fatal(err)
}
```

## Related Projects

- [OpenDataHub Collectors Repository](https://github.com/noi-techpark/opendatahub-collectors)

## Contact

OpenDataHub Team - help@opendatahub.com


