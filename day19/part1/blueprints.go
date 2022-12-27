package main

var testBlueprints = []*Blueprint{
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 2}, obsidianRobot: Cost{ore: 3, clay: 14}, geodeRobot: Cost{ore: 2, obsidian: 7}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 8}, geodeRobot: Cost{ore: 3, obsidian: 12}},
}

var mainBlueprints = []*Blueprint{
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 9}, geodeRobot: Cost{ore: 3, obsidian: 9}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 4, clay: 20}, geodeRobot: Cost{ore: 4, obsidian: 8}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 2, clay: 16}, geodeRobot: Cost{ore: 2, obsidian: 9}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 20}, geodeRobot: Cost{ore: 4, obsidian: 16}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 16}, geodeRobot: Cost{ore: 2, obsidian: 15}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 2}, obsidianRobot: Cost{ore: 2, clay: 20}, geodeRobot: Cost{ore: 2, obsidian: 14}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 3, clay: 7}, geodeRobot: Cost{ore: 3, obsidian: 20}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 3, clay: 14}, geodeRobot: Cost{ore: 4, obsidian: 15}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 7}, geodeRobot: Cost{ore: 2, obsidian: 7}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 2, clay: 11}, geodeRobot: Cost{ore: 2, obsidian: 19}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 20}, geodeRobot: Cost{ore: 2, obsidian: 12}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 20}, geodeRobot: Cost{ore: 2, obsidian: 8}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 3, clay: 14}, geodeRobot: Cost{ore: 4, obsidian: 9}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 18}, geodeRobot: Cost{ore: 3, obsidian: 8}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 2, clay: 9}, geodeRobot: Cost{ore: 3, obsidian: 15}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 11}, geodeRobot: Cost{ore: 2, obsidian: 16}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 13}, geodeRobot: Cost{ore: 3, obsidian: 15}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 16}, geodeRobot: Cost{ore: 3, obsidian: 20}},
	{oreRobot: Cost{ore: 2}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 3, clay: 19}, geodeRobot: Cost{ore: 4, obsidian: 8}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 4, clay: 16}, geodeRobot: Cost{ore: 2, obsidian: 15}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 7}, geodeRobot: Cost{ore: 2, obsidian: 19}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 2, clay: 14}, geodeRobot: Cost{ore: 3, obsidian: 17}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 4, clay: 8}, geodeRobot: Cost{ore: 2, obsidian: 8}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 7}, geodeRobot: Cost{ore: 4, obsidian: 17}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 16}, geodeRobot: Cost{ore: 3, obsidian: 9}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 4, clay: 15}, geodeRobot: Cost{ore: 4, obsidian: 9}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 2, clay: 20}, geodeRobot: Cost{ore: 4, obsidian: 7}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 3}, obsidianRobot: Cost{ore: 3, clay: 17}, geodeRobot: Cost{ore: 4, obsidian: 8}},
	{oreRobot: Cost{ore: 3}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 3, clay: 12}, geodeRobot: Cost{ore: 3, obsidian: 17}},
	{oreRobot: Cost{ore: 4}, clayRobot: Cost{ore: 4}, obsidianRobot: Cost{ore: 4, clay: 5}, geodeRobot: Cost{ore: 2, obsidian: 10}},
}
