# mdb-cli

cli written in go

### Commands, subcommands, flags

| command       |            Description            |
| :------------ | :-------------------------------: |
| -v, --version |        returns cli version        |
| -h, --help    |                 s                 |
| -s, --scan    |                 a                 |
| -i, --info    |                 a                 |
| -l, --list    | lists movies available in library |
| -r, --refresh |     refresh metadata of media     |
| -u, --update  |     refresh metadata of media     |

modifiers

| command | Description |
| :------ | :---------: |
| --force |             |

Scan folder,

`mdb --scan <path>`

`mdb --scan /user/username/Documents/movies`

Add folder to config

`mdb --scan /user/username/Documents/movies --save`

Auth
set token

`go run main.go command arg1 arg2`

`go build`

TMDB_URL=https://api.themoviedb.org/3 TMDB_API_KEY=test go run main.go

```text
mdb-cli/
├── cmd/                # Cobra command definitions
│   ├── root.go         # Entry point & global flags
│   ├── auth.go         # Login / Logout commands
│   ├── scan.go         # Directory scanning logic
│   └── sync.go         # Cloud sync commands
├── internal/           # Private application code
│   ├── scanner/        # Logic for FS traversal & regex matching
│   ├── metadata/       # API client for TMDB/MDB Gateway
│   └── config/         # Viper setup for ~/.mdb.yaml
├── main.go             # Minimal entry point
└── go.mod              # Dependencies

```

Sample by Plex:

Plex Media Scanner (c) 2010-2014 Plex Development Team.

-h, --help Display this message.
-v, --verbose Show more output.
-p, --progress Show special progress output.
--log-file-suffix Specify suffix for log file.

Actions:

-r, --refresh Refresh the metadata. Deprecated
-a, --analyze Analyze media information.
--analyze-deeply Fully read and perform deep media analysis.
-b, --index Generate a media index file. (Video Preview Thumbnails) Deprecated
-s, --scan Scan for new media. Deprecated
-i, --info Get information.
-l, --list List.
-g, --generate Regenerate thumbnails/fanart.
-t, --tree Show a section tree.
-w, --reset Delete all media out of a section.
-n, --add-section --type <type:1,2,8> --agent --location --lang Add a new section.
-D, --del-section Delete a section.

Items to which actions apply:

-c, --section A library section ID.
-o, --item An item ID.
-d, --directory A directory path.
-f, --file A file.

Modifiers to actions:

-x, --force Force an operation (e.g. refresh).
--no-thumbs Do not regenerate thumbs when analyzing.
--chapter-thumbs-only Only generate chapter thumbnails during generate pass
--thumbOffset Percent offset into video for thumbnail image generated during media analysis.
--artOffset Percent offset into video for fanart image generated during media analysis.
