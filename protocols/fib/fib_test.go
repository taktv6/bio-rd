package fib

import (
	"testing"

	bnet "github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/routingtable/vrf"

	"github.com/bio-routing/bio-rd/route"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f, err := New(nil)
	assert.Error(t, err, "expected error")
	assert.Nil(t, f, "expected nil")
}

func TestComparePfxPath(t *testing.T) {
	v, _ := vrf.New("inet.0", 0)
	f, _ := New(v)

	tests := []struct {
		name                string
		left                map[bnet.Prefix][]*route.FIBPath
		right               []*route.PrefixPathsPair
		inLeftButNotInRight bool
		expected            []*route.PrefixPathsPair
	}{
		{
			name: "in left but not in right",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 253),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: true,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 253),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
		{
			name: "in right but not in left",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 253),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: false,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 253),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
		{
			name: "left filled, right no paths, only in fib=true",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx:   bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: make([]*route.FIBPath, 0),
				},
			},
			inLeftButNotInRight: true,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
		{
			name: "left filled, right no paths, only in fib=false",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx:   bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: make([]*route.FIBPath, 0),
				},
			},
			inLeftButNotInRight: false,
			expected: []*route.PrefixPathsPair{
				{
					Pfx:   bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: make([]*route.FIBPath, 0),
				},
			},
		},
		{
			name: "left no paths, right filled, only in fib=true",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): make([]*route.FIBPath, 0),
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: true,
			expected: []*route.PrefixPathsPair{
				{
					Pfx:   bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: make([]*route.FIBPath, 0),
				},
			},
		},
		{
			name: "left no paths, right filled, only in fib=false",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): make([]*route.FIBPath, 0),
			},
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: false,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
		{
			name: "left yes, right nil, only in fib=true",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right:               make([]*route.PrefixPathsPair, 0),
			inLeftButNotInRight: true,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
		{
			name: "left yes, right nil, only in fib=false",
			left: map[bnet.Prefix][]*route.FIBPath{
				bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24): {
					{
						Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
						NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
						Protocol: route.ProtoKernel,
						Priority: 1,
						Table:    254,
						Type:     1,
					},
				},
			},
			right:               make([]*route.PrefixPathsPair, 0),
			inLeftButNotInRight: false,
			expected:            make([]*route.PrefixPathsPair, 0),
		},
		{
			name: "left nil, right yes, only in fib=true",
			left: make(map[bnet.Prefix][]*route.FIBPath, 0),
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: true,
			expected:            make([]*route.PrefixPathsPair, 0),
		},
		{
			name: "left nil, right yes, only in fib=false",
			left: make(map[bnet.Prefix][]*route.FIBPath, 0),
			right: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
			inLeftButNotInRight: false,
			expected: []*route.PrefixPathsPair{
				{
					Pfx: bnet.NewPfx(bnet.IPv4FromOctets(10, 0, 0, 1), 24),
					Paths: []*route.FIBPath{
						{
							Src:      bnet.IPv4FromOctets(10, 0, 0, 1),
							NextHop:  bnet.IPv4FromOctets(10, 0, 0, 254),
							Protocol: route.ProtoKernel,
							Priority: 1,
							Table:    254,
							Type:     1,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		f.pathTable = test.left

		expected := f.compareFibPfxPath(test.right, test.inLeftButNotInRight)
		assert.Equalf(t, test.expected, expected, test.name)
	}
}