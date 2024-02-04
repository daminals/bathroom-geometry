![github repo badge: frontend](https://img.shields.io/badge/Frontend-Svelte-Red) ![github repo badge: frontend](https://img.shields.io/badge/Frontend-TypeScript-Blue) ![github repo badge: backend](https://img.shields.io/badge/Backend-Go-8A2BE2) ![github repo badge: backend](https://img.shields.io/badge/Testing-Python-Blue) 

# SBU Bathroom Finder
Welcome to SBU Bathroom Finder, your go-to WebApp for locating the nearest and best bathrooms on Stony Brook University's campus! Tailor your restroom search based on gender, accessibility, ratings, and amenities. While currently designed for SBU, our open-source approach invites community members from other universities to contribute and create maps for their campuses using our built-in map creator. Explore and find the perfect bathroom with ease!
This project was made for [Hopper Hacks 2024](https://hopperhacks2024.devpost.com/) and we hope you enjoy our project!
![image](https://github.com/daminals/bathroom-geometry/assets/6569519/904633e8-8931-4af9-af4b-982ef79ef8c8)

# Inspiration 
As residents of NYC, we've all faced the challenge of locating clean and accessible bathrooms in a bustling city. Recognizing the widespread nature of this issue, we decided to tackle it on a smaller scale starting with the University, we all know and love, SBU! Our collective knowledge about the campus bathrooms serves as a proof of concept, inviting others to contribute and adapt the solution to their university environments.   

Additionally, our shared experience with the Computational Geometry class here at SBU has made us excited to try our hand at implementing what we've learned in class. 

# What it Does
Using A* and Point Sampling algorithms, the SBU Bathroom Finder enables users to locate the nearest restroom by generating a color-coded grid. On the Gallery and Viewer pages, users encounter a grid displaying a Voronoi diagram, where bathrooms serve as hotspots, each denoted by a unique color to signify proximity. Meanwhile, the walls are delineated in black.

After either logging in or signing up, users unlock even more features. For example, they have access to the map editor, where they can select a location on Google Maps, insert a grid, and mark all the walls and bathrooms. Upon saving the map, users can press the "compute geometry" button to generate a Voronoi diagram for it. Users can also rate existing restrooms on Stony Brook's campus based on accessibility, cleanliness, and other features. Finally, they can log out when they're finished using the application.

# How we Built it

## Backend  
Deciding on our implementation and what exactly we wanted to achieve was pretty tough. We floated around multiple ideas but eventually settled on creating a Voronoi diagram of some sort. Thus, our extensive research began on implementations of Voronoi diagrams that had already been accomplished. 

The problem was though, that the Voronoi diagram implementations we found in Golang did not account for walls, which was problematic given that we couldn't just have people phase through the walls and so we started trying to implement our own. 

This first led us to the [Jump Flooding Algorithm](https://en.wikipedia.org/wiki/Jump_flooding_algorithm) and we watched [this](https://www.youtube.com/watch?v=AT0jTugdi0M) excellent YouTube video explaining the concept, but while this algorithm was cool to learn about, it was out of the scope of this project to create a customized Jump Flooding Algorithm that would account for obstacles. 

We settled on implementing the Voronoi diagram in an unconventional but more naive approach, taking advantage of A* pathfinding, Distance Heuristics, and some clever splitting of our input into tiles represented by a 2D array.  

With the A* pathfinding algorithm we had to implement our version as many Golang implementations wanted to take in a string representation which would not be compatible with creating Voronoi diagrams. 

Overall, we did end up overcoming most of the challenges we faced in our Backend, though this doesn't mean it's perfect and further optimizations/changes will be mentioned in future sections. So, for now, let's move on to the front-end challenges.

## Frontend
For the front end of the SBU Bathroom Finder, we utilized Svelte as our primary framework. Svelte is a modern JavaScript framework that focuses on performance and simplicity. We were able to build a reactive user interactive as a result, with various pages enabling the user to log in, sign up, view a gallery, edit, and log out. Users are also able to complete forms and interact with images and grids. Grids are fully interactive, with users being able to mark both walls and restrooms.

To make the map interactions fully possible, we also used the Google Maps API. This allows the Voronoi diagrams created by the backend to be displayed as a grid over a map of Stony Brook's campus. Users are also able to scroll through Google Maps to pick a location for a new bathroom grid.

TypeScript was used to store user information from signing up, logging in, and completing forms.

# Challenges we ran into
SBU Bathroom Finder was an interesting and tall project for us to tackle as we didn't have much experience in Golang or Svelte and we had never done any pathfinding or projects that required complex distance calculations before. Doing this in 24 hours didn't make it much easier. Implementing our versions of the A* and Point Sampling algorithm was quite difficult as we had no experience coding these before, but eventually, we got an approximation of the Voronoi diagram which is pretty close to what we wanted.  

For reference here's what our code was outputting when we were having difficulties and it broke: 
<img src="https://github.com/daminals/bathroom-geometry/assets/107337676/bc2fe815-447c-467a-a986-675fd2754aba" width="600" height="600">


Here's what it was outputting once it was fixed:  

<img src="https://cdn.discordapp.com/attachments/1181082647833890876/1203723383288762368/Screenshot_2024-02-04_at_10.26.55_AM.png?ex=65d221c2&is=65bfacc2&hm=9a90f6dc25f6910add98bda01c167e7b06eac275b6136a9f7ac147946bca9804&" width="600" height="600">


# Accomplishments that we're proud of
Our team is proud that we made finding bathrooms more accessible at Stony Brook University. We are also happy we got to implement what we learned in our Computational Geometry class in a software project.

# What we learned
No one on our team had experience with Svelte before, so we all familiarized ourselves with a new JavaScript framework. We also all used the Google Maps API for the first time during this project, so we learned how to integrate this both into a Svelte frontend and our Golang backend.
We also learned how to implement an A* star and point sampling algorithms into our backend.

# What's next for SBU Bathroom Finder
We want to implement a more largescale bathroom finder, that is easily adaptable to metropolitan and campus areas. We also want to improve user functionality by having an improved rating experience that is better integrated with our main Gallery and Viewer.

If we had more time, we would have written a Jump Flooding Algorithm with the following customizations, based on a few stack overflow posts (notably [this one](https://stackoverflow.com/questions/73255352/creating-a-voroni-diagram-with-arbitrary-boundaries)):
  - If the tile is colored with the boundary color, discard consideration for recoloring automatically
  - Implement A* for the dist() check which respects the walls

Also, in the future, we would like to set up a proper database like MongoDB or SQL. Overall, making community-sourced maps more accessible and faster to access as the project scales.  
  
