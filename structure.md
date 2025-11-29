go_web_server/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/                # Private application code
│   ├── handlers/           # HTTP handlers
│   ├── middleware/         # HTTP middleware
│   ├── models/             # Data models
│   ├── database/           # Database logic
│   └── config/             # Configuration
├── pkg/                     # Public library code (if needed)
│   └── ...
├── api/                     # API definitions (OpenAPI, protobuf, etc.)
├── web/                     # Static files, templates (if applicable)
├── migrations/              # Database migrations (if applicable)
├── scripts/                 # Build/deployment scripts
├── go.mod
├── go.sum
└── README.md
