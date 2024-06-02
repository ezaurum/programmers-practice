package bestalbulm

import "sort"

type Song struct {
	ID    int
	Genre string
	Play  int
}

// solution https://school.programmers.co.kr/learn/courses/30/lessons/42579
func solution(genres []string, plays []int) []int {
	playByGenre := map[string]int{}
	songsByGenre := map[string][]Song{}
	for i, g := range genres {
		p := plays[i]
		s := Song{
			ID:    i,
			Genre: g,
			Play:  p,
		}
		playByGenre[g] = playByGenre[g] + p
		songsByGenre[g] = append(songsByGenre[g], s)
	}
	var genreNames []string
	for g, songs := range songsByGenre {
		genreNames = append(genreNames, g)
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Play > songs[j].Play
		})
		if len(songs) > 2 {
			songs = songs[:2]
		}
		songsByGenre[g] = songs
	}
	sort.Slice(genreNames, func(i, j int) bool {
		return playByGenre[genreNames[i]] > playByGenre[genreNames[j]]
	})

	var r []int
	for _, g := range genreNames {
		songs := songsByGenre[g]
		for _, s := range songs {
			r = append(r, s.ID)
		}
	}

	return r
}
