# Welcome to the Merico Technical Interview for Dev Lake

## Intro

Dev Lake aims to help dev teams improve their productivity by drawing insights from data. It is a data platform that collects, analyzes, and visualizes data from various engineering tools including code hosting (e.g. Gitlab), project management (e.g. Jira), and build system (e.g. Jenkins).

This takehome interview will give you a tour of Dev Lake and walk you through the process of supporting a new data source. It should take you 90 - 120 mins to complete once you have the project set up locally. Please email us your zipped code within *48 hours*.

## Getting Started

1. First, set up dev lake as a developer by following the instructions in README's [Developer Setup](https://github.com/merico-dev/lake#dev-setup) section.
2. If you collect Gitlab data successfully, you should see data showing up in the "Development" section of the "Delivery Velocity" Grafana dashboard. Other charts may still show "No Data" since they need data from other data sources than Gitlab.
3. Check out our [Grafana doc](https://github.com/merico-dev/lake/blob/main/docs/GRAFANA.md) on how to customize existing charts or add your own charts. Essentially, you just need to write SQL queries to pull data from Dev Lake's MySQL database.
4. Now you're good to go!


## Requirements

1. Implement a new Dev Lake plugin that collects data from PokeApi (https://pokeapi.co). You can find the interface for a plugin [here](https://github.com/merico-dev/lake/blob/main/plugins/core/plugin.go). Feel free to use existing plugins as reference.
2. This plugin would collect all Pokemon with their Pokedex number, name, height, and weight into the database.
3. Implement a Grafana dashboard for showcasing the data you collected from Pokeapi, specifically:
    3a. A chart showing the total number of Pokemon in the database (a single number).
    3b. A chart showing the medium weight of all Pokemon (a single number).
    3c. A chart showing the top 10 Pokemon in terms of height (A bar chart or a chart of your own choice).
4. (Bonus) Utilize multiple threads/process to accelerate the data collection process and measure the speedup.

## Evaluation

What we expect from your solution:

1. A well-written README for your newly created plugin
2. Well-documented and clean code
3. Code that handles errors and edge cases
4. Code that is performant and efficient

## Contact

If you have any questions, please email Jonathan at joncodo@merico.dev.