package haikunator

import (
	"fmt"
	"math/rand"
)

var (
	ADJECTIVES = []string{"aged", "ancient", "autumn", "billowing", "bitter", "black", "blue", "bold",
  "broad", "broken", "calm", "cold", "cool", "crimson", "curly", "damp",
  "dark", "dawn", "delicate", "divine", "dry", "empty", "falling", "fancy",
  "flat", "floral", "fragrant", "frosty", "gentle", "green", "hidden", "holy",
  "icy", "jolly", "late", "lingering", "little", "lively", "long", "lucky",
  "misty", "morning", "muddy", "mute", "nameless", "noisy", "odd", "old",
  "orange", "patient", "plain", "polished", "proud", "purple", "quiet", "rapid",
  "raspy", "red", "restless", "rough", "round", "royal", "shiny", "shrill",
  "shy", "silent", "small", "snowy", "soft", "solitary", "sparkling", "spring",
  "square", "steep", "still", "summer", "super", "sweet", "throbbing", "tight",
  "tiny", "twilight", "wandering", "weathered", "white", "wild", "winter", "wispy",
  "withered", "yellow", "young"}
	NOUNS = []string{"art", "band", "bar", "base", "bird", "block", "boat", "bonus",
  "bread", "breeze", "brook", "bush", "butterfly", "cake", "cell", "cherry",
  "cloud", "credit", "darkness", "dawn", "dew", "disk", "dream", "dust",
  "feather", "field", "fire", "firefly", "flower", "fog", "forest", "frog",
  "frost", "glade", "glitter", "grass", "hall", "hat", "haze", "heart",
  "hill", "king", "lab", "lake", "leaf", "limit", "math", "meadow",
  "mode", "moon", "morning", "mountain", "mouse", "mud", "night", "paper",
  "pine", "poetry", "pond", "queen", "rain", "recipe", "resonance", "rice",
  "river", "salad", "scene", "sea", "shadow", "shape", "silence", "sky",
  "smoke", "snow", "snowflake", "sound", "star", "sun", "sunset",
  "surf", "term", "thunder", "tooth", "tree", "truth", "union", "unit",
  "violet", "voice", "water", "waterfall", "wave", "wildflower", "wind", "wood"}
)

type Name interface {
	Haikunate() string
	Size() int
}

type RandomName struct {
	r *rand.Rand
}

func (r RandomName) Haikunate() string {
	return fmt.Sprintf("%v-%v", ADJECTIVES[r.r.Intn(len(ADJECTIVES))], NOUNS[r.r.Intn(len(NOUNS))])
}

func (r RandomName) Size() int {
	return len(ADJECTIVES) * len(NOUNS)
}

func New(seed int64) Name {
	name := RandomName{rand.New(rand.New(rand.NewSource(99)))}
	name.r.Seed(seed)
	return name
}
