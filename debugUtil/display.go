package debugUtil

// 显示代码
type Code int

const (
	Code_Default    Code = 0
	Code_Hightlight Code = 1
	Code_Underline  Code = 4
	Code_Blink      Code = 5
	Code_Reverse    Code = 7
	Code_Invisible  Code = 8
)

// 前景色
type ForegroundColor int

const (
	Foreground_Black  ForegroundColor = 30
	Foreground_Red    ForegroundColor = 31
	Foreground_Green  ForegroundColor = 32
	Foreground_Yellow ForegroundColor = 33
	Foreground_Blue   ForegroundColor = 34
	Foreground_Purple ForegroundColor = 35
	Foreground_Cyan   ForegroundColor = 36
	Foreground_White  ForegroundColor = 37
)

// 背景色
type BackgroundColor int

const (
	BackgroundColor_Black  BackgroundColor = 40
	BackgroundColor_Red    BackgroundColor = 41
	BackgroundColor_Green  BackgroundColor = 42
	BackgroundColor_Yellow BackgroundColor = 43
	BackgroundColor_Blue   BackgroundColor = 44
	BackgroundColor_Purple BackgroundColor = 45
	BackgroundColor_Cyan   BackgroundColor = 46
	BackgroundColor_White  BackgroundColor = 47
)
