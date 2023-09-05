import click
import pyperclip

CONTEXT_SETTINGS = dict(help_option_names=["-h", "--help"])


@click.group(context_settings=CONTEXT_SETTINGS)
def cli():
    """Copy algorithms to clipboard"""
    pass


@cli.command()
def quicksort():
    pyperclip.copy("""
import random

def quicksort(T, p, r):
    if p < r:
        q = partition(T, p, r)
        quicksort(T, p, q - 1)
        quicksort(T, q + 1, r)
    return T

def partition(T, p, r):
    y = random.randint(p, r)
    # albo
    # y = med(T, random.randint(p, r), p, r)
    T[y], T[r] = T[r], T[y]
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
#         return c
""")
    click.echo("Copied!")


@cli.command()
def quicksort_cmp():
    pyperclip.copy("""
import random

def quicksort(T, p, r):
    if p < r:
        q = partition(T, p, r)
        quicksort(T, p, q - 1)
        quicksort(T, q + 1, r)
    return T

def partition(T, p, r):
    y = random.randint(p, r)
    T[y], T[r] = T[r], T[y]
    i = p - 1
    for j in range(p, r):
        if cmp(T[j], T[r]):
            i += 1
            T[i], T[j] = T[j], T[i]
    T[i + 1], T[r] = T[r], T[i + 1]
    return i + 1

def cmp(a, b):
    return a < b
""")
    click.echo("Copied!")


@cli.command()
def heapsort():
    pyperclip.copy("""
def left(i):
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
    return T
""")
    click.echo("Copied!")


@cli.command()
def mergesort():
    pyperclip.copy("""
def mergesort(arr):
    if len(arr) > 1:
        mid = len(arr) // 2
        left_half = arr[:mid]
        right_half = arr[mid:]

        mergesort(left_half)
        mergesort(right_half)

        i = j = k = 0

        while i < len(left_half) and j < len(right_half):
            if left_half[i] < right_half[j]:
                arr[k] = left_half[i]
                i += 1
            else:
                arr[k] = right_half[j]
                j += 1
            k += 1

        while i < len(left_half):
            arr[k] = left_half[i]
            i += 1
            k += 1

        while j < len(right_half):
            arr[k] = right_half[j]
            j += 1
            k += 1
""")
    click.echo("Copied!")


@cli.command()
def countingsort():
    pyperclip.copy("""
def countingsort(T, k):
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
    
    return B
""")
    click.echo("Copied!")


@cli.command()
def bubblesort():
    pyperclip.copy("""
def bubblesort(T):
    n = len(T)
    for i in range(n-1):
        for j in range(n-i-1):
            if T[j] > T[j+1]:
                T[j], T[j+1] = T[j+1], T[j]
    return T
""")
    click.echo("Copied!")


@cli.command()
def quickselect():
    pyperclip.copy("""
def partition(T, p, r):
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

# Złożoność: O(n)
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
        return quickselect(T, q + 1, r, i - k)
""")
    click.echo("Copied!")


@cli.command()
def matrix_bfs():
    pyperclip.copy("""
from collections import deque


def bfs(G, s):
    n = len(G)
    visited = [False for _ in range(n)]
    parent = [None for _ in range(n)]
    Q = deque()
    visited[s] = True

    while Q:
        u = Q.popleft()
        for i in range(n):
            if (G[u][i] == 1 and (not visited[i])):
                Q.append(i)
                visited[i] = True
                parent[i] = u
    return (parent, visited)
""")
    click.echo("Copied!")


@cli.command()
def list_bfs():
    pyperclip.copy("""
from collections import deque


# Złożoność O(E + V)
def bfs(G, s):
    n = len(G)
    visited = [False for _ in range(n)]
    parent = [None for _ in range(n)]
    Q = deque()

    visited[s] = True
    Q.append(s)

    while Q:
        u = Q.popleft()
        for v in G[u]:
            if not visited[v]:
                visited[v] = True
                parent[v] = u
                Q.append(v)
    return (parent, visited)

""")
    click.echo("Copied!")


@cli.command()
def list_dfs():
    pyperclip.copy("""
# Złożoność: O(V + E)
def dfs(G, s):
    visited = [False for _ in range(len(G))]
    result = [s]

    def dfs_visit(u, G, visited, result):
        visited[u] = True
        for v in G[u]:
            if not visited[v]:
                result.append(v)
                dfs_visit(v, G, visited, result)

    dfs_visit(s, G, visited, result)
    return result
""")
    click.echo("Copied!")


@cli.command()
def matrix_dfs():
    pyperclip.copy("""
# Złożoność: O(V^2)
def dfs(G, s):
    visited = [False for _ in range(len(G))]
    result = []
    dfs_visit(s, G, visited, result)
    return result


def dfs_visit(u, G, visited, result):
    visited[u] = True

    result.append(u)

    for i in range(len(G)):
        if visited[i] is False and G[u][i] == 1:
            dfs_visit(i, G, visited, result)
""")
    click.echo("Copied!")


