package main

import (
        "google.golang.org/appengine"
        "fmt"
        "time"
        "log"
        "strconv"
        "os"
      
)


func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
    
        appengine.Main()
    start := time.Now()

    fmt.Printf("1\n")
    fmt.Printf("2\n")
    fmt.Printf("3\n")
    fmt.Printf("4\n")

        const name, age = "Kim", 22
    n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

    // The n and err return values from Fprintln are
    // those returned by the underlying io.Writer.
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
    }
    fmt.Println(n, "bytes written.")

    elapsed := time.Since(start)
    log.Printf("Process took %s", elapsed)

    go forever()
    select {} // block forever

}


func stringToInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    return i
}

func forever() {
    for {
        fmt.Printf("%v+\n", time.Now())
        time.Sleep(time.Second)
    }
}

func stringToFloat64(str string) float64 {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    return f
}


func Add(num1, num2 int) int {
        return num1 + num2
}

func Subtract(num1, num2 int) int {
        return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
        return num1 * num2
}

func Divide(num1, num2 float64) float64 {
        return num1 / num2
}

/*

func main() {
    p := fmt.Println
//We’ll start by getting the current time.

    now := time.Now()
    p(now)
//You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)
//You can extract the various components of the time value as expected.

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
//The Monday-Sunday Weekday is also available.

    p(then.Weekday())
//These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
//The Sub methods returns a Duration representing the interval between two times.

    diff := now.Sub(then)
    p(diff)
//We can compute the length of the duration in various units.

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
//You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.

    p(then.Add(diff))
    p(then.Add(-diff))
}
*/

