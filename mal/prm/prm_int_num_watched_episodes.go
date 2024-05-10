package prm

import "net/url"

// MARK: NumEpisodesWatched
// NumEpisodesWatched is an option that can update the number of episodes
// watched of an anime in the user's list.
type NumEpisodesWatched int

func (n NumEpisodesWatched) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("num_watched_episodes", itoa(int(n)))
}

func (n NumEpisodesWatched) Val(v int) NumEpisodesWatched { return NumEpisodesWatched(v) }
