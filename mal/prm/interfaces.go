package prm

import (
	"net/url"
	"strconv"
)

var itoa = strconv.Itoa

// MARK: Details Option
// DetailsOption is an option specific for the anime and manga details methods.
type DetailsOption interface {
	DetailsApply(v *url.Values)
}

type DetailsOptionProvider struct {
	Fields
}

// MARK: OptionalParam Option
// common.OptionalParam is implemented by types that can be used as options in most methods
// such as Limit, Offset and Fields.
type OptionalParam interface {
	Apply(v *url.Values)
}

type OptionalParamProvider struct {
	Fields
	Limit
	NSFW
	Offset
}

// MARK: SeasonalAnime Option
// SeasonalAnimeOption are options specific to the Service.Seasonal method.
// +Fields
type SeasonalAnimeOption interface {
	SeasonalAnimeApply(v *url.Values)
}

type SeasonalAnimeOptionProvider struct {
	// Required
	AnimeSeason

	// Optional
	Fields
	// Optional
	Limit
	// Optional
	NSFW
	// Optional
	Offset
	// Optional
	SortSeasonalAnime
}

// MARK: UpdateMyAnimeListStatus Option
// UpdateMyAnimeListStatusOption are options specific to the
// Service.UpdateMyListStatus method.
type UpdateMyAnimeListStatusOption interface {
	UpdateMyAnimeListStatusApply(fnSet *url.Values)
}

type UpdateMyAnimeListStatusOptionProvider struct {
	StartDate
	FinishDate
	AnimeStatus
	Score
	NumEpisodesWatched
	NumTimesRewatched
	IsRewatching
	RewatchValue
	Priority
	Tags
	Comments
}

// MARK: MyInfoOption Option
// MyInfoOption are options specific to the User.MyInfo method.
type MyInfoOption interface {
	MyInfoApply(v *url.Values)
}

type MyInfoOptionProvider struct {
	Fields
}

// MARK: User MangaList Option
// MangaListOption are options specific to the UserService.MangaList method.
type MangaListOption interface {
	MangaListApply(v *url.Values)
}

type MangaListOptionProvider struct {
	SortMangaList
	Fields
	Limit
	MangaStatus
	NSFW
	Offset
}

// MARK: User AnimeList Option
// AnimeListOption are options specific to the UserService.AnimeList method.
type AnimeListOption interface {
	AnimeListApply(v *url.Values)
}

type AnimeListOptionProvider struct {
	SortAnimeList
	AnimeStatus
	Fields
	Limit
	NSFW
	Offset
}

// MARK: UpdateMyMangaListStatus Option
// UpdateMyMangaListStatusOption are options specific to the
// MangaService.UpdateMyListStatus method.
type UpdateMyMangaListStatusOption interface {
	UpdateMyMangaListStatusApply(v *url.Values)
}

type UpdateMyMangaListStatusOptionProvider struct {
	Priority
	Tags
	Comments
	FinishDate
	IsRereading
	MangaStatus
	NumChaptersRead
	NumTimesReread
	NumVolumesRead
	RereadValue
	Score
	StartDate
}

// MARK: Topics Option
// TopicsOption are options specific to the ForumService.Topics method.
type TopicsOption interface {
	TopicsApply(v *url.Values)
}

type TopicsOptionProvider struct {
	BoardID
	SubboardID
	SortTopics sortTopics
	Query
	TopicUserName
	UserName
}

// A PagingOption includes the Limit and Offset options which are used for
// controlling pagination in results.
type PagingOption interface {
	PagingApply(v *url.Values)
}

type PagingOptionProvider struct {
	// Optional
	Limit
	// Optional
	Offset
}
