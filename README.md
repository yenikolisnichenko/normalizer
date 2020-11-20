# CSV normalizer

## What it is

A tool that reads a CSV formatted file on `stdin` and
emits a normalized CSV formatted file on `stdout`.

Normalized, in this case, means:

* The entire CSV is in the UTF-8 character set.
* The `Timestamp` column will be in RFC3339 format and converted from US/Pacific time to US/Eastern.
* All `ZIP` codes should be formatted as 5 digits. If there are less
  than 5 digits, padded by 0 if necessary.
* The `FullName` column will be converted to uppercase. This will be
  able to handle non-English names.
* The `Address` column will be passed through as is.
* The `FooDuration` and `BarDuration` columns will be the
  total number of seconds expressed in floating point format, no rounding will happen.
* The `TotalDuration` column will be the value of `TotalDuration` with the sum of
  `FooDuration` and `BarDuration`.
* The `Notes` column is free form text input by end-users passed as is.

## Assumptions:

* The input document is in UTF-8
* Any times that are missing timezone information are in US/Pacific
* TODO: Any invalid character will be replaced with the Unicode Replacement
Character. If that replacement makes data invalid (for example,
because it turns a date field into something unparseable) a 
warning to `stderr` will be displayed and the row will be dropped from the output

## How to run it

After building binary run:

```sh
./normalizer sample.csv output.csv
```

