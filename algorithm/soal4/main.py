nTrip = 100
totalTurnedOn = 0
for i in range(1,nTrip+1):
     turnedOn = 0
     for j in range(1,100):
       if (j % i == 0):
         turnedOn += 1
     print("Andrew turned on ", turnedOn + 1, " lamp on his " + str(i) + "-trip")
     totalTurnedOn += turnedOn + 1

print("Andrew turned on total lamp of : " ,totalTurnedOn)