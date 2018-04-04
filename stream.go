package main

type Stream struct {
	
	id            string
	user_id		  string
	game_id       string
	language      string
	pagination    string
	started_at    string
	thumbnail_url string
	title         string
	stream_type   string "json:type"
	viewer_count  int
	live_type     string

}

type Streams []Stream
