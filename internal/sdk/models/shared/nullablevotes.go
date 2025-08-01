// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// NullableVotes - Votes model
type NullableVotes struct {
	// Whether or not the current actor has voted negatively
	Disliked *bool `json:"disliked,omitempty"`
	Dislikes *int  `json:"dislikes,omitempty"`
	// Whether or not the current actor has voted positively
	Liked *bool `json:"liked,omitempty"`
	Likes *int  `json:"likes,omitempty"`
	// Whether or not the current actor has voted
	Voted *bool `json:"voted,omitempty"`
}

func (o *NullableVotes) GetDisliked() *bool {
	if o == nil {
		return nil
	}
	return o.Disliked
}

func (o *NullableVotes) GetDislikes() *int {
	if o == nil {
		return nil
	}
	return o.Dislikes
}

func (o *NullableVotes) GetLiked() *bool {
	if o == nil {
		return nil
	}
	return o.Liked
}

func (o *NullableVotes) GetLikes() *int {
	if o == nil {
		return nil
	}
	return o.Likes
}

func (o *NullableVotes) GetVoted() *bool {
	if o == nil {
		return nil
	}
	return o.Voted
}
