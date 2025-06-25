#!/usr/bin/env python3
"""
Symphony API Database Population Script
For University Database Course Presentation

This script populates all three databases:
- PostgreSQL: Users, Posts, Communities, Chats
- Neo4j: User relationships, friendships, genre preferences
- MongoDB: Artists, Songs, Playlists

Usage: python3 populate_databases.py
"""

import requests
import json
import random
import time
from datetime import datetime, timedelta
import sys

# Configuration
API_BASE_URL = "http://localhost:8080"
NUM_USERS = 50
NUM_ARTISTS = 30
NUM_SONGS = 100
NUM_PLAYLISTS = 40
NUM_POSTS = 200
NUM_COMMUNITIES = 15
NUM_CHATS = 80

# Sample data for realistic generation
FIRST_NAMES = [
    "João", "Maria", "Pedro", "Ana", "Carlos", "Lucia", "Fernando", "Isabela",
    "Rafael", "Camila", "Lucas", "Julia", "Gabriel", "Beatriz", "Matheus",
    "Sofia", "Thiago", "Mariana", "Diego", "Carolina", "André", "Amanda",
    "Bruno", "Letícia", "Ricardo", "Bianca", "Felipe", "Natalia", "Alexandre",
    "Vanessa", "Daniel", "Priscila", "Marcelo", "Tatiana", "Roberto", "Renata"
]

LAST_NAMES = [
    "Silva", "Santos", "Oliveira", "Souza", "Rodrigues", "Ferreira", "Almeida",
    "Pereira", "Lima", "Gomes", "Costa", "Ribeiro", "Martins", "Carvalho",
    "Alves", "Lopes", "Soares", "Fernandes", "Vieira", "Barbosa", "Rocha",
    "Dias", "Nascimento", "Andrade", "Moreira", "Nunes", "Mendes", "Freitas",
    "Cardoso", "Correia", "Melo", "Cavalcanti", "Castro", "Araujo", "Cunha"
]

MUSIC_GENRES = [
    "Rock", "Pop", "Hip Hop", "Jazz", "Classical", "Electronic", "Country",
    "R&B", "Reggae", "Blues", "Folk", "Metal", "Punk", "Indie", "Alternative",
    "Funk", "Samba", "Bossa Nova", "MPB", "Forró", "Axé", "Pagode", "Sertanejo"
]

ARTIST_NAMES = [
    "The Beatles", "Queen", "Pink Floyd", "Led Zeppelin", "Rolling Stones",
    "Bob Dylan", "David Bowie", "Elvis Presley", "Michael Jackson", "Madonna",
    "Prince", "Stevie Wonder", "Aretha Franklin", "James Brown", "Ray Charles",
    "John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr",
    "Freddie Mercury", "Roger Waters", "Jimmy Page", "Mick Jagger",
    "Keith Richards", "Eric Clapton", "Jimi Hendrix", "Janis Joplin",
    "Joni Mitchell", "Bob Marley", "Tupac Shakur", "Notorious B.I.G.",
    "Eminem", "Jay-Z", "Kanye West", "Drake", "Taylor Swift", "Adele",
    "Ed Sheeran", "Bruno Mars", "Lady Gaga", "Beyoncé", "Rihanna",
    "Justin Timberlake", "Coldplay", "Radiohead", "Nirvana", "Pearl Jam",
    "Soundgarden", "Alice in Chains", "Stone Temple Pilots", "Red Hot Chili Peppers"
]

COUNTRIES = ["BR", "US", "UK", "CA", "AU", "DE", "FR", "IT", "ES", "JP", "KR", "MX", "AR", "CL", "CO"]

COMMUNITY_NAMES = [
    "Rock Lovers", "Jazz Enthusiasts", "Classical Music Society", "Hip Hop Nation",
    "Electronic Music Fans", "Country Music Community", "Blues Brothers",
    "Metalheads United", "Pop Music Fans", "Indie Music Collective",
    "Funk & Soul", "Samba & Bossa Nova", "MPB Brasil", "Forró Dance Club",
    "Axé Music Fans", "Pagode & Samba", "Sertanejo Country", "Reggae Vibes",
    "Punk Rockers", "Alternative Music", "Folk Music Circle", "R&B Soul",
    "Gospel Music", "World Music", "Instrumental Music", "Acoustic Sessions",
    "Cover Bands", "Songwriters Guild", "Music Producers", "DJ Community",
    "Music Critics", "Vinyl Collectors", "Concert Goers", "Music Students"
]

