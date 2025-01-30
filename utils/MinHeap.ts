const MinHeap = {
  siftDown(h: [number, ...any[]][], i = 0, v = h[i]) {
    if (i < h.length) {
      let k = v[0];
      while (1) {
        let j = i * 2 + 1;
        if (j + 1 < h.length && h[j][0] > h[j + 1][0]) j++;
        if (j >= h.length || k <= h[j][0]) break;
        h[i] = h[j];
        i = j;
      }
      h[i] = v;
    }
  },
  heapify(h: [number, ...any[]][]) {
    for (let i = h.length >> 1; i--; ) this.siftDown(h, i);
    return h;
  },
  pop(h: [number, ...any[]][]) {
    return this.exchange(h, h.pop());
  },
  exchange(h: [number, ...any[]][], v: [number, ...any[]] | undefined) {
    if (!h.length) return v;
    let w = h[0];
    this.siftDown(h, 0, v);
    return w;
  },
  push(h: [number, ...any[]][], v: [number, ...any[]]) {
    let k = v[0],
      i = h.length,
      j;
    while ((j = (i - 1) >> 1) >= 0 && k < h[j][0]) {
      h[i] = h[j];
      i = j;
    }
    h[i] = v;
    return h;
  },
};

export default MinHeap;
