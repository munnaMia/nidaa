# nidaa
A social platform for good hearts...

In this project i follow Domain Driven Design (DDD) principles and Clean Architecture to build a scalable and maintainable structure. The project is built using Go programming language and follows best practices for software development.

I know this this is also an unfinished project but i will keep updating it as i learn more about DDD and Clean Architecture. when i have time i will also add more features to this project. but this isn'a an commitment to finish it, this is just a learning project for me.

## Installation and Setup
```bash
# Clone the repository
git clone

# move into the project directory
cd nidaa

# Copy the example environment file and create a new .env file
cp .env.example .env

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

// messaging feature. architecture diagram

[ User A (Client) ]                                        [ User B (Client) ]
         │                                                           ▲
         │ 1. Sends WS Message {"recipient_id": 2, content: "hi"}    │ 4. Delivered via WS
         ▼                                                           │
┌─────────────────────────────────────────────────────────────────────────┐
│                          Go Backend Service                             │
│                                                                         │
│  ┌───────────────────────┐             ┌─────────────────────────────┐  │
│  │  WebSocket Handler    │ ──────────► │ WebSocket Hub / Manager     │  │
│  └───────────────────────┘             │ (Maps UserID -> WS Client)  │  │
│              │                         └─────────────────────────────┘  │
│              ▼                                                          │
│  ┌───────────────────────┐                                              │
│  │   Message UseCase     │ ──► Save message to Postgres                 │
│  └───────────────────────┘                                              │
└─────────────────────────────────────────────────────────────────────────┘