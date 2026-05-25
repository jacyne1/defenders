package main

type WaveMobConfig struct {
	MobType int
	Weight  int // how frequent the mobs of that type spawns
}

type WaveConfig struct {
	MobPool           []WaveMobConfig
	TotalToSpawn      int
	SpawnRate         float32
	IntermissionDelay float32
}

type LevelConfig struct {
	LevelNumber int
	Waves       []WaveConfig
}

var GameCampaign = []LevelConfig{
	{
		LevelNumber: 1,
		Waves: []WaveConfig{
			// Wave 1: Just basic aliens
			{
				MobPool:           []WaveMobConfig{{MobType: MobBasic, Weight: 100}},
				TotalToSpawn:      5,
				SpawnRate:         1.5,
				IntermissionDelay: 3.0, // 3 seconds after clearing
			},
			// Wave 2: Mix of 70% basic and 30% ZigZag
			{
				MobPool: []WaveMobConfig{
					{MobType: MobBasic, Weight: 70},
					{MobType: MobZigZag, Weight: 30},
				},
				TotalToSpawn:      12,
				SpawnRate:         0.8,
				IntermissionDelay: 4.0,
			},
		},
	},
	{
		LevelNumber: 2,
		Waves: []WaveConfig{

			{
				MobPool: []WaveMobConfig{
					{MobType: MobBasic, Weight: 50},
					{MobType: MobZigZag, Weight: 50},
				},
				TotalToSpawn:      15,
				SpawnRate:         0.6,
				IntermissionDelay: 5.0,
			},
		},
	},
}
