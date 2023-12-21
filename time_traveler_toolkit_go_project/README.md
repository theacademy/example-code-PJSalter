# ⏰ Go Time Traveler Toolkit 📆

This Go project provides a comprehensive set of functions to effortlessly handle date and time manipulations using Go's time package. These functions enable users to perform various operations, including calculating dates, determining differences, and navigating through time seamlessly!

## About the Project

Have you ever wished for an easy way to manage dates and times in your code? Look no further! The Go Time Traveler Toolkit empowers developers with powerful date and time functionalities. From basic operations like getting today's date to more complex tasks such as calculating time differences between dates, this toolkit offers an intuitive set of functions for all your temporal needs.

## Key Functions

- `GetTodaysDate()`: Fetches today's date.
- `GetLaterDateByDays(date time.Time, x int)`: Calculates the date x days after the input date.
- `GetPreviousDateByWeeks(date time.Time, x int)`: Determines the date x weeks before the input date.
- `GetTimeDifference(date1, date2 time.Time)`: Computes the time difference between two dates in a user-readable format.

## Usage

Utilize these powerful functions in your Go projects with ease:

### 1. Installation

Ensure Go is installed and import the toolkit package:

```go
import "time_traveler_toolkit/internal/time_travel"
```

### Running the code:

```go
go run main.go

```

### Testing:

```go
go test
```
