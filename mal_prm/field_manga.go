package mal_prm

import qs "github.com/dmji/qs"

//go:generate go run github.com/dmji/go-stringer@latest -type=MangaFieldType -trimprefix=MangaFieldType -output field_manga_string.go -nametransform=snake_case_lower -fromstringgenfn -marshaljson -marshalqs -marshalqspkg=github.com/dmji/qs -outputtransform=snake_case_lower

type MangaFieldType uint8

const (
	MangaFieldTypeID                MangaFieldType = iota // "id" + argJoin(p...)
	MangaFieldTypeTitle                                   //  "title" + argJoin(p...)
	MangaFieldTypeMainPicture                             //  "main_picture" + argJoin(p...)
	MangaFieldTypeAlternativeTitles                       //  "alternative_titles" + argJoin(p...)
	MangaFieldTypeStartDate                               //  "start_date" + argJoin(p...)
	MangaFieldTypeSynopsis                                //  "synopsis" + argJoin(p...)
	MangaFieldTypeMean                                    //  "mean" + argJoin(p...)
	MangaFieldTypeRank                                    //  "rank" + argJoin(p...)
	MangaFieldTypePopularity                              //  "popularity" + argJoin(p...)
	MangaFieldTypeNumListUsers                            //  "num_list_users" + argJoin(p...)
	MangaFieldTypeNumScoringUsers                         //  "num_scoring_users" + argJoin(p...)
	MangaFieldTypeNsfw                                    //  "nsfw" + argJoin(p...)
	MangaFieldTypeCreatedAt                               //  "created_at" + argJoin(p...)
	MangaFieldTypeUpdatedAt                               //  "updated_at" + argJoin(p...)
	MangaFieldTypeMediaType                               //  "media_type" + argJoin(p...)
	MangaFieldTypeStatus                                  //  "status" + argJoin(p...)
	MangaFieldTypeGenres                                  //  "genres" + argJoin(p...)
	MangaFieldTypeMyListStatus                            //  "my_list_status" + argJoin(p...)
	MangaFieldTypeNumVolumes                              //  "num_volumes" + argJoin(p...)
	MangaFieldTypeNumChapters                             //  "num_chapters" + argJoin(p...)
	MangaFieldTypeAuthors                                 //  "authors" + argJoin(p...)
	MangaFieldTypePictures                                //  "pictures" + argJoin(p...)
	MangaFieldTypeBackground                              //  "background" + argJoin(p...)
	MangaFieldTypeRelatedAnime                            //  "related_anime" + argJoin(p...)
	MangaFieldTypeRelatedManga                            //  "related_manga" + argJoin(p...)
	MangaFieldTypeRecommendations                         //  "recommendations" + argJoin(p...)
	MangaFieldTypeSerialization                           //  "serialization" + argJoin(p...)

	// User Specific
	MangaFieldTypeListStatus
	MangaFieldTypeNode
)

func (f MangaFieldType) MangaField(args ...string) MangaField {
	return MangaField{
		Field: f,
		Args:  args,
	}
}

type MangaField struct {
	Field MangaFieldType
	Args  []string
}

func (e MangaField) MarshalQS(opts *qs.MarshalOptions) ([]string, error) {
	return []string{e.Field.String() + argJoin(e.Args...)}, nil
}
