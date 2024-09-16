package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	playerCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "snake_ladder_player_count",
			Help: "Total no of players and name",
		},
	 []string{"name"},
	)
)

func IncreamentPlayerCount(name string) {
	playerCount.WithLabelValues(name).Inc()
}

func init() {
	prometheus.MustRegister(playerCount)
}


