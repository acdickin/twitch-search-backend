package main

type Stream struct {
	game_id       string
	id            string
	language      string
	pagination    string
	started_at    string
	thumbnail_url string
	title         string
	stream_type   string "json:type"
	user_id       string
	viewer_count  int
}

type Streams []Stream
