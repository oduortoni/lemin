# How to split the project

### folders

+ **vars/** - contains all the global variables
+ **models/** - contains all the types i.e structs

### complete files

+ **processpaths/findants.go** and **processpaths/findants_test.go** - given an adjacency list or matrix i.e vars.Colony, this function finds all the paths in the universe that begin at start to end room
+ **ants/moveants.go** and **ants/moveants_test.go** - given the most optimal paths, the function takes care of printing the moves that each ant takes from start to end, i.e it is responsible for the final print out e.g L1-2 L2-3 ...
+ **ants/assignants.go** and **ants/assignants_test.go** - given the most optimal paths, the function takes care of destributing the ants to each path. It uses an algorithm that ensures that traffic jams wont occur based on the summation of the number of rooms and ants on that given path


### utils/utils.go and utils/utils_test.go

Contains so many functions. We have to split them into several files based off functions

#### functions in utils.go

- **SliceContainsString** and **TestSliceContainsString** - tests whether a string is contained in a slice of strings
- **SliceContainsSlice** and **TestSliceContainsSlice** - tests whether a slice contains, a slice
- **ValidCoordinates** and **TestValidCoordinates**
- **ValidRoomConnection** and **TestValidRoomConnection**
- **ProcessNumberOfAnts** and **TestProcessNumberOfAnts**
- **GetRoom** and **TestGetRoom**
- **StoreRoom** and **TestStoreRoom**
- **ValidColonyRooms** and **TestValidColonyRooms**
- **StoreConnectedRooms** and **TestConnectedRooms**
- **HasStartAndEnd** and **TestHasStartAndEnd**
- **MaxTurns** and **TestMaxTurns**

### UNDOCUMENTED

- **processpaths/optimalpaths.go**
- **utils/utils.go** - there are three functions in there I did not test


### FINALLY

- **main.go**
- **README.md** - it is still empty and therefore needs to be written