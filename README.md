# Command Vault
Is a small CLI application whose main purpose is to allow the user to save commands in a list. The user can also delete a command or list all commands.

## How to use

### Adding a new command to the list

- make cv/a args="cp -r dir_to_copy/ new_copy_dir/|-|Linux|-|copy entire directories"

OR

- ./main -a ="cp -r dir_to_copy/ new_copy_dir/|-|Linux|-|copy entire directories"

**Note 1**: 3 parameters are needed

| Arguments         | Example                                               |
| ----------------- | ---------------------------------------------------------------- |
| The command itself      | docker ps |
| The command category    | Docker |
| The command description | This command is used to list the running containers |

**Note 2**: The characters |-| are used to separate arguments

### Listing all commands

- make cv/l

OR

- ./main -l

### Removing a command

- make cv/d id=1

OR

- ./main -d=1

**Note**: *The ID shown in the table must be used as an index for removing the command*

### Search for command

- make cv/scom args=docker OR make cv/scom args="docker run"

OR

- ./main -scom docker OR ./main -scom "docker run"

**Note**: *To search by compound word, use double quotes*

### Search for category

- make cv/scat args=Linux OR make cv/scat args="Linux Mint"

OR

- ./main -scat Linux OR ./main -scat "Linux Mint"

**Note**: *To search by compound word, use double quotes*

### Search for description

- make cv/sdes args=running OR make cv/sdes args="running container"

OR

- ./main -sdes running OR ./main -sdes "running container"

**Note**: *To search by compound word, use double quotes*

## Screenshots
![image](https://user-images.githubusercontent.com/27534241/183266869-608e0225-e1c5-49b9-a756-e65589fc8ae1.png)

## Improvements

- Add the ability for the user to choose the characters to split the arguments
- Add the ability for the user to search by commands and categories
- Add the ability to copy a command to the clipboard

