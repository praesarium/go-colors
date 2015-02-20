package colors

const (

	// foreground colors
	FG_BLACK   = "\x1b[30m"
	FG_RED	   = "\x1b[31m"
	FG_GREEN   = "\x1b[32m"
	FG_YELLOW  = "\x1b[33m"
	FG_BLUE	   = "\x1b[34m"
	FG_MAGENTA = "\x1b[35m"
	FG_CYAN	   = "\x1b[36m"
	FG_WHITE   = "\x1b[37m"

	// background colors
	BG_BLACK   = "\x1b[40m"
	BG_RED	   = "\x1b[41m"
	BG_GREEN   = "\x1b[42m"
	BG_YELLOW  = "\x1b[43m"
	BG_BLUE    = "\x1b[44m"
	BG_MAGENTA = "\x1b[45m"
	BG_CYAN    = "\x1b[46m"
	BG_WHITE   = "\x1b[47m"

	// extra effects
	EFT_RESET 	   = "\x1b[0m"
	EFT_BRIGHT 	   = "\x1b[1m"
	EFT_DIM 	   = "\x1b[2m"
	EFT_UNDERSCORE = "\x1b[4m"
	EFT_BLINK 	   = "\x1b[5m"
	EFT_REVERSE	   = "\x1b[7m"
	EFT_HIDDEN	   = "\x1b[8m"

)

func HttpMethod() string {

	switch {
		case method == "OPTIONS": {
			return FG_WHITE
		}
		case method == "HEAD": 	  {
			return BG_MAGENTA
		}
		case method == "GET":	  {
			return BG_BLUE
		}
		case method == "PUT": 	  {
			return BG_YELLOW
		}
		case method == "POST": 	  {
			return BG_CYAN
		}
		case method == "DELETE":  {
			return BG_RED
		}
		case method == "PATCH":   {
			return BG_GREEN
		}
		default: {
			return EFT_RESET
		}
	}
}

func HttpStatus() string {

	switch {
		case code >= 200 && code <= 299: {
			return FG_GREEN
		}
		case code >= 300 && code <= 399: {
			return FG_CYAN
		}
		case code >= 400 && code <= 499: {
			return FG_YELLOW
		}
		default: {
			return FG_RED
		}
	}
}
