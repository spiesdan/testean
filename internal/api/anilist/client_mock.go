package anilist

import (
	"context"
	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/seanime-app/seanime/internal/test_utils"
	"github.com/seanime-app/seanime/internal/util"
	"log"
	"os"
)

// This file contains helper functions for testing the anilist package

func TestGetMockAnilistClientWrapper() AnilistClient {
	return NewMockClientWrapper()
}

// MockClientWrapper is a mock implementation of the AnilistClient, used for tests.
// It uses the real implementation of the AnilistClient to make requests then populates a cache with the results.
// This is to avoid making repeated requests to the AniList API during tests but still have realistic data.
type MockClientWrapper struct {
	realClientWrapper AnilistClient
	logger            *zerolog.Logger
}

func NewMockClientWrapper() *MockClientWrapper {
	return &MockClientWrapper{
		realClientWrapper: NewClientWrapper(test_utils.ConfigData.Provider.AnilistJwt),
		logger:            util.NewLogger(),
	}
}

func (cw *MockClientWrapper) BaseAnimeByMalID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*BaseAnimeByMalID, error) {
	file, err := os.Open(test_utils.GetTestDataPath("BaseAnimeByMalID"))
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [BaseAnimeByMalID]: %d", *id)
			ret, err := cw.realClientWrapper.BaseAnimeByMalID(context.Background(), id)
			if err != nil {
				return nil, err
			}
			data, err := json.Marshal([]*BaseAnimeByMalID{ret})
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(test_utils.GetTestDataPath("BaseAnimeByMalID"), data, 0644)
			if err != nil {
				log.Fatal(err)
			}
			return ret, nil
		}
	}

	var media []*BaseAnimeByMalID
	err = json.NewDecoder(file).Decode(&media)
	if err != nil {
		log.Fatal(err)
	}
	var ret *BaseAnimeByMalID
	for _, m := range media {
		if m.GetMedia().ID == *id {
			ret = m
			break
		}
	}

	if ret == nil {
		cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [BaseAnimeByMalID]: %d", *id)
		ret, err := cw.realClientWrapper.BaseAnimeByMalID(context.Background(), id)
		if err != nil {
			return nil, err
		}
		media = append(media, ret)
		data, err := json.Marshal(media)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(test_utils.GetTestDataPath("BaseAnimeByMalID"), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return ret, nil
	}

	cw.logger.Trace().Msgf("MockClientWrapper: CACHE HIT [BaseAnimeByMalID]: %d", *id)
	return ret, nil
}

func (cw *MockClientWrapper) BaseAnimeByID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*BaseAnimeByID, error) {
	file, err := os.Open(test_utils.GetTestDataPath("BaseAnimeByID"))
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [BaseAnimeByID]: %d", *id)
			baseAnime, err := cw.realClientWrapper.BaseAnimeByID(context.Background(), id)
			if err != nil {
				return nil, err
			}
			data, err := json.Marshal([]*BaseAnimeByID{baseAnime})
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(test_utils.GetTestDataPath("BaseAnimeByID"), data, 0644)
			if err != nil {
				log.Fatal(err)
			}
			return baseAnime, nil
		}
	}

	var media []*BaseAnimeByID
	err = json.NewDecoder(file).Decode(&media)
	if err != nil {
		log.Fatal(err)
	}
	var baseAnime *BaseAnimeByID
	for _, m := range media {
		if m.GetMedia().ID == *id {
			baseAnime = m
			break
		}
	}

	if baseAnime == nil {
		cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [BaseAnimeByID]: %d", *id)
		baseAnime, err := cw.realClientWrapper.BaseAnimeByID(context.Background(), id)
		if err != nil {
			return nil, err
		}
		media = append(media, baseAnime)
		data, err := json.Marshal(media)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(test_utils.GetTestDataPath("BaseAnimeByID"), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return baseAnime, nil
	}

	cw.logger.Trace().Msgf("MockClientWrapper: CACHE HIT [BaseAnimeByID]: %d", *id)
	return baseAnime, nil
}