POST_TEXTS = [
    "Just discovered this amazing new artist! 🎵",
    "Can't stop listening to this song! 🔥",
    "What's everyone's favorite album right now?",
    "Going to a concert tonight! So excited! 🎤",
    "This band changed my life! 💫",
    "New music Friday! What should I listen to?",
    "Remembering this classic song today 🎶",
    "Music is the universal language of the soul",
    "Just finished recording my first song! 🎸",
    "Who else loves this genre?",
    "This song brings back so many memories",
    "Music therapy is real! ��",
    "What's your go-to song when you're feeling down?",
    "This artist deserves more recognition!",
    "Music festival season is here! ��",
    "Learning to play guitar! Any tips? 🎸",
    "This album is a masterpiece! 👑",
    "Music connects us all 🌍",
    "What's the best concert you've ever been to?",
    "This song gets me through tough times ��",
    "Music is my escape from reality",
    "Who else is obsessed with this band?",
    "This song is stuck in my head! 🎵",
    "Music has the power to heal ❤️",
    "What's your favorite music decade?",
    "This artist is underrated!",
    "Music brings people together ��",
    "What's your current playlist?",
    "This song makes me want to dance! 💃",
    "Music is life! 🎼",
    "Who else loves live music?",
    "This band is incredible live!",
    "Music is the soundtrack of our lives 🎬"
]

def generate_random_user():
    """Generate a random user with realistic data"""
    first_name = random.choice(FIRST_NAMES)
    last_name = random.choice(LAST_NAMES)
    username = f"{first_name.lower()}{last_name.lower()}{random.randint(1, 999)}"
    email = f"{username}@example.com"
    birth_year = random.randint(1980, 2005)
    birth_month = random.randint(1, 12)
    birth_day = random.randint(1, 28)
    birth_date = f"{birth_year}-{birth_month:02d}-{birth_day:02d}T00:00:00Z"
    
    return {
        "username": username,
        "fullname": f"{first_name} {last_name}",
        "email": email,
        "telephone": f"+55{random.randint(10, 99)}{random.randint(10000000, 99999999)}",
        "birth_date": birth_date
    }

def generate_random_artist():
    """Generate a random artist with realistic data"""
    name = random.choice(ARTIST_NAMES)
    country = random.choice(COUNTRIES)
    genres = random.sample(MUSIC_GENRES, random.randint(1, 3))
    
    return {
        "name": name,
        "description": f"Amazing {genres[0].lower()} artist from {country}",
        "country": country,
        "biography": f"{name} is a talented musician known for their unique style and innovative approach to music. They have been active in the music industry for over a decade and have released multiple successful albums.",
        "genres": genres,
        "id_spotify": f"spotify_artist_{random.randint(1000000, 9999999)}",
        "image_url": f"https://example.com/artists/{name.lower().replace(' ', '_')}.jpg"
    }

def generate_random_song(artist_id):
    """Generate a random song with realistic data"""
    titles = [
        "Midnight Dreams", "Electric Love", "Ocean Waves", "City Lights",
        "Mountain High", "Desert Wind", "Starry Night", "Golden Hour",
        "Silver Moon", "Crystal Clear", "Deep Blue", "Fire and Ice",
        "Thunder Road", "Silent Echo", "Wild Heart", "Gentle Soul",
        "Brave New World", "Ancient Times", "Future Days", "Present Moment",
        "Lost in Time", "Found in Love", "Breaking Free", "Coming Home",
        "Rising Sun", "Setting Moon", "Morning Light", "Evening Star"
    ]
    
    albums = [
        "First Album", "Second Chance", "New Beginnings", "The Journey",
        "Life Stories", "Dreams and Reality", "Sound of Silence",
        "Voice of the People", "Heart and Soul", "Mind and Body",
        "Spirit of Music", "Rhythm of Life", "Melody of Love",
        "Harmony of Nature", "Symphony of Emotions", "Concerto of Dreams"
    ]
    
    return {
        "title": random.choice(titles),
        "duration": random.randint(120, 360),
        "artist_id": artist_id,
        "genre": random.choice(MUSIC_GENRES),
        "release_year": random.randint(1990, 2024),
        "album": random.choice(albums),
        "id_spotify": f"spotify_song_{random.randint(1000000, 9999999)}",
        "url_spotify": f"https://open.spotify.com/track/{random.randint(1000000000000000000, 9999999999999999999)}"
    }

