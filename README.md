# TerminalChat

## TCP based Chat application in terminal

### Run server
`go run .`

### Connect to server
 - Linux : `telnet localhost 8080`

### Basic Commands
 - Set your name : `/alias {name}`
   - Sets new visible name as {name}, default was Anonymous
 - View all rooms : `/rooms`
   - Shows a list of live rooms
 - Join or Create a room : `/join {room-name}`
   - Joins a room {room-name} if exists or creates a new one.
 - Send message : `/msg {your message to the room}`
   - Your message will be broadcasted in the room
 - Quit : `/quit`
   - Ends the connections from server
