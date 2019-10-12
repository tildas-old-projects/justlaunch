# Just Launch The Game

WIP Minecraft launcher written in Go

# Building

Requirements: 
- [Go 1.13 or above](https://golang.org)
- [Ultralight latest](https://github.com/ultralight-ux/ultralight#getting-the-latest-sdk)
- [Yarn](https://yarnpkg.com)
- [fileb0x](https://github.com/UnnoTed/fileb0x)

1. Extract the libraries (located in the `bin` folder inside the archive) from Ultralight's SDK into the repository. The other stuff can be safely deleted.
2. Go into the `public` folder and run `yarn` to download the needed files for the frontend.
3. Go back into the root of the repo and run `go generate` to generate the frontend webfiles. (or: `fileb0x b0x.yml`)
4. Run `go build` to get a `justlaunch` executable.

**Note:** The executable is dynamically linked, so you'll either need:
- the Ultralight libraries installed in your path
- change the place where the OS searches for the libraries

I personally use the 2nd option: `env LD_LIBRARY_PATH=bin ./justlaunch`

**TODO:** Only works on Linux, figure out what Windows and Mac do