// AnimeCollection
//   - Set userName to nil to use the boilerplate AnimeCollection
//   - Set userName to a specific username to fetch and cache
func (cw *MockClientWrapper) AnimeCollection(ctx context.Context, userName *string, interceptors ...clientv2.RequestInterceptor) (*AnimeCollection, error) {

	if userName == nil {
		file, err := os.Open(test_utils.GetDataPath("BoilerplateAnimeCollection"))
		defer file.Close()

		var ret *AnimeCollection
		err = json.NewDecoder(file).Decode(&ret)
		if err != nil {
			log.Fatal(err)
		}

		cw.logger.Trace().Msgf("MockClientWrapper: Using [BoilerplateAnimeCollection]")
		return ret, nil
	}

	file, err := os.Open(test_utils.GetTestDataPath("AnimeCollection"))
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [AnimeCollection]: %s", *userName)
			ret, err := cw.realClientWrapper.AnimeCollection(context.Background(), userName)
			if err != nil {
				return nil, err
			}
			data, err := json.Marshal(ret)
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(test_utils.GetTestDataPath("AnimeCollection"), data, 0644)
			if err != nil {
				log.Fatal(err)
			}
			return ret, nil
		}
	}

	var ret *AnimeCollection
	err = json.NewDecoder(file).Decode(&ret)
	if err != nil {
		log.Fatal(err)
	}

	if ret == nil {
		cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [AnimeCollection]: %s", *userName)
		ret, err := cw.realClientWrapper.AnimeCollection(context.Background(), userName)
		if err != nil {
			return nil, err
		}
		data, err := json.Marshal(ret)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(test_utils.GetTestDataPath("AnimeCollection"), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return ret, nil
	}

	cw.logger.Trace().Msgf("MockClientWrapper: CACHE HIT [AnimeCollection]: %s", *userName)
	return ret, nil

}

func (cw *MockClientWrapper) AnimeCollectionWithRelations(ctx context.Context, userName *string, interceptors ...clientv2.RequestInterceptor) (*AnimeCollectionWithRelations, error) {

	if userName == nil {
		file, err := os.Open(test_utils.GetDataPath("BoilerplateAnimeCollectionWithRelations"))
		defer file.Close()

		var ret *AnimeCollectionWithRelations
		err = json.NewDecoder(file).Decode(&ret)
		if err != nil {
			log.Fatal(err)
		}

		cw.logger.Trace().Msgf("MockClientWrapper: Using [BoilerplateAnimeCollectionWithRelations]")
		return ret, nil
	}

	file, err := os.Open(test_utils.GetTestDataPath("AnimeCollectionWithRelations"))
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [AnimeCollectionWithRelations]: %s", *userName)
			ret, err := cw.realClientWrapper.AnimeCollectionWithRelations(context.Background(), userName)
			if err != nil {
				return nil, err
			}
			data, err := json.Marshal(ret)
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(test_utils.GetTestDataPath("AnimeCollectionWithRelations"), data, 0644)
			if err != nil {
				log.Fatal(err)
			}
			return ret, nil
		}
	}

	var ret *AnimeCollectionWithRelations
	err = json.NewDecoder(file).Decode(&ret)
	if err != nil {
		log.Fatal(err)
	}

	if ret == nil {
		cw.logger.Warn().Msgf("MockClientWrapper: CACHE MISS [AnimeCollectionWithRelations]: %s", *userName)
		ret, err := cw.realClientWrapper.AnimeCollectionWithRelations(context.Background(), userName)
		if err != nil {
			return nil, err
		}
		data, err := json.Marshal(ret)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(test_utils.GetTestDataPath("AnimeCollectionWithRelations"), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return ret, nil
	}

	cw.logger.Trace().Msgf("MockClientWrapper: CACHE HIT [AnimeCollectionWithRelations]: %s", *userName)
	return ret, nil

}

