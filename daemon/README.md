# justlaunch daemon

This directory contains code for the justlaunch daemon.

The purpose of this software is for the CLI to communicate with to install things (i.e. the justlaunch CLI is just a client).

It handles all the magic, including installation, extraction, launching, and much much more. Basically if you removed this out of the equation, nothing would work.

All of this should fit in a light package. After all I'm developing this on a computer with 4gb of RAM.

# How it works (hopefully)

The daemon will be sent messages through a unix socket. (Windows isn't focused on, y'all have Twitch anyways.)
Why not HTTP? Local HTTP servers are a security risk.

# Commands/the specification
The cilent and daemon will share a unix socket (I think that works, right?) as a 2 way communication method.

Glossary:
- **Profile**: A directory where launcher specific files are stored for configuration as well as the modpack's own data.
- **File ID**: CurseForge assigns each uploaded file an ID that can be used to refer to the file. Used in downloads, etc.
- **Modpack ID**: A set of numbers that CurseForge assigns a modpack. Generally useless IMO, but I have to deal with it.
- **Modpack**: Generally, a modpack is considered a collection of mods and mod configurations that change Minecraft.


The client can send commands to the daemon socket:
- `install <modpack-id> <file-id>` Creates a profile and extracts the modpack to a folder. File ID is latest by default.
- `update <profile-name>` Checks for updates on a profile's assigned modpack.
- `uninstall <profile-name>` Removes all traces of a profile and its modpack.
- `import <zip-file>` Imports a Curse-format ZIP file containing a modpack.
- `search <type> <term>` Searches CurseForge for relevant mods/modpacks.
- `launch <profile-name>` Attempts to launch the modpack in a profile.

The server can send responses back through the socket:
- `ok` Simply a acknowledgement; "I'm on my way!"
- `general-error` Some programming error happened. In this case, check the logs for a stack trace.
- `waiting` Installation is still happening.
- `permission-denied` The daemon cannot access the requested directory for installation.
- `done` It's ready to go.
