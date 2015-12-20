package readableid

import (
    "io"
    "fmt"

    "math/big"
    "crypto/rand"
)

var animals = []string{ "ape", "baboon", "badger", "bat", "bear", "bird", "bobcat", "bulldog", "bullfrog", "cat", "catfish", "cheetah", "chicken", "chipmunk", "cobra", "cougar", "cow", "crab", "deer", "dingo", "dodo", "dog", "dolphin", "donkey", "dragon", "dragonfly", "duck", "eagle", "earwig", "eel", "elephant", "emu", "falcon", "fireant", "firefox", "fish", "fly", "fox", "frog", "gecko", "goat", "goose", "grasshopper", "horse", "hound", "husky", "impala", "insect", "jellyfish", "kangaroo", "ladybug", "liger", "lion", "lionfish", "lizard", "mayfly", "mole", "monkey", "moose", "moth", "mouse", "mule", "newt", "octopus", "otter", "owl", "panda", "panther", "parrot", "penguin", "pig", "puma", "pug", "quail", "rabbit", "rat", "rattlesnake", "robin", "seahorse", "sheep", "shrimp", "skunk", "sloth", "snail", "snake", "squid", "starfish", "stingray", "swan", "termite", "tiger", "treefrog", "turkey", "turtle", "vampirebat", "walrus", "warthog", "wasp", "wolverine", "wombat", "yak", "zebra", }

var adjectives = []string{ "afraid", "ancient", "angry", "average", "bad", "big", "bitter", "black", "blue", "brave", "breezy", "bright", "brown", "calm", "chatty", "chilly", "clever", "cold", "cowardly", "cuddly", "curly", "curvy", "dangerous", "dry", "dull", "empty", "evil", "fast", "fat", "fluffy", "foolish", "fresh", "friendly", "funny", "fuzzy", "gentle", "giant", "good", "great", "green", "grumpy", "happy", "hard", "heavy", "helpless", "honest", "horrible", "hot", "hungry", "itchy", "jolly", "kind", "lazy", "light", "little", "loud", "lovely", "lucky", "massive", "mean", "mighty", "modern", "moody", "nasty", "neat", "nervous", "new", "nice", "odd", "old", "orange", "ordinary", "perfect", "pink", "plastic", "polite", "popular", "pretty", "proud", "purple", "quick", "quiet", "rare", "red", "rotten", "rude", "selfish", "serious", "shaggy", "sharp", "short", "shy", "silent", "silly", "slimy", "slippery", "smart", "smooth", "soft", "sour", "spicy", "splendid", "spotty", "stale", "strange", "strong", "stupid", "sweet", "swift", "tall", "tame", "tasty", "tender", "terrible", "thin", "tidy", "tiny", "tough", "tricky", "ugly", "unlucky", "warm", "weak", "wet", "white", "wicked", "wise", "witty", "wonderful", "yellow", "young", }

// getRandomInt generates a random integer which fulfils the following inequality,
// where n is the generated value.
//
// 0 <= n < max
func getRandomInt(r io.Reader, max int) (int, error) {
    bigMax := big.NewInt(int64(max))

    // probably not a good idea to disregard the error
    bigRandom, error := rand.Int(r, bigMax)

    if error != nil {
        return 0, error
    }

    intRandom := int((*bigRandom).Int64())

    return intRandom, nil
}

// generate picks a random item from a list
func generate(reader io.Reader, list []string) (string, error) {
    index, error := getRandomInt(reader, len(list))

    if error != nil {
        return "", error
    }

    return list[index], nil
}

// GenID generates a random string of the form [adjective]-[noun]-[random int]
func GenID(reader io.Reader, max int) string {
    adjective, _ := generate(reader, adjectives)
    animal, _ := generate(reader, animals)
    number, _ := getRandomInt(reader, max)

    return fmt.Sprintf("%v-%v-%v", adjective, animal, number)
}