def generate_random_playlist(user_id, song_ids):
    """Generate a random playlist with realistic data"""
    playlist_names = [
        "My Favorites", "Workout Mix", "Chill Vibes", "Party Time",
        "Study Music", "Road Trip", "Late Night", "Morning Coffee",
        "Rainy Day", "Sunny Afternoon", "Weekend Vibes", "Holiday Spirit",
        "Summer Hits", "Winter Warmth", "Spring Awakening", "Autumn Colors",
        "Rock Classics", "Jazz Lounge", "Hip Hop Beats", "Electronic Dreams",
        "Country Roads", "Blues Night", "Folk Tales", "Metal Mayhem",
        "Pop Hits", "Indie Gems", "Alternative Rock", "R&B Soul",
        "Reggae Vibes", "Classical Masterpieces", "World Music", "Acoustic Sessions"
    ]
    
    playlist_songs = []
    if song_ids:
        num_songs = random.randint(1, min(10, len(song_ids)))
        selected_songs = random.sample(song_ids, num_songs)
        for i, song_id in enumerate(selected_songs):
            playlist_songs.append({
                "song_id": song_id,
                "order": i + 1
            })
    
    return {
        "name": random.choice(playlist_names),
        "description": f"A curated collection of amazing music",
        "user_id": user_id,
        "public": random.choice([True, False]),
        "id_spotify": f"spotify_playlist_{random.randint(1000000, 9999999)}",
        "title": random.choice(playlist_names),
        "image_url": f"https://example.com/playlists/playlist_{random.randint(1, 100)}.jpg",
        "songs": playlist_songs
    }

def generate_random_post(user_id):
    """Generate a random post with realistic data"""
    return {
        "user_id": user_id,
        "text": random.choice(POST_TEXTS),
        "url_foto": f"https://example.com/posts/post_{random.randint(1, 100)}.jpg"
    }

def generate_random_community():
    """Generate a random community with realistic data"""
    name = random.choice(COMMUNITY_NAMES)
    return {
        "community_name": name,
        "description": f"A community for {name.lower()} enthusiasts to share and discover music together."
    }

def make_request(method, url, data=None):
    """Make HTTP request with error handling"""
    try:
        if method == "GET":
            response = requests.get(url)
        elif method == "POST":
            response = requests.post(url, json=data, headers={"Content-Type": "application/json"})
        
        if response.status_code in [200, 201]:
            return response.json() if response.content else None
        else:
            print(f"❌ Error {response.status_code}: {response.text}")
            return None
    except requests.exceptions.RequestException as e:
        print(f"❌ Request failed: {e}")
        return None

def populate_postgresql():
    """Populate PostgreSQL database with users, posts, communities, and chats"""
    print("\n🗄️  Populating PostgreSQL Database...")
    
    users = []
    communities = []
    posts = []
    chats = []
    
    # Create users
    print("👥 Creating users...")
    for i in range(NUM_USERS):
        user_data = generate_random_user()
        response = make_request("POST", f"{API_BASE_URL}/api/user/create", user_data)
        if response:
            # Store the original user data since the API only returns a success message
            user_data["id"] = i + 1  # We'll use the index as ID for now
            users.append(user_data)
            print(f"✅ Created user: {user_data['username']}")
        time.sleep(0.1)  # Rate limiting
    
    # Create communities
    print("\n🏘️  Creating communities...")
    for i in range(NUM_COMMUNITIES):
        community_data = generate_random_community()
        response = make_request("POST", f"{API_BASE_URL}/api/community/create", community_data)
        if response:
            communities.append(community_data)
            print(f"✅ Created community: {community_data['community_name']}")
        time.sleep(0.1)
    
    # Add users to communities
    print("\n👥 Adding users to communities...")
    for community in communities:
        num_members = random.randint(5, min(20, len(users)))
        selected_users = random.sample(users, num_members)
        for user in selected_users:
            data = {
                "username": user["username"],
                "community_name": community["community_name"]
            }
            make_request("POST", f"{API_BASE_URL}/api/community/add_user", data)
        time.sleep(0.1)
    
    # Create posts
    print("\n📝 Creating posts...")
    for i in range(NUM_POSTS):
        user = random.choice(users)
        post_data = generate_random_post(user["id"])
        response = make_request("POST", f"{API_BASE_URL}/api/create-post", post_data)
        if response:
            posts.append(response)
            print(f"✅ Created post {i+1}/{NUM_POSTS}")
        time.sleep(0.1)
    
    # Create chats and friendships
    print("\n�� Creating chats and friendships...")
    for i in range(NUM_CHATS):
        user1, user2 = random.sample(users, 2)
        
        # Create friendship in Neo4j
        friendship_data = {
            "username1": user1["username"],
            "username2": user2["username"]
        }
        make_request("POST", f"{API_BASE_URL}/api/user/create_friendship", friendship_data)
        
        # Create chat
        chat_data = {
            "username1": user1["username"],
            "username2": user2["username"]
        }
        response = make_request("POST", f"{API_BASE_URL}/api/chat/create", chat_data)
        if response:
            chats.append(response)
            print(f"✅ Created chat {i+1}/{NUM_CHATS}")
        time.sleep(0.1)
    
    # Add genre preferences to users
    print("\n🎵 Adding genre preferences...")
    for user in users:
        num_genres = random.randint(1, 5)
        selected_genres = random.sample(MUSIC_GENRES, num_genres)
        for genre in selected_genres:
            data = {
                "username": user["username"],
                "genre_name": genre
            }
            make_request("POST", f"{API_BASE_URL}/api/user/like_genre", data)
        time.sleep(0.1)
    
    return users, communities, posts, chats

