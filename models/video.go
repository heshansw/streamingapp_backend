package models

type Video struct {
	VideoId            int    `json:"video_id" gorm:"primary_key"`
	VideoPath          string `json:"video_path"`
	VideoStatus        int    `json:"status"`
	VideoConvertedPath string `json:"video_path_converted"`
}