type TestModifyAnimeCollectionEntryInput struct {
	Status            *MediaListStatus
	Progress          *int
	Score             *float64
	AiredEpisodes     *int
	NextAiringEpisode *BaseAnime_NextAiringEpisode
}

// TestModifyAnimeCollectionEntry will modify an entry in the fetched anime collection.
// This is used to fine-tune the anime collection for testing purposes.
//
// Example: Setting a specific progress in case the origin anime collection has no progress
func TestModifyAnimeCollectionEntry(ac *AnimeCollection, mId int, input TestModifyAnimeCollectionEntryInput) *AnimeCollection {
	if ac == nil {
		panic("AnimeCollection is nil")
	}

	lists := ac.GetMediaListCollection().GetLists()

	removedFromList := false
	var rEntry *AnimeCollection_MediaListCollection_Lists_Entries

	// Move the entry to the correct list
	if input.Status != nil {
		for _, list := range lists {
			if list.Status == nil {
				continue
			}
			if list.Entries == nil {
				continue
			}
			entries := list.GetEntries()
			for idx, entry := range entries {
				if entry.GetMedia().ID == mId {
					if *list.Status != *input.Status {
						removedFromList = true
						entries = append(entries[:idx], entries[idx+1:]...)
						rEntry = entry
						break
					}
				}
			}
		}
		if removedFromList {
			for _, list := range lists {
				if list.Status == nil {
					continue
				}
				if *list.Status == *input.Status {
					if list.Entries == nil {
						list.Entries = make([]*AnimeCollection_MediaListCollection_Lists_Entries, 0)
					}
					list.Entries = append(list.Entries, rEntry)
					break
				}
			}
		}
	}

out:
	for _, list := range lists {
		entries := list.GetEntries()
		for _, entry := range entries {
			if entry.GetMedia().ID == mId {
				if input.Status != nil {
					entry.Status = input.Status
				}
				if input.Progress != nil {
					entry.Progress = input.Progress
				}
				if input.Score != nil {
					entry.Score = input.Score
				}
				if input.AiredEpisodes != nil {
					entry.Media.Episodes = input.AiredEpisodes
				}
				if input.NextAiringEpisode != nil {
					entry.Media.NextAiringEpisode = input.NextAiringEpisode
				}
				break out
			}
		}
	}

	return ac
}

//
// WILL NOT IMPLEMENT
//

func (cw *MockClientWrapper) UpdateMediaListEntry(ctx context.Context, mediaID *int, status *MediaListStatus, scoreRaw *int, progress *int, startedAt *FuzzyDateInput, completedAt *FuzzyDateInput, interceptors ...clientv2.RequestInterceptor) (*UpdateMediaListEntry, error) {
	cw.logger.Debug().Int("mediaId", *mediaID).Msg("anilist: Updating media list entry")
	return &UpdateMediaListEntry{}, nil
}

func (cw *MockClientWrapper) UpdateMediaListEntryProgress(ctx context.Context, mediaID *int, progress *int, status *MediaListStatus, interceptors ...clientv2.RequestInterceptor) (*UpdateMediaListEntryProgress, error) {
	cw.logger.Debug().Int("mediaId", *mediaID).Msg("anilist: Updating media list entry progress")
	return &UpdateMediaListEntryProgress{}, nil
}

func (cw *MockClientWrapper) DeleteEntry(ctx context.Context, mediaListEntryID *int, interceptors ...clientv2.RequestInterceptor) (*DeleteEntry, error) {
	cw.logger.Debug().Int("entryId", *mediaListEntryID).Msg("anilist: Deleting media list entry")
	return &DeleteEntry{}, nil
}

func (cw *MockClientWrapper) AnimeDetailsByID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*AnimeDetailsByID, error) {
	cw.logger.Debug().Int("mediaId", *id).Msg("anilist: Fetching anime details")
	return cw.realClientWrapper.AnimeDetailsByID(ctx, id, interceptors...)
}

func (cw *MockClientWrapper) CompleteAnimeByID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*CompleteAnimeByID, error) {
	cw.logger.Debug().Int("mediaId", *id).Msg("anilist: Fetching complete media")
	return cw.realClientWrapper.CompleteAnimeByID(ctx, id, interceptors...)
}