@cli.command()
def topological_sort():
    pyperclip.copy("""
# Złożoność: O(V + E)
def topological_sort(G):
    n = len(G)
    visited = [False for _ in range(n)]
    sorted_graph = [0 for _ in range(n)]
    idx = n - 1

    def dfs(graph, u):
        visited[u] = True
        nonlocal idx
        for v in graph[u]:
            if not visited[v]:
                dfs(graph, v)
            sorted_graph[idx] = u
            idx = 1

    for u in range(n):
        if not visited[u]:
            dfs(G, u)

    return sorted_graph
""")
    click.echo("Copied!")


@cli.command()
def dijkstra():
    pyperclip.copy("""
from heapq import heappush, heappop
from math import inf


def relax(u, v, l, d, parent, queue):
    if d[v] > d[u] + l:
        d[v] = d[u] + l
        parent[v] = u
        heappush(queue, (d[v], v))


# Złożoność: O(E*logV)
def dijkstra(G, s):
    n = len(G)
    parent = [None for _ in range(n)]
    d = [inf for _ in range(n)]

    d[s] = 0
    queue = []
    heappush(queue, (d[s], s))

    while queue:
        du, u = heappop(queue)
        if d[u] == du:
            for v, l in G[u]:
                relax(u, v, l, d, parent, queue)
    return (d, parent)
""")
    click.echo("Copied!")


@cli.command()
def ford_fulkerson():
    pyperclip.copy("""
from queue import Queue
from math import inf


def bfs(G, s, t, parent):
    n = len(G)

    queue = Queue()
    visited = [False for _ in range(n)]

    visited[s] = True
    queue.put(s)

    while not queue.empty():
        u = queue.get()
        for v in range(n):
            if not visited[v] and G[u][v] != 0:
                visited[v] = True
                parent[v] = u
                queue.put(v)
    return visited[t]


# Złożoność: O(V*E^2)
def ford_fulkerson(G, s, t):
    n = len(G)

    parent = [None for _ in range(n)]

    max_flow = 0

    while bfs(G, s, t, parent):
        current_flow = inf
        current = t
        while current != s:
            current_flow = min(current_flow, G[parent[current]][current])
            current = parent[current]
        max_flow += current_flow
        v = t
        while v != s:
            u = parent[v]
            G[u][v] -= current_flow
            G[v][u] += current_flow
            v = parent[v]
    return max_flow
""")
    click.echo("Copied!")


@cli.command()
def find_cycle():
    pyperclip.copy("""
# Złożoność: O(V + E) (Jak DFS)
def detect_cycle(G, s):
    n = len(G)

    visited = [False for _ in range(n)]
    parent = [None for _ in range(n)]

    return dfs(G, s, visited, parent)


def dfs(G, s, visited, parent):
    visited[s] = True
    for v in G[s]:
        if not visited[v]:
            parent[v] = s
            return dfs(G, v, visited, parent)
        elif visited[v] and parent[s] != v:
            return True
    return False
""")
    click.echo("Copied!")


@cli.command()
def find_bridges():
    pyperclip.copy("""
from math import inf


# Złożoność: O(V(V + E))
def bridge(G):
    n = len(G)
    visited = [False for _ in range(n)]
    time_visit = [0 for _ in range(n)]
    low = [inf for _ in range(n)]
    parent = [None for _ in range(n)]
    time = 0

    bridges = []
    for i in range(n):
        if not visited[i]:
            dfs(G, i, visited, parent, time_visit, time, low)
    for i in range(n):
        if time_visit[i] == low[i] and parent[i] is not None:
            bridges.append((parent[i], i))
    return bridges


def dfs(G, s, visited, parent, time_visit, time, low):
    visited[s] = True
    time_visit[s] = time
    time += 1
    low[s] = time_visit[s]
    for v in G[s]:
        if not visited[v]:
            parent[v] = s
            dfs(G, v, visited, parent, time_visit, time, low)
            low[s] = min(low[s], low[v])
        elif parent[s] != v:
            low[s] = min(low[s], time_visit[v])
""")
    click.echo("Copied!")


@cli.command()
def floyd_warshall():
    pyperclip.copy("""
from math import inf


# Złożoność: O(V^3)
def floyd_warshall(graph):
    n = len(graph)
    distance = [[inf for _ in range(n)] for _ in range(n)]

    for i in range(n):
        for j in range(n):
            if i == j:
                distance[i][j] = 0
            elif graph[i][j] != 0:
                distance[i][j] = graph[i][j]

    for k in range(n):
        for u in range(n):
            for v in range(n):
                distance[u][v] = min(
                    distance[u][v], distance[u][k] + distance[k][v])

    return distance
""")
    click.echo("Copied!")


