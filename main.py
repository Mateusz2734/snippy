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
        return quickselect(T, q + 1, r, i - k)""",
    "matrix_bfs": """def BFS(G, s):
    n = len(G)
    visited = [False for _ in range(n)]
    Q = deque()
    visited[s] = True

    while Q:
        vis = Q.popleft()
        print(vis, end = ' ')
        for i in range(n):
            if (G[vis][i] == 1 and
                    (not visited[i])):
                Q.append(i)
                visited[i] = True
    return visited""",
    "list_bfs": """def bfs(G, s):
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
    return (parent, visited)""",
    "topological_sort": """# Złożoność: O(V + E)
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
    return sorted_graph""",
    "dijkstra": """#Złożoność: O(V*E*logV)
    def relax(u, v, l, d, parent, queue: PriorityQueue):
    if d[v] > d[u] + l:
        d[v] = d[u] + l
        parent[v] = u
        queue.put((d[v], v))


def dijkstra(G, s):
    n = len(G)
    parent = [None for _ in range(n)]
    d = [inf for _ in range(n)]

    d[s] = 0
    queue = PriorityQueue()
    queue.put((d[s], s))

    while queue:
        du, u = queue.get()
        if d[u] == du:
            for v, l in G[u]:
                relax(u, v, l, d, parent, queue)
    return (d, parent)""",
    "ford_fulkerson": """# Złożoność: O(E*max_flow)
    # Jeśli korzystamy z BFS, to O(V*E^2)
    def find_path(G, s, t):
    # BFS lub DFS
    pass

def min_weight(G, path):
    w = G[path[0]][path[1]]
    for i in range(1, len(path) - 1):
        w = min(w, G[path[i]][path[i + 1]])
    return w

def update_weights(G, path):
    w = min_weight(G, path)
    for i in range(len(path) - 1):
        G[path[i]][path[i + 1]] -= w
        G[path[i + 1]][path[i]] += w

def ford_fulkerson(M, s, t):
    n = len(M)
    G = deepcopy(M)
    
    flow = 0
    my_path = find_path(G, s, t)
    while my_path:
        flow += min_weight(G, my_path)
        update_weights(G, my_path)

        my_path = find_path(G, s, t)
    return flow""",
    "list_dfs": """# Złożoność: O(V + E)
    def DFS(G, s):
    visited = [False for _ in range(len(G))]
    result = [s]

    def dfs_visit(u, G, visited, result):
        visited[u] = True 
        for v in G[u]:
            if not visited[v]:
                result.append(v)
                dfs_visit(v, G, visited, result)

    dfs_visit(s, G, visited, result)
    return result""",
    "matrix_dfs": """#Złożoność: O(V^2)
    def DFS(G, s):
    visited = [False for _ in range(len(G))]
    result = []
    dfs_visit(s, G, visited, result)
    return result


def dfs_visit(u, G, visited, result):
    visited[u] = True
    result.append(u)
    for i in range(len(G)):
        if visited[i] is False and G[u][i] == 1:
            dfs_visit(i, G, visited, result)""",
    "find_cycle": """# Złożoność: O(V + E) (Jak DFS)
def detect_cycle(G, s):
    visited = [False for _ in range(len(G))]
    parent = [None for _ in range(len(G))]
    return dfs(G, s, visited, parent)


def dfs(G, s, visited, parent):
    visited[s] = True
    for v in G[s]:
        if not visited[v]:
            parent[v] = s
            return dfs(G, v, visited, parent)
        elif visited[v] and parent[s] != v:
            return True
    return False""",
    "find_bridges": """# Złożoność: O(V(V + E))
from math import inf


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
""",
    "floyd_warshall": """# Złożoność: O(V^3)
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
                distance[u][v] = min(distance[u][v], distance[u][k] + distance[k][v])

    return distance""",
    "bellman_ford": """# Złożoność: O(V*E)
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
    """,
    "list_prim": """# Złożoność: O(ElogV)
from queue import PriorityQueue
from math import inf


def list_prim(G, s):
    n = len(G)
    queue = PriorityQueue()

    parent = [None for _ in range(n)]
    distance = [inf for _ in range(n)]
    visited = [False for _ in range(n)]

    distance[s] = 0
    visited[s] = True
    queue.put((0, s))

    while not queue.empty():
        _, u = queue.get()
        visited[u] = True
        for v in G[u]:
            if distance[v[0]] > v[1] and not visited[v[0]]:
                parent[v[0]] = u
                distance[v[0]] = v[1]
                queue.put((distance[v[0]], v[0]))

    result = []
    for i in range(len(parent)):
        if parent[i] is not None:
            result.append((i, parent[i], distance[i]))
    return result""",
    "matrix_prim": """# Złożoność: O(V^2)
from queue import PriorityQueue
from math import inf


def matrix_prim(G, s):
    n = len(G)

    queue = PriorityQueue()
    parent = [None for _ in range(n)]
    distance = [inf for _ in range(n)]
    visited = [False for _ in range(n)]
    
    distance[s] = 0
    visited[s] = True
    queue.put((0, s))
    
    while not queue.empty():
        _, u = queue.get()
        visited[u] = True
        for i in range(len(G)):
            if G[u][i] != 0 and distance[i] > G[u][i] and not visited[i]:
                parent[i] = u
                distance[i] = G[u][i]
                queue.put((distance[i], i))

    result = []
    for i in range(len(parent)):
        if parent[i] is not None:
            result.append((i, parent[i], distance[i]))
    return result""",
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


@cli.command()
def matrix_bfs():
    pyperclip.copy(to_copy["matrix_bfs"])
    click.echo("Copied!")


@cli.command()
def list_bfs():
    pyperclip.copy(to_copy["list_bfs"])
    click.echo("Copied!")


@cli.command()
def list_dfs():
    pyperclip.copy(to_copy["list_dfs"])
    click.echo("Copied!")


@cli.command()
def matrix_dfs():
    pyperclip.copy(to_copy["matrix_dfs"])
    click.echo("Copied!")


@cli.command()
def topological_sort():
    pyperclip.copy(to_copy["topological_sort"])
    click.echo("Copied!")


@cli.command()
def dijkstra():
    pyperclip.copy(to_copy["dijkstra"])
    click.echo("Copied!")


@cli.command()
def ford_fulkerson():
    pyperclip.copy(to_copy["ford_fulkerson"])
    click.echo("Copied!")


@cli.command()
def find_cycle():
    pyperclip.copy(to_copy["find_cycle"])
    click.echo("Copied!")


@cli.command()
def find_bridges():
    pyperclip.copy(to_copy["find_bridges"])
    click.echo("Copied!")


@cli.command()
def floyd_warshall():
    pyperclip.copy(to_copy["floyd_warshall"])
    click.echo("Copied!")


@cli.command()
def bellman_ford():
    pyperclip.copy(to_copy["bellman_ford"])
    click.echo("Copied!")


@cli.command()
def list_prim():
    pyperclip.copy(to_copy["list_prim"])
    click.echo("Copied!")


@cli.command()
def matrix_prim():
    pyperclip.copy(to_copy["matrix_prim"])
    click.echo("Copied!")
