package posts

import (
	"testing"
)

type titleTest struct {
	in, title, body string
}

func TestExtractTitle(t *testing.T) {
	tests := []titleTest{
		{`# Hello World
This is my post`, "Hello World", "This is my post"},
		{"No title", "", "No title"},
		{`Not explicit title

It's not explicit.
Yep.`, "", `Not explicit title

It's not explicit.
Yep.`},
		{"# Only a title", "", "# Only a title"},
	}

	for _, test := range tests {
		title, body := ExtractTitle(test.in)
		if title != test.title {
			t.Fatalf("Wanted title '%s', got '%s'", test.title, title)
		}
		if body != test.body {
			t.Fatalf("Wanted body '%s', got '%s'", test.body, body)
		}
	}
}

func TestPostLede(t *testing.T) {
	text := map[string]string{
		"早安。跨出舒適圈，才能前往":                                  "早安。",
		"早安。This is my post. It is great.":               "早安。",
		"Hello. 早安。":                                     "Hello.",
		"Sup? Everyone says punctuation is punctuation.": "Sup?",
		"Humans are humans, and society is full of good and bad actors. Technology, at the most fundamental level, is a neutral tool that can be used by either to meet any ends. ": "Humans are humans, and society is full of good and bad actors.",
		`Online Domino Is Must For Everyone

		Do you want to understand how to play poker online?`: "Online Domino Is Must For Everyone",
		`おはようございます

		私は日本から帰ったばかりです。`: "おはようございます",
		"Hello, we say, おはよう. We say \"good morning\"": "Hello, we say, おはよう.",
	}

	c := 1
	for i, o := range text {
		if s := PostLede(i, true); s != o {
			t.Errorf("#%d: Got '%s' from '%s'; expected '%s'", c, s, i, o)
		}
		c++
	}
}

func TestTruncToWord(t *testing.T) {
	text := map[string]string{
		"Можливо, ми можемо використовувати інтернет-інструменти, щоб виготовити якийсь текст, який би міг бути і на, і в кінцевому підсумку, буде скорочено, тому що це тривало так довго.": "Можливо, ми можемо використовувати інтернет-інструменти, щоб виготовити якийсь",
		"早安。This is my post. It is great. It is a long post that is great that is a post that is great.":                                                                     "早安。This is my post. It is great. It is a long post that is great that is a post",
		"Sup? Everyone says punctuation is punctuation.":                                                                                                                     "Sup? Everyone says punctuation is punctuation.",
		"I arrived in Japan six days ago. Tired from a 10-hour flight after a night-long layover in Calgary, I wandered wide-eyed around Narita airport looking for an ATM.": "I arrived in Japan six days ago. Tired from a 10-hour flight after a night-long",
	}

	c := 1
	for i, o := range text {
		if s, _ := TruncToWord(i, 80); s != o {
			t.Errorf("#%d: Got '%s' from '%s'; expected '%s'", c, s, i, o)
		}
		c++
	}
}
