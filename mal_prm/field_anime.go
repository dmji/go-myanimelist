package mal_prm

import qs "github.com/dmji/qs"

//go:generate go run github.com/dmji/go-stringer@latest -type=AnimeFieldType -trimprefix=AnimeFieldType -output field_anime_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

type AnimeFieldType uint8

const (
	AnimeFieldTypeID                     AnimeFieldType = iota // "id" + argJoin(p...)
	AnimeFieldTypeTitle                                        // "title" + argJoin(p...)
	AnimeFieldTypeMainPicture                                  // "main_picture" + argJoin(p...)
	AnimeFieldTypeAlternativeTitles                            // "alternative_titles" + argJoin(p...)
	AnimeFieldTypeStartDate                                    // "start_date" + argJoin(p...)
	AnimeFieldTypeEndDate                                      // "end_date" + argJoin(p...)
	AnimeFieldTypeSynopsis                                     // "synopsis" + argJoin(p...)
	AnimeFieldTypeMean                                         // "mean" + argJoin(p...)
	AnimeFieldTypeRank                                         // "rank" + argJoin(p...)
	AnimeFieldTypePopularity                                   // "popularity" + argJoin(p...)
	AnimeFieldTypeNumListUsers                                 // "num_list_users" + argJoin(p...)
	AnimeFieldTypeNumScoringUsers                              // "num_scoring_users" + argJoin(p...)
	AnimeFieldTypeNSFW                                         // "nsfw" + argJoin(p...)
	AnimeFieldTypeCreatedAt                                    // "created_at" + argJoin(p...)
	AnimeFieldTypeUpdatedAt                                    // "updated_at" + argJoin(p...)
	AnimeFieldTypeMediaType                                    // "media_type" + argJoin(p...)
	AnimeFieldTypeStatus                                       // "status" + argJoin(p...)
	AnimeFieldTypeGenres                                       // "genres" + argJoin(p...)
	AnimeFieldTypeMyListStatus                                 // "my_list_status" + argJoin(p...)
	AnimeFieldTypeNumEpisodes                                  // "num_episodes" + argJoin(p...)
	AnimeFieldTypeStartSeason                                  // "start_season" + argJoin(p...)
	AnimeFieldTypeBroadcast                                    // "broadcast" + argJoin(p...)
	AnimeFieldTypeSource                                       // "source" + argJoin(p...)
	AnimeFieldTypeAverageEpisodeDuration                       // "average_episode_duration" + argJoin(p...)
	AnimeFieldTypeRating                                       // "rating" + argJoin(p...)
	AnimeFieldTypePictures                                     // "pictures" + argJoin(p...)
	AnimeFieldTypeBackground                                   // "background" + argJoin(p...)
	AnimeFieldTypeRelatedAnime                                 // "related_anime" + argJoin(p...)
	AnimeFieldTypeRelatedManga                                 // "related_manga" + argJoin(p...)
	AnimeFieldTypeRecommendations                              // "recommendations" + argJoin(p...)
	AnimeFieldTypeStudios                                      // "studios" + argJoin(p...)
	AnimeFieldTypeStatistics                                   // "statistics" + argJoin(p...)

	// User Specific
	AnimeFieldTypeListStatus
	AnimeFieldTypeNode
)

func (f AnimeFieldType) AnimeField(args ...string) AnimeField {
	return AnimeField{
		Field: f,
		Args:  args,
	}
}

type AnimeField struct {
	Field AnimeFieldType
	Args  []string
}

func (e AnimeField) MarshalQS(opts *qs.MarshalOptions) ([]string, error) {
	return []string{e.Field.String() + argJoin(e.Args...)}, nil
}
