package emojis

type Emoji struct {
	Emoji   string
	Code    string
	Unicode string
}

var (
	Smile         = Emoji{Emoji: "😊", Code: ":smile:", Unicode: ""}
	Heart         = Emoji{Emoji: "❤️", Code: ":heart:", Unicode: ""}
	Candy         = Emoji{Emoji: "🍬", Code: ":candy:", Unicode: ""}
	Lollipop      = Emoji{Emoji: "🍭", Code: ":lollipop:", Unicode: ""}
	Cupcake       = Emoji{Emoji: "🧁", Code: ":cupcake:", Unicode: ""}
	Bingsoo       = Emoji{Emoji: "🍧", Code: ":shaved_ice:", Unicode: ""}
	ShavedIce     = Emoji{Emoji: "🍧", Code: ":shaved_ice:", Unicode: ""}
	FireHeart     = Emoji{Emoji: "❤️‍🔥", Code: ":fire_heart:", Unicode: ""}
	Duck          = Emoji{Emoji: "🦆", Code: ":duck:", Unicode: ""}
	Chick         = Emoji{Emoji: "🐥", Code: ":chick:", Unicode: ""}
	HatchingChick = Emoji{Emoji: "🐣", Code: ":hatching_chick:", Unicode: ""}
)
