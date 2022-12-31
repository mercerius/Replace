# Replace
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