def populate_mongodb():
    """Populate MongoDB database with artists, songs, and playlists"""
    print("\n�� Populating MongoDB Database...")
    
    artists = []
    songs = []
    playlists = []
    
    # Create artists
    print("🎤 Creating artists...")
    for i in range(NUM_ARTISTS):
        artist_data = generate_random_artist()
        response = make_request("POST", f"{API_BASE_URL}/artists", artist_data)
        if response:
            artists.append(response)
            print(f"✅ Created artist: {artist_data['name']}")
        time.sleep(0.1)
    
    # Create songs
    print("\n🎵 Creating songs...")
    for i in range(NUM_SONGS):
        artist = random.choice(artists)
        song_data = generate_random_song(artist["_id"])
        response = make_request("POST", f"{API_BASE_URL}/songs", song_data)
        if response:
            songs.append(response)
            print(f"✅ Created song {i+1}/{NUM_SONGS}: {song_data['title']}")
        time.sleep(0.1)
    
    # Create playlists
    print("\n�� Creating playlists...")
    song_ids = [song["_id"] for song in songs]
    for i in range(NUM_PLAYLISTS):
        user_id = f"user_{random.randint(1, NUM_USERS)}"
        playlist_data = generate_random_playlist(user_id, song_ids)
        response = make_request("POST", f"{API_BASE_URL}/playlists", playlist_data)
        if response:
            playlists.append(response)
            print(f"✅ Created playlist {i+1}/{NUM_PLAYLISTS}: {playlist_data['name']}")
        time.sleep(0.1)
    
    return artists, songs, playlists

def main():
    """Main function to populate all databases"""
    print("🎵 Symphony API Database Population Script")
    print("=" * 50)
    print(f"Target API: {API_BASE_URL}")
    print(f"Users to create: {NUM_USERS}")
    print(f"Artists to create: {NUM_ARTISTS}")
    print(f"Songs to create: {NUM_SONGS}")
    print(f"Playlists to create: {NUM_PLAYLISTS}")
    print(f"Posts to create: {NUM_POSTS}")
    print(f"Communities to create: {NUM_COMMUNITIES}")
    print(f"Chats to create: {NUM_CHATS}")
    print("=" * 50)
    
    # Check if API is running
    try:
        response = requests.get(f"{API_BASE_URL}/")
        if response.status_code != 200:
            print("❌ API is not responding properly. Please make sure the server is running.")
            sys.exit(1)
    except requests.exceptions.RequestException:
        print("❌ Cannot connect to API. Please make sure the server is running on localhost:8080")
        sys.exit(1)
    
    print("✅ API is running and accessible")
    
    # Populate databases
    start_time = time.time()
    
    # PostgreSQL + Neo4j (users, posts, communities, chats, friendships, genre preferences)
    users, communities, posts, chats = populate_postgresql()
    
    # MongoDB (artists, songs, playlists)
    artists, songs, playlists = populate_mongodb()
    
    end_time = time.time()
    
    # Summary
    print("\n" + "=" * 50)
    print("🎉 DATABASE POPULATION COMPLETED!")
    print("=" * 50)
    print(f"⏱️  Total time: {end_time - start_time:.2f} seconds")
    print(f"👥 Users created: {len(users)}")
    print(f"🏘️  Communities created: {len(communities)}")
    print(f"📝 Posts created: {len(posts)}")
    print(f"💬 Chats created: {len(chats)}")
    print(f"�� Artists created: {len(artists)}")
    print(f"🎵 Songs created: {len(songs)}")
    print(f"�� Playlists created: {len(playlists)}")
    print("\n📊 Database Summary:")
    print("   • PostgreSQL: Users, Posts, Communities, Chats")
    print("   • Neo4j: User relationships, Friendships, Genre preferences")
    print("   • MongoDB: Artists, Songs, Playlists")
    print("\n🚀 Your Symphony API is now populated with realistic data!")
    print("   Ready for your university presentation! 🎓")

if __name__ == "__main__":
    main()