# justlaunch daemon

This directory contains code for the justlaunch daemon.

The purpose of this software is for the CLI to communicate with to install things (i.e. the justlaunch CLI is just a client).

It handles all the magic, including installation, extraction, launching, and much much more. Basically if you removed this out of the equation, nothing would work.

All of this should fit in a light package. After all I'm developing this on a computer with 4gb of RAM.

# How it works (hopefully)

The daemon will be sent messages through a unix socket. (Windows isn't focused on, y'all have Twitch anyways.)
Why not HTTP? Local HTTP servers are a security risk.
