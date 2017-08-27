package models

import (
	"fmt"
	"path"
	"time"

	"github.com/astaxie/beego"

	vlc "github.com/jteeuwen/go-vlc"
)

type AudioPlayer struct {
	instance *vlc.Instance
	player   *vlc.Player
	media    *vlc.Media
	evt      *vlc.EventManager
	err      error
}

type AudioPlayerResponse struct {
	Music  string
	Status string
}

func (ap *AudioPlayer) Play(file string) (*AudioPlayerResponse, error) {
	// Load the VLC engine. 
	if ap.instance, ap.err = vlc.New([]string{"--no-video"}); ap.err != nil {
		//fmt.Fprintf(os.Stderr, "[e] New(): %v", ap.err)
		beego.Critical(ap.err)
		return nil, ap.err
	}

	//defer ap.instance.Release()

	// Create a new media item from a file.
	if ap.media, ap.err = ap.instance.OpenMediaFile(path.Join(beego.AppConfig.String("soundrepo"), file)); ap.err != nil {
		//fmt.Fprintf(os.Stderr, "[e] OpenMediaFile(): %v", ap.err)
		beego.Critical(ap.err)
		return nil, ap.err
	}
	// Create a player for the created media.
	if ap.player, ap.err = ap.media.NewPlayer(); ap.err != nil {
		//fmt.Fprintf(os.Stderr, "[e] NewPlayer(): %v", ap.err)
		ap.media.Release()
		beego.Critical(ap.err)
		return nil, ap.err
	}

	//defer ap.player.Release()

	// We don't need the media anymore, now that we have the player.
	ap.media.Release()
	ap.media = nil

	// get an event manager for our player.
	if ap.evt, ap.err = ap.player.Events(); ap.err != nil {
		//fmt.Fprintf(os.Stderr, "[e] Events(): %v", ap.err)
		beego.Critical(ap.err)
		return nil, ap.err
	}

	// Be notified when the player stops playing.
	// This is just to demonstrate usage of event callbacks.
	//ap.evt.Attach(vlc.MediaPlayerStopped, handler, "wahey!")

	// Play the video.
	ap.err = ap.player.Play()
	if ap.err != nil {
		beego.Critical(ap.err)
		return nil, ap.err
	}

	// Wait some amount of time for the media to start playing
	// TODO: Implement proper callbacks for getting the state of the media
	time.Sleep(1 * time.Second)

	// If the media played is a live stream the length will be 0
	length, err := ap.player.Length()
	beego.Debug("Music length: ", length)
	if err != nil || length == 0 {
		length = 1000 * 60
	}

	go func(ap *AudioPlayer, length int64) {
		time.Sleep(time.Duration(length) * time.Millisecond)
		// Stop and free the media player
		ap.player.Stop()
		ap.player.Release()
		ap.instance.Release()
		TurnOnAllArduinos()
	}(ap, length)

	// Give the player 10 seconds of play time.
	//time.Sleep(10 * time.Second)

	// Stop playing.
	//ap.player.Stop()

	return &AudioPlayerResponse{Status: "Playing", Music: file}, nil
}

func handler(evt *vlc.Event, data interface{}) {
	fmt.Printf("[i] %s occurred: %s\n", evt.Type, data.(string))
}

