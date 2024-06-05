import numpy as np

# Data from table
data = np.array([
    [0, 0, 0, 0, 0, 4],
    [0, 0, 0, 0, 0, 3],
    [0, 5, 2, 0, 0, 4],
    [0, 11, 7, 14, 10, 4],
    [5, 4, 32, 7, 0, 0],
    [17, 0, 0, 0, 0, 0]
])

P, Q = 0, 0
rows, cols = data.shape

for i in range(rows):
    for j in range(i+1, rows):
        for k in range(cols):
            for l in range(k+1, cols):
                P += data[i, k] * data[j, l]
                Q += data[i, l] * data[j, k]

print("Total Concordant pairs (P):", P)
print("Total Discordant pairs (Q):", Q)
