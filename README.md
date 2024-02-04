![github repo badge: frontend](https://img.shields.io/badge/Frontend-Svelte-Red) ![github repo badge: frontend](https://img.shields.io/badge/Frontend-TypeScript-Blue) ![github repo badge: backend](https://img.shields.io/badge/Backend-Go-8A2BE2) ![github repo badge: backend](https://img.shields.io/badge/Backend-Python-Orange) 

# SBU Bathroom Finder
Welcome to SBU Bathroom Finder, your go-to WebApp for locating the nearest and best bathrooms on Stony Brook University's campus! Tailor your restroom search based on gender, accessibility, ratings, and amenities. While currently designed for SBU, our open-source approach invites community members from other universities to contribute and create maps for their campuses using our built-in map creator. Explore and find the perfect bathroom with ease!
This project was made for [Hopper Hacks 2024](https://hopperhacks2024.devpost.com/) and we hope you enjoy our project!  

# Inspiration 
As residents of NYC, we've all faced the challenge of locating clean and accessible bathrooms in a bustling city. Recognizing the widespread nature of this issue, we decided to tackle it on a smaller scale starting with the the University, we all know and love, SBU! Our collective knowledge about the campus bathrooms serves as a proof of concept, inviting others to contribute and adapt the solution to their own university environments.   

Additionally, our shared experience with the Computational Geometry class here at SBU has made us excited to try our hand at implementing what we've learned in class. 

# Challenges & Lessons   
SBU Bathroom Finder was an interesting and tall project for us to tackle as we didn't have much experience in Golang or Svelte and we had never done any pathfinding or projects that required complex distance calculations before. Doing this in 24 hours didn't make it much easier. However, as a result, we learned a lot and had fun working on the project! Here are some highlights of our challenges and what we learned: 

## Backend  
Deciding our implementation and what exactly we wanted to achieve was pretty tough. We floated around multiple ideas but eventually settled on creating a Voronoi diagram of some sorts. Thus, our extensive research began on implementations of Voronoi diagrams that had already been accomplished. 

The problem was though, that the Voronoi diagram implementations we found in Golang did not account for walls, which was problematic given that we couldn't just have people phase through the walls and so we started trying to implement our own. 

This first led us to the [Jump Flooding Algorithm](https://en.wikipedia.org/wiki/Jump_flooding_algorithm) and we watched [this](https://www.youtube.com/watch?v=AT0jTugdi0M) excellent youtube video explaining the concept, but while this algorithm was cool to learn about, it was out of the scope of this project to create a customized Jump Flood Algorithm that would account for obstacles. 

We settled on implementing the Voronoi diagram in an unconventional way, taking advantage of A* pathfinding, Distance Hueristics, and some clever splitting of our input into tiles represented by a 2D array.  

With the A* pathfinding algorithm we had to implement our own version as many Golang implementations wanted to take in a string representation which would not be compatible with creating Voronoi diagrams. 

Implementing our own versions of these algorithms was quite difficult as we had no experience coding these before, but eventually it worked until we had to deal with coding endpoints for the part of the team working on the frontend. The I/O to JSON files and their formatting was a bit tricky to decide on and figure out in Golang as the sytax was quite different from languages like python or java. 

Overall, we did end up overcoming most of the challenges we faced in our Backend, though this doesn't mean it's perfect, our A* implementation is a very good approximation but not exact, and further optimizations/changes will be mentioned in future sections. So, for now, let's move onto the Frontend challenges.  

## Frontend

# Where do we go from here?   

When it comes to the future of SBU Bathroom Finder, there are several things that we wish we could've done if not for the time limit of 24 hours. I/O to JSON Files and loading previously made maps could be made more efficient with the use of a database like MongoDB or SQL. In the use of a database we could also take advantage of showing many more pre-made maps that are contributed from the community. 

In terms of our A* implmenetation and how we created the Voronoi there may be other data structures or algorithms we could've explored to optimize the time and space complexity, but our implementation runs quite well regardless.

# Technical Details 




