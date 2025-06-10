#!/bin/bash

set -e

timestamp=$(date +%s)
community_name="test_community_$timestamp"
username="test_user_$timestamp"

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

    if [ "$http_code" -eq 200 ]; then
        echo "✅ Success: $description, Response: $http_body"
    else
        echo "❌ Failed: $description (HTTP $http_code), Response: $http_body"
        exit 1
    fi
}

post_and_assert "http://localhost:8080/api/community/create" "{
    \"community_name\": \"$community_name\",
    \"description\": \"test\"
}" "Create community"

post_and_assert "http://localhost:8080/api/community/get_by_name" "{
    \"community_name\": \"$community_name\"
}" "Get community by name"

post_and_assert "http://localhost:8080/api/user/create" "{
    \"username\": \"$username\",
    \"fullname\": \"John Doe 2\",
    \"email\": \"${username}@example.com\",
    \"telephone\": \"123456789\",
    \"birth_date\": \"2002-01-01T00:00:00Z\"
}" "Create user"

post_and_assert "http://localhost:8080/api/user/get_by_username" "{
    \"username\": \"$username\"
}" "Get user by username"

post_and_assert "http://localhost:8080/api/community/add_user" "{
    \"username\": \"$username\",
    \"community_name\": \"$community_name\"
}" "Add user to community"

post_and_assert "http://localhost:8080/api/community/list_users" "{
    \"community_name\": \"$community_name\"
}" "List users in community"

post_and_assert "http://localhost:8080/api/user/list_communities" "{
    \"username\": \"$username\"
}" "List user communities"

post_and_assert "http://localhost:8080/api/create-post" '{
    "user_id": 1, "text": "Hello world", "url_foto": "image.jpg"
}' "Create post"

post_and_assert "http://localhost:8080/api/get-post-by-id?post_id=1" "{}" "Get post"

post_and_assert "http://localhost:8080/api/get-posts-by-user-id?user_id=1" "{}" "List user posts"

echo "🎉 All tests passed successfully!"
