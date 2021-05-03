
import math
import os
import random
import re
import sys

def countingValleys(steps, path):
    # Write your code here
    countValley = 0
    altitude = 0
    for step in path:
        if step=='U':
            altitude += 1
            if (altitude == 0) :
                countValley += 1
        else:
            altitude -= 1
    return countValley
            



if __name__ == '__main__':    
    steps = int(input())
    
    path = str(input())
    
    if (len(path)>steps):
        print("the length of steps is more than given. will compute based on",steps, "first char")
        path = path[0:steps] 

    result = countingValleys(steps, path)
    print("Output : ", result)
