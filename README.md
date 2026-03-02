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

`go run --version=1`

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
