import click
import pyperclip

CONTEXT_SETTINGS = dict(help_option_names=["-h", "--help"])

to_copy = {
    "quicksort": """def quicksort(T, p, r):
    if p < r:
        q = partition(T, p, r)
        quicksort(T, p, q - 1)
        quicksort(T, q + 1, r)
    return T

def partition(T, p, r):
    # y = random.randint(p, r)
    # albo
    # y = med(random.randint(p, r), p, r)
    # T[y], T[r] = T[r], T[y]
    x = T[r]
    i = p - 1
    for j in range(p, r):
        if T[j] <= x:
            i += 1
            T[i], T[j] = T[j], T[i]
    T[i + 1], T[r] = T[r], T[i + 1]
    return i + 1

# def med(T, a, b, c):
#     if (T[b] < T[a] and T[a] < T[c]) or (T[c] < T[a] and T[a] < T[b]):
#         return a
#     if (T[a] < T[b] and T[b] < T[c]) or (T[c] < T[b] and T[b] < T[a]):
#         return b
#     else:
#         return c""",
    "heapsort": """def left(i):
    return 2*i + 1

def right(i):
    return 2*i + 2

def parent(i):
    return (i-1)//2

def heapify(T, i, n):
    l = left(i)
    r = right(i)
    max_ind = i
    if l < n and T[l] > T[max_ind]:
        max_ind = l
    if r < n and T[r] > T[max_ind]:
        max_ind = r
    if max_ind != i:
        T[i], T[max_ind] = T[max_ind], T[i]
        heapify(T, max_ind, n)

def build_heap(T):
    n = len(T)
    for i in range(parent(n-1), -1, -1):
        heapify(T, i, n)

def heapsort(T):
    n = len(T)
    build_heap(T)
    for i in range(n-1, 0, -1):
        T[0], T[i] = T[i], T[0]
        heapify(T, 0, i)
    return T""",
    "mergesort": """def mergesort(T, l, r):  
    if l >= r:  
        return  
    m = (l + r) // 2  
    mergesort(T, l, m)  
    mergesort(T, m + 1, r)  
    T = merge(T, l, r, m)  
    return T
  
def merge(T, l, r, middle):  
    L = T[l:middle + 1]  
    R = T[middle+1:r+1]  
    id_L = 0  
    id_R = 0  
    sorted = l  
    while id_L < len(L) and id_R < len(R):  
        if L[id_L][0] <= R[id_R][0]:  
            T[sorted] = L[id_L]  
            id_L = id_L + 1   
        else:  
            T[sorted] = R[id_R]  
            id_R = id_R + 1     
        sorted = sorted + 1  
   
    while id_L < len(L):  
        T[sorted] = L[id_L]  
        id_L = id_L + 1  
        sorted = sorted + 1  
  
    while id_R < len(R):  
        T[sorted] = R[id_R]  
        id_R = id_R + 1  
        sorted = sorted + 1  
    return T""",
    "countingsort": """def countingsort(T, k):
    n = len(T)
    C = [0 for _ in range(k)]
    B = [0 for _ in range(n)]

    for i in range(n):
        C[T[i]] += 1

    for i in range(1, k):
        C[i] += C[i-1]

    for i in range(n-1, -1, -1):
        B[C[T[i]]-1] = T[i]
        C[T[i]] -= 1
    
    return B""",
    "bubblesort": """def bubblesort(T):
    n = len(T)
    for i in range(n-1):
        for j in range(n-i-1):
            if T[j] > T[j+1]:
                T[j], T[j+1] = T[j+1], T[j]
    return T""",
    "quickselect": """def partition(T, p, r):
    # y = random.randint(p, r)
    # albo
    # y = med(random.randint(p, r), p, r)
    # T[y], T[r] = T[r], T[y]
    x = T[r]
    i = p - 1
    for j in range(p, r):
        if T[j] >= x:
            i += 1
            T[i], T[j] = T[j], T[i]
    T[i + 1], T[r] = T[r], T[i + 1]
    return i + 1

# def med(T, a, b, c):
#     if (T[b] < T[a] and T[a] < T[c]) or (T[c] < T[a] and T[a] < T[b]):
#         return a
#     if (T[a] < T[b] and T[b] < T[c]) or (T[c] < T[b] and T[b] < T[a]):
#         return b
#     else:
#         return c

# k-ty największy element
def quickselect(T, p, r, i):
    if p == r:
        print(T)
        return T[p]
    q = partition(T, p, r)
    k = q - p + 1
    if i == k:
        print(T)
        return T[q]
    elif i < k:
        return quickselect(T, p, q - 1, i)
    else:
        return quickselect(T, q + 1, r, i - k)"""
}


@click.group(context_settings=CONTEXT_SETTINGS)
def cli():
    """Copy algorithms to clipboard"""
    pass


@cli.command()
def quicksort():
    pyperclip.copy(to_copy["quicksort"])
    click.echo("Copied!")


@cli.command()
def heapsort():
    pyperclip.copy(to_copy["heapsort"])
    click.echo("Copied!")


@cli.command()
def mergesort():
    pyperclip.copy(to_copy["mergesort"])
    click.echo("Copied!")


@cli.command()
def countingsort():
    pyperclip.copy(to_copy["countingsort"])
    click.echo("Copied!")


@cli.command()
def bubblesort():
    pyperclip.copy(to_copy["bubblesort"])
    click.echo("Copied!")

@cli.command()
def quickselect():
    pyperclip.copy(to_copy["quickselect"])
    click.echo("Copied!")
