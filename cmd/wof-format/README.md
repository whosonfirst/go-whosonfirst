## wof-format

Formats a Who's On First GeoJSON feature or checks its formatting.

```
$> wof-format -h
Usage: wof-format [-check] [input]

Either provide a WOF feature on stdin:
cat input.geojson | wof-format > output.geojson

Or provide a path to a WOF feature as the first argument:
wof-format input.geojson > output.geojson

Optional arguments:

  -check
    	exits silently with a non-zero status code if the file is not correctly formatted
  -output string
    	write the output to a file instead of stdout
  -overwrite
    	overwrite the input file with the formatted output
```
