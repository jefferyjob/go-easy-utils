package emojiUtil

import "fmt"

func ExampleEncodeEmojiUnicode() {
	e := "[\\u1F60E]"
	res := DecodeEmojiUnicode(e)

	fmt.Println(res)

	// Output:
	// 😎
}

func ExampleDecodeEmojiUnicode() {
	e := "😂"
	res := EncodeEmojiUnicode(e)

	fmt.Println(res)

	// Output:
	// [\u1f602]
}
