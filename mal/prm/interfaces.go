package prm

import (
	"net/url"
	"strconv"
)

var itoa = strconv.Itoa

// MARK: Details Option
// DetailsOption is an option specific for the anime and manga details methods.
// +Fields
type DetailsOption interface {
	DetailsApply(v *url.Values)
}

// MARK: OptionalParam Option
// common.OptionalParam is implemented by types that can be used as options in most methods
// such as Limit, Offset and Fields.
// +OptionFunc
// +Fields
// +Limit
// +NSFW
// +Offset
type OptionalParam interface {
	Apply(v *url.Values)
}

// MARK: SeasonalAnime Option
// SeasonalAnimeOption are options specific to the Service.Seasonal method.
// +Fields
// +Limit
// +NSFW
// +Offset
// +SortSeasonalAnime
type SeasonalAnimeOption interface {
	SeasonalAnimeApply(v *url.Values)
}

// MARK: UpdateMyAnimeListStatus Option
// UpdateMyAnimeListStatusOption are options specific to the
// Service.UpdateMyListStatus method.
// +StartDate
// +FinishDate
// +AnimeStatus
// +Score
// +NumEpisodesWatched
// +NumTimesRewatched
// +IsRewatching
// +RewatchValue
// +Priority
// +Tags
// +Comments
type UpdateMyAnimeListStatusOption interface {
	UpdateMyAnimeListStatusApply(fnSet *url.Values)
}

// MARK: MyInfoOption Option
// MyInfoOption are options specific to the User.MyInfo method.
// +Fields
type MyInfoOption interface {
	MyInfoApply(v *url.Values)
}

// MARK: User MangaList Option
// MangaListOption are options specific to the UserService.MangaList method.
// +SortMangaList
// +Fields
// +Limit
// +MangaStatus
// +NSFW
// +Offset
type MangaListOption interface {
	MangaListApply(v *url.Values)
}

// MARK: User AnimeList Option
// AnimeListOption are options specific to the UserService.AnimeList method.
// +SortAnimeList
// +AnimeStatus
// +Fields
// +Limit
// +NSFW
// +Offset
type AnimeListOption interface {
	AnimeListApply(v *url.Values)
}

// MARK: UpdateMyMangaListStatus Option
// UpdateMyMangaListStatusOption are options specific to the
// MangaService.UpdateMyListStatus method.
// +Priority
// +Tags
// +Comments
// +FinishDate
// +IsRereading
// +MangaStatus
// +NumChaptersRead
// +NumTimesReread
// +NumVolumesRead
// +RereadValue
// +Score
// +StartDate
type UpdateMyMangaListStatusOption interface {
	UpdateMyMangaListStatusApply(v *url.Values)
}

// MARK: Topics Option
// TopicsOption are options specific to the ForumService.Topics method.
// +BoardID
// +SubboardID
// +sortTopics
// +Query
// +TopicUserName
// +UserName
type TopicsOption interface {
	TopicsApply(v *url.Values)
}

// A PagingOption includes the Limit and Offset options which are used for
// controlling pagination in results.
// +Limit
// +Offset
type PagingOption interface {
	PagingApply(v *url.Values)
}
