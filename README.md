# test
test
# test

## Date and Time

This project uses [ISO 8601](https://www.iso.org/standard/70907.html) for all
date and time values.

### Calendar date

| Format | Example | Meaning |
| --- | --- | --- |
| `YYYY-MM-DD` | `2026-10-26` | Gregorian calendar date |

### Ordinal date

| Format | Example | Meaning |
| --- | --- | --- |
| `YYYY-DDD` | `2026-299` | Day of the year (001 = Jan 1) |

### Week date

| Format | Example | Meaning |
| --- | --- | --- |
| `YYYY-Www` | `2026-W44` | Week number (week 01 is the week containing the first Thursday of the year) |
| `YYYY-Www-D` | `2026-W44-1` | Week date with day (1 = Monday) |

### Time

| Format | Example | Meaning |
| --- | --- | --- |
| `hh` | `14` | Hour |
| `hh:mm` | `14:03` | Hour and minute |
| `hh:mm:ss` | `14:03:54` | Hour, minute, and second |
| `hh:mm:ss.sss` | `14:03:54.123` | Fractional seconds (comma or period separator) |

### Time zone

| Format | Example | Meaning |
| --- | --- | --- |
| `Z` | `2026-10-26T14:03:54Z` | UTC |
| `±hh:mm` | `2026-10-26T14:03:54+02:00` | Offset from UTC |

### Combined date and time

The date and time are separated by `T`.

| Format | Example |
| --- | --- |
| `YYYY-MM-DDThh:mm:ssZ` | `2026-10-26T14:03:54Z` |
| `YYYY-MM-DDThh:mm:ss±hh:mm` | `2026-10-26T14:03:54+02:00` |
| `YYYY-MM-DDThh:mm:ss.sssZ` | `2026-10-26T14:03:54.123Z` |

### Durations and intervals

| Format | Example | Meaning |
| --- | --- | --- |
| `PnYnMnDTnHnMnS` | `P3Y6M4DT12H30M5S` | Duration |
| `PnW` | `P2W` | Duration in weeks |
| `YYYY-MM-DD/YYYY-MM-DD` | `2026-01-01/2026-12-31` | Interval |

### Notes

- Times are written in 24-hour notation; midnight is `00:00:00`.
- A date without a time means the whole day.
- `Z` and numeric offsets are interchangeable; `Z` means UTC.
