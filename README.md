Redis Clone - In-Memory Database

A simplified Redis-like in-memory database built from scratch in Go, featuring data persistence and RESP protocol support.

🎯 Project Overview
This project is a learning-focused implementation of a Redis-like database that demonstrates the fundamental concepts behind in-memory data storage, network protocols, and data persistence. It's designed to help developers understand how databases work at a low level.
✨ Features

In-Memory Data Storage: Fast key-value storage for strings and hashes
RESP Protocol Parser: Full implementation of Redis Serialization Protocol for client-server communication
Concurrent Connection Handling: Uses Go routines to handle multiple client connections simultaneously
Data Persistence: Append-Only File (AOF) implementation for data durability
Crash Recovery: Automatic data restoration from AOF on server restart

📁 Project Structure
.
├── aof.go       # Append-Only File persistence implementation
├── handler.go   # Command handlers and business logic
├── main.go      # Server initialization and connection management
└── resp.go      # RESP protocol parser
🚀 Getting Started
Prerequisites

Go 1.16 or higher

Installation
bash# Clone the repository
git clone https://github.com/yourusername/redis-clone.git
cd redis-clone

# Build the project
go build -o redis-server

# Run the server
./redis-server
Usage
Connect to the server using any Redis client or telnet:
bash# Using redis-cli
redis-cli -p 6379

# Using telnet
telnet localhost 6379
Supported Commands

SET key value - Set a string value
GET key - Get a string value
HSET key field value - Set a hash field
HGET key field - Get a hash field
Additional hash commands...

🛠️ How It Works
1. RESP Protocol
The server implements Redis Serialization Protocol (RESP) to communicate with clients, parsing commands and formatting responses.
2. Concurrent Handling
Go routines enable the server to handle multiple client connections simultaneously without blocking.
3. Data Persistence
All write operations are appended to an AOF file, ensuring data durability. On startup, the server replays the AOF to restore state.
4. In-Memory Storage
Data is stored in Go's native data structures (maps) for fast read/write operations.

