package main

var Rules = []struct {
	character   string
	replacement string
}{
	{
		character:   " ",
		replacement: "_",
	},
	{
		character:   "/",
		replacement: "_",
	},
	{
		character:   ",",
		replacement: "_",
	},
	{
		character:   "D’",
		replacement: "",
	},
	{
		character:   "D'",
		replacement: "",
	},
	{
		character:   "'",
		replacement: "",
	},
	{
		character:   "’",
		replacement: "",
	},
	{
		character:   "É",
		replacement: "E",
	},
	{
		character:   "È",
		replacement: "E",
	},
	{
		character:   "Ô",
		replacement: "O",
	},
	{
		character:   "&",
		replacement: "",
	},
	{
		character:   "(",
		replacement: "",
	},
	{
		character:   ")",
		replacement: "",
	},
	{
		character:   "__",
		replacement: "_",
	},
	{
		character:   "-",
		replacement: "_",
	},
	{
		character:   ".",
		replacement: "",
	},
}
