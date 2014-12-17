# fsnoop

fsnoop snoops on your files

ran on its own with no arguments, fsnoop will report all file modification events.

here's a sample session run from `/tmp/snoop`.  Lines that start with `#` were
run in a separate shell:

```
# ls
open /tmp/snoop
close /tmp/snoop
close-nowrite /tmp/snoop

# touch test
create /tmp/snoop/test
open /tmp/snoop/test
attrib /tmp/snoop/test
close /tmp/snoop/test
close-write /tmp/snoop/test

# echo hi > test
modify /tmp/snoop/test
open /tmp/snoop/test
modify /tmp/snoop/test
close /tmp/snoop/test
close-write /tmp/snoop/test

# echo append >> test
open /tmp/snoop/test
modify /tmp/snoop/test
close /tmp/snoop/test
close-write /tmp/snoop/test

# mv test test.1
moved-from /tmp/snoop/test
move /tmp/snoop/test -> /tmp/snoop/test.1
moved-to /tmp/snoop/test.1

# gzip test.1
open /tmp/snoop/test.1
create /tmp/snoop/test.1.gz
open /tmp/snoop/test.1.gz
access /tmp/snoop/test.1
modify /tmp/snoop/test.1.gz
close /tmp/snoop/test.1
close-nowrite /tmp/snoop/test.1
attrib /tmp/snoop/test.1.gz
attrib /tmp/snoop/test.1.gz
attrib /tmp/snoop/test.1.gz
close /tmp/snoop/test.1.gz
close-write /tmp/snoop/test.1.gz
delete /tmp/snoop/test.1

# rm test.1.gz
delete /tmp/snoop/test.1.gz
```

of course, this is not as accurate as actually reading a program or stracing
it, but if the ordering of the file modification events is all you care about,
then fsnoop acts as a useful filter.

this thing only works on linux.

basically i reinvented a poorman's version of
[inotifywait](https://github.com/rvoicilas/inotify-tools).
