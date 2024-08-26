package main

import (
    "github.com/Chris-Coleongco/blog/api"
)

func main() {
    server := api.New_Api_Server(":8080")
    server.Run()

}


// go through your entire github concurrently getting all ur data into a golang struct for each git repo


// display the readme.mds as the blog text.

// for images / video files demoing certain repos, access an api endpoint to get the data into your Chris-Coleongco.github.io site


