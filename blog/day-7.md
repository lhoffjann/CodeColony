I now program this project for about a week.

# Intention of this project
The intention of this project is to write a simulation and not a game.

The idea is copied https://github.com/jurajmajerik/rides

In the end you should see one or (hopefully) more creeps moving between a home base and a harvestable resource.

tech-wise i use go for the backend. for the front end i am currently unsure.



# Current Problems
currently I am trying to figure out how to implement A* to have Pathfinding for the creep. For that it seems that I need a better representation of the world it is moving in. Well, I need any representation. Currently the representation is a infinite grid and there is no way to know where something is place on this grid without iterating over every exisiting entity.

It may make sense to introduce a `world` struct which holds basic information about the world.
and also can return every free neighbor for a given node.
