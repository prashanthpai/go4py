#!/usr/bin/env python

# START OMIT
d = {'name': 'koti', 'age': 28}
d['weight'] = 71.6

# Get only keys
for k in d.keys():
    print k

# Get only values
for v in d.values():
    print v

# Get both keys and values
for k, v in d.items():
    print k, v

del d['age']
# END OMIT
