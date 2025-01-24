package errors

import (
	"fmt"
	"log"
	"net/http"
)

const help = "Please make a new issue if you think this is a bug: https://lichess-tui/issues."

// Handles the request response. If the request response status
// code is >= 400 and < 600, then the application crashes with
// a unique error screen message for each status code and writes
// a log message with the contents of the request and the contents
// of the response gotten.
//
// Otherwise the function returns 0.
//
// Not every status code is supported. If the function hits an
// unsupported status code, it returns 1.
//
// Supported status codes:
//    400 - bad request.
//    401 - unauthorized.
//    404 - not found.
//    408 - request timeout.
//    418 - teapot.
//    429 - too many requests.
//    500 - internal server error.
//    501 - not implemented.
//    502 - bad gateway.
func HandleRequestResponse(req *http.Request, res *http.Response, err error) int {
  var status = res.StatusCode;
  if status < 400 && err == nil{
    return 0
  }

  switch status {
  case http.StatusBadRequest: {
    fmt.Println("  _  _    ___   ___    ____            _                                  _ ")
    fmt.Println(" | || |  / _ \\ / _ \\  | __ )  __ _  __| |  _ __ ___  __ _ _   _  ___  ___| |_ ")
    fmt.Println(" | || |_| | | | | | | |  _ \\ / _` |/ _` | | '__/ _ \\/ _` | | | |/ _ \\/ __| __|")
    fmt.Println(" |__   _| |_| | |_| | | |_) | (_| | (_| | | | |  __/ (_| | |_| |  __/\\__ \\ |_ ")
    fmt.Println("    |_|  \\___/ \\___/  |____/ \\__,_|\\__,_| |_|  \\___|\\__, |\\__,_|\\___||___/\\__|")
    fmt.Println("                                                       |_|                    ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusUnauthorized: {
    fmt.Println("  _  _    ___  _   _   _                   _   _                _             _ ")
    fmt.Println(" | || |  / _ \\/ | | | | |_ __   __ _ _   _| |_| |__   ___  _ __(_)_______  __| |")
    fmt.Println(" | || |_| | | | | | | | | '_ \\ / _` | | | | __| '_ \\ / _ \\| '__| |_  / _ \\/ _` |")
    fmt.Println(" |__   _| |_| | | | |_| | | | | (_| | |_| | |_| | | | (_) | |  | |/ /  __/ (_| |")
    fmt.Println("    |_|  \\___/|_|  \\___/|_| |_|\\__,_|\\__,_|\\__|_| |_|\\___/|_|  |_/___\\___|\\__,_|")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusNotFound: {
    fmt.Println("  _  _    ___  _  _    __        ___                                                                          _           ___ ")
    fmt.Println(" | || |  / _ \\| || |   \\ \\      / / |__   ___     __ _ _ __ ___   _   _  ___  _   _  __      ____ _ _ __ _ __(_) ___  _ _|__ \\")
    fmt.Println(" | || |_| | | | || |_   \\ \\ /\\ / /| '_ \\ / _ \\   / _` | '__/ _ \\ | | | |/ _ \\| | | | \\ \\ /\\ / / _` | '__| '__| |/ _ \\| '__|/ /")
    fmt.Println(" |__   _| |_| |__   _|   \\ V  V / | | | | (_) | | (_| | | |  __/ | |_| | (_) | |_| |  \\ V  V / (_| | |  | |  | | (_) | |  |_| ")
    fmt.Println("    |_|  \\___/   |_|      \\_/\\_/  |_| |_|\\___/   \\__,_|_|  \\___|  \\__, |\\___/ \\__,_|   \\_/\\_/ \\__,_|_|  |_|  |_|\\___/|_|  (_) ")
    fmt.Println("                                                                  |___/                                                       ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusRequestTimeout: {
    fmt.Println("  _  _    ___   ___    _____ _                    _               _   ")
    fmt.Println(" | || |  / _ \\ ( _ )  |_   _(_)_ __ ___   ___  __| |   ___  _   _| |_ ")
    fmt.Println(" | || |_| | | |/ _ \\    | | | | '_ ` _ \\ / _ \\/ _` |  / _ \\| | | | __|")
    fmt.Println(" |__   _| |_| | (_) |   | | | | | | | | |  __/ (_| | | (_) | |_| | |_ ")
    fmt.Println("    |_|  \\___/ \\___/    |_| |_|_| |_| |_|\\___|\\__,_|  \\___/ \\__,_|\\__|")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  // Lol, you must do really hard job to achieve
  // this one.
  case http.StatusTeapot: {
    fmt.Println("  _  _   _  ___    ___ _   _       _               _   _                    _             _ _  ___ ")
    fmt.Println(" | || | / |( _ )  |_ _| |_( )___  | |_ ___  __ _  | |_(_)_ __ ___   ___    (_)_ __  _ __ (_) ||__ \\")
    fmt.Println(" | || |_| |/ _ \\   | || __|// __| | __/ _ \\/ _` | | __| | '_ ` _ \\ / _ \\   | | '_ \\| '_ \\| | __|/ /")
    fmt.Println(" |__   _| | (_) |  | || |_  \\__ \\ | ||  __/ (_| | | |_| | | | | | |  __/_  | | | | | | | | | |_|_| ")
    fmt.Println("    |_| |_|\\___/  |___|\\__| |___/  \\__\\___|\\__,_|  \\__|_|_| |_| |_|\\___( ) |_|_| |_|_| |_|_|\\__(_) ")
    fmt.Println("                                                                       |/                          ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusTooManyRequests: {
    fmt.Println("  _  _  ____   ___    _____                                                                          _       ")
    fmt.Println(" | || ||___ \\ / _ \\  |_   _|__   ___    _ __ ___   __ _ _ __  _   _   _ __ ___  __ _ _   _  ___  ___| |_ ___ ")
    fmt.Println(" | || |_ __) | (_) |   | |/ _ \\ / _ \\  | '_ ` _ \\ / _` | '_ \\| | | | | '__/ _ \\/ _` | | | |/ _ \\/ __| __/ __|")
    fmt.Println(" |__   _/ __/ \\__, |   | | (_) | (_) | | | | | | | (_| | | | | |_| | | | |  __/ (_| | |_| |  __/\\__ \\ |_\\__ \\")
    fmt.Println("    |_||_____|  /_/    |_|\\___/ \\___/  |_| |_| |_|\\__,_|_| |_|\\__, | |_|  \\___|\\__, |\\__,_|\\___||___/\\__|___/")
    fmt.Println("                                                              |___/               |_|                        ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }

  case http.StatusInternalServerError: {
    fmt.Println("  ____   ___   ___    ___       _                        _                                                           ")
    fmt.Println(" | ___| / _ \\ / _ \\  |_ _|_ __ | |_ ___ _ __ _ __   __ _| |  ___  ___ _ ____   _____ _ __    ___ _ __ _ __ ___  _ __ ")
    fmt.Println(" |___ \\| | | | | | |  | || '_ \\| __/ _ \\ '__| '_ \\ / _` | | / __|/ _ \\ '__\\ \\ / / _ \\ '__|  / _ \\ '__| '__/ _ \\| '__|")
    fmt.Println("  ___) | |_| | |_| |  | || | | | ||  __/ |  | | | | (_| | | \\__ \\  __/ |   \\ V /  __/ |    |  __/ |  | | | (_) | |   ")
    fmt.Println(" |____/ \\___/ \\___/  |___|_| |_|\\__\\___|_|  |_| |_|\\__,_|_| |___/\\___|_|    \\_/ \\___|_|     \\___|_|  |_|  \\___/|_|   ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusNotImplemented: {
    fmt.Println("  ____   ___  _   _   _       _     _                 _                           _           _ ")
    fmt.Println(" | ___| / _ \\/ | | \\ | | ___ | |_  (_)_ __ ___  _ __ | | ___ _ __ ___   ___ _ __ | |_ ___  __| |")
    fmt.Println(" |___ \\| | | | | |  \\| |/ _ \\| __| | | '_ ` _ \\| '_ \\| |/ _ \\ '_ ` _ \\ / _ \\ '_ \\| __/ _ \\/ _` |")
    fmt.Println("  ___) | |_| | | | |\\  | (_) | |_  | | | | | | | |_) | |  __/ | | | | |  __/ | | | ||  __/ (_| |")
    fmt.Println(" |____/ \\___/|_| |_| \\_|\\___/ \\__| |_|_| |_| |_| .__/|_|\\___|_| |_| |_|\\___|_| |_|\\__\\___|\\__,_|")
    fmt.Println("                                               |_|                                              ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  case http.StatusBadGateway: {
    fmt.Println("  ____   ___ ____    ____            _               _                           ")
    fmt.Println(" | ___| / _ \\___ \\  | __ )  __ _  __| |   __ _  __ _| |_ _____      ____ _ _   _ ")
    fmt.Println(" |___ \\| | | |__) | |  _ \\ / _` |/ _` |  / _` |/ _` | __/ _ \\ \\ /\\ / / _` | | | |")
    fmt.Println("  ___) | |_| / __/  | |_) | (_| | (_| | | (_| | (_| | ||  __/\\ V  V / (_| | |_| |")
    fmt.Println(" |____/ \\___/_____| |____/ \\__,_|\\__,_|  \\__, |\\__,_|\\__\\___| \\_/\\_/ \\__,_|\\__, |")
    fmt.Println("                                         |___/                             |___/ ")

    fmt.Println(help)
    log.Fatalf("The request %v got response %v\nError: %v", req, res, err)
  }
  }
  // I couldn't trick the compiler to accept the default
  // case return of 1, so I had to do it here.
  return 1
}

func RequestError(err error) {
	fmt.Println(help)
	log.Fatalf("Error making request: %v", err)
}
