package actions

import (
	"github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/route"
)

type AddLargeCommunityAction struct {
	communities *route.LargeCommunities
}

func NewAddLargeCommunityAction(coms *route.LargeCommunities) *AddLargeCommunityAction {
	return &AddLargeCommunityAction{
		communities: coms,
	}
}

func (a *AddLargeCommunityAction) Do(p net.Prefix, pa *route.Path) Result {
	if pa.BGPPath == nil || len(*a.communities) == 0 {
		return Result{Path: pa}
	}

	modified := pa.Copy()
	if modified.BGPPath.LargeCommunities == nil {
		modified.BGPPath.LargeCommunities = &route.LargeCommunities{}
	}

	*modified.BGPPath.LargeCommunities = append(*modified.BGPPath.LargeCommunities, *a.communities...)
	return Result{Path: modified}
}
