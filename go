import tweepy
import feedgen.feed

# Twitter API credentials
consumer_key = 'your_consumer_key'
consumer_secret = 'your_consumer_secret'
access_token = 'your_access_token'
access_secret = 'your_access_secret'

# Twitter search query
search_query = 'Consensus OR EOS OR Antelope Coalition'

# Create a feed generator
feed = feedgen.feed.FeedGenerator()
feed.title('Twitter Search Results')
feed.link(href='https://twitter.com/search?q=' + search_query)
feed.description('RSS feed of Twitter search results for ' + search_query)

# Connect to the Twitter API
auth = tweepy.OAuthHandler(consumer_key, consumer_secret)
auth.set_access_token(access_token, access_secret)
api = tweepy.API(auth)

# Search for tweets and add to feed
tweets = api.search(q=search_query, lang='en')
for tweet in tweets:
    feed.add_entry(title=tweet.user.name, link='https://twitter.com/' + tweet.user.screen_name, description=tweet.text, pubdate=tweet.created_at)

# Generate the RSS feed
rss_feed = feed.atom_str()

# Print the RSS feed
print(rss_feed)
