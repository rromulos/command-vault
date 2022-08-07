# Command Vault
Is a small CLI application whose main purpose is to allow the user to save commands in a list. The user can also delete a command or list all commands.

## How to use

| Command           | Description                                               |
| ----------------- | ---------------------------------------------------------------- |
| ./cv -a         | Adds a new command |
| ./cv -l         | List all commands |
| ./cv -d=id      | Delete a command. Id parameter represents the value shown on the ID column |
| ./cv -scom "value"      | Search by command |
| ./cv -scat "value"     | Search by category |
| ./cv -sdes "value"     | Search by description |

**Note**: You can search without enclosing the values in double quotes, but the value cannot represent a composite value
## Screenshots
![image](https://user-images.githubusercontent.com/27534241/183266869-608e0225-e1c5-49b9-a756-e65589fc8ae1.png)

## Improvements

- Add the ability to copy a command to the clipboard
- Add the ability for the user search by ID

