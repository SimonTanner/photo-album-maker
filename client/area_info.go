package client

type AreaInfo struct {
	Data []struct {
		ResultType   string `json:"result_type"`
		ResultObject struct {
			LocationID     string `json:"location_id"`
			Name           string `json:"name"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			NumReviews     string `json:"num_reviews"`
			Timezone       string `json:"timezone"`
			LocationString string `json:"location_string"`
			Photo          struct {
				Images struct {
					Small struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"small"`
					Thumbnail struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"thumbnail"`
					Original struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"original"`
					Large struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"large"`
					Medium struct {
						Width  string `json:"width"`
						URL    string `json:"url"`
						Height string `json:"height"`
					} `json:"medium"`
				} `json:"images"`
				IsBlessed     bool        `json:"is_blessed"`
				UploadedDate  string      `json:"uploaded_date"`
				Caption       string      `json:"caption"`
				ID            string      `json:"id"`
				HelpfulVotes  string      `json:"helpful_votes"`
				PublishedDate string      `json:"published_date"`
				User          interface{} `json:"user"`
			} `json:"photo"`
			Awards []struct {
				AwardType string `json:"award_type"`
				Year      string `json:"year"`
				Images    struct {
					Small string `json:"small"`
					Large string `json:"large"`
				} `json:"images"`
				Categories  []string `json:"categories"`
				DisplayName string   `json:"display_name"`
			} `json:"awards"`
			DoubleclickZone    string `json:"doubleclick_zone"`
			PreferredMapEngine string `json:"preferred_map_engine"`
			GeoType            string `json:"geo_type"`
			CategoryCounts     struct {
				Attractions struct {
					Activities  string `json:"activities"`
					Attractions string `json:"attractions"`
					Nightlife   string `json:"nightlife"`
					Shopping    string `json:"shopping"`
					Total       string `json:"total"`
				} `json:"attractions"`
				Restaurants struct {
					Total string `json:"total"`
				} `json:"restaurants"`
				Accommodations struct {
					Hotels  string `json:"hotels"`
					BbsInns string `json:"bbs_inns"`
					Others  string `json:"others"`
					Total   string `json:"total"`
				} `json:"accommodations"`
				Neighborhoods string `json:"neighborhoods"`
				Airports      string `json:"airports"`
			} `json:"category_counts"`
			NearbyAttractions      []interface{} `json:"nearby_attractions"`
			Description            string        `json:"description"`
			IsLocalizedDescription bool          `json:"is_localized_description"`
			WebURL                 string        `json:"web_url"`
			Ancestors              []struct {
				Subcategory []struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"subcategory"`
				Name       string `json:"name"`
				Abbrv      string `json:"abbrv"`
				LocationID string `json:"location_id"`
			} `json:"ancestors"`
			Category struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"category"`
			Subcategory []struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"subcategory"`
			IsJfyEnabled           bool          `json:"is_jfy_enabled"`
			NearestMetroStation    []interface{} `json:"nearest_metro_station"`
			GeoDescription         string        `json:"geo_description"`
			HasRestaurantCoverpage bool          `json:"has_restaurant_coverpage"`
			HasAttractionCoverpage bool          `json:"has_attraction_coverpage"`
			HasCuratedShoppingList bool          `json:"has_curated_shopping_list"`
		} `json:"result_object"`
		Scope       string `json:"scope"`
		IsTopResult bool   `json:"is_top_result"`
	} `json:"data"`
	Metadata struct {
		Scope string `json:"scope"`
	} `json:"metadata"`
	Sort []struct {
		FilterKey              string `json:"filter_key"`
		Label                  string `json:"label"`
		LocaleIndependentLabel string `json:"locale_independent_label"`
		Selected               bool   `json:"selected"`
	} `json:"sort"`
	PartialContent bool `json:"partial_content"`
	Tracking       struct {
		SearchID string `json:"search_id"`
	} `json:"tracking"`
	Paging struct {
		Results      string `json:"results"`
		TotalResults string `json:"total_results"`
	} `json:"paging"`
}
