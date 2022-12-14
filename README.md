# About Command Vault
Command Vault is a small CLI application whose main purpose is to allow the user to save commands in a list. The user can also delete a command, list all commands, search by a command, search by category, or even search by a description.

Below you can check all the commands available so far.

## How to use

| Command           | Description                                               |
| ----------------- | ---------------------------------------------------------------- |
| ./cv -a         | Adds a new command |
| ./cv -l         | List all commands |
| ./cv -d=id      | Delete a command. Id parameter represents the value shown on the ID column |
| ./cv -scom "value"      | Search by command |
| ./cv -scat "value"     | Search by category |
| ./cv -sdes "value"     | Search by description |
| ./cv -v     | Shows the version |

**Note**: You can search without enclosing the values in double quotes, but the value cannot represent a composite value.
## Screenshots
![image](https://user-images.githubusercontent.com/27534241/183315162-e8027a6c-e7f8-43b0-bffb-5c51d53b0d8e.png)


## Improvements

- Add the ability to copy a command to the clipboard.
- Add the ability for the user search by ID.

## Notes
- So far the executable is only available for Unix systems, feel free to build for Windows.

## Current Version
- 1.1.3
