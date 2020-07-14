package ring

import (
	"testing"
)

func TestRing(t *testing.T) {
	r := NewRing(1)
	t.Log(r)
	r.Insert(2)
	r = r.Next()
	r.Insert(3)
	r = r.Next()
	r.Insert(4)
	r = r.Next()
	r.Insert(5)
	r = r.Next()
	r.Insert(6)
	r = r.Next()
	r.Insert(7)
	r = r.Next()
	r.Insert(8)
	r = r.Next()
	r.Insert(9)
	r = r.Next().Next()
	t.Log("添加结束")

	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())
	r = r.Next()
	t.Logf("地址：%p => 内容：%d \n", r, r.Value())

	p := r
	for {
		p = p.Next()
		t.Log(p.Value())
		if p == r {
			t.Log("相等")
			p.Next().Next().Next().Next().Next().Remove()
			t.Log("删除结束")
			break
		}
	}
	t.Log("==================================")
	for {
		p = p.Next()
		t.Log(p.Value())
		if p == r {
			break
		}
	}
	t.Log("==================================")
}
