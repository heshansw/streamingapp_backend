package api

import (
	"backendapi/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/redis/go-redis/v9"
)

func UploadVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, handler, errReq := r.FormFile("file")

	if errReq != nil {
		http.Error(w, "Error Occured", http.StatusBadRequest)
		return
	}

	defer file.Close()

	errMkdir := os.MkdirAll("./uploads", os.ModePerm)
	if errMkdir != nil {
		http.Error(w, "Upload Location Error", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

	dst, errFile := os.Create(filePath)

	if errFile != nil {
		http.Error(w, "Upload Location Error", http.StatusBadRequest)
		return
	}

	defer dst.Close()

	_, errCopy := io.Copy(dst, file)

	if errCopy != nil {
		http.Error(w, "Copy Upload Location Error", http.StatusBadRequest)
		return
	}

	var client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	ctx := context.Background()
	errQueue := client.LPush(ctx, "video-convert", filePath).Err()

	if errQueue != nil {
		http.Error(w, "Queue Error", http.StatusBadRequest)
		return
	}

	videoNew := &models.Video{
		VideoPath:   filePath,
		VideoStatus: 0,
	}

	models.DB.Create(videoNew)

	if videoNew.VideoId > 0 {
		json.NewEncoder(w).Encode(videoNew)
	} else {
		http.Error(w, "Error Occured", http.StatusBadRequest)
	}
}
