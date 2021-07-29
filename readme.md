# Photo Album Maker

## Description
An attempt tp generate descriptions/titles for a holiday photo album

## Running the app

To run the app with the csv files supplied simply change the file number to 1, 2, or 3 in the file main.go at this line: `data, err := reader.ReadCSV("./reader/test_data/1.csv")`, save the file & enter in the terminal `go run main.go` from the root folder. This will generate a json file named "{city/state}-{country}.json" e.g. "Nevada-United States.json".


## Thought Process

### Starting point

My inital starting point was looking into what data could be retrieved from the various apis pertaining to geocoding, in order to firstly find out the location of where the photo was taken. Upon testing this and seeing the kind of data returned, which was disappointingly just a street address, I looked at what other apis could give more info about the places. This didn't offer much so I started thinking about how I would want to have captions/titles and what would remind me about a particular trip. So I thought a wider bit of info about the place & the kind of description you'd get in a travel guide would be really nice and hopefully mean the user would reminisce and remember all the things that they maybe didn't even take a photograph of.

Since time was limited I started thinking about how I personally organise trips in my mind when I think back, and for me at least I tend to remember days. This started me thinking about how I wanted to organise the data & describe the trip in terms of each day. E.g. How many different places/areas did I visit, what were those areas like, how many photos did I take? This is what inspired me to start structing the data by top level data about the holiday, summarising the trip and reminding the user how long the trip was, the different cities they visited and a bit about the country. Then I tried to organise the photos by day and look at the different areas they were in and find a description from trip advisor about each area.

### Hurdles

The difficulty with all this given the time constraint, was to sift through a large amount of data and get the relevant parts. On top of this I started to realise the complexity in structuring data in the way that I had chosen, and implement something in 4 hours! It was also hard to find apis that would offer the kind of data that was easy to find the kind of data I was looking for. E.g. one idea was to look at the history of an area and include facts like this, however this didn't pan out.

### What I learnt

I think the mistake I made at the beginning was overcomplicating the data structure. I came at the problem from a top down, thinking about the data around a trip and working my way down. I think a better way would have been to start with a photo, get info for that and start composing a title, working from there up.

### If I'd had more time...

I'd really loved to have started getting data on the weather, so I could compose a message based upon say the temperature if it was really hot, or cold, or if it was sunny & clear, or wet and windy etc. I also wanted to use the data around attractions to try and tailor a message about whether say it was bustling, great for shopping, or quiet & tranquil. I'd also rethink the organisation of the data returned, however I think this is always more complex in go.