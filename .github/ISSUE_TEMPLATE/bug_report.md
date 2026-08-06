---
name: Bug report
about: Create a report to help us improve remedy
title: ''
labels: bug
assignees: elzorrorebelde

---

**Environment**

- remedy version:
- Git repository URL (if possible):
- contents of `.remedy-run`:

```
.remedy-run contents to insert here
```

- latest revision of `.remedy-run` (`git --no-pager log --format='%H' -1 -- .remedy-run`): `edit this`
- contents of JSON configuration (default name: `remedy.json`):

```
JSON configuration contents to insert here
```

- contents of the license header template:

```
header template contents to insert here
```

**Protips**

_the following only applies to duplicated or malformed header issues_

`remedy` can be quite sensitive to small differences between the header template and the actual headers in source files, especially at the first run!

Here a few things to check before reporting an issue:

- [ ] the configuration JSON file and header template have not been changed since `.remedy-run` last revision
- [ ] there are no differences between the affected source file's header and the header in configured template file

**Describe the bug**
A clear and concise description of what the bug is.
If it concerns duplicated/malformed headers, please include the full contents of at least one affected source file:

```
source file contents to insert here
```

**To Reproduce**
Steps to reproduce the behavior:

1. Go to '...'
2. Click on '....'
3. Scroll down to '....'

**Expected behavior**
A clear and concise description of what you expected to happen.

**Additional context**
Add any other context about the problem here.
