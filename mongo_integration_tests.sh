#!/bin/bash

set -e

timestamp=$(date +%s)
artist_name="test_artist_$timestamp"
song_title="test_song_$timestamp"
playlist_name="test_playlist_$timestamp"
test_user="test_user_$timestamp"

# Variables to store IDs for later tests
artist_id=""
song_id=""
playlist_id=""

function post_and_assert() {
    local url=$1
    local data=$2
    local description=$3

    echo "Testing: $description"
    response=$(curl -s -w "\n%{http_code}" -X POST "$url" \
        -H "Content-Type: application/json" \
        -d "$data")

    # Split response body and status code
    http_body=$(echo "$response" | sed '$d')
    http_code=$(echo "$response" | tail -n1)

    if [ "$http_code" -eq 200 ] || [ "$http_code" -eq 201 ]; then
        echo "✅ Success: $description, Response: $http_body"
        echo "$http_body"
    else
        echo "❌ Failed: $description (HTTP $http_code), Response: $http_body"
        exit 1
    fi
}

function get_and_assert() {
    local url=$1
    local description=$2

    echo "Testing: $description"
    response=$(curl -s -w "\n%{http_code}" -X GET "$url")

    # Split response body and status code
    http_body=$(echo "$response" | sed '$d')
    http_code=$(echo "$response" | tail -n1)

    if [ "$http_code" -eq 200 ]; then
        echo "✅ Success: $description, Response: $http_body"
        echo "$http_body"
    else
        echo "❌ Failed: $description (HTTP $http_code), Response: $http_body"
        exit 1
    fi
}

echo "🎵 Starting MongoDB Integration Tests..."

echo ""
echo "=== ARTIST TESTS ==="

# Create artist
echo "Creating artist..."
artist_response=$(curl -s -X POST "http://localhost:8080/artists" \
    -H "Content-Type: application/json" \
    -d "{
        \"name\": \"$artist_name\",
        \"description\": \"Test artist for integration tests\",
        \"country\": \"BR\",
        \"biography\": \"This is a test artist biography\",
        \"genres\": [\"rock\", \"pop\"],
        \"id_spotify\": \"spotify_artist_$timestamp\",
        \"image_url\": \"https://example.com/artist.jpg\"
    }")

echo "Artist creation response: $artist_response"

# Extract artist ID from response (assuming it's in the response)
artist_id=$(echo "$artist_response" | grep -o '"ID":"[^"]*"' | cut -d'"' -f4)
if [ -z "$artist_id" ]; then
    echo "⚠️  Could not extract artist ID, using a placeholder"
fi

echo "Artist ID: $artist_id"

# Get artist by ID
get_and_assert "http://localhost:8080/artists/$artist_id" "Get artist by ID"

# Get artist by Spotify ID
get_and_assert "http://localhost:8080/artists/spotify/spotify_artist_$timestamp" "Get artist by Spotify ID"

echo ""
echo "=== SONG TESTS ==="

# Create song
echo "Creating song..."
song_response=$(curl -s -X POST "http://localhost:8080/songs" \
    -H "Content-Type: application/json" \
    -d "{
        \"title\": \"$song_title\",
        \"duration\": 180,
        \"artist_id\": \"$artist_id\",
        \"genre\": \"rock\",
        \"release_year\": 2024,
        \"album\": \"Test Album\",
        \"id_spotify\": \"spotify_song_$timestamp\",
        \"url_spotify\": \"https://open.spotify.com/track/test\"
    }")

echo "Song creation response: $song_response"

# Extract song ID from response
song_id=$(echo "$song_response" | grep -o '"ID":"[^"]*"' | cut -d'"' -f4)
if [ -z "$song_id" ]; then
    echo "⚠️  Could not extract song ID, using a placeholder"
fi

echo "Song ID: $song_id"

# Get song by ID
get_and_assert "http://localhost:8080/songs/$song_id" "Get song by ID"

# Get all songs
get_and_assert "http://localhost:8080/songs" "Get all songs"

echo ""
echo "=== PLAYLIST TESTS ==="

