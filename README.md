# TopMusicStreaming

This Golang program scrapes streaming platforms to generate the **top 100** most listened-to songs. The ranking is based on the average of the positions of a track on **Spotify**, **Apple Music** and **Deezer**. When it is launched, it goes to the different platforms to retrieve the rankings. An algorithm sorts the data and generates a final file in Json format.

## Link

[music.mathi3u.com](https://music.mathi3u.com/)

## Exemple

### Data :

|                      	| Track 1    | Track 2    | Track 3    | Track 4    | Track 5    |
|----------------------	|---------	|---------	|---------	|---------	|---------	|
| Spotify Position        | 1        | 2        | 3        | 10        | 12        |
| Apple Music Position    | 7        | 2        | 4        | 13        | 1        |
| Deezer Position        | 3        | 1        | 6        | 16        | 18        |
| **Average Position**  |**3.65**   | **1.67**    | **4.33**  | **13**    | **10.33**    |

### Final ranking :

1. Track 2
2. Track 1
3. Track 3
4. Track 5
5. Track 4

## Usage

#### Clone the repository
```
git clone https://github.com/mathieumarcelino/TopMusicStreamingAPI.git
```

#### Create a `.env` file
By default the `ENV` is set to **prod**, to run the program in a **local** environment create a `.env` file and add the
following:

```shell
ENV="local"
```

#### Run program
```
make run
```

#### Example API request

```
http://localhost:9990/?country=fr
```

#### Country Options

- ww (Worldwide 🌍)
- us (United States 🇺🇸)
- fr (France 🇫🇷)
- uk (United Kingdom 🇬🇧)
- jp (Japan 🇯🇵)
- kr (South Korea 🇰🇷)
- tr (Turkey 🇹🇷)
- de (Germany 🇩🇪)
- es (Spain 🇪🇸)
- pt (Portugal 🇵🇹)
- it (Italy 🇮🇹)