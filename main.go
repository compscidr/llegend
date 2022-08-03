package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// https://www.sohamkamani.com/golang/how-to-build-a-web-application/
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.GET("/", handler)
	r.Run() // listen on 0.0.0.0:8080
}

func handler(c *gin.Context) {
	apiKey := os.Getenv("ACCESS_TOKEN")
	athleteID, _ := strconv.Atoi(os.Getenv("ATHLETE"))
	mapsKey := os.Getenv("GMAPS")

	log.Println("Using apiKey " + apiKey + " athleteID: " + strconv.Itoa(athleteID))

	segments := getLLSegments(apiKey, athleteID)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"segments": segments,
		"gmaps":    mapsKey,
	})
}

type SegmentData struct {
	ID          int
	LL          bool
	EffortsAway int // if LL = true, efforts for someone to catch-up, otherwise, efforts to become
}

// got this from getting the response, turning it into a string, and then pasting into here: https://mholt.github.io/json-to-go/
// guide from: https://dev.to/billylkc/parse-json-api-response-in-go-10ng
type StarredSegmentResponse []struct {
	ID               int         `json:"id"`
	ResourceState    int         `json:"resource_state"`
	Name             string      `json:"name"`
	ActivityType     string      `json:"activity_type"`
	Distance         float64     `json:"distance"`
	AverageGrade     float64     `json:"average_grade"`
	MaximumGrade     float64     `json:"maximum_grade"`
	ElevationHigh    float64     `json:"elevation_high"`
	ElevationLow     float64     `json:"elevation_low"`
	StartLatlng      []float64   `json:"start_latlng"`
	EndLatlng        []float64   `json:"end_latlng"`
	ElevationProfile interface{} `json:"elevation_profile"`
	ClimbCategory    int         `json:"climb_category"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	Country          string      `json:"country"`
	Private          bool        `json:"private"`
	Hazardous        bool        `json:"hazardous"`
	Starred          bool        `json:"starred"`
	PrTime           int         `json:"pr_time,omitempty"`
	AthletePrEffort  struct {
		ID             int64     `json:"id"`
		ActivityID     int64     `json:"activity_id"`
		ElapsedTime    int       `json:"elapsed_time"`
		Distance       float64   `json:"distance"`
		StartDate      time.Time `json:"start_date"`
		StartDateLocal time.Time `json:"start_date_local"`
		IsKom          bool      `json:"is_kom"`
	} `json:"athlete_pr_effort,omitempty"`
	StarredDate time.Time `json:"starred_date"`
}

func getLLSegments(apiKey string, atheleteID int) []SegmentData {
	var segments []SegmentData
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/segments/starred?page=1&per_page=2", nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	var results StarredSegmentResponse
	if err := json.Unmarshal(body, &results); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Loop through all the segments and extract their IDs
	for _, result := range results {
		segment := getSegmentData(apiKey, result.ID, atheleteID)
		segments = append(segments, segment)
	}

	return segments
}

type DetailedSegmentResponse []struct {
	Category         string `json:"category"`
	AnalyticsContext struct {
		SegmentID                 int    `json:"segment_id"`
		LeaderboardFilterType     string `json:"leaderboard_filter_type"`
		LocalLegendAthleteID      int    `json:"local_legend_athlete_id"`
		LocalLegendEffortCount    int    `json:"local_legend_effort_count"`
		ViewingAthleteEffortCount int    `json:"viewing_athlete_effort_count"`
		ViewingAthleteBucket      int    `json:"viewing_athlete_bucket"`
		TotalAthleteCount         int    `json:"total_athlete_count"`
		TotalEffortCount          int    `json:"total_effort_count"`
		TotalBuckets              int    `json:"total_buckets"`
	} `json:"analytics_context"`
	WindowText string `json:"window_text"`
	Segment    struct {
		ID            int     `json:"id"`
		Name          string  `json:"name"`
		ActivityType  string  `json:"activity_type"`
		Distance      float64 `json:"distance"`
		AverageGrade  float64 `json:"average_grade"`
		ElevationHigh float64 `json:"elevation_high"`
		ElevationLow  float64 `json:"elevation_low"`
		Map           struct {
			Polyline string      `json:"polyline"`
			LatLng   [][]float64 `json:"lat_lng"`
		} `json:"map"`
	} `json:"segment"`
	ElevationProfile string `json:"elevation_profile"`
	StaticMaps       struct {
		OneX string `json:"1x"`
		TwoX string `json:"2x"`
	} `json:"static_maps"`
	Histogram struct {
		XLabels []struct {
			Label       string `json:"label"`
			BucketIndex int    `json:"bucket_index"`
		} `json:"x_labels"`
		YLabels []struct {
			Label    string  `json:"label"`
			Position float64 `json:"position"`
			DrawLine bool    `json:"draw_line"`
		} `json:"y_labels"`
		AthleteBuckets []struct {
			AthleteID   int `json:"athlete_id"`
			BucketIndex int `json:"bucket_index"`
		} `json:"athlete_buckets"`
		BucketValues []float64 `json:"bucket_values"`
		Buckets      []struct {
			Value       float64 `json:"value"`
			EffortsText struct {
				Text     string `json:"text"`
				Emphasis []struct {
					StartIndex int `json:"start_index"`
					Length     int `json:"length"`
				} `json:"emphasis"`
			} `json:"efforts_text"`
		} `json:"buckets"`
		Footer struct {
			Icon        string `json:"icon"`
			IconColor   string `json:"icon_color"`
			Text        string `json:"text"`
			Destination string `json:"destination"`
		} `json:"footer"`
	} `json:"histogram"`
	LocalLegend struct {
		AthleteID          int    `json:"athlete_id"`
		Title              string `json:"title"`
		Profile            string `json:"profile"`
		BadgeTypeID        int    `json:"badge_type_id"`
		EffortDescription  string `json:"effort_description"`
		MayorEffortCount   int    `json:"mayor_effort_count"`
		ShowSeeYourResults bool   `json:"show_see_your_results"`
		YourEffortsText    struct {
			Text     string `json:"text"`
			Emphasis []struct {
				StartIndex int `json:"start_index"`
				Length     int `json:"length"`
			} `json:"emphasis"`
		} `json:"your_efforts_text"`
	} `json:"local_legend"`
	OverallEfforts struct {
		TotalAthletes string `json:"total_athletes"`
		TotalEfforts  string `json:"total_efforts"`
		TotalDistance string `json:"total_distance"`
	} `json:"overall_efforts"`
	YourEfforts struct {
		EffortDescription string `json:"effort_description"`
		TotalEfforts      string `json:"total_efforts"`
		TotalDistance     string `json:"total_distance"`
	} `json:"your_efforts"`
	Leaderboards []struct {
		Type           string `json:"type"`
		OverallEfforts struct {
			TotalAthletes string `json:"total_athletes"`
			TotalEfforts  string `json:"total_efforts"`
			TotalDistance string `json:"total_distance"`
		} `json:"overall_efforts"`
		Entries []struct {
			Rank           int    `json:"rank"`
			AthleteID      int    `json:"athlete_id"`
			Name           string `json:"name"`
			Profile        string `json:"profile"`
			BadgeTypeID    int    `json:"badge_type_id"`
			EffortCount    int    `json:"effort_count"`
			IsLocalLegend  bool   `json:"is_local_legend"`
			Destination    string `json:"destination"`
			LastEffortText string `json:"last_effort_text"`
		} `json:"entries"`
		Footer struct {
			Icon        string `json:"icon"`
			IconColor   string `json:"icon_color"`
			Text        string `json:"text"`
			Destination string `json:"destination"`
		} `json:"footer"`
		AnalyticsContext struct {
			TopFollowerAthleteID   interface{} `json:"top_follower_athlete_id"`
			TopFollowerEffortCount interface{} `json:"top_follower_effort_count"`
			FollowingAthleteCount  int         `json:"following_athlete_count"`
			FollowingEffortCount   int         `json:"following_effort_count"`
		} `json:"analytics_context"`
	} `json:"leaderboards"`
	OptIn struct {
		AthleteOptedIn bool   `json:"athlete_opted_in"`
		Text           string `json:"text"`
		Icon           string `json:"icon"`
		IconColor      string `json:"icon_color"`
		LearnMore      struct {
			Title string `json:"title"`
			Rows  []struct {
				Title       string `json:"title"`
				Destination string `json:"destination"`
			} `json:"rows"`
		} `json:"learn_more"`
		Privacy struct {
			Title              string `json:"title"`
			ButtonText         string `json:"button_text"`
			Body               string `json:"body"`
			ActionConfirmation struct {
				Title   string `json:"title"`
				Body    string `json:"body"`
				Cancel  string `json:"cancel"`
				Confirm string `json:"confirm"`
				Action  string `json:"action"`
			} `json:"action_confirmation"`
		} `json:"privacy"`
	} `json:"opt_in"`
}

func getSegmentData(apiKey string, segmentID int, athleteID int) SegmentData {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.strava.com/api/v3/segments/"+strconv.Itoa(segmentID)+"/local_legend", nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	var results DetailedSegmentResponse
	if err := json.Unmarshal(body, &results); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	isLL := results[0].LocalLegend.AthleteID == athleteID
	fmt.Println("LL: " + strconv.Itoa(results[0].LocalLegend.AthleteID))
	effortDiff := 0
	if isLL {
		totalEfforts, _ := strconv.Atoi(results[0].YourEfforts.TotalEfforts)
		effortDiff = totalEfforts - results[0].LocalLegend.MayorEffortCount
		fmt.Println("My efforts: " + strconv.Itoa(totalEfforts) + " Best efforts: " + strconv.Itoa(results[0].LocalLegend.MayorEffortCount) + " Diff: " + strconv.Itoa(effortDiff))
	} else {
		totalEfforts, _ := strconv.Atoi(results[0].YourEfforts.TotalEfforts)
		effortDiff = results[0].LocalLegend.MayorEffortCount - totalEfforts
		fmt.Println("My efforts: " + strconv.Itoa(totalEfforts) + " Best efforts: " + strconv.Itoa(results[0].LocalLegend.MayorEffortCount) + " Diff: " + strconv.Itoa(effortDiff))
	}

	return SegmentData{
		segmentID,
		isLL,
		effortDiff,
	}
}
