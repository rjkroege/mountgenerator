This is  a tool to scan the list of named disks mounted on the system and
then generates a `systemd` mount unit for each of the mounted disks.

Build the tool and install somewhere convenient:

```
upx mountgenerator
rclone copy mountgenerator gda:boot-tools-liqui-org/linux/amd64
```
