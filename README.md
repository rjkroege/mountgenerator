This is  a tool to scan the list of named disks mounted on the system and
then generates a `systemd` mount unit for each of the mounted disks.

Build the tool, and then copy it to the cloud. Make sure to make the
tool public. Maybe the entire bucket should be public.

```
upx mountgenerator
gsutil cp mountgenerator gs://boot-tools-liqui-org
```
