# Chapter 1

- What is Go: history, companies, examples 
- Install, setup, GOPATH 
- IDE for Go: vim, goglang, atom 
- Packages: repositories, awesome-go, dependency manager 
- Project structure: layout, naming 
- Build: make, cross-platform 
- Run: locally, docker 
- What to read to understand basic syntaxes 

Download [Slides](https://www.slideshare.net/secret/mWiifR3xs0F4mH).

## Homework

See the slides to get the task description.

Google Places Ids:
- Alesund, Norway - `ChIJS_tHEEfaFkYR6o_AO-5iJeg`
- Bergen, Norway - `ChIJd312ZkkNOUYRCAretD6gQp4`
- Oslo, Norway - `ChIJOfBn8mFuQUYRmh4j019gkn4`
- Lillehammer, Norway - `ChIJjSi196diakYRNO1bStM6JzI`
- Berlin, Germany - `ChIJAVkDPzdOqEcRcDteW0YgIQQ`

## Execution

Build, run, and verify.

```bash
make

build/homework1 \
    -p "ChIJS_tHEEfaFkYR6o_AO-5iJeg" \
    -p "ChIJd312ZkkNOUYRCAretD6gQp4" \
    -p "ChIJOfBn8mFuQUYRmh4j019gkn4" \
    -p "ChIJjSi196diakYRNO1bStM6JzI" \
    -p "ChIJAVkDPzdOqEcRcDteW0YgIQQ" \
    -o "build/out.json" \
    -k "GOOGLE API KEY"

cat build/out.json
```
