package model

type Region struct {
	Name string `json:"name"`
	//Url			string		`json:"url"`
	LocId int `json:"lid"`
	//Counts		float64		`json:"cnt"`
}

type RegionResponse struct {
	Result RegionResponseLocations `json:"result"`
}

type RegionResponseLocations struct {
	Locations []*RegionResponseLocation `json:"locations"`
}

type RegionResponseLocation struct {
	Id    int               `json:"id"`
	Names map[string]string `json:"names"`
}

func (regionResp *RegionResponse) ToRegionSlice() []*Region {
	regionSlice := make([]*Region, 0, len(regionResp.Result.Locations))

	for _, loc := range regionResp.Result.Locations {
		region := Region{}
		region.Name = loc.Names["1"]
		region.LocId = loc.Id
		regionSlice = append(regionSlice, &region)
	}

	return regionSlice
}

type RegionExt struct {
	Name   string  `json:"name"`
	Url    string  `json:"url"`
	LocId  int     `json:"lid"`
	Counts float64 `json:"cnt"`
}
