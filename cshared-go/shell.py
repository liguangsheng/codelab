# -*- coding: utf-8 -*-

from ctypes import *

lib = CDLL("./app.so")
print(lib.Add(3,5))
lib.foo()

# import ipdb; ipdb.set_trace()

