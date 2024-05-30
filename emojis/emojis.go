package emojis

type Emoji struct {
	Emoji   string
	Code    string
	Unicode string
}

var (
	Smile         = Emoji{Emoji: "😊", Code: ":smile:", Unicode: ""}
	Heart         = Emoji{Emoji: "❤️", Code: ":heart:", Unicode: ""}
	Candy         = Emoji{Emoji: "🍬", Code: ":candy:", Unicode: "1F367"}
	Lollipop      = Emoji{Emoji: "🍭", Code: ":lollipop:", Unicode: "1F36D"}
	Cupcake       = Emoji{Emoji: "🧁", Code: ":cupcake:", Unicode: "1F9C1"}
	Bingsoo       = Emoji{Emoji: "🍧", Code: ":shaved_ice:", Unicode: "1F367"}
	ShavedIce     = Emoji{Emoji: "🍧", Code: ":shaved_ice:", Unicode: "1F367"}
	FireHeart     = Emoji{Emoji: "❤️‍🔥", Code: ":fire_heart:", Unicode: ""}
	Duck          = Emoji{Emoji: "🦆", Code: ":duck:", Unicode: ""}
	Chick         = Emoji{Emoji: "🐥", Code: ":chick:", Unicode: ""}
	HatchingChick = Emoji{Emoji: "🐣", Code: ":hatching_chick:", Unicode: ""}
)