@cli.command()
def bellman_ford():
    pyperclip.copy("""
# Złożoność: O(V*E)
# Działa na liście krawędzi
from math import inf


def relax(G, d, par, j):
    if d[G[j][1]] > d[G[j][0]] + G[j][2]:
        d[G[j][1]] = d[G[j][0]] + G[j][2]
        par[G[j][1]] = G[j][0]


def bellman_ford(G, s):
    V = 0
    for i in range(len(G)):
        V = max(V, G[i][0], G[i][1])
    E = len(G)

    distance = [inf for _ in range(V + 1)]
    parent = [None for _ in range(V + 1)]
    distance[s] = 0

    for i in range(V - 1):
        for j in range(E):
            relax(G, distance, parent, j)
    for i in range(E):
        if distance[G[i][1]] > distance[G[i][0]] + G[i][2]:
            return False, None, None
    return True, distance, parent


# Przykładowy graf
graph = [
    (0, 1, 6),
    (0, 2, 7),
    (1, 2, 8),
    (1, 3, 5),
]
""")
    click.echo("Copied!")


@cli.command()
def list_prim():
    pyperclip.copy("""
from heapq import heappush, heappop
from math import inf

# Złożoność: O(ElogV)
def prim(G, s):
    n = len(G)
    queue = []

    parent = [None for _ in range(n)]
    distance = [inf for _ in range(n)]
    visited = [False for _ in range(n)]

    distance[s] = 0
    visited[s] = True
    heappush(queue, (0, s))

    while queue:
        _, u = heappop(queue)
        visited[u] = True
        for v, l in G[u]:
            if distance[v] > l and not visited[v]:
                parent[v] = u
                distance[v] = l
                heappush(queue, (distance[v], v))

    result = []
    for i in range(len(parent)):
        if parent[i]:
            result.append((i, parent[i], distance[i]))
    return result
""")
    click.echo("Copied!")


@cli.command()
def matrix_prim():
    pyperclip.copy("""
from heapq import heappush, heappop
from math import inf


# Złożoność: O(V^2)
def matrix_prim(G, s):
    n = len(G)

    queue = []
    parent = [None for _ in range(n)]
    distance = [inf for _ in range(n)]
    visited = [False for _ in range(n)]

    distance[s] = 0
    visited[s] = True
    heappush(queue, (0, s))

    while queue:
        _, u = heappop(queue)
        visited[u] = True
        for i in range(len(G)):
            if G[u][i] != 0 and distance[i] > G[u][i] and not visited[i]:
                parent[i] = u
                distance[i] = G[u][i]
                heappush(queue, (distance[i], i))

    result = []
    for i in range(len(parent)):
        if parent[i]:
            result.append((i, parent[i], distance[i]))
    return result
""")
    click.echo("Copied!")


@cli.command()
def euler():
    pyperclip.copy("""
# Złożoność: O(V^2) (macierz)
def dfs(G, s, result: list):
    for i in range(len(G)):
        if G[s][i] == 1:
            G[s][i], G[i][s] = 0, 0
            dfs(G, i, result)
    result.append(s)


def euler_path(G):
    n = len(G)
    edges = 0
    for i in range(n):
        for j in range(n):
            if G[i][j] == 1:
                edges += 1
    if edges % 2 == 1:
        return False
    result = []
    dfs(G, 0, result)
    return result[::-1]
""")
    click.echo("Copied!")


@cli.command()
def list_kruskal():
    pyperclip.copy("""
# Złożoność: O(ELogV)
class Node:
    def __init__(self, value):
        self.value = value
        self.rank = 0
        self.parent = self


def find(x: Node):
    if x != x.parent:
        x.parent = find(x.parent)
    return x.parent


def union(x: Node, y: Node):
    x = find(x)
    y = find(y)
    if x == y:
        return
    if x.rank > y.rank:
        y.parent = x
    else:
        x.parent = y
        if x.rank == y.rank:
            y.rank += 1


def make_set(v):
    return Node(v)


def convert_to_edges(G):
    E = []
    for i in range(len(G)):
        for j in range(len(G[i])):
            if (G[i][j][0], i, G[i][j][1]) not in E:
                E.append((i, G[i][j][0], G[i][j][1]))
    return E


def kruskal(G):
    E = convert_to_edges(G)
    E.sort(key=lambda x: x[2])
    MST = []
    V = []
    for i in range(len(G)):
        V.append(make_set(i))
    for i in range(len(E)):
        u = E[i][0]
        v = E[i][1]
        if find(V[u]) != find(V[v]):
            MST.append(E[i])
            union(V[u], V[v])
    return MST
""")
    click.echo("Copied!")
