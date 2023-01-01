# Replace
Given a text file of names and the "sector" folder from your game install (or a copy of it), this program will replace the names of every sector within that folder with what it finds in the name text file. It expects the names to be line-separated in that file.
## Usage
`replace "names/to/use.txt" "../Astrox Imperium/Astrox Imperium_Data/MOD/sectors/" [flags]`

## Flags
`-h, --help     help for replace`

`-r, --random   Shuffles names randomly if set.`

### This program has no input checking

I recommend you place a copy of the sectors folder and the names.txt in the same place as the executable, and run it that way.

## Build
`go build`
Produces a 'replace.exe,' unless given an output name. This has only been tested on Windows.
