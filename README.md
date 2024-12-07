# Baker

A simple file backup utility.

Baker's automates a simple, yet incredibly common task for Linux users:

```Bash
# Creating a backup of a file
cp file.txt file.txt.bak
# Restoring the backup
mv -f file.txt.bak file.txt
# Creating additional backups
cp file.txt file.txt.bak2 # etc...
```

Any terminal user knows this is horribly annoying, but in a lot of cases, setting up an entire git repository or some other backup tool just to keep track of one or two backups is wholly unnecessary.
