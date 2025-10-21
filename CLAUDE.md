# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ktncord-go is an abstraction wrapper around the [discordgo](https://github.com/bwmarrin/discordgo) library. It provides a cleaner, interface-based API for building Discord bots in Go with simplified access to common Discord operations.

## Development Commands

### Build and Test
```bash
# Build the module
go build

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Get dependencies
go mod download

# Tidy dependencies
go mod tidy

# Generate documentation
go doc -all
```

## Architecture

### Core Design Pattern

The codebase follows a **wrapper pattern** with strict separation between interfaces and implementations:

- **Interfaces** (`interfaces.go`): Define contracts starting with `I` prefix (e.g., `IDiscordUnit`, `IDiscordUserUnit`)
- **Types** (`types.go`): Concrete struct implementations (e.g., `DiscordUnit`, `DiscordUserUnit`)
- **Implementations**: Spread across domain-specific files (discord.go, user.go, message.go, etc.)

### Unit System

Every Discord entity is wrapped in a "Unit" that holds:
1. A reference to the parent `DiscordUnit` (`discord` field)
2. A reference to the underlying `discordgo` object (`session`, `user`, `message`, etc.)

This creates a consistent hierarchy where all units can access the root Discord session through their `Discord()` method.

**Example hierarchy:**
```
DiscordUnit (root)
  └─> DiscordInteractionUnit
        └─> references DiscordUnit via discord field
  └─> DiscordMessageUnit
        └─> references DiscordUnit via discord field
        └─> can get DiscordChannelUnit via Channel()
              └─> can get DiscordGuildUnit via Guild()
```

### Data Type Conversion

The package provides bidirectional conversion between Discord's native types and ktncord types:

- **datatypes.go**: Defines ktncord-specific types (e.g., `DiscordMessageSend`, `DiscordEmbed`)
- **datatypesimpl.go**: Contains `Build()` methods to convert ktncord types → discordgo types, and `Prepare*()` functions to convert discordgo types → ktncord types

This allows users to work with cleaner types while maintaining full compatibility with the underlying discordgo library.

### File Organization

- `interfaces.go` - All interface definitions
- `types.go` - All struct type definitions
- `discord.go` - Root DiscordUnit implementation and session management
- `interaction.go` - Slash command and interaction handling
- `message.go` - Message operations
- `channel.go` - Channel operations
- `guild.go` - Guild operations
- `user.go` - User operations
- `datatypes.go` - Custom ktncord data structures
- `datatypesimpl.go` - Conversion logic between ktncord and discordgo types

## Key Patterns

### Creating a Discord Bot

```go
// Create the Discord unit
discord, err := CreateDiscordUnit(token)

// Register event handlers
discord.OnSlashCommand(func(d IDiscordUnit, i IDiscordInteractionUnit) {
    // Handle slash commands
})

discord.OnMessageCreate(func(d IDiscordUnit, m IDiscordMessageUnit) {
    // Handle messages
})

// Start the bot with slash commands
discord.Start(commands)
defer discord.Stop()
```

### Unit Creation Pattern

Units are created either:
1. Through factory methods on `DiscordUnit` (e.g., `GetUser()`, `GetChannel()`, `GetGuild()`)
2. Through event handlers that provide units (e.g., `OnSlashCommand` provides `IDiscordInteractionUnit`)
3. Through methods on other units (e.g., `message.Channel()` returns `IDiscordChannelUnit`)

### Accessing Native Objects

All units provide a `Native()` method to access the underlying discordgo object when needed:
```go
nativeSession := discord.Session()
nativeUser := user.Native()
nativeMessage := message.Native()
```

## Important Notes

- The package uses Go 1.25.1+ features (see go.mod)
- All interfaces use the `I` prefix convention
- Methods follow Go naming conventions with proper documentation
- Error handling returns `error` types, not panics
- The package depends on `github.com/ktnuity/ktnuitygo` for utility functions like `AsRef()`
