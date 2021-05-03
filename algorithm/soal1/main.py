#!/bin/python

import math
import os
import random
import re
import sys

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    match = 0
    for i in ar:
        count = 0
        for j in ar:
            if i == j :
                count += 1
        match += int(count / 2)
        ar = [x for x in ar if x != i]
    return match

if __name__ == '__main__':

    n = int(input())

    ar = input().rstrip().split()
    
    if (len(ar)>n):
        print("the length of array is more than given. will compute based on",n, "first index")
        newAr = []
        for i in range(0,n):
            newAr.append(ar[i])
        ar = newAr

    result = sockMerchant(n, ar)

    print("Output : " , result)