func (cw *MockClientWrapper) ListAnime(ctx context.Context, page *int, search *string, perPage *int, sort []*MediaSort, status []*MediaStatus, genres []*string, averageScoreGreater *int, season *MediaSeason, seasonYear *int, format *MediaFormat, isAdult *bool, interceptors ...clientv2.RequestInterceptor) (*ListAnime, error) {
	cw.logger.Debug().Msg("anilist: Fetching media list")
	return cw.realClientWrapper.ListAnime(ctx, page, search, perPage, sort, status, genres, averageScoreGreater, season, seasonYear, format, isAdult, interceptors...)
}

func (cw *MockClientWrapper) ListRecentAnime(ctx context.Context, page *int, perPage *int, airingAtGreater *int, airingAtLesser *int, interceptors ...clientv2.RequestInterceptor) (*ListRecentAnime, error) {
	cw.logger.Debug().Msg("anilist: Fetching recent media list")
	return cw.realClientWrapper.ListRecentAnime(ctx, page, perPage, airingAtGreater, airingAtLesser, interceptors...)
}

func (cw *MockClientWrapper) GetViewer(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetViewer, error) {
	cw.logger.Debug().Msg("anilist: Fetching viewer")
	return cw.realClientWrapper.GetViewer(ctx, interceptors...)
}

func (cw *MockClientWrapper) MangaCollection(ctx context.Context, userName *string, interceptors ...clientv2.RequestInterceptor) (*MangaCollection, error) {
	cw.logger.Debug().Msg("anilist: Fetching manga collection")
	return cw.realClientWrapper.MangaCollection(ctx, userName, interceptors...)
}

func (cw *MockClientWrapper) SearchBaseManga(ctx context.Context, page *int, perPage *int, sort []*MediaSort, search *string, status []*MediaStatus, interceptors ...clientv2.RequestInterceptor) (*SearchBaseManga, error) {
	cw.logger.Debug().Msg("anilist: Searching manga")
	return cw.realClientWrapper.SearchBaseManga(ctx, page, perPage, sort, search, status, interceptors...)
}

func (cw *MockClientWrapper) BaseMangaByID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*BaseMangaByID, error) {
	cw.logger.Debug().Int("mediaId", *id).Msg("anilist: Fetching manga")
	return cw.realClientWrapper.BaseMangaByID(ctx, id, interceptors...)
}

func (cw *MockClientWrapper) MangaDetailsByID(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*MangaDetailsByID, error) {
	cw.logger.Debug().Int("mediaId", *id).Msg("anilist: Fetching manga details")
	return cw.realClientWrapper.MangaDetailsByID(ctx, id, interceptors...)
}

func (cw *MockClientWrapper) ListManga(ctx context.Context, page *int, search *string, perPage *int, sort []*MediaSort, status []*MediaStatus, genres []*string, averageScoreGreater *int, startDateGreater *string, startDateLesser *string, format *MediaFormat, isAdult *bool, interceptors ...clientv2.RequestInterceptor) (*ListManga, error) {
	cw.logger.Debug().Msg("anilist: Fetching manga list")
	return cw.realClientWrapper.ListManga(ctx, page, search, perPage, sort, status, genres, averageScoreGreater, startDateGreater, startDateLesser, format, isAdult, interceptors...)
}

func (cw *MockClientWrapper) StudioDetails(ctx context.Context, id *int, interceptors ...clientv2.RequestInterceptor) (*StudioDetails, error) {
	cw.logger.Debug().Int("studioId", *id).Msg("anilist: Fetching studio details")
	return cw.realClientWrapper.StudioDetails(ctx, id, interceptors...)
}

func (cw *MockClientWrapper) ViewerStats(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*ViewerStats, error) {
	cw.logger.Debug().Msg("anilist: Fetching stats")
	return cw.realClientWrapper.ViewerStats(ctx, interceptors...)
}
