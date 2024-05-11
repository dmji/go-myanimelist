package mal_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionAnimeFields(t *testing.T) {
	af := prm.AnimeFields{}
	tests := []struct {
		name string
		in   []string
		out  string
	}{
		{
			name: "ID",
			in: []string{
				af.ID(),
			},
			out: "fields=id",
		},
		{
			name: "Title",
			in: []string{
				af.Title(),
			},
			out: "fields=title",
		},
		{
			name: "MainPicture",
			in: []string{
				af.MainPicture(),
			},
			out: "fields=main_picture",
		},
		{
			name: "AlternativeTitles",
			in: []string{
				af.AlternativeTitles(),
			},
			out: "fields=alternative_titles",
		},
		{
			name: "StartDate",
			in: []string{
				af.StartDate(),
			},
			out: "fields=start_date",
		},
		{
			name: "EndDate",
			in: []string{
				af.EndDate(),
			},
			out: "fields=end_date",
		},
		{
			name: "Synopsis",
			in: []string{
				af.Synopsis(),
			},
			out: "fields=synopsis",
		},
		{
			name: "Mean",
			in: []string{
				af.Mean(),
			},
			out: "fields=mean",
		},
		{
			name: "Rank",
			in: []string{
				af.Rank(),
			},
			out: "fields=rank",
		},
		{
			name: "Popularity",
			in: []string{
				af.Popularity(),
			},
			out: "fields=popularity",
		},
		{
			name: "NumListUsers",
			in: []string{
				af.NumListUsers(),
			},
			out: "fields=num_list_users",
		},
		{
			name: "NumScoringUsers",
			in: []string{
				af.NumScoringUsers(),
			},
			out: "fields=num_scoring_users",
		},
		{
			name: "NSFW",
			in: []string{
				af.NSFW(),
			},
			out: "fields=nsfw",
		},
		{
			name: "CreatedAt",
			in: []string{
				af.CreatedAt(),
			},
			out: "fields=created_at",
		},
		{
			name: "UpdatedAt",
			in: []string{
				af.UpdatedAt(),
			},
			out: "fields=updated_at",
		},
		{
			name: "MediaType",
			in: []string{
				af.MediaType(),
			},
			out: "fields=media_type",
		},
		{
			name: "Status",
			in: []string{
				af.Status(),
			},
			out: "fields=status",
		},
		{
			name: "Genres",
			in: []string{
				af.Genres(),
			},
			out: "fields=genres",
		},
		{
			name: "MyListStatus",
			in: []string{
				af.MyListStatus(),
			},
			out: "fields=my_list_status",
		},
		{
			name: "NumEpisodes",
			in: []string{
				af.NumEpisodes(),
			},
			out: "fields=num_episodes",
		},
		{
			name: "StartSeason",
			in: []string{
				af.StartSeason(),
			},
			out: "fields=start_season",
		},
		{
			name: "Broadcast",
			in: []string{
				af.Broadcast(),
			},
			out: "fields=broadcast",
		},
		{
			name: "Source",
			in: []string{
				af.Source(),
			},
			out: "fields=source",
		},
		{
			name: "AverageEpisodeDuration",
			in: []string{
				af.AverageEpisodeDuration(),
			},
			out: "fields=average_episode_duration",
		},
		{
			name: "Rating",
			in: []string{
				af.Rating(),
			},
			out: "fields=rating",
		},
		{
			name: "Pictures",
			in: []string{
				af.Pictures(),
			},
			out: "fields=pictures",
		},
		{
			name: "Background",
			in: []string{
				af.Background(),
			},
			out: "fields=background",
		},
		{
			name: "RelatedAnime",
			in: []string{
				af.RelatedAnime(),
			},
			out: "fields=related_anime",
		},
		{
			name: "RelatedManga",
			in: []string{
				af.RelatedManga(),
			},
			out: "fields=related_manga",
		},
		{
			name: "Recommendations",
			in: []string{
				af.Recommendations(),
			},
			out: "fields=recommendations",
		},
		{
			name: "Studios",
			in: []string{
				af.Studios(),
			},
			out: "fields=studios",
		},
		{
			name: "Statistics",
			in: []string{
				af.Statistics(),
			},
			out: "fields=statistics",
		},
		{
			name: "ID Single Arg",
			in: []string{
				af.ID("arg"),
			},
			out: "fields=id{arg}",
		},
		{
			name: "Title Single Arg",
			in: []string{
				af.Title("arg"),
			},
			out: "fields=title{arg}",
		},
		{
			name: "MainPicture Single Arg",
			in: []string{
				af.MainPicture("arg"),
			},
			out: "fields=main_picture{arg}",
		},
		{
			name: "AlternativeTitles Single Arg",
			in: []string{
				af.AlternativeTitles("arg"),
			},
			out: "fields=alternative_titles{arg}",
		},
		{
			name: "StartDate Single Arg",
			in: []string{
				af.StartDate("arg"),
			},
			out: "fields=start_date{arg}",
		},
		{
			name: "EndDate Single Arg",
			in: []string{
				af.EndDate("arg"),
			},
			out: "fields=end_date{arg}",
		},
		{
			name: "Synopsis Single Arg",
			in: []string{
				af.Synopsis("arg"),
			},
			out: "fields=synopsis{arg}",
		},
		{
			name: "Mean Single Arg",
			in: []string{
				af.Mean("arg"),
			},
			out: "fields=mean{arg}",
		},
		{
			name: "Rank Single Arg",
			in: []string{
				af.Rank("arg"),
			},
			out: "fields=rank{arg}",
		},
		{
			name: "Popularity Single Arg",
			in: []string{
				af.Popularity("arg"),
			},
			out: "fields=popularity{arg}",
		},
		{
			name: "NumListUsers Single Arg",
			in: []string{
				af.NumListUsers("arg"),
			},
			out: "fields=num_list_users{arg}",
		},
		{
			name: "NumScoringUsers Single Arg",
			in: []string{
				af.NumScoringUsers("arg"),
			},
			out: "fields=num_scoring_users{arg}",
		},
		{
			name: "NSFW Single Arg",
			in: []string{
				af.NSFW("arg"),
			},
			out: "fields=nsfw{arg}",
		},
		{
			name: "CreatedAt Single Arg",
			in: []string{
				af.CreatedAt("arg"),
			},
			out: "fields=created_at{arg}",
		},
		{
			name: "UpdatedAt Single Arg",
			in: []string{
				af.UpdatedAt("arg"),
			},
			out: "fields=updated_at{arg}",
		},
		{
			name: "MediaType Single Arg",
			in: []string{
				af.MediaType("arg"),
			},
			out: "fields=media_type{arg}",
		},
		{
			name: "Status Single Arg",
			in: []string{
				af.Status("arg"),
			},
			out: "fields=status{arg}",
		},
		{
			name: "Genres Single Arg",
			in: []string{
				af.Genres("arg"),
			},
			out: "fields=genres{arg}",
		},
		{
			name: "MyListStatus Single Arg",
			in: []string{
				af.MyListStatus("arg"),
			},
			out: "fields=my_list_status{arg}",
		},
		{
			name: "NumEpisodes Single Arg",
			in: []string{
				af.NumEpisodes("arg"),
			},
			out: "fields=num_episodes{arg}",
		},
		{
			name: "StartSeason Single Arg",
			in: []string{
				af.StartSeason("arg"),
			},
			out: "fields=start_season{arg}",
		},
		{
			name: "Broadcast Single Arg",
			in: []string{
				af.Broadcast("arg"),
			},
			out: "fields=broadcast{arg}",
		},
		{
			name: "Source Single Arg",
			in: []string{
				af.Source("arg"),
			},
			out: "fields=source{arg}",
		},
		{
			name: "AverageEpisodeDuration Single Arg",
			in: []string{
				af.AverageEpisodeDuration("arg"),
			},
			out: "fields=average_episode_duration{arg}",
		},
		{
			name: "Rating Single Arg",
			in: []string{
				af.Rating("arg"),
			},
			out: "fields=rating{arg}",
		},
		{
			name: "Pictures Single Arg",
			in: []string{
				af.Pictures("arg"),
			},
			out: "fields=pictures{arg}",
		},
		{
			name: "Background Single Arg",
			in: []string{
				af.Background("arg"),
			},
			out: "fields=background{arg}",
		},
		{
			name: "RelatedAnime Single Arg",
			in: []string{
				af.RelatedAnime("arg"),
			},
			out: "fields=related_anime{arg}",
		},
		{
			name: "RelatedManga Single Arg",
			in: []string{
				af.RelatedManga("arg"),
			},
			out: "fields=related_manga{arg}",
		},
		{
			name: "Recommendations Single Arg",
			in: []string{
				af.Recommendations("arg"),
			},
			out: "fields=recommendations{arg}",
		},
		{
			name: "Studios Single Arg",
			in: []string{
				af.Studios("arg"),
			},
			out: "fields=studios{arg}",
		},
		{
			name: "Statistics Single Arg",
			in: []string{
				af.Statistics("arg"),
			},
			out: "fields=statistics{arg}",
		},
		{
			name: "ID Two Args",
			in: []string{
				af.ID("arg1", "arg2"),
			},
			out: "fields=id{arg1,arg2}",
		},
		{
			name: "Title Two Args",
			in: []string{
				af.Title("arg1", "arg2"),
			},
			out: "fields=title{arg1,arg2}",
		},
		{
			name: "MainPicture Two Args",
			in: []string{
				af.MainPicture("arg1", "arg2"),
			},
			out: "fields=main_picture{arg1,arg2}",
		},
		{
			name: "AlternativeTitles Two Args",
			in: []string{
				af.AlternativeTitles("arg1", "arg2"),
			},
			out: "fields=alternative_titles{arg1,arg2}",
		},
		{
			name: "StartDate Two Args",
			in: []string{
				af.StartDate("arg1", "arg2"),
			},
			out: "fields=start_date{arg1,arg2}",
		},
		{
			name: "EndDate Two Args",
			in: []string{
				af.EndDate("arg1", "arg2"),
			},
			out: "fields=end_date{arg1,arg2}",
		},
		{
			name: "Synopsis Two Args",
			in: []string{
				af.Synopsis("arg1", "arg2"),
			},
			out: "fields=synopsis{arg1,arg2}",
		},
		{
			name: "Mean Two Args",
			in: []string{
				af.Mean("arg1", "arg2"),
			},
			out: "fields=mean{arg1,arg2}",
		},
		{
			name: "Rank Two Args",
			in: []string{
				af.Rank("arg1", "arg2"),
			},
			out: "fields=rank{arg1,arg2}",
		},
		{
			name: "Popularity Two Args",
			in: []string{
				af.Popularity("arg1", "arg2"),
			},
			out: "fields=popularity{arg1,arg2}",
		},
		{
			name: "NumListUsers Two Args",
			in: []string{
				af.NumListUsers("arg1", "arg2"),
			},
			out: "fields=num_list_users{arg1,arg2}",
		},
		{
			name: "NumScoringUsers Two Args",
			in: []string{
				af.NumScoringUsers("arg1", "arg2"),
			},
			out: "fields=num_scoring_users{arg1,arg2}",
		},
		{
			name: "NSFW Two Args",
			in: []string{
				af.NSFW("arg1", "arg2"),
			},
			out: "fields=nsfw{arg1,arg2}",
		},
		{
			name: "CreatedAt Two Args",
			in: []string{
				af.CreatedAt("arg1", "arg2"),
			},
			out: "fields=created_at{arg1,arg2}",
		},
		{
			name: "UpdatedAt Two Args",
			in: []string{
				af.UpdatedAt("arg1", "arg2"),
			},
			out: "fields=updated_at{arg1,arg2}",
		},
		{
			name: "MediaType Two Args",
			in: []string{
				af.MediaType("arg1", "arg2"),
			},
			out: "fields=media_type{arg1,arg2}",
		},
		{
			name: "Status Two Args",
			in: []string{
				af.Status("arg1", "arg2"),
			},
			out: "fields=status{arg1,arg2}",
		},
		{
			name: "Genres Two Args",
			in: []string{
				af.Genres("arg1", "arg2"),
			},
			out: "fields=genres{arg1,arg2}",
		},
		{
			name: "MyListStatus Two Args",
			in: []string{
				af.MyListStatus("arg1", "arg2"),
			},
			out: "fields=my_list_status{arg1,arg2}",
		},
		{
			name: "NumEpisodes Two Args",
			in: []string{
				af.NumEpisodes("arg1", "arg2"),
			},
			out: "fields=num_episodes{arg1,arg2}",
		},
		{
			name: "StartSeason Two Args",
			in: []string{
				af.StartSeason("arg1", "arg2"),
			},
			out: "fields=start_season{arg1,arg2}",
		},
		{
			name: "Broadcast Two Args",
			in: []string{
				af.Broadcast("arg1", "arg2"),
			},
			out: "fields=broadcast{arg1,arg2}",
		},
		{
			name: "Source Two Args",
			in: []string{
				af.Source("arg1", "arg2"),
			},
			out: "fields=source{arg1,arg2}",
		},
		{
			name: "AverageEpisodeDuration Two Args",
			in: []string{
				af.AverageEpisodeDuration("arg1", "arg2"),
			},
			out: "fields=average_episode_duration{arg1,arg2}",
		},
		{
			name: "Rating Two Args",
			in: []string{
				af.Rating("arg1", "arg2"),
			},
			out: "fields=rating{arg1,arg2}",
		},
		{
			name: "Pictures Two Args",
			in: []string{
				af.Pictures("arg1", "arg2"),
			},
			out: "fields=pictures{arg1,arg2}",
		},
		{
			name: "Background Two Args",
			in: []string{
				af.Background("arg1", "arg2"),
			},
			out: "fields=background{arg1,arg2}",
		},
		{
			name: "RelatedAnime Two Args",
			in: []string{
				af.RelatedAnime("arg1", "arg2"),
			},
			out: "fields=related_anime{arg1,arg2}",
		},
		{
			name: "RelatedManga Two Args",
			in: []string{
				af.RelatedManga("arg1", "arg2"),
			},
			out: "fields=related_manga{arg1,arg2}",
		},
		{
			name: "Recommendations Two Args",
			in: []string{
				af.Recommendations("arg1", "arg2"),
			},
			out: "fields=recommendations{arg1,arg2}",
		},
		{
			name: "Studios Two Args",
			in: []string{
				af.Studios("arg1", "arg2"),
			},
			out: "fields=studios{arg1,arg2}",
		},
		{
			name: "Statistics Two Args",
			in: []string{
				af.Statistics("arg1", "arg2"),
			},
			out: "fields=statistics{arg1,arg2}",
		},
		{
			name: "ID+Title Two Args",
			in: []string{
				af.ID("arg1", "arg2"),
				af.Title("arg1", "arg2"),
			},
			out: "fields=id{arg1,arg2},title{arg1,arg2}",
		},
		{
			name: "Title+ID Two Args",
			in: []string{
				af.Title("arg1", "arg2"),
				af.ID("arg1", "arg2"),
			},
			out: "fields=title{arg1,arg2},id{arg1,arg2}",
		},
		{
			name: "ID+Title",
			in: []string{
				af.ID(),
				af.Title(),
			},
			out: "fields=id,title",
		},
		{
			name: "ID+Title+MainPicture",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
			},
			out: "fields=id,title,main_picture",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
			},
			out: "fields=id,title,main_picture,alternative_titles",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures+Background",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures+Background+RelatedAnime",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
				af.RelatedAnime(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures+Background+RelatedAnime+RelatedManga",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
				af.RelatedAnime(),
				af.RelatedManga(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures+Background+RelatedAnime+RelatedManga+Recommendations",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
				af.RelatedAnime(),
				af.RelatedManga(),
				af.Recommendations(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga,recommendations",
		},
		{
			name: "ID+Title+MainPicture+AlternativeTitles+StartDate+EndDate+Synopsis+Mean+Rank+Popularity+NumListUsers+NumScoringUsers+NSFW+CreatedAt+UpdatedAt+MediaType+Status+Genres+MyListStatus+NumEpisodes+StartSeason+Broadcast+Source+AverageEpisodeDuration+Rating+Pictures+Background+RelatedAnime+RelatedManga+Recommendations+Studios",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
				af.RelatedAnime(),
				af.RelatedManga(),
				af.Recommendations(),
				af.Studios(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga,recommendations,studios",
		},
		{
			name: "All",
			in: []string{
				af.ID(),
				af.Title(),
				af.MainPicture(),
				af.AlternativeTitles(),
				af.StartDate(),
				af.EndDate(),
				af.Synopsis(),
				af.Mean(),
				af.Rank(),
				af.Popularity(),
				af.NumListUsers(),
				af.NumScoringUsers(),
				af.NSFW(),
				af.CreatedAt(),
				af.UpdatedAt(),
				af.MediaType(),
				af.Status(),
				af.Genres(),
				af.MyListStatus(),
				af.NumEpisodes(),
				af.StartSeason(),
				af.Broadcast(),
				af.Source(),
				af.AverageEpisodeDuration(),
				af.Rating(),
				af.Pictures(),
				af.Background(),
				af.RelatedAnime(),
				af.RelatedManga(),
				af.Recommendations(),
				af.Studios(),
				af.Statistics(),
			},
			out: "fields=id,title,main_picture,alternative_titles,start_date,end_date,synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at,updated_at,media_type,status,genres,my_list_status,num_episodes,start_season,broadcast,source,average_episode_duration,rating,pictures,background,related_anime,related_manga,recommendations,studios,statistics",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := url.Values{}
			prm.NewFields(tt.in...).Apply(&v)
			got := v.Encode()
			got = strings.ReplaceAll(got, "%2C", ",")
			got = strings.ReplaceAll(got, "%7B", "{")
			got = strings.ReplaceAll(got, "%7D", "}")
			want := tt.out

			if got != want {
				t.Errorf("AnimeFields expected '%s', got '%s'", want, got)
			}
		})
	}
}
