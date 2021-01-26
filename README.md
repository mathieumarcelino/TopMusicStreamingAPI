# TopMusicStreaming

Hi ! This program in Golang generates the **top 100** most listened to music on streaming platforms. The ranking is based on the average of the positions of a track on **Spotify**, **Apple Music** and **Deezer**. When it is launched, it goes to the different platforms to retrieve the rankings. An algorithm sorts the data and generates a final file in Json format.

## Exemple :

### Data :

|                      	| Track 1 	| Track 2 	| Track 3 	| Track 4 	| Track 5 	|
|----------------------	|---------	|---------	|---------	|---------	|---------	|
| Spotify Position     	| 1       	| 2       	| 3       	| 10      	| 12      	|
| Apple Music Position 	| 7       	| 2       	| 4       	| 13      	| 1       	|
| Deezer Position      	| 3       	| 1       	| 6       	| 16      	| 18      	|
| **Average Position**  |**3.65**   | **1.67** 	| **4.33**  | **13**    | **10.33**	|


### Final ranking :

1. Track 2
2. Track 1
3. Track 3
4. Track 5
5. Track 4

## Usage

```
git clone https://github.com/mathieumarcelino/topmusicstreaming
```

```
go run main.go
```

Visit [localhost:9990](http://localhost:9990)

