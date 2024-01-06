package internals

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	assets "github.com/peixotoleonardo/tetris/assets/audio"
)

type Audio struct {
	player       *audio.Player
	audioContext *audio.Context
}

func (a *Audio) Update() error {
	if a.player != nil {
		return nil
	}

	if a.audioContext == nil {
		a.audioContext = audio.NewContext(sampleRate)
	}

	oggS, err := vorbis.DecodeWithoutResampling(bytes.NewReader(assets.Music_ogg))

	if err != nil {
		return err
	}

	s := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)
	a.player, err = a.audioContext.NewPlayer(s)

	if err != nil {
		return err
	}

	a.player.Play()

	return nil
}
