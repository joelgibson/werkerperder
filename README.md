# Werkerperder

https://werkerperder.com/, an automatic translation of Wikipedia into [ermahgerd](https://knowyourmeme.com/memes/ermahgerd)-speak.


## The story

While tutoring on a high-school summer camp back in 2013 or so, my friend [Ben Taylor](https://github.com/taybenlor) showed me a website he'd made that automatically translated his friend's very professional-looking profile into meme-speak.
I thought this was absolutely hilarious (most probably due to sleep deprivation at the time), stole Ben's code (some of which he assumedly pilfered from elsewhere --- the furthest back I can trace this dank collection of regular expresisons is to the [ermahgerd translator](https://ermahgerd.jmillerdesign.com/) by [Justin Miller](https://github.com/jmillerdesign)), and repurposed it to do the same thing for Wikipedia. This was also hilarious. I then forgot about the code for several years.

The other day while waiting for a midnight lecture to start (I was virtually attending an overseas conference, with a timezone slip of about 10 hours), I found the code again, ran it, and again (in my sleep-deprived state) thought that it was hilarious.
A friend encouraged me to post the code --- here it is.

All credit to Ben, Justin, and internet meme culture more broadly.


## To run

To translate a word list in command-line mode:

    cat werds | go run ermehgerd.go -stdio

To launch a server:

    go run ermehgerd.go -bind localhost:1234
