# test
test
# test

# Date Time (ISO 8601)

This project represents date and time values using the
[ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard.

## Examples

| Meaning          | ISO 8601                    |
|------------------|-----------------------------|
| Date             | `2026-08-21`                |
| Date and time    | `2026-08-21T18:47:27Z`      |
| Date and time + offset | `2026-08-21T18:47:27+02:00` |
| Week date        | `2026-W34-5`                |
| Ordinal date     | `2026-233`                  |
| Time only        | `18:47:27`                  |
| Date and time + fraction | `2026-08-21T18:47:27.123Z` |
| Basic format date-time | `20260821T184727Z` |
| Leap second | `2026-06-30T23:59:60Z` |
| Expanded year | `+002026-08-21` |

## Format

- Dates use the format `YYYY-MM-DD`.
- Combined date-time strings use `YYYY-MM-DDTHH:MM:SS` with an optional timezone
  designator (`Z` for UTC or an offset like `+02:00`).
- A decimal fraction of a second may follow the seconds, separated by `.`
  (e.g. `2026-08-21T18:47:27.123Z`).
- Times are expressed in 24-hour clock format; time-only values use
  `HH:MM:SS`.
- Week dates use the format `YYYY-Www-D`, where `Www` is the ISO week number
  (`01`–`53`) and `D` is the weekday number (`1`–`7`, Monday through Sunday).
- Ordinal dates use the format `YYYY-DDD`, where `DDD` is the day of the year
  (`001`–`366`).
- The examples above use the extended format (components separated by `-`, `:`,
  and `T`). The basic format omits all separators (e.g. `20260821T184727Z`);
  do not mix the two in a single value.
- The `T` between date and time may be replaced by a space, but only by mutual
  agreement between producer and consumer; always emit `T`.
- Leap seconds are represented as `23:59:60` within the day they occur
  (e.g. `2026-06-30T23:59:60Z`).
- Years outside `0000`–`9999` require an expanded representation with an
  explicit sign and six digits (e.g. `+002026-08-21`).

## Notes

- Always use UTC (`Z`) for storing and exchanging timestamps.
- Convert to local time only at the presentation layer.
