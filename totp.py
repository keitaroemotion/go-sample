#!/usr/bin/env python

import hmac
import hashlib
import array
import time
import unittest

def HOTP(K, C, digits=6):
    C_bytes = _long_to_byte_array(C)
    hmac_sha_512 = hmac.new(key=K, msg=C_bytes, digestmod=hashlib.sha512).hexdigest()
    return Truncate(hmac_sha_512)[-digits:]

def TOTP(K, digits=6, window=30):
    C = long(time.time() / window)
    return HOTP(K, C, digits=digits)

def Truncate(hmac_sha_512):
    offset = int(hmac_sha_512[-1], 16)
    binary = int(hmac_sha_512[(offset * 2):((offset * 2) + 8)], 16) & 0x7fffffff
    return str(binary)

def _long_to_byte_array(long_num):
    byte_array = array.array('B')
    for i in reversed(range(0, 8)):
        byte_array.insert(0, long_num & 0xff)
        long_num >>= 8
    return byte_array

key_string = 'tobasojyo@gmail.comHENNGECHALLENGE003'
value = TOTP(key_string)
print(value)
