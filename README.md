
## How to use

### Adding a new command to the list

- make cv/a args="cp -r dir_to_copy/ new_copy_dir/|-|Linux|-|copy entire directories"

**Note 1**: 3 parameters are needed

| Arguments         | Example                                               |
| ----------------- | ---------------------------------------------------------------- |
| The command itself      | docker ps |
| The command category    | Docker |
| The command description | This command is used to list the running containers |

**Note 2**: The characters |-| are used to separate arguments

### Listing all commands

- make cv/l

### Removing a command

- make cv/d id=1

**Note**: *The ID shown in the table must be used as an index for removing the command*

## Screenshots
![image](https://user-images.githubusercontent.com/27534241/183266869-608e0225-e1c5-49b9-a756-e65589fc8ae1.png)

## Improvements

- Add the ability for the user to choose the characters to split the arguments
- Add the ability for the user to search by commands and categories
