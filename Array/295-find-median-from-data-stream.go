type MedianFinder struct {
    hpMax Heap
    hpMin Heap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{
        hpMax: NewMaxHeap(),
        hpMin: NewMinHeap(),
    }
}


func (this *MedianFinder) AddNum(num int)  {
    l := this.Len()
    if l&1 == 0 {
        if this.hpMin.Count() > 0 && this.hpMin.TopHeap() < num {
            this.hpMin.PushBack(num)
            num = this.hpMin.PopHead()
        }
        this.hpMax.PushBack(num)
    } else {
        if this.hpMax.Count() > 0 && this.hpMax.TopHeap() > num {
            this.hpMax.PushBack(num)
            num = this.hpMax.PopHead()
        }
        this.hpMin.PushBack(num)
    }
    return
}


func (this *MedianFinder) FindMedian() float64 {
    l := this.Len()
    if l == 0 {
        return 0
    }
    // if this.hpMin.Count() > 2 {
    //     return float64(this.hpMin.GetData(2))
    // }
    if l&1 == 0 {
        return (float64(this.hpMax.TopHeap()) + float64(this.hpMin.TopHeap()))/2
    } else {
        return float64(this.hpMax.TopHeap())
    }
    // return -1.0
}

func (this *MedianFinder) Len() int {
    return this.hpMax.Count() + this.hpMin.Count()
}

type Heap interface {
    PushBack(num int)
    PopHead() int
    TopHeap() int
    Swap(i, j int)
    Count() int
    GetData(i int) int
}


type MaxHeap struct {
    count int
    data []int
}


func NewMaxHeap() Heap {
    return &MaxHeap{
        count: 0,
        data: []int{0},
    }
}

func (hp *MaxHeap) Count() int {
    return hp.count
}


func (hp *MaxHeap) GetData(i int) int {
    return hp.data[i]
}


func (hp *MaxHeap) PushBack(num int) {
    if len(hp.data) < hp.count+2 {
        hp.data = append(hp.data, 0)
    }
    hp.count++
    hp.data[hp.count] = num
    for i := hp.count; i >= 1; {
        if i/2 >= 1 && hp.data[i/2] < hp.data[i] {
            hp.Swap(i, i/2)
            i /= 2
        } else {
            break
        }
    }
    return
}


func (hp *MaxHeap) PopHead() int {
    ret := hp.data[1]
    hp.data[1] = hp.data[hp.count]
    hp.count--
    i, pos := 1, 1
    for {
        pos = i
        if i*2 <= hp.count && hp.data[i*2] > hp.data[i] {
            pos = i*2
        }
        if i*2+1 <= hp.count && hp.data[i*2+1] > hp.data[pos] {
            pos = i*2+1
        }
        if i != pos {
            hp.Swap(i, pos)
        } else {
            break
        }
        i = pos
    }
    return ret
}

func (hp *MaxHeap) TopHeap() int {
    return hp.data[1]
}

func (hp *MaxHeap) Swap(i, j int) {
    hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}


type MinHeap struct {
    count int
    data []int
}


func NewMinHeap() Heap {
    return &MinHeap{
        count: 0,
        data: []int{0},
    }
}


func (hp *MinHeap) GetData(i int) int {
    return hp.data[i]
}


func (hp *MinHeap) Count() int {
    return hp.count
}


func (hp *MinHeap) PushBack(num int) {
    if len(hp.data) < hp.count+2 {
        hp.data = append(hp.data, 0)
    }
    hp.count++
    hp.data[hp.count] = num
    for i := hp.count; i >= 1; {
        if i/2 >= 1 && hp.data[i/2] > hp.data[i] {
            hp.Swap(i, i/2)
            i /= 2
        } else {
            break
        }
    }
    return
}


func (hp *MinHeap) PopHead() int {
    ret := hp.data[1]
    hp.data[1] = hp.data[hp.count]
    hp.count--
    i, pos := 1, 1
    for {
        pos = i
        if i*2 <= hp.count && hp.data[i*2] < hp.data[i] {
            pos = i*2
        }
        if i*2+1 <= hp.count && hp.data[i*2+1] < hp.data[pos] {
            pos = i*2+1
        }
        if i != pos {
            hp.Swap(i, pos)
        } else {
            break
        }
        i = pos
    }
    return ret
}

func (hp *MinHeap) TopHeap() int {
    return hp.data[1]
}

func (hp *MinHeap) Swap(i, j int) {
    hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}



/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */


























