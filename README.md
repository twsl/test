# test
test
# test

# Date Time (ISO 8601)

This project represents date and time values using the ISO 8601 standard.

## Examples

| Meaning          | ISO 8601                    |
|------------------|-----------------------------|
| Date             | `2026-08-21`                |
| Date and time    | `2026-08-21T18:47:27Z`      |
| Date and time + offset | `2026-08-21T18:47:27+02:00` |
| Week date        | `2026-W34-5`                |
| Ordinal date     | `2026-233`                  |

## Format

- Dates use the format `YYYY-MM-DD`.
- Combined date-time strings use `YYYY-MM-DDTHH:MM:SS` with an optional timezone
  designator (`Z` for UTC or an offset like `+02:00`).
- Times are expressed in 24-hour clock format.

## Notes

- Always use UTC (`Z`) for storing and exchanging timestamps.
- Convert to local time only at the presentation layer.
