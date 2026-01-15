# web-song

<div align="center">
<img src="https://github.com/LeoScripts/web-song/raw/main/.gitassets/2.jpeg" width="350" />

<!--<div data-badges>
    <img src="https://img.shields.io/github/stars/LeoScripts/web-song?style=for-the-badge" alt="GitHub stars" />
    <img src="https://img.shields.io/github/forks/LeoScripts/web-song?style=for-the-badge" alt="GitHub forks" />
    <img src="https://img.shields.io/github/issues/LeoScripts/web-song?style=for-the-badge" alt="GitHub issues" />
</div>-->

</div>

Web-Song is an audio player, created to meet some personal needs 😂😂😂😂😂😂 (listening to my music)!! I wanted to listen to the music from my homelab server so I had the idea to create it, the initial idea is quite simple, but I started thinking about other implementations for this project
let's see where it goes 😂😂😂😂😂


```diff
- attention

! You are fully responsible for your actions, 
! and I do not assume responsibility for any improper practices 
! or endorse such behaviors.

+ use responsibly
```

## with docker

##### Requirements:
- Docker and Docker Compose

1. Clone this repository:
```sh
   git clone https://github.com/LeoScripts/web-song
```
2. Access the project directory:
```sh
   cd web-song
```
3. Configure environment variables:
   You will need to create a `.env` file with the same environment variables listed in the `.env.example` file which should be filled with the corresponding environment variables exemplified in the `.env.example` file.

4. runing app
```
docker compose up -d --build
```

5. Access the project at [http://localhost:7880](http://localhost:7880).


## Local

##### Requirements:
- Go installed

##### Execution:
1. Clone this repository:
```sh
   git clone https://github.com/LeoScripts/web-song
```
2. Access the project directory:
```sh
   cd web-song
```
3. Install dependencies with the following command in the project root folder:
```sh
    go mod tidy
```
4. Configure environment variables:
   You will need to create a `.env` file with the same environment variables listed in the `.env.example` file which should be filled with the corresponding environment variables exemplified in the `.env.example` file.
5. change to file `main.go` 
> main.go
```go
// uncomment this code
err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


// modify this line ( mediaDir = "/app/media" ) to
mediaDir = os.Getenv("MEDIA_DIR")
```
6. Start the application by running the command `go run main.go` in your application's root folder. This command will start all projects in your application.
7. Access the project at [http://localhost:7880](http://localhost:7880).

##### 🗒️ Project Features 🗒️
- music playback from a predetermined folder

![](https://github.com/LeoScripts/web-song/raw/main/.gitassets/2.jpeg)


##### 💎 Useful Links 💎
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose)
- [Html](https://html.com/)
- [Css](https://css3.com/)
- [Javascript](https://www.javascript.com/)
