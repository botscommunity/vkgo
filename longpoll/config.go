package longpoll

import (
	"github.com/botscommunity/vkgo/scene"
)

func (lp *Longpoll) SetTimeout(time int) *Longpoll {
	lp.Timeout = time
	return lp
}

func (lp *Longpoll) SetScenes(scenes *scene.Scenes) *Longpoll {
	lp.Scenes = scenes
	return lp
}
