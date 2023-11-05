package longpoll

import (
	"github.com/botscommunity/vkgo/scene"
)

func (lp *LongPoll) SetTimeout(time int) *LongPoll {
	lp.Timeout = time
	return lp
}

func (lp *LongPoll) SetScenes(scenes *scene.Scenes) *LongPoll {
	lp.Scenes = scenes
	return lp
}
