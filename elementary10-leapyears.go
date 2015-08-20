package elementary10

import "fmt"
import "time"

func IsLeapYear(year int) bool {
    if year % 4 == 0 {
        if year % 400 == 0 {
            return true
        }
        if year % 100 != 0 {
            return true
        }
    }
    return false
}

func main() {
    curryear := time.Now().Year()
    for year := curryear; year <= curryear + 20; year=year+1 {
        if IsLeapYear(year) {
            fmt.Printf("%v\n", year)
        }
    }
}

