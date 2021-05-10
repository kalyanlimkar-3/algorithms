import math

n = 2
for i in range(0, int(math.pow(2, n))):
    div = i
    s = ""

    j = n
    while(j > 0):
        rem = int(div % 3)
        div = int(div/3)

        s = s+"F" if(not(rem)) else s+"T"
        j -= 1

    print(s)
