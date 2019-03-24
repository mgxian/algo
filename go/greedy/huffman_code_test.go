package greedy

import "testing"

func TestHuffCode(t *testing.T) {
	s := "hitomiamfineandyoufinethankyouhellowhatisthwmatteronceuponatimealittlegirltriedtomakealivingbysellingmatchesinthestreetitwasnewyearseveandthesnowedstreetsweredeserted"
	aChars := NewChars(s)
	t.Log(aChars)
}
