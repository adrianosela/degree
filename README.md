# degree

Just a small Go program I wrote to make sure I don't have outstanding UBC Computer Engineering degree requirements.

The `/docs` subdirectory has documentation with which I've produced the requirements CSV file, as well as tagged the courses I took with the specific requirement they are supposed to fulfill 

### Usage:

Note that the file names are hardcoded in the program.

```
$ go run main.go
requirement "oral and written communication" satisfied with [ENGL 112, CPEN 281]
requirement "impact of technology on society" satisfied with [APSC 377]
requirement "humanities and social sciences" satisfied with [GERM 100, PSYC 102]
requirement "free electives" satisfied with [EOSC 110, FRST 231]
requirement "professionalism, ethics, equity, and law" satisfied with [APSC 450]
requirement "engineering economics and project management" satisfied with [CPEN 481]
requirement "science electives" satisfied with [ASTR 200]
requirement "CPEN breadth" satisfied with [ELEC 331, CPEN 321]
requirement "CPEN advanced" satisfied with [CPEN 412, CPEN 442, CPEN 421]
requirement "technical electives" satisfied with [CPSC 322, CPSC 314]
ALL REQUIREMENTS SATISFIED!
```