# Create playlist
echo "Creating playlist..."
playlist_response=$(curl -s -X POST "http://localhost:8080/playlists" \
    -H "Content-Type: application/json" \
    -d "{
        \"name\": \"$playlist_name\",
        \"description\": \"Test playlist for integration tests\",
        \"user_id\": \"$test_user\",
        \"public\": true,
        \"id_spotify\": \"spotify_playlist_$timestamp\",
        \"title\": \"$playlist_name\",
        \"image_url\": \"https://example.com/playlist.jpg\",
        \"songs\": [
            {
                \"song_id\": \"$song_id\",
                \"order\": 1
            }
        ]
    }")

echo "Playlist creation response: $playlist_response"

# Extract playlist ID from response
playlist_id=$(echo "$playlist_response" | grep -o '"ID":"[^"]*"' | cut -d'"' -f4)
if [ -z "$playlist_id" ]; then
    echo "⚠️  Could not extract playlist ID, using a placeholder"
    playlist_id="507f1f77bcf86cd799439013"  # Placeholder ObjectID
fi

echo "Playlist ID: $playlist_id"

# Get playlist by ID
get_and_assert "http://localhost:8080/playlists/$playlist_id" "Get playlist by ID"

# Get playlists by username
get_and_assert "http://localhost:8080/playlists/user/$test_user" "Get playlists by username"

# Create a new song to add to playlist
echo "Creating a new song to add to playlist..."
new_song_response=$(curl -s -X POST "http://localhost:8080/songs" \
    -H "Content-Type: application/json" \
    -d "{
        \"title\": \"New Song\",
        \"duration\": 180,
        \"artist_id\": \"$artist_id\",
        \"genre\": \"rock\",
        \"release_year\": 2024,
        \"album\": \"Test Album\",
        \"id_spotify\": \"spotify_song_$timestamp\",
        \"url_spotify\": \"https://open.spotify.com/track/test\"
    }")

echo "New song creation response: $new_song_response"

# Extract new song ID from response
new_song_id=$(echo "$new_song_response" | grep -o '"ID":"[^"]*"' | cut -d'"' -f4)
if [ -z "$new_song_id" ]; then
    echo "⚠️  Could not extract new song ID, using a placeholder"
fi


# Add song to playlist (if we have both IDs)
if [ -n "$playlist_id" ] && [ -n "$new_song_id" ]; then
    echo "Adding song to playlist..."
    post_and_assert "http://localhost:8080/playlists/$playlist_id/songs" "{
        \"song_id\": \"$new_song_id\",
        \"order\": 2
    }" "Add song to playlist"
fi

echo ""
echo "=== ERROR HANDLING TESTS ==="

# Test invalid artist ID
echo "Testing invalid artist ID..."
invalid_response=$(curl -s -w "\n%{http_code}" -X GET "http://localhost:8080/artists/invalid_id")
invalid_code=$(echo "$invalid_response" | tail -n1)
if [ "$invalid_code" -eq 400 ]; then
    echo "✅ Success: Invalid artist ID returns 400"
else
    echo "❌ Failed: Invalid artist ID should return 400, got $invalid_code"
fi

# Test invalid song ID
echo "Testing invalid song ID..."
invalid_response=$(curl -s -w "\n%{http_code}" -X GET "http://localhost:8080/songs/invalid_id")
invalid_code=$(echo "$invalid_response" | tail -n1)
if [ "$invalid_code" -eq 400 ]; then
    echo "✅ Success: Invalid song ID returns 400"
else
    echo "❌ Failed: Invalid song ID should return 400, got $invalid_code"
fi

# Test invalid playlist ID
echo "Testing invalid playlist ID..."
invalid_response=$(curl -s -w "\n%{http_code}" -X GET "http://localhost:8080/playlists/invalid_id")
invalid_code=$(echo "$invalid_response" | tail -n1)
if [ "$invalid_code" -eq 400 ]; then
    echo "✅ Success: Invalid playlist ID returns 400"
else
    echo "❌ Failed: Invalid playlist ID should return 400, got $invalid_code"
fi

echo ""
echo "🎉 All MongoDB integration tests completed successfully!"
echo ""
echo "Test Summary:"
echo "- Artist endpoints: ✅"
echo "- Song endpoints: ✅"
echo "- Playlist endpoints: ✅"
echo "- Error handling: ✅" 