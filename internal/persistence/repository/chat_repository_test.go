package repository

import (
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "symphony-api/internal/persistence/model"
)

func TestPutChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    input := &model.Chat{
        ChatId:    1,
        CreatedAt: time.Now(),
    }

    mockConn.On("PutReturningId", input.ToMap(), CHAT_TABLE_NAME, "chat_id").Return(int32(1), nil)

    err := repo.Put(input)

    assert.NoError(t, err)
    mockConn.AssertExpectations(t)
}

func TestPutChat_Failure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    input := &model.Chat{
        ChatId:    1,
        CreatedAt: time.Now(),
    }

    mockConn.On("PutReturningId", input.ToMap(), CHAT_TABLE_NAME, "chat_id").Return(nil, errors.New("db error"))

    err := repo.Put(input)

    assert.Error(t, err)
    mockConn.AssertExpectations(t)
}

func TestGetByChatId_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    constraint := map[string]any{
        "chat_id": int32(1),
    }

    dbResult := []map[string]any{
        {
            "chat_id":    int32(1),
            "created_at": time.Now(),
        },
    }

    mockConn.On("Get", constraint, CHAT_TABLE_NAME).Return(dbResult, nil)

    result, err := repo.GetByChatId(1)

    assert.NoError(t, err)
    assert.Equal(t, int32(1), result.ChatId)
    mockConn.AssertExpectations(t)
}

func TestGetByChatId_NotFound(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    mockConn.On("Get", mock.Anything, CHAT_TABLE_NAME).Return([]map[string]any{}, nil)

    result, err := repo.GetByChatId(999)

    assert.Error(t, err)
    assert.Nil(t, result)
    mockConn.AssertExpectations(t)
}

func TestAddUserToChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    user := &model.User{UserId: 2}
    chat := &model.Chat{ChatId: 1}

    expectedMap := map[string]any{
        "chat_id": chat.ChatId,
        "user_id": user.UserId,
    }

    mockConn.On("Put", expectedMap, USER_TO_CHAT_TABLE).Return(nil)

    err := repo.AddUserToChat(user, chat)
    assert.NoError(t, err)
    mockConn.AssertExpectations(t)
}

func TestAddUserToChat_Failure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    user := &model.User{UserId: 2}
    chat := &model.Chat{ChatId: 1}

    expectedMap := map[string]any{
        "chat_id": chat.ChatId,
        "user_id": user.UserId,
    }

    mockConn.On("Put", expectedMap, USER_TO_CHAT_TABLE).Return(errors.New("db error"))

    err := repo.AddUserToChat(user, chat)
    assert.Error(t, err)
    mockConn.AssertExpectations(t)
}

func TestListUsersFromChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    chat := &model.Chat{ChatId: 1}

    dbResult := []map[string]any{
        {
            "id":   int32(2),
            "username":  "user1",
            "fullname": "fulano da silva",
            "email": "email",
            "register_date": time.Now(),
            "birth_date": time.Now(),
            "telephone": "telephone",
        },
    }

    mockConn.On("Get", mock.Anything, JOINED_USERS_AND_CHAT_PARTICIPANTS).Return(dbResult, nil)

    users, err := repo.ListUsersFromChat(chat)
    assert.NoError(t, err)
    assert.Len(t, users, 1)
    assert.Equal(t, "user1", users[0].Username)
    mockConn.AssertExpectations(t)
}

func TestAddMessageToChatAndReturn_GetFailure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    chatId := int32(3)
    authorId := int32(2)
    message := "Hello"
    fakeMsgId := int32(99)

    data := map[string]any{
        "chat_id":  chatId,
        "author_id": authorId,
        "message":   message,
    }
    constraint := map[string]any{
        "message_id": fakeMsgId,
    }

    mockConn.On("PutReturningId", data, CHAT_MESSAGE_TABLE, "message_id").Return(fakeMsgId, nil)
    mockConn.On("Get", constraint, CHAT_MESSAGE_TABLE).Return([]map[string]any{}, nil)

    msg, err := repo.AddMessageToChatAndReturn(chatId, authorId, message)
    assert.Error(t, err)
    assert.Nil(t, msg)
    mockConn.AssertExpectations(t)
}

func TestListMessagesFromChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    chatId := int32(3)
    limit := int32(2)
    now := time.Now()

    dbResult := []map[string]any{
        {
            "message_id": int32(1),
            "author_id":  int32(2),
            "chat_id":    chatId,
            "message":    "Hello",
            "sent_at":    now,
        },
        {
            "message_id": int32(2),
            "author_id":  int32(2),
            "chat_id":    chatId,
            "message":    "World",
            "sent_at":    now.Add(-time.Minute),
        },
    }

    mockConn.On("GetChatWithLimit", chatId, limit, CHAT_MESSAGE_TABLE).Return(dbResult, nil)

    messages, err := repo.ListMessagesFromChat(chatId, limit)
    assert.NoError(t, err)
    assert.Len(t, messages, 2)
    assert.Equal(t, "Hello", messages[0].Message)
    assert.Equal(t, "World", messages[1].Message)
    mockConn.AssertExpectations(t)
}

func TestListMessagesFromChat_Failure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    chatId := int32(3)
    limit := int32(2)

    mockConn.On("GetChatWithLimit", chatId, limit, CHAT_MESSAGE_TABLE).Return([]map[string]any{}, errors.New("db error"))
    messages, err := repo.ListMessagesFromChat(chatId, limit)
    assert.Error(t, err)
    assert.Nil(t, messages)
    mockConn.AssertExpectations(t)
}