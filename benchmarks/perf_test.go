package benchmarks_test

import (
	"os"
	"strconv"
	"testing"

	// "go.uber.org/zap".

	vksdkApi "github.com/SevereCloud/vksdk/api"
	vkgoApi "github.com/botscommunity/vkgo/api"
)

func Benchmark_VKGO(b *testing.B) {
	groupID, err := strconv.Atoi(os.Getenv("GROUP_ID"))
	if err != nil {
		b.Fatal(err)
	}

	bot, err := vkgoApi.NewBot(os.Getenv("TOKEN"), groupID)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("getUser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if user := bot.GetUser(542439242); user.Error != nil {
				b.Fatal(user.Error.Message)
			}
		}
	})
}

func Benchmark_VKSDK(b *testing.B) {
	bot := vksdkApi.NewVK(os.Getenv("TOKEN"))

	b.Run("getUser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := bot.UsersGet(vksdkApi.Params{
				"user_ids": 542439242,
			}); err != nil {
				b.Fatal(err)
			}
		}
	})
}
